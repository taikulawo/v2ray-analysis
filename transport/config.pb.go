package transport

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_transport_internet "v2ray.com/core/transport/internet"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Global transport settings. This affects all type of connections that go through V2Ray.
type Config struct {
	TransportSettings []*v2ray_core_transport_internet.TransportConfig `protobuf:"bytes,1,rep,name=transport_settings,json=transportSettings" json:"transport_settings,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetTransportSettings() []*v2ray_core_transport_internet.TransportConfig {
	if m != nil {
		return m.TransportSettings
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/transport/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x29, 0x4a, 0xcc, 0x2b,
	0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x81, 0x29, 0x2b, 0x4a, 0xd5, 0x83, 0x2b, 0x91, 0xd2, 0xc3, 0xa9, 0x39, 0x33,
	0xaf, 0x24, 0xb5, 0x28, 0x2f, 0x15, 0xd5, 0x14, 0xa5, 0x74, 0x2e, 0x36, 0x67, 0x30, 0x5f, 0x28,
	0x96, 0x4b, 0x08, 0xae, 0x38, 0xbe, 0x38, 0xb5, 0xa4, 0x24, 0x33, 0x2f, 0xbd, 0x58, 0x82, 0x51,
	0x81, 0x59, 0x83, 0xdb, 0x48, 0x4f, 0x0f, 0x9b, 0x65, 0x7a, 0x30, 0x23, 0xf5, 0x42, 0x60, 0x42,
	0x10, 0xb3, 0x82, 0x04, 0xe1, 0x6a, 0x82, 0xa1, 0x06, 0x39, 0xd9, 0x71, 0x49, 0x24, 0xe7, 0xe7,
	0x62, 0x35, 0x27, 0x80, 0x31, 0x8a, 0x13, 0xce, 0x59, 0xc5, 0x24, 0x12, 0x66, 0x14, 0x94, 0x58,
	0xa9, 0xe7, 0x0c, 0x52, 0x03, 0x37, 0x38, 0x89, 0x0d, 0xec, 0x5e, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x92, 0xcd, 0x15, 0x9b, 0x1e, 0x01, 0x00, 0x00,
}
