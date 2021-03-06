// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package tracing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Link_Type int32

const (
	Link_TYPE_UNSPECIFIED   Link_Type = 0
	Link_CHILD_LINKED_SPAN  Link_Type = 1
	Link_PARENT_LINKED_SPAN Link_Type = 2
)

var Link_Type_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "CHILD_LINKED_SPAN",
	2: "PARENT_LINKED_SPAN",
}

var Link_Type_value = map[string]int32{
	"TYPE_UNSPECIFIED":   0,
	"CHILD_LINKED_SPAN":  1,
	"PARENT_LINKED_SPAN": 2,
}

func (x Link_Type) String() string {
	return proto.EnumName(Link_Type_name, int32(x))
}

func (Link_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2, 0}
}

type Span struct {
	Id        uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TraceId   string               `protobuf:"bytes,2,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Operation string               `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	Start     *timestamp.Timestamp `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	End       *timestamp.Timestamp `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
	// Types that are valid to be assigned to Reference:
	//	*Span_FollowsFrom
	//	*Span_ChildOf
	Reference            isSpan_Reference  `protobuf_oneof:"reference"`
	Tags                 map[string]string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Links                *Links            `protobuf:"bytes,9,opt,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Span) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Span) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *Span) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Span) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

type isSpan_Reference interface {
	isSpan_Reference()
}

type Span_FollowsFrom struct {
	FollowsFrom uint64 `protobuf:"varint,6,opt,name=follows_from,json=followsFrom,proto3,oneof"`
}

type Span_ChildOf struct {
	ChildOf uint64 `protobuf:"varint,7,opt,name=child_of,json=childOf,proto3,oneof"`
}

func (*Span_FollowsFrom) isSpan_Reference() {}

func (*Span_ChildOf) isSpan_Reference() {}

func (m *Span) GetReference() isSpan_Reference {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *Span) GetFollowsFrom() uint64 {
	if x, ok := m.GetReference().(*Span_FollowsFrom); ok {
		return x.FollowsFrom
	}
	return 0
}

func (m *Span) GetChildOf() uint64 {
	if x, ok := m.GetReference().(*Span_ChildOf); ok {
		return x.ChildOf
	}
	return 0
}

func (m *Span) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Span) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Span) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Span_FollowsFrom)(nil),
		(*Span_ChildOf)(nil),
	}
}

type Spans struct {
	Spans                []*Span  `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Spans) Reset()         { *m = Spans{} }
func (m *Spans) String() string { return proto.CompactTextString(m) }
func (*Spans) ProtoMessage()    {}
func (*Spans) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

func (m *Spans) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spans.Unmarshal(m, b)
}
func (m *Spans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spans.Marshal(b, m, deterministic)
}
func (m *Spans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spans.Merge(m, src)
}
func (m *Spans) XXX_Size() int {
	return xxx_messageInfo_Spans.Size(m)
}
func (m *Spans) XXX_DiscardUnknown() {
	xxx_messageInfo_Spans.DiscardUnknown(m)
}

var xxx_messageInfo_Spans proto.InternalMessageInfo

func (m *Spans) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type Link struct {
	TraceId              []byte    `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId               []byte    `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	Type                 Link_Type `protobuf:"varint,3,opt,name=type,proto3,enum=tracing.Link_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}

func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *Link) GetSpanId() []byte {
	if m != nil {
		return m.SpanId
	}
	return nil
}

func (m *Link) GetType() Link_Type {
	if m != nil {
		return m.Type
	}
	return Link_TYPE_UNSPECIFIED
}

type Links struct {
	Link                 []*Link  `protobuf:"bytes,1,rep,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Links) Reset()         { *m = Links{} }
func (m *Links) String() string { return proto.CompactTextString(m) }
func (*Links) ProtoMessage()    {}
func (*Links) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}

func (m *Links) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Links.Unmarshal(m, b)
}
func (m *Links) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Links.Marshal(b, m, deterministic)
}
func (m *Links) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Links.Merge(m, src)
}
func (m *Links) XXX_Size() int {
	return xxx_messageInfo_Links.Size(m)
}
func (m *Links) XXX_DiscardUnknown() {
	xxx_messageInfo_Links.DiscardUnknown(m)
}

var xxx_messageInfo_Links proto.InternalMessageInfo

func (m *Links) GetLink() []*Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func init() {
	proto.RegisterEnum("tracing.Link_Type", Link_Type_name, Link_Type_value)
	proto.RegisterType((*Span)(nil), "tracing.Span")
	proto.RegisterMapType((map[string]string)(nil), "tracing.Span.TagsEntry")
	proto.RegisterType((*Spans)(nil), "tracing.Spans")
	proto.RegisterType((*Link)(nil), "tracing.Link")
	proto.RegisterType((*Links)(nil), "tracing.Links")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0xac, 0xbf, 0x9a, 0xf8, 0x39, 0x44, 0xe1, 0xa9, 0xd0, 0x25, 0x20, 0x11, 0x5c, 0x84, 0x22,
	0xa8, 0x5c, 0x14, 0x0e, 0x20, 0x6e, 0xa5, 0x71, 0x55, 0xab, 0x51, 0x88, 0x36, 0xe6, 0xc0, 0xc9,
	0x72, 0xe3, 0xb5, 0xb1, 0xea, 0x78, 0x2d, 0x7b, 0x0b, 0xca, 0x5f, 0xe0, 0xcf, 0xf0, 0x17, 0xd1,
	0xae, 0x93, 0x56, 0x11, 0x07, 0x6e, 0xfb, 0x66, 0xc6, 0xeb, 0x79, 0x33, 0x0b, 0x8e, 0xd8, 0x54,
	0xac, 0xf1, 0xaa, 0x9a, 0x0b, 0x8e, 0x1d, 0x51, 0xc7, 0xab, 0xbc, 0xcc, 0x86, 0x2f, 0x33, 0xce,
	0xb3, 0x82, 0x9d, 0x29, 0xf8, 0xe6, 0x2e, 0x3d, 0x13, 0xf9, 0x9a, 0x35, 0x22, 0x5e, 0x57, 0xad,
	0xd2, 0xfd, 0x6d, 0x80, 0xb9, 0xac, 0xe2, 0x12, 0xfb, 0xa0, 0xe7, 0x09, 0xd1, 0x46, 0xda, 0xd8,
	0xa4, 0x7a, 0x9e, 0xe0, 0x33, 0xe8, 0xca, 0x4b, 0x58, 0x94, 0x27, 0x44, 0x1f, 0x69, 0x63, 0x9b,
	0xaa, 0x4b, 0x59, 0x90, 0xe0, 0x0b, 0xb0, 0x79, 0xc5, 0xea, 0x58, 0xe4, 0xbc, 0x24, 0x86, 0xe2,
	0x1e, 0x00, 0x7c, 0x0f, 0x56, 0x23, 0xe2, 0x5a, 0x10, 0x73, 0xa4, 0x8d, 0x9d, 0xc9, 0xd0, 0x6b,
	0x2d, 0x78, 0x3b, 0x0b, 0x5e, 0xb8, 0xb3, 0x40, 0x5b, 0x21, 0x9e, 0x82, 0xc1, 0xca, 0x84, 0x58,
	0xff, 0xd5, 0x4b, 0x19, 0x9e, 0x40, 0x2f, 0xe5, 0x45, 0xc1, 0x7f, 0x35, 0x51, 0x5a, 0xf3, 0x35,
	0x39, 0x94, 0x96, 0xaf, 0x0e, 0xa8, 0xb3, 0x45, 0x2f, 0x6b, 0xbe, 0xc6, 0xe7, 0xd0, 0x5d, 0xfd,
	0xc8, 0x8b, 0x24, 0xe2, 0x29, 0xe9, 0x6c, 0x05, 0x1d, 0x85, 0x7c, 0x4d, 0xf1, 0x1d, 0x98, 0x22,
	0xce, 0x1a, 0xd2, 0x1d, 0x19, 0x63, 0x67, 0x72, 0xec, 0x6d, 0xc3, 0xf2, 0x64, 0x0e, 0x5e, 0x18,
	0x67, 0x8d, 0x5f, 0x8a, 0x7a, 0x43, 0x95, 0x08, 0x5f, 0x83, 0x55, 0xe4, 0xe5, 0x6d, 0x43, 0x6c,
	0x65, 0xaf, 0x7f, 0xaf, 0x9e, 0x49, 0x94, 0xb6, 0xe4, 0xf0, 0x23, 0xd8, 0xf7, 0x1f, 0xe2, 0x00,
	0x8c, 0x5b, 0xb6, 0x51, 0x59, 0xda, 0x54, 0x1e, 0xf1, 0x08, 0xac, 0x9f, 0x71, 0x71, 0xc7, 0xb6,
	0x49, 0xb6, 0xc3, 0x67, 0xfd, 0x93, 0xf6, 0xc5, 0x01, 0xbb, 0x66, 0x29, 0xab, 0x59, 0xb9, 0x62,
	0xee, 0x29, 0x58, 0xd2, 0x43, 0x83, 0x27, 0x60, 0x35, 0xf2, 0x40, 0x34, 0x65, 0xf1, 0xd1, 0x9e,
	0x45, 0xda, 0x72, 0xee, 0x1f, 0x0d, 0x4c, 0x69, 0x62, 0xaf, 0x2a, 0xf9, 0xd3, 0xde, 0x43, 0x55,
	0xc7, 0xd0, 0x91, 0xe2, 0x5d, 0x89, 0x3d, 0x7a, 0x28, 0xc7, 0x20, 0xc1, 0x37, 0x60, 0xca, 0x07,
	0xa3, 0xea, 0xeb, 0x4f, 0x70, 0x6f, 0x2b, 0x2f, 0xdc, 0x54, 0x8c, 0x2a, 0xde, 0xbd, 0x06, 0x53,
	0x4e, 0x78, 0x04, 0x83, 0xf0, 0xfb, 0xc2, 0x8f, 0xbe, 0xcd, 0x97, 0x0b, 0xff, 0x22, 0xb8, 0x0c,
	0xfc, 0xe9, 0xe0, 0x00, 0x9f, 0xc0, 0xe3, 0x8b, 0xab, 0x60, 0x36, 0x8d, 0x66, 0xc1, 0xfc, 0xda,
	0x9f, 0x46, 0xcb, 0xc5, 0xf9, 0x7c, 0xa0, 0xe1, 0x53, 0xc0, 0xc5, 0x39, 0xf5, 0xe7, 0xe1, 0x1e,
	0xae, 0xbb, 0x6f, 0xc1, 0x52, 0xa9, 0xe1, 0x2b, 0x30, 0x65, 0x6e, 0xff, 0xac, 0x27, 0x59, 0xaa,
	0xa8, 0x9b, 0x43, 0xd5, 0xff, 0x87, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x23, 0xbc, 0xa8,
	0xd8, 0x02, 0x00, 0x00,
}
