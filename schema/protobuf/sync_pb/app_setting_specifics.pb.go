// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app_setting_specifics.proto

package sync_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Properties of app setting sync objects; just an extension setting.
type AppSettingSpecifics struct {
	ExtensionSetting     *ExtensionSettingSpecifics `protobuf:"bytes,1,opt,name=extension_setting,json=extensionSetting" json:"extension_setting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AppSettingSpecifics) Reset()         { *m = AppSettingSpecifics{} }
func (m *AppSettingSpecifics) String() string { return proto.CompactTextString(m) }
func (*AppSettingSpecifics) ProtoMessage()    {}
func (*AppSettingSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fc883edb4e2657f, []int{0}
}

func (m *AppSettingSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppSettingSpecifics.Unmarshal(m, b)
}
func (m *AppSettingSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppSettingSpecifics.Marshal(b, m, deterministic)
}
func (m *AppSettingSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppSettingSpecifics.Merge(m, src)
}
func (m *AppSettingSpecifics) XXX_Size() int {
	return xxx_messageInfo_AppSettingSpecifics.Size(m)
}
func (m *AppSettingSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_AppSettingSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_AppSettingSpecifics proto.InternalMessageInfo

func (m *AppSettingSpecifics) GetExtensionSetting() *ExtensionSettingSpecifics {
	if m != nil {
		return m.ExtensionSetting
	}
	return nil
}

func init() {
	proto.RegisterType((*AppSettingSpecifics)(nil), "sync_pb.AppSettingSpecifics")
}

func init() {
	proto.RegisterFile("app_setting_specifics.proto", fileDescriptor_4fc883edb4e2657f)
}

var fileDescriptor_4fc883edb4e2657f = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x2c, 0x28, 0x88,
	0x2f, 0x4e, 0x2d, 0x29, 0xc9, 0xcc, 0x4b, 0x8f, 0x2f, 0x2e, 0x48, 0x4d, 0xce, 0x4c, 0xcb, 0x4c,
	0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0xae, 0xcc, 0x4b, 0x8e, 0x2f, 0x48,
	0x92, 0x52, 0x4c, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xc3, 0xa5, 0x56, 0x29, 0x8d,
	0x4b, 0xd8, 0xb1, 0xa0, 0x20, 0x18, 0x22, 0x1b, 0x0c, 0x93, 0x14, 0xf2, 0xe7, 0x12, 0xc4, 0xd0,
	0x2b, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa4, 0x07, 0x35, 0x5e, 0xcf, 0x15, 0xa6, 0x02,
	0x5d, 0x7b, 0x90, 0x40, 0x2a, 0x9a, 0x94, 0x93, 0x36, 0x97, 0x6a, 0x7e, 0x51, 0xba, 0x5e, 0x72,
	0x46, 0x51, 0x7e, 0x6e, 0x66, 0x69, 0xae, 0x5e, 0x72, 0x7e, 0x6e, 0x41, 0x7e, 0x5e, 0x6a, 0x5e,
	0x49, 0x31, 0xd8, 0x38, 0x88, 0x6b, 0x92, 0xf3, 0x73, 0x3c, 0x98, 0x03, 0x18, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x4a, 0x58, 0x62, 0x53, 0xde, 0x00, 0x00, 0x00,
}