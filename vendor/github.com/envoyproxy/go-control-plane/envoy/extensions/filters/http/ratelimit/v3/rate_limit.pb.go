// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/ratelimit/v3/rate_limit.proto

package envoy_extensions_filters_http_ratelimit_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/ratelimit/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type RateLimit struct {
	Domain                         string                     `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Stage                          uint32                     `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	RequestType                    string                     `protobuf:"bytes,3,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	Timeout                        *duration.Duration         `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	FailureModeDeny                bool                       `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	RateLimitedAsResourceExhausted bool                       `protobuf:"varint,6,opt,name=rate_limited_as_resource_exhausted,json=rateLimitedAsResourceExhausted,proto3" json:"rate_limited_as_resource_exhausted,omitempty"`
	RateLimitService               *v3.RateLimitServiceConfig `protobuf:"bytes,7,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}                   `json:"-"`
	XXX_unrecognized               []byte                     `json:"-"`
	XXX_sizecache                  int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf645dbc1fee5d54, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitedAsResourceExhausted() bool {
	if m != nil {
		return m.RateLimitedAsResourceExhausted
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v3.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.extensions.filters.http.ratelimit.v3.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/ratelimit/v3/rate_limit.proto", fileDescriptor_bf645dbc1fee5d54)
}

var fileDescriptor_bf645dbc1fee5d54 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd4, 0x3c,
	0x10, 0x56, 0xb6, 0xdd, 0xdd, 0xd6, 0xfd, 0x7f, 0x28, 0xbe, 0x10, 0x2a, 0xb1, 0xa4, 0xe5, 0xb2,
	0xda, 0x83, 0x0d, 0xbb, 0x07, 0x10, 0x9c, 0x08, 0xe5, 0x82, 0x8a, 0x54, 0x05, 0xee, 0x91, 0xbb,
	0x9e, 0x4d, 0x2d, 0x65, 0xed, 0x60, 0x3b, 0xd1, 0xe6, 0x0d, 0x10, 0x8f, 0xc0, 0xcb, 0x70, 0xe2,
	0xa5, 0x7a, 0x42, 0xb1, 0x9d, 0x8d, 0xda, 0x13, 0x37, 0xcf, 0xcc, 0xf7, 0x7d, 0xe3, 0x99, 0x6f,
	0xd0, 0x7b, 0x90, 0x8d, 0x6a, 0x29, 0xec, 0x2c, 0x48, 0x23, 0x94, 0x34, 0x74, 0x23, 0x4a, 0x0b,
	0xda, 0xd0, 0x5b, 0x6b, 0x2b, 0xaa, 0x99, 0x85, 0x52, 0x6c, 0x85, 0xa5, 0xcd, 0xca, 0x05, 0xb9,
	0x8b, 0x48, 0xa5, 0x95, 0x55, 0x78, 0xe1, 0xc8, 0x64, 0x20, 0x93, 0x40, 0x26, 0x1d, 0x99, 0xec,
	0xc9, 0xa4, 0x59, 0x9d, 0xbd, 0xf4, 0x8d, 0xd6, 0x4a, 0x6e, 0x44, 0xf1, 0x40, 0xb7, 0x34, 0x5e,
	0xf0, 0x6c, 0x56, 0x28, 0x55, 0x94, 0x40, 0x5d, 0x74, 0x53, 0x6f, 0x28, 0xaf, 0x35, 0xb3, 0x42,
	0xc9, 0x50, 0x3f, 0xaf, 0x79, 0xc5, 0x28, 0x93, 0x52, 0x59, 0x97, 0x36, 0xb4, 0x01, 0xdd, 0x75,
	0x16, 0xb2, 0x08, 0x90, 0xa7, 0x0d, 0x2b, 0x05, 0x67, 0x16, 0x68, 0xff, 0xf0, 0x85, 0x8b, 0xdf,
	0x07, 0xe8, 0x38, 0x63, 0x16, 0xae, 0xba, 0xb6, 0xf8, 0x05, 0x9a, 0x70, 0xb5, 0x65, 0x42, 0xc6,
	0x51, 0x12, 0xcd, 0x8f, 0xd3, 0xe9, 0x5d, 0x7a, 0xa8, 0x47, 0x49, 0x94, 0x85, 0x34, 0x7e, 0x8e,
	0xc6, 0xc6, 0xb2, 0x02, 0xe2, 0x51, 0x12, 0xcd, 0xff, 0x77, 0xf5, 0xc5, 0x28, 0x46, 0x99, 0xcf,
	0xe2, 0x73, 0xf4, 0x9f, 0x86, 0xef, 0x35, 0x18, 0x9b, 0xdb, 0xb6, 0x82, 0xf8, 0xa0, 0x53, 0xc9,
	0x4e, 0x42, 0xee, 0x5b, 0x5b, 0x01, 0x5e, 0xa1, 0xa9, 0x15, 0x5b, 0x50, 0xb5, 0x8d, 0x0f, 0x93,
	0x68, 0x7e, 0xb2, 0x7c, 0x46, 0xfc, 0x78, 0xa4, 0x1f, 0x8f, 0x5c, 0x86, 0xf1, 0xb2, 0x1e, 0x89,
	0x17, 0xe8, 0xc9, 0x86, 0x89, 0xb2, 0xd6, 0x90, 0x6f, 0x15, 0x87, 0x9c, 0x83, 0x6c, 0xe3, 0x71,
	0x12, 0xcd, 0x8f, 0xb2, 0xc7, 0xa1, 0xf0, 0x45, 0x71, 0xb8, 0x04, 0xd9, 0xe2, 0xcf, 0xe8, 0x62,
	0xb0, 0x04, 0x78, 0xce, 0x4c, 0xae, 0xc1, 0xa8, 0x5a, 0xaf, 0x21, 0x87, 0xdd, 0x2d, 0xab, 0x8d,
	0x05, 0x1e, 0x4f, 0x1c, 0x79, 0xa6, 0xfb, 0xd1, 0x81, 0x7f, 0x30, 0x59, 0x80, 0x7d, 0xea, 0x51,
	0x58, 0x20, 0x3c, 0x68, 0xe5, 0x06, 0x74, 0x23, 0xd6, 0x10, 0x4f, 0xdd, 0xbf, 0x5f, 0x13, 0xef,
	0xb3, 0xf7, 0xee, 0x9e, 0xad, 0x64, 0xbf, 0xd1, 0xaf, 0x9e, 0xf2, 0xd1, 0x61, 0xd2, 0xa3, 0xbb,
	0x74, 0xfc, 0x33, 0x1a, 0x9d, 0x46, 0xd9, 0xa9, 0x7e, 0x80, 0x78, 0xf7, 0xe6, 0xd7, 0x9f, 0x1f,
	0xb3, 0x25, 0x7a, 0x75, 0x4f, 0xd4, 0x1f, 0xce, 0x70, 0x37, 0xe1, 0xce, 0x9a, 0xe5, 0xd0, 0x21,
	0xbd, 0x42, 0x6f, 0x85, 0xf2, 0x7f, 0xa9, 0xb4, 0xda, 0xb5, 0xe4, 0xdf, 0xcf, 0x2f, 0x7d, 0xb4,
	0x97, 0xb9, 0xee, 0x96, 0x7f, 0x1d, 0xdd, 0x4c, 0x9c, 0x0b, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xfb, 0xfb, 0x84, 0x80, 0x02, 0x03, 0x00, 0x00,
}
