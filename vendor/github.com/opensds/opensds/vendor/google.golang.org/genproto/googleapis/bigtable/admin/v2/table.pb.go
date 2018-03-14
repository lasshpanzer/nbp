// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/bigtable/admin/v2/table.proto

package admin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf4 "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Possible timestamp granularities to use when keeping multiple versions
// of data in a table.
type Table_TimestampGranularity int32

const (
	// The user did not specify a granularity. Should not be returned.
	// When specified during table creation, MILLIS will be used.
	Table_TIMESTAMP_GRANULARITY_UNSPECIFIED Table_TimestampGranularity = 0
	// The table keeps data versioned at a granularity of 1ms.
	Table_MILLIS Table_TimestampGranularity = 1
)

var Table_TimestampGranularity_name = map[int32]string{
	0: "TIMESTAMP_GRANULARITY_UNSPECIFIED",
	1: "MILLIS",
}
var Table_TimestampGranularity_value = map[string]int32{
	"TIMESTAMP_GRANULARITY_UNSPECIFIED": 0,
	"MILLIS": 1,
}

func (x Table_TimestampGranularity) String() string {
	return proto.EnumName(Table_TimestampGranularity_name, int32(x))
}
func (Table_TimestampGranularity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0, 0}
}

// Defines a view over a table's fields.
type Table_View int32

const (
	// Uses the default view for each method as documented in its request.
	Table_VIEW_UNSPECIFIED Table_View = 0
	// Only populates `name`.
	Table_NAME_ONLY Table_View = 1
	// Only populates `name` and fields related to the table's schema.
	Table_SCHEMA_VIEW Table_View = 2
	// Populates all fields.
	Table_FULL Table_View = 4
)

var Table_View_name = map[int32]string{
	0: "VIEW_UNSPECIFIED",
	1: "NAME_ONLY",
	2: "SCHEMA_VIEW",
	4: "FULL",
}
var Table_View_value = map[string]int32{
	"VIEW_UNSPECIFIED": 0,
	"NAME_ONLY":        1,
	"SCHEMA_VIEW":      2,
	"FULL":             4,
}

func (x Table_View) String() string {
	return proto.EnumName(Table_View_name, int32(x))
}
func (Table_View) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 1} }

// A collection of user data indexed by row, column, and timestamp.
// Each table is served using the resources of its parent cluster.
type Table struct {
	// (`OutputOnly`)
	// The unique name of the table. Values are of the form
	// `projects/<project>/instances/<instance>/tables/[_a-zA-Z0-9][-_.a-zA-Z0-9]*`.
	// Views: `NAME_ONLY`, `SCHEMA_VIEW`, `FULL`
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (`CreationOnly`)
	// The column families configured for this table, mapped by column family ID.
	// Views: `SCHEMA_VIEW`, `FULL`
	ColumnFamilies map[string]*ColumnFamily `protobuf:"bytes,3,rep,name=column_families,json=columnFamilies" json:"column_families,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// (`CreationOnly`)
	// The granularity (e.g. `MILLIS`, `MICROS`) at which timestamps are stored in
	// this table. Timestamps not matching the granularity will be rejected.
	// If unspecified at creation time, the value will be set to `MILLIS`.
	// Views: `SCHEMA_VIEW`, `FULL`
	Granularity Table_TimestampGranularity `protobuf:"varint,4,opt,name=granularity,enum=google.bigtable.admin.v2.Table_TimestampGranularity" json:"granularity,omitempty"`
}

func (m *Table) Reset()                    { *m = Table{} }
func (m *Table) String() string            { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()               {}
func (*Table) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Table) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Table) GetColumnFamilies() map[string]*ColumnFamily {
	if m != nil {
		return m.ColumnFamilies
	}
	return nil
}

func (m *Table) GetGranularity() Table_TimestampGranularity {
	if m != nil {
		return m.Granularity
	}
	return Table_TIMESTAMP_GRANULARITY_UNSPECIFIED
}

// A set of columns within a table which share a common configuration.
type ColumnFamily struct {
	// Garbage collection rule specified as a protobuf.
	// Must serialize to at most 500 bytes.
	//
	// NOTE: Garbage collection executes opportunistically in the background, and
	// so it's possible for reads to return a cell even if it matches the active
	// GC expression for its family.
	GcRule *GcRule `protobuf:"bytes,1,opt,name=gc_rule,json=gcRule" json:"gc_rule,omitempty"`
}

func (m *ColumnFamily) Reset()                    { *m = ColumnFamily{} }
func (m *ColumnFamily) String() string            { return proto.CompactTextString(m) }
func (*ColumnFamily) ProtoMessage()               {}
func (*ColumnFamily) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *ColumnFamily) GetGcRule() *GcRule {
	if m != nil {
		return m.GcRule
	}
	return nil
}

// Rule for determining which cells to delete during garbage collection.
type GcRule struct {
	// Garbage collection rules.
	//
	// Types that are valid to be assigned to Rule:
	//	*GcRule_MaxNumVersions
	//	*GcRule_MaxAge
	//	*GcRule_Intersection_
	//	*GcRule_Union_
	Rule isGcRule_Rule `protobuf_oneof:"rule"`
}

func (m *GcRule) Reset()                    { *m = GcRule{} }
func (m *GcRule) String() string            { return proto.CompactTextString(m) }
func (*GcRule) ProtoMessage()               {}
func (*GcRule) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type isGcRule_Rule interface {
	isGcRule_Rule()
}

type GcRule_MaxNumVersions struct {
	MaxNumVersions int32 `protobuf:"varint,1,opt,name=max_num_versions,json=maxNumVersions,oneof"`
}
type GcRule_MaxAge struct {
	MaxAge *google_protobuf4.Duration `protobuf:"bytes,2,opt,name=max_age,json=maxAge,oneof"`
}
type GcRule_Intersection_ struct {
	Intersection *GcRule_Intersection `protobuf:"bytes,3,opt,name=intersection,oneof"`
}
type GcRule_Union_ struct {
	Union *GcRule_Union `protobuf:"bytes,4,opt,name=union,oneof"`
}

func (*GcRule_MaxNumVersions) isGcRule_Rule() {}
func (*GcRule_MaxAge) isGcRule_Rule()         {}
func (*GcRule_Intersection_) isGcRule_Rule()  {}
func (*GcRule_Union_) isGcRule_Rule()         {}

func (m *GcRule) GetRule() isGcRule_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *GcRule) GetMaxNumVersions() int32 {
	if x, ok := m.GetRule().(*GcRule_MaxNumVersions); ok {
		return x.MaxNumVersions
	}
	return 0
}

func (m *GcRule) GetMaxAge() *google_protobuf4.Duration {
	if x, ok := m.GetRule().(*GcRule_MaxAge); ok {
		return x.MaxAge
	}
	return nil
}

func (m *GcRule) GetIntersection() *GcRule_Intersection {
	if x, ok := m.GetRule().(*GcRule_Intersection_); ok {
		return x.Intersection
	}
	return nil
}

func (m *GcRule) GetUnion() *GcRule_Union {
	if x, ok := m.GetRule().(*GcRule_Union_); ok {
		return x.Union
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GcRule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GcRule_OneofMarshaler, _GcRule_OneofUnmarshaler, _GcRule_OneofSizer, []interface{}{
		(*GcRule_MaxNumVersions)(nil),
		(*GcRule_MaxAge)(nil),
		(*GcRule_Intersection_)(nil),
		(*GcRule_Union_)(nil),
	}
}

func _GcRule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GcRule)
	// rule
	switch x := m.Rule.(type) {
	case *GcRule_MaxNumVersions:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.MaxNumVersions))
	case *GcRule_MaxAge:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MaxAge); err != nil {
			return err
		}
	case *GcRule_Intersection_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Intersection); err != nil {
			return err
		}
	case *GcRule_Union_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Union); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GcRule.Rule has unexpected type %T", x)
	}
	return nil
}

func _GcRule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GcRule)
	switch tag {
	case 1: // rule.max_num_versions
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Rule = &GcRule_MaxNumVersions{int32(x)}
		return true, err
	case 2: // rule.max_age
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf4.Duration)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_MaxAge{msg}
		return true, err
	case 3: // rule.intersection
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GcRule_Intersection)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_Intersection_{msg}
		return true, err
	case 4: // rule.union
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GcRule_Union)
		err := b.DecodeMessage(msg)
		m.Rule = &GcRule_Union_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GcRule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GcRule)
	// rule
	switch x := m.Rule.(type) {
	case *GcRule_MaxNumVersions:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.MaxNumVersions))
	case *GcRule_MaxAge:
		s := proto.Size(x.MaxAge)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GcRule_Intersection_:
		s := proto.Size(x.Intersection)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GcRule_Union_:
		s := proto.Size(x.Union)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// A GcRule which deletes cells matching all of the given rules.
type GcRule_Intersection struct {
	// Only delete cells which would be deleted by every element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Intersection) Reset()                    { *m = GcRule_Intersection{} }
func (m *GcRule_Intersection) String() string            { return proto.CompactTextString(m) }
func (*GcRule_Intersection) ProtoMessage()               {}
func (*GcRule_Intersection) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2, 0} }

func (m *GcRule_Intersection) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// A GcRule which deletes cells matching any of the given rules.
type GcRule_Union struct {
	// Delete cells which would be deleted by any element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *GcRule_Union) Reset()                    { *m = GcRule_Union{} }
func (m *GcRule_Union) String() string            { return proto.CompactTextString(m) }
func (*GcRule_Union) ProtoMessage()               {}
func (*GcRule_Union) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2, 1} }

func (m *GcRule_Union) GetRules() []*GcRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterType((*Table)(nil), "google.bigtable.admin.v2.Table")
	proto.RegisterType((*ColumnFamily)(nil), "google.bigtable.admin.v2.ColumnFamily")
	proto.RegisterType((*GcRule)(nil), "google.bigtable.admin.v2.GcRule")
	proto.RegisterType((*GcRule_Intersection)(nil), "google.bigtable.admin.v2.GcRule.Intersection")
	proto.RegisterType((*GcRule_Union)(nil), "google.bigtable.admin.v2.GcRule.Union")
	proto.RegisterEnum("google.bigtable.admin.v2.Table_TimestampGranularity", Table_TimestampGranularity_name, Table_TimestampGranularity_value)
	proto.RegisterEnum("google.bigtable.admin.v2.Table_View", Table_View_name, Table_View_value)
}

func init() { proto.RegisterFile("google/bigtable/admin/v2/table.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x7d, 0x6b, 0xda, 0x5e,
	0x14, 0x36, 0x4d, 0xb4, 0xbf, 0x1e, 0xfb, 0x6b, 0xc3, 0x5d, 0xff, 0x70, 0x52, 0x36, 0x27, 0xdb,
	0x90, 0xc1, 0x12, 0xb0, 0x65, 0xec, 0x7d, 0xd8, 0x36, 0xd6, 0x80, 0x3a, 0x89, 0x2f, 0xa3, 0x63,
	0x10, 0xae, 0xe9, 0xed, 0xe5, 0xd2, 0xdc, 0x1b, 0xc9, 0x8b, 0xab, 0xdf, 0x62, 0xdf, 0x6c, 0x5f,
	0x69, 0xe4, 0x26, 0x32, 0xdb, 0x55, 0x1c, 0xfb, 0xcb, 0x73, 0xcf, 0x79, 0x9e, 0xe7, 0xbc, 0x1a,
	0x78, 0x4a, 0x83, 0x80, 0xfa, 0xc4, 0x9c, 0x32, 0x1a, 0xe3, 0xa9, 0x4f, 0x4c, 0x7c, 0xc9, 0x99,
	0x30, 0xe7, 0x4d, 0x53, 0x3e, 0x8d, 0x59, 0x18, 0xc4, 0x01, 0xaa, 0x64, 0x28, 0x63, 0x89, 0x32,
	0x24, 0xca, 0x98, 0x37, 0xab, 0x87, 0x39, 0x1f, 0xcf, 0x98, 0x89, 0x85, 0x08, 0x62, 0x1c, 0xb3,
	0x40, 0x44, 0x19, 0xaf, 0xfa, 0x28, 0x8f, 0xca, 0xd7, 0x34, 0xb9, 0x32, 0x2f, 0x93, 0x50, 0x02,
	0xf2, 0xf8, 0xe3, 0xbb, 0xf1, 0x98, 0x71, 0x12, 0xc5, 0x98, 0xcf, 0x32, 0x40, 0xfd, 0xa7, 0x0a,
	0xc5, 0x51, 0x9a, 0x11, 0x21, 0xd0, 0x04, 0xe6, 0xa4, 0xa2, 0xd4, 0x94, 0xc6, 0x8e, 0x23, 0x6d,
	0xf4, 0x0d, 0xf6, 0xbd, 0xc0, 0x4f, 0xb8, 0x70, 0xaf, 0x30, 0x67, 0x3e, 0x23, 0x51, 0x45, 0xad,
	0xa9, 0x8d, 0x72, 0xf3, 0xc8, 0x58, 0x57, 0xb0, 0x21, 0xd5, 0x8c, 0x53, 0x49, 0x6b, 0xe7, 0x2c,
	0x4b, 0xc4, 0xe1, 0xc2, 0xd9, 0xf3, 0x6e, 0x39, 0xd1, 0x04, 0xca, 0x34, 0xc4, 0x22, 0xf1, 0x71,
	0xc8, 0xe2, 0x45, 0x45, 0xab, 0x29, 0x8d, 0xbd, 0xe6, 0xf1, 0x26, 0xe5, 0xd1, 0xb2, 0x83, 0xf3,
	0xdf, 0x5c, 0x67, 0x55, 0xa8, 0xca, 0xe0, 0xc1, 0x3d, 0xe9, 0x91, 0x0e, 0xea, 0x35, 0x59, 0xe4,
	0xfd, 0xa5, 0x26, 0x7a, 0x0f, 0xc5, 0x39, 0xf6, 0x13, 0x52, 0xd9, 0xaa, 0x29, 0x8d, 0x72, 0xf3,
	0xf9, 0xfa, 0xd4, 0x2b, 0x7a, 0x0b, 0x27, 0x23, 0xbd, 0xdd, 0x7a, 0xad, 0xd4, 0x6d, 0x38, 0xb8,
	0xaf, 0x1e, 0xf4, 0x0c, 0x9e, 0x8c, 0xec, 0x9e, 0x35, 0x1c, 0xb5, 0x7a, 0x03, 0xf7, 0xdc, 0x69,
	0xf5, 0xc7, 0xdd, 0x96, 0x63, 0x8f, 0x2e, 0xdc, 0x71, 0x7f, 0x38, 0xb0, 0x4e, 0xed, 0xb6, 0x6d,
	0x9d, 0xe9, 0x05, 0x04, 0x50, 0xea, 0xd9, 0xdd, 0xae, 0x3d, 0xd4, 0x95, 0x7a, 0x1b, 0xb4, 0x09,
	0x23, 0xdf, 0xd1, 0x01, 0xe8, 0x13, 0xdb, 0xfa, 0x72, 0x07, 0xf9, 0x3f, 0xec, 0xf4, 0x5b, 0x3d,
	0xcb, 0xfd, 0xdc, 0xef, 0x5e, 0xe8, 0x0a, 0xda, 0x87, 0xf2, 0xf0, 0xb4, 0x63, 0xf5, 0x5a, 0x6e,
	0x8a, 0xd5, 0xb7, 0xd0, 0x7f, 0xa0, 0xb5, 0xc7, 0xdd, 0xae, 0xae, 0xd5, 0x6d, 0xd8, 0x5d, 0xad,
	0x16, 0xbd, 0x81, 0x6d, 0xea, 0xb9, 0x61, 0xe2, 0x67, 0xab, 0x2d, 0x37, 0x6b, 0xeb, 0xdb, 0x3c,
	0xf7, 0x9c, 0xc4, 0x27, 0x4e, 0x89, 0xca, 0xdf, 0xfa, 0x0f, 0x15, 0x4a, 0x99, 0x0b, 0xbd, 0x00,
	0x9d, 0xe3, 0x1b, 0x57, 0x24, 0xdc, 0x9d, 0x93, 0x30, 0x4a, 0x4f, 0x50, 0xca, 0x15, 0x3b, 0x05,
	0x67, 0x8f, 0xe3, 0x9b, 0x7e, 0xc2, 0x27, 0xb9, 0x1f, 0x1d, 0xc3, 0x76, 0x8a, 0xc5, 0x74, 0x39,
	0xd8, 0x87, 0xcb, 0x8c, 0xcb, 0x33, 0x34, 0xce, 0xf2, 0x33, 0xed, 0x14, 0x9c, 0x12, 0xc7, 0x37,
	0x2d, 0x4a, 0xd0, 0x10, 0x76, 0x99, 0x88, 0x49, 0x18, 0x11, 0x2f, 0x8d, 0x54, 0x54, 0x49, 0x7d,
	0xb9, 0xa9, 0x58, 0xc3, 0x5e, 0x21, 0x75, 0x0a, 0xce, 0x2d, 0x11, 0xf4, 0x11, 0x8a, 0x89, 0x48,
	0xd5, 0xb4, 0x4d, 0x1b, 0xce, 0xd5, 0xc6, 0x22, 0x93, 0xc9, 0x68, 0xd5, 0x36, 0xec, 0xae, 0xea,
	0xa3, 0x57, 0x50, 0x4c, 0x27, 0x99, 0xf6, 0xae, 0xfe, 0xd5, 0x28, 0x33, 0x78, 0xf5, 0x13, 0x14,
	0xa5, 0xf2, 0xbf, 0x0a, 0x9c, 0x94, 0x40, 0x4b, 0x8d, 0x93, 0x6b, 0x38, 0xf4, 0x02, 0xbe, 0x96,
	0x75, 0x02, 0xf2, 0x4f, 0x32, 0x48, 0xe7, 0x3c, 0x50, 0xbe, 0x7e, 0xc8, 0x71, 0x34, 0xf0, 0xb1,
	0xa0, 0x46, 0x10, 0x52, 0x93, 0x12, 0x21, 0xb7, 0x60, 0x66, 0x21, 0x3c, 0x63, 0xd1, 0x9f, 0xdf,
	0xa6, 0x77, 0xd2, 0x98, 0x96, 0x24, 0xf2, 0xe8, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xe3,
	0x1b, 0xd9, 0xc4, 0x04, 0x00, 0x00,
}
