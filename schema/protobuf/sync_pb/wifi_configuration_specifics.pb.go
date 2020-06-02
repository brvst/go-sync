// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wifi_configuration_specifics.proto

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

type WifiConfigurationSpecifics_SecurityType int32

const (
	WifiConfigurationSpecifics_SECURITY_TYPE_UNSPECIFIED WifiConfigurationSpecifics_SecurityType = 0
	WifiConfigurationSpecifics_SECURITY_TYPE_NONE        WifiConfigurationSpecifics_SecurityType = 1
	WifiConfigurationSpecifics_SECURITY_TYPE_WEP         WifiConfigurationSpecifics_SecurityType = 2
	WifiConfigurationSpecifics_SECURITY_TYPE_PSK         WifiConfigurationSpecifics_SecurityType = 3
)

var WifiConfigurationSpecifics_SecurityType_name = map[int32]string{
	0: "SECURITY_TYPE_UNSPECIFIED",
	1: "SECURITY_TYPE_NONE",
	2: "SECURITY_TYPE_WEP",
	3: "SECURITY_TYPE_PSK",
}

var WifiConfigurationSpecifics_SecurityType_value = map[string]int32{
	"SECURITY_TYPE_UNSPECIFIED": 0,
	"SECURITY_TYPE_NONE":        1,
	"SECURITY_TYPE_WEP":         2,
	"SECURITY_TYPE_PSK":         3,
}

func (x WifiConfigurationSpecifics_SecurityType) Enum() *WifiConfigurationSpecifics_SecurityType {
	p := new(WifiConfigurationSpecifics_SecurityType)
	*p = x
	return p
}

func (x WifiConfigurationSpecifics_SecurityType) String() string {
	return proto.EnumName(WifiConfigurationSpecifics_SecurityType_name, int32(x))
}

func (x *WifiConfigurationSpecifics_SecurityType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecifics_SecurityType_value, data, "WifiConfigurationSpecifics_SecurityType")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecifics_SecurityType(value)
	return nil
}

func (WifiConfigurationSpecifics_SecurityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 0}
}

type WifiConfigurationSpecifics_AutomaticallyConnectOption int32

const (
	WifiConfigurationSpecifics_AUTOMATICALLY_CONNECT_UNSPECIFIED WifiConfigurationSpecifics_AutomaticallyConnectOption = 0
	WifiConfigurationSpecifics_AUTOMATICALLY_CONNECT_DISABLED    WifiConfigurationSpecifics_AutomaticallyConnectOption = 1
	WifiConfigurationSpecifics_AUTOMATICALLY_CONNECT_ENABLED     WifiConfigurationSpecifics_AutomaticallyConnectOption = 2
)

var WifiConfigurationSpecifics_AutomaticallyConnectOption_name = map[int32]string{
	0: "AUTOMATICALLY_CONNECT_UNSPECIFIED",
	1: "AUTOMATICALLY_CONNECT_DISABLED",
	2: "AUTOMATICALLY_CONNECT_ENABLED",
}

var WifiConfigurationSpecifics_AutomaticallyConnectOption_value = map[string]int32{
	"AUTOMATICALLY_CONNECT_UNSPECIFIED": 0,
	"AUTOMATICALLY_CONNECT_DISABLED":    1,
	"AUTOMATICALLY_CONNECT_ENABLED":     2,
}

func (x WifiConfigurationSpecifics_AutomaticallyConnectOption) Enum() *WifiConfigurationSpecifics_AutomaticallyConnectOption {
	p := new(WifiConfigurationSpecifics_AutomaticallyConnectOption)
	*p = x
	return p
}

func (x WifiConfigurationSpecifics_AutomaticallyConnectOption) String() string {
	return proto.EnumName(WifiConfigurationSpecifics_AutomaticallyConnectOption_name, int32(x))
}

func (x *WifiConfigurationSpecifics_AutomaticallyConnectOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecifics_AutomaticallyConnectOption_value, data, "WifiConfigurationSpecifics_AutomaticallyConnectOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecifics_AutomaticallyConnectOption(value)
	return nil
}

func (WifiConfigurationSpecifics_AutomaticallyConnectOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 1}
}

type WifiConfigurationSpecifics_IsPreferredOption int32

const (
	WifiConfigurationSpecifics_IS_PREFERRED_UNSPECIFIED WifiConfigurationSpecifics_IsPreferredOption = 0
	WifiConfigurationSpecifics_IS_PREFERRED_DISABLED    WifiConfigurationSpecifics_IsPreferredOption = 1
	WifiConfigurationSpecifics_IS_PREFERRED_ENABLED     WifiConfigurationSpecifics_IsPreferredOption = 2
)

var WifiConfigurationSpecifics_IsPreferredOption_name = map[int32]string{
	0: "IS_PREFERRED_UNSPECIFIED",
	1: "IS_PREFERRED_DISABLED",
	2: "IS_PREFERRED_ENABLED",
}

var WifiConfigurationSpecifics_IsPreferredOption_value = map[string]int32{
	"IS_PREFERRED_UNSPECIFIED": 0,
	"IS_PREFERRED_DISABLED":    1,
	"IS_PREFERRED_ENABLED":     2,
}

func (x WifiConfigurationSpecifics_IsPreferredOption) Enum() *WifiConfigurationSpecifics_IsPreferredOption {
	p := new(WifiConfigurationSpecifics_IsPreferredOption)
	*p = x
	return p
}

func (x WifiConfigurationSpecifics_IsPreferredOption) String() string {
	return proto.EnumName(WifiConfigurationSpecifics_IsPreferredOption_name, int32(x))
}

func (x *WifiConfigurationSpecifics_IsPreferredOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecifics_IsPreferredOption_value, data, "WifiConfigurationSpecifics_IsPreferredOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecifics_IsPreferredOption(value)
	return nil
}

func (WifiConfigurationSpecifics_IsPreferredOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 2}
}

type WifiConfigurationSpecifics_MeteredOption int32

const (
	WifiConfigurationSpecifics_METERED_OPTION_UNSPECIFIED WifiConfigurationSpecifics_MeteredOption = 0
	WifiConfigurationSpecifics_METERED_OPTION_NO          WifiConfigurationSpecifics_MeteredOption = 1
	WifiConfigurationSpecifics_METERED_OPTION_YES         WifiConfigurationSpecifics_MeteredOption = 2
	// Allows the device to use heuristics to determine if network is metered.
	WifiConfigurationSpecifics_METERED_OPTION_AUTO WifiConfigurationSpecifics_MeteredOption = 3
)

var WifiConfigurationSpecifics_MeteredOption_name = map[int32]string{
	0: "METERED_OPTION_UNSPECIFIED",
	1: "METERED_OPTION_NO",
	2: "METERED_OPTION_YES",
	3: "METERED_OPTION_AUTO",
}

var WifiConfigurationSpecifics_MeteredOption_value = map[string]int32{
	"METERED_OPTION_UNSPECIFIED": 0,
	"METERED_OPTION_NO":          1,
	"METERED_OPTION_YES":         2,
	"METERED_OPTION_AUTO":        3,
}

func (x WifiConfigurationSpecifics_MeteredOption) Enum() *WifiConfigurationSpecifics_MeteredOption {
	p := new(WifiConfigurationSpecifics_MeteredOption)
	*p = x
	return p
}

func (x WifiConfigurationSpecifics_MeteredOption) String() string {
	return proto.EnumName(WifiConfigurationSpecifics_MeteredOption_name, int32(x))
}

func (x *WifiConfigurationSpecifics_MeteredOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecifics_MeteredOption_value, data, "WifiConfigurationSpecifics_MeteredOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecifics_MeteredOption(value)
	return nil
}

func (WifiConfigurationSpecifics_MeteredOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 3}
}

type WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption int32

const (
	WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_UNSPECIFIED WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption = 0
	WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_DISABLED    WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption = 1
	// Use a Proxy Auto-config(PAC) Url, set in proxy_url
	WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_AUTOMATIC WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption = 2
	// Uses Web Proxy Auto-Discovery Protocol (WPAD) to discover the proxy
	// settings using DHCP/DNS.
	WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_AUTODISCOVERY WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption = 3
	// User sets proxy_url, proxy_port, and whitelisted_domains manually.
	WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_MANUAL WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption = 4
)

var WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_name = map[int32]string{
	0: "PROXY_OPTION_UNSPECIFIED",
	1: "PROXY_OPTION_DISABLED",
	2: "PROXY_OPTION_AUTOMATIC",
	3: "PROXY_OPTION_AUTODISCOVERY",
	4: "PROXY_OPTION_MANUAL",
}

var WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_value = map[string]int32{
	"PROXY_OPTION_UNSPECIFIED":   0,
	"PROXY_OPTION_DISABLED":      1,
	"PROXY_OPTION_AUTOMATIC":     2,
	"PROXY_OPTION_AUTODISCOVERY": 3,
	"PROXY_OPTION_MANUAL":        4,
}

func (x WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption) Enum() *WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption {
	p := new(WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption)
	*p = x
	return p
}

func (x WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption) String() string {
	return proto.EnumName(WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_name, int32(x))
}

func (x *WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_value, data, "WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption(value)
	return nil
}

func (WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 0, 0}
}

type WifiConfigurationSpecifics struct {
	// SSID encoded to hex, letters should be upper case and 0x prefix should be
	// omitted. For example, ssid "network" would be provided as "6E6574776F726B".
	HexSsid      []byte                                   `protobuf:"bytes,1,opt,name=hex_ssid,json=hexSsid" json:"hex_ssid,omitempty"`
	SecurityType *WifiConfigurationSpecifics_SecurityType `protobuf:"varint,2,opt,name=security_type,json=securityType,enum=sync_pb.WifiConfigurationSpecifics_SecurityType" json:"security_type,omitempty"`
	// The passphrase can be ASCII, UTF-8, or a string of hex digits.
	Passphrase           []byte                                                 `protobuf:"bytes,3,opt,name=passphrase" json:"passphrase,omitempty"`
	AutomaticallyConnect *WifiConfigurationSpecifics_AutomaticallyConnectOption `protobuf:"varint,4,opt,name=automatically_connect,json=automaticallyConnect,enum=sync_pb.WifiConfigurationSpecifics_AutomaticallyConnectOption" json:"automatically_connect,omitempty"`
	IsPreferred          *WifiConfigurationSpecifics_IsPreferredOption          `protobuf:"varint,5,opt,name=is_preferred,json=isPreferred,enum=sync_pb.WifiConfigurationSpecifics_IsPreferredOption" json:"is_preferred,omitempty"`
	Metered              *WifiConfigurationSpecifics_MeteredOption              `protobuf:"varint,6,opt,name=metered,enum=sync_pb.WifiConfigurationSpecifics_MeteredOption" json:"metered,omitempty"`
	ProxyConfiguration   *WifiConfigurationSpecifics_ProxyConfiguration         `protobuf:"bytes,7,opt,name=proxy_configuration,json=proxyConfiguration" json:"proxy_configuration,omitempty"`
	// List of DNS servers to be used.  Up to 4.
	CustomDns []string `protobuf:"bytes,8,rep,name=custom_dns,json=customDns" json:"custom_dns,omitempty"`
	// This is represented by the UNIX timestamp, ms since epoch.
	LastUpdateTimestamp  *int64   `protobuf:"varint,9,opt,name=last_update_timestamp,json=lastUpdateTimestamp" json:"last_update_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WifiConfigurationSpecifics) Reset()         { *m = WifiConfigurationSpecifics{} }
func (m *WifiConfigurationSpecifics) String() string { return proto.CompactTextString(m) }
func (*WifiConfigurationSpecifics) ProtoMessage()    {}
func (*WifiConfigurationSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0}
}

func (m *WifiConfigurationSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiConfigurationSpecifics.Unmarshal(m, b)
}
func (m *WifiConfigurationSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiConfigurationSpecifics.Marshal(b, m, deterministic)
}
func (m *WifiConfigurationSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiConfigurationSpecifics.Merge(m, src)
}
func (m *WifiConfigurationSpecifics) XXX_Size() int {
	return xxx_messageInfo_WifiConfigurationSpecifics.Size(m)
}
func (m *WifiConfigurationSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiConfigurationSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_WifiConfigurationSpecifics proto.InternalMessageInfo

func (m *WifiConfigurationSpecifics) GetHexSsid() []byte {
	if m != nil {
		return m.HexSsid
	}
	return nil
}

func (m *WifiConfigurationSpecifics) GetSecurityType() WifiConfigurationSpecifics_SecurityType {
	if m != nil && m.SecurityType != nil {
		return *m.SecurityType
	}
	return WifiConfigurationSpecifics_SECURITY_TYPE_UNSPECIFIED
}

func (m *WifiConfigurationSpecifics) GetPassphrase() []byte {
	if m != nil {
		return m.Passphrase
	}
	return nil
}

func (m *WifiConfigurationSpecifics) GetAutomaticallyConnect() WifiConfigurationSpecifics_AutomaticallyConnectOption {
	if m != nil && m.AutomaticallyConnect != nil {
		return *m.AutomaticallyConnect
	}
	return WifiConfigurationSpecifics_AUTOMATICALLY_CONNECT_UNSPECIFIED
}

func (m *WifiConfigurationSpecifics) GetIsPreferred() WifiConfigurationSpecifics_IsPreferredOption {
	if m != nil && m.IsPreferred != nil {
		return *m.IsPreferred
	}
	return WifiConfigurationSpecifics_IS_PREFERRED_UNSPECIFIED
}

func (m *WifiConfigurationSpecifics) GetMetered() WifiConfigurationSpecifics_MeteredOption {
	if m != nil && m.Metered != nil {
		return *m.Metered
	}
	return WifiConfigurationSpecifics_METERED_OPTION_UNSPECIFIED
}

func (m *WifiConfigurationSpecifics) GetProxyConfiguration() *WifiConfigurationSpecifics_ProxyConfiguration {
	if m != nil {
		return m.ProxyConfiguration
	}
	return nil
}

func (m *WifiConfigurationSpecifics) GetCustomDns() []string {
	if m != nil {
		return m.CustomDns
	}
	return nil
}

func (m *WifiConfigurationSpecifics) GetLastUpdateTimestamp() int64 {
	if m != nil && m.LastUpdateTimestamp != nil {
		return *m.LastUpdateTimestamp
	}
	return 0
}

type WifiConfigurationSpecifics_ProxyConfiguration struct {
	ProxyOption *WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption `protobuf:"varint,1,opt,name=proxy_option,json=proxyOption,enum=sync_pb.WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption" json:"proxy_option,omitempty"`
	// Only set if PROXY_OPTION_AUTOMATIC or PROXY_OPTION_MANUAL.
	ProxyUrl *string `protobuf:"bytes,2,opt,name=proxy_url,json=proxyUrl" json:"proxy_url,omitempty"`
	// Only set if PROXY_OPTION_MANUAL.
	ProxyPort *int32 `protobuf:"varint,3,opt,name=proxy_port,json=proxyPort" json:"proxy_port,omitempty"`
	// Only set if PROXY_OPTION_MANUAL.
	WhitelistedDomains   []string `protobuf:"bytes,4,rep,name=whitelisted_domains,json=whitelistedDomains" json:"whitelisted_domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration) Reset() {
	*m = WifiConfigurationSpecifics_ProxyConfiguration{}
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration) String() string {
	return proto.CompactTextString(m)
}
func (*WifiConfigurationSpecifics_ProxyConfiguration) ProtoMessage() {}
func (*WifiConfigurationSpecifics_ProxyConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 0}
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration.Unmarshal(m, b)
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration.Marshal(b, m, deterministic)
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration.Merge(m, src)
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration) XXX_Size() int {
	return xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration.Size(m)
}
func (m *WifiConfigurationSpecifics_ProxyConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_WifiConfigurationSpecifics_ProxyConfiguration proto.InternalMessageInfo

func (m *WifiConfigurationSpecifics_ProxyConfiguration) GetProxyOption() WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption {
	if m != nil && m.ProxyOption != nil {
		return *m.ProxyOption
	}
	return WifiConfigurationSpecifics_ProxyConfiguration_PROXY_OPTION_UNSPECIFIED
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration) GetProxyUrl() string {
	if m != nil && m.ProxyUrl != nil {
		return *m.ProxyUrl
	}
	return ""
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration) GetProxyPort() int32 {
	if m != nil && m.ProxyPort != nil {
		return *m.ProxyPort
	}
	return 0
}

func (m *WifiConfigurationSpecifics_ProxyConfiguration) GetWhitelistedDomains() []string {
	if m != nil {
		return m.WhitelistedDomains
	}
	return nil
}

func init() {
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecifics_SecurityType", WifiConfigurationSpecifics_SecurityType_name, WifiConfigurationSpecifics_SecurityType_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecifics_AutomaticallyConnectOption", WifiConfigurationSpecifics_AutomaticallyConnectOption_name, WifiConfigurationSpecifics_AutomaticallyConnectOption_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecifics_IsPreferredOption", WifiConfigurationSpecifics_IsPreferredOption_name, WifiConfigurationSpecifics_IsPreferredOption_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecifics_MeteredOption", WifiConfigurationSpecifics_MeteredOption_name, WifiConfigurationSpecifics_MeteredOption_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption", WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_name, WifiConfigurationSpecifics_ProxyConfiguration_ProxyOption_value)
	proto.RegisterType((*WifiConfigurationSpecifics)(nil), "sync_pb.WifiConfigurationSpecifics")
	proto.RegisterType((*WifiConfigurationSpecifics_ProxyConfiguration)(nil), "sync_pb.WifiConfigurationSpecifics.ProxyConfiguration")
}

func init() {
	proto.RegisterFile("wifi_configuration_specifics.proto", fileDescriptor_532bc6a80c7070ba)
}

var fileDescriptor_532bc6a80c7070ba = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xe1, 0x6e, 0xe2, 0x46,
	0x10, 0xc7, 0x6b, 0x9c, 0x6b, 0xc2, 0xc0, 0x9d, 0x9c, 0xe5, 0xb8, 0x3a, 0xb4, 0x89, 0x38, 0xa4,
	0x93, 0x90, 0x2a, 0xd1, 0x16, 0xa9, 0xfd, 0x58, 0x89, 0x98, 0x3d, 0xd5, 0x3a, 0xb0, 0x2d, 0xdb,
	0xf4, 0x8e, 0x4f, 0x96, 0x6b, 0x96, 0xb0, 0x12, 0xf6, 0xae, 0xbc, 0x8b, 0x12, 0x5e, 0xa0, 0x1f,
	0xfa, 0x0a, 0x7d, 0xcb, 0x3e, 0x41, 0x65, 0x1b, 0x1a, 0x1b, 0x88, 0x84, 0xfa, 0xd1, 0xff, 0xdf,
	0xcc, 0xfc, 0x77, 0x76, 0x76, 0x64, 0xe8, 0x3d, 0xd2, 0x25, 0x0d, 0x22, 0x96, 0x2c, 0xe9, 0xc3,
	0x26, 0x0d, 0x25, 0x65, 0x49, 0x20, 0x38, 0x89, 0xe8, 0x92, 0x46, 0x62, 0xc0, 0x53, 0x26, 0x19,
	0xba, 0x14, 0xdb, 0x24, 0x0a, 0xf8, 0x1f, 0x1d, 0x8d, 0x24, 0x51, 0xba, 0xe5, 0x59, 0x50, 0x81,
	0x7a, 0xff, 0x34, 0xa0, 0xf3, 0x99, 0x2e, 0xa9, 0x51, 0x2e, 0xe0, 0xed, 0xf3, 0xd1, 0x0d, 0x5c,
	0xad, 0xc8, 0x53, 0x20, 0x04, 0x5d, 0xe8, 0x4a, 0x57, 0xe9, 0x37, 0xdd, 0xcb, 0x15, 0x79, 0xf2,
	0x04, 0x5d, 0xa0, 0x19, 0xbc, 0x16, 0x24, 0xda, 0xa4, 0x54, 0x6e, 0x03, 0xb9, 0xe5, 0x44, 0xaf,
	0x75, 0x95, 0xfe, 0x9b, 0xe1, 0x8f, 0x83, 0x9d, 0xd9, 0xe0, 0xe5, 0xb2, 0x03, 0x6f, 0x97, 0xe8,
	0x6f, 0x39, 0x71, 0x9b, 0xa2, 0xf4, 0x85, 0xee, 0x00, 0x78, 0x28, 0x04, 0x5f, 0xa5, 0xa1, 0x20,
	0xba, 0x9a, 0x7b, 0x96, 0x14, 0x24, 0xa0, 0x1d, 0x6e, 0x24, 0x8b, 0x43, 0x49, 0xa3, 0x70, 0xbd,
	0xde, 0x66, 0xad, 0x27, 0x24, 0x92, 0xfa, 0x45, 0x6e, 0xff, 0xeb, 0x39, 0xf6, 0xa3, 0x72, 0x01,
	0xa3, 0xc8, 0xb7, 0xf3, 0x5b, 0x71, 0xdf, 0x86, 0x27, 0x18, 0xfa, 0x02, 0x4d, 0x2a, 0x02, 0x9e,
	0x92, 0x25, 0x49, 0x53, 0xb2, 0xd0, 0x5f, 0xe5, 0x5e, 0x3f, 0x9f, 0xe3, 0x65, 0x0a, 0x67, 0x9f,
	0xb6, 0xb3, 0x68, 0xd0, 0x67, 0x09, 0x7d, 0x82, 0xcb, 0x98, 0x48, 0x92, 0x15, 0xfd, 0x3a, 0x2f,
	0xfa, 0xd3, 0x39, 0x45, 0xa7, 0x45, 0xca, 0xae, 0xe0, 0xbe, 0x02, 0x7a, 0x80, 0x16, 0x4f, 0xd9,
	0xd3, 0xb6, 0xfa, 0x1c, 0xf4, 0xcb, 0xae, 0xd2, 0x6f, 0x0c, 0x7f, 0x39, 0xa7, 0xb0, 0x93, 0xa5,
	0x57, 0x98, 0x8b, 0xf8, 0x91, 0x86, 0x6e, 0x01, 0xa2, 0x8d, 0x90, 0x2c, 0x0e, 0x16, 0x89, 0xd0,
	0xaf, 0xba, 0x6a, 0xbf, 0xee, 0xd6, 0x0b, 0x65, 0x9c, 0x08, 0x34, 0x84, 0xf6, 0x3a, 0x14, 0x32,
	0xd8, 0xf0, 0x45, 0x28, 0x49, 0x20, 0x69, 0x4c, 0x84, 0x0c, 0x63, 0xae, 0xd7, 0xbb, 0x4a, 0x5f,
	0x75, 0x5b, 0x19, 0x9c, 0xe5, 0xcc, 0xdf, 0xa3, 0xce, 0x9f, 0x2a, 0xa0, 0x63, 0x77, 0x44, 0xa0,
	0x59, 0xb4, 0xc4, 0xf2, 0x5e, 0xf3, 0x47, 0xf8, 0x66, 0x78, 0xff, 0xff, 0x7a, 0x29, 0xa4, 0xfd,
	0x18, 0xf8, 0xf3, 0x07, 0xfa, 0x16, 0xea, 0x85, 0xcd, 0x26, 0x5d, 0xe7, 0x0f, 0xb9, 0xee, 0x5e,
	0xe5, 0xc2, 0x2c, 0x5d, 0x67, 0xdd, 0x16, 0x90, 0xb3, 0x54, 0xe6, 0x4f, 0xf2, 0x95, 0x5b, 0x84,
	0x3b, 0x2c, 0x95, 0xe8, 0x07, 0x68, 0x3d, 0xae, 0xa8, 0x24, 0x6b, 0x2a, 0x24, 0x59, 0x04, 0x0b,
	0x16, 0x87, 0x34, 0x11, 0xfa, 0x45, 0x7e, 0x2b, 0xa8, 0x84, 0xc6, 0x05, 0xe9, 0xfd, 0xad, 0x40,
	0xa3, 0x74, 0x12, 0xf4, 0x1d, 0xe8, 0x8e, 0x6b, 0x7f, 0x99, 0x07, 0xb6, 0xe3, 0x9b, 0xb6, 0x15,
	0xcc, 0x2c, 0xcf, 0xc1, 0x86, 0xf9, 0xd1, 0xc4, 0x63, 0xed, 0x2b, 0x74, 0x03, 0xed, 0x0a, 0x1d,
	0x9b, 0xde, 0xe8, 0x7e, 0x82, 0xc7, 0x9a, 0x82, 0x3a, 0xf0, 0xae, 0x82, 0x46, 0x33, 0xdf, 0x9e,
	0x8e, 0x7c, 0xd3, 0xd0, 0x6a, 0xe8, 0x0e, 0x3a, 0x47, 0x6c, 0x6c, 0x7a, 0x86, 0xfd, 0x3b, 0x76,
	0xe7, 0x9a, 0x8a, 0xbe, 0x81, 0x56, 0x85, 0x4f, 0x47, 0xd6, 0x6c, 0x34, 0xd1, 0x2e, 0x7a, 0x02,
	0x9a, 0xe5, 0xf5, 0x44, 0xb7, 0x70, 0xe3, 0x61, 0x63, 0xe6, 0x9a, 0xfe, 0x3c, 0xf0, 0xe7, 0x0e,
	0x3e, 0x38, 0xde, 0x3b, 0x40, 0x55, 0x6c, 0xd9, 0x16, 0xd6, 0x14, 0xd4, 0x86, 0xeb, 0xaa, 0xfe,
	0x19, 0x3b, 0x5a, 0xed, 0x58, 0x76, 0xbc, 0x4f, 0x9a, 0xda, 0xfb, 0x4b, 0x81, 0xce, 0xcb, 0x5b,
	0x89, 0x3e, 0xc0, 0xfb, 0xff, 0x7a, 0x1b, 0x4d, 0x26, 0xf3, 0xc0, 0xb0, 0x2d, 0x0b, 0x1b, 0xfe,
	0xc1, 0x59, 0x7a, 0x70, 0x77, 0x3a, 0xac, 0x74, 0x67, 0xef, 0xe1, 0xf6, 0x74, 0x0c, 0xb6, 0x8a,
	0x90, 0x5a, 0x6f, 0x09, 0xd7, 0x47, 0x5b, 0x9b, 0x0d, 0xc9, 0xf4, 0x02, 0xc7, 0xc5, 0x1f, 0xb1,
	0xeb, 0xe2, 0xf1, 0xf1, 0x90, 0x2a, 0xb4, 0x64, 0xa8, 0xc3, 0xdb, 0x0a, 0x7a, 0xf6, 0x79, 0x84,
	0xd7, 0x95, 0x45, 0xce, 0x66, 0x36, 0xc5, 0x3e, 0xce, 0xa2, 0x4e, 0x3e, 0x85, 0x36, 0x5c, 0x1f,
	0x70, 0xcb, 0xd6, 0x94, 0x6c, 0x04, 0x07, 0xf2, 0x1c, 0x7b, 0x5a, 0x2d, 0x1b, 0xf1, 0x81, 0x9e,
	0x75, 0xae, 0xa9, 0xf7, 0xdf, 0xc3, 0x07, 0x96, 0x3e, 0x0c, 0xa2, 0x55, 0xca, 0x62, 0xba, 0x89,
	0x07, 0x11, 0x8b, 0x39, 0x4b, 0x48, 0x22, 0x45, 0xbe, 0x57, 0xc5, 0xaf, 0x21, 0x62, 0xeb, 0xdf,
	0x54, 0x47, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x80, 0x15, 0x55, 0x7d, 0x61, 0x06, 0x00, 0x00,
}