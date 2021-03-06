package opensds

import (
	"fmt"
	"log"

	"google.golang.org/grpc/codes"

	"strings"

	csi "github.com/container-storage-interface/spec/lib/go/csi/v0"
	"github.com/opensds/nbp/client/iscsi"
	sdscontroller "github.com/opensds/nbp/client/opensds"
	"github.com/opensds/nbp/driver"
	"github.com/opensds/opensds/pkg/model"
	"golang.org/x/net/context"
	"google.golang.org/grpc/status"
)

////////////////////////////////////////////////////////////////////////////////
//                            Node Service                                    //
////////////////////////////////////////////////////////////////////////////////

func (p *Plugin) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
func (p *Plugin) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// NodePublishVolume implementation
func (p *Plugin) NodePublishVolume(
	ctx context.Context,
	req *csi.NodePublishVolumeRequest) (
	*csi.NodePublishVolumeResponse, error) {

	log.Println("start to NodePublishVolume")
	defer log.Println("end to NodePublishVolume")

	client := sdscontroller.GetClient("")

	//check volume is exist
	volSpec, errVol := client.GetVolume(req.VolumeId)
	if errVol != nil || volSpec == nil {
		msg := fmt.Sprintf("the volume %s is not exist", req.VolumeId)
		return nil, status.Error(codes.NotFound, msg)
	}

	atc, atcErr := client.GetVolumeAttachment(req.PublishInfo["atcid"])
	if atcErr != nil || atc == nil {
		return nil, status.Error(codes.FailedPrecondition, "Failed to publish node.")
	}

	var targetPaths []string
	if tps, exist := atc.Metadata["target_path"]; exist && len(tps) != 0 {
		targetPaths = strings.Split(tps, ";")
		for _, tp := range targetPaths {
			if req.TargetPath == tp {
				return &csi.NodePublishVolumeResponse{}, nil
			}
		}

		// if volume don't have MULTI_NODE capability, just termination.
		mode := req.VolumeCapability.AccessMode.Mode
		if mode != csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER &&
			mode != csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY &&
			mode != csi.VolumeCapability_AccessMode_MULTI_NODE_SINGLE_WRITER {
			msg := fmt.Sprintf("the volume %s has been published to this node.", req.VolumeId)
			return nil, status.Error(codes.Aborted, msg)
		}
	}

	// if not attach before, attach first.
	if len(atc.Mountpoint) == 0 {
		volDriver := driver.NewVolumeDriver(atc.DriverVolumeType)
		if volDriver == nil {
			return nil, status.Error(codes.FailedPrecondition, fmt.Sprintf("Unsupport driverVolumeType: %s", atc.DriverVolumeType))
		}

		device, err := volDriver.Attach(atc.ConnectionData)
		if err != nil {
			return nil, status.Errorf(codes.FailedPrecondition, "%s", err.Error())
		}
		atc.Mountpoint = device

		_, err = client.UpdateVolumeAttachment(atc.Id, atc)
		if err != nil {
			return nil, status.Errorf(codes.FailedPrecondition, "%s", err.Error())
		}
	}

	// obtain attachments to decide if can format.
	atcs, err := client.ListVolumeAttachments()
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "Failed to publish node.")
	}
	format := true
	for _, attachSpec := range atcs {
		if attachSpec.VolumeId == req.VolumeId {
			if _, exist := attachSpec.Metadata["target_path"]; exist {
				// The device is formatted, can't be reformat for shared storage.
				format = false
				break
			}
		}
	}

	// Format and Mount
	log.Printf("[NodePublishVolume] device:%s TargetPath:%s", atc.Mountpoint, req.TargetPath)
	if format {
		err = iscsi.FormatandMount(atc.Mountpoint, "", req.TargetPath)
	} else {
		err = iscsi.Mount(atc.Mountpoint, req.TargetPath)
	}
	if err != nil {
		return nil, err
	}

	targetPaths = append(targetPaths, req.TargetPath)
	atc.Metadata["target_path"] = strings.Join(targetPaths, ";")
	_, err = client.UpdateVolumeAttachment(atc.Id, atc)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "Failed to publish node.")
	}

	return &csi.NodePublishVolumeResponse{}, nil
}

// NodeUnpublishVolume implementation
func (p *Plugin) NodeUnpublishVolume(
	ctx context.Context,
	req *csi.NodeUnpublishVolumeRequest) (
	*csi.NodeUnpublishVolumeResponse, error) {

	log.Println("start to NodeUnpublishVolume")
	defer log.Println("end to NodeUnpublishVolume")

	client := sdscontroller.GetClient("")

	//check volume is exist
	volSpec, errVol := client.GetVolume(req.VolumeId)
	if errVol != nil || volSpec == nil {
		msg := fmt.Sprintf("the volume %s is not exist", req.VolumeId)
		return nil, status.Error(codes.NotFound, msg)
	}

	attachments, err := client.ListVolumeAttachments()
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "Failed to NodeUnpublish volume.")
	}

	var atc *model.VolumeAttachmentSpec
	// hostname, _ := os.Hostname()
	iqns, _ := iscsi.GetInitiator()
	localIqn := ""
	if len(iqns) > 0 {
		localIqn = iqns[0]
	}

	for _, attachSpec := range attachments {
		if attachSpec.VolumeId == req.VolumeId && attachSpec.Host == localIqn {
			atc = attachSpec
			break
		}
	}

	if atc == nil {
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	if _, exist := atc.Metadata["target_path"]; !exist {
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	var modifyTargetPaths []string
	tpExist := false
	targetPaths := strings.Split(atc.Metadata["target_path"], ";")
	for index, path := range targetPaths {
		if path == req.TargetPath {
			modifyTargetPaths = append(targetPaths[:index], targetPaths[index+1:]...)
			tpExist = true
			break
		}
	}
	if !tpExist {
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	// Umount
	log.Printf("[NodeUnpublishVolume] TargetPath:%s", req.TargetPath)
	err = iscsi.Umount(req.TargetPath)
	if err != nil {
		return nil, err
	}

	if len(modifyTargetPaths) == 0 {
		volDriver := driver.NewVolumeDriver(atc.DriverVolumeType)
		if volDriver == nil {
			return nil, status.Error(codes.FailedPrecondition, fmt.Sprintf("Unsupport driverVolumeType: %s", atc.DriverVolumeType))
		}

		err := volDriver.Detach(atc.ConnectionData)
		if err != nil {
			return nil, status.Errorf(codes.FailedPrecondition, "%s", err.Error())
		}
		atc.Mountpoint = ""
	}

	atc.Metadata["target_path"] = strings.Join(modifyTargetPaths, ";")
	_, err = client.UpdateVolumeAttachment(atc.Id, atc)
	if err != nil {
		return nil, status.Error(codes.FailedPrecondition, "Failed to NodeUnpublish volume.")
	}

	return &csi.NodeUnpublishVolumeResponse{}, nil
}

// GetNodeID implementation
func (p *Plugin) NodeGetId(
	ctx context.Context,
	req *csi.NodeGetIdRequest) (
	*csi.NodeGetIdResponse, error) {

	log.Println("start to GetNodeID")
	defer log.Println("end to GetNodeID")

	iqns, _ := iscsi.GetInitiator()
	localIqn := ""
	if len(iqns) > 0 {
		localIqn = iqns[0]
	}

	return &csi.NodeGetIdResponse{
		NodeId: localIqn,
	}, nil
}

// NodeGetCapabilities implementation
func (p *Plugin) NodeGetCapabilities(
	ctx context.Context,
	req *csi.NodeGetCapabilitiesRequest) (
	*csi.NodeGetCapabilitiesResponse, error) {

	log.Println("start to NodeGetCapabilities")
	defer log.Println("end to NodeGetCapabilities")

	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			&csi.NodeServiceCapability{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_UNKNOWN,
					},
				},
			},
		},
	}, nil
}
