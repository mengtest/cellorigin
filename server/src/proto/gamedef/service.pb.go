// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

package gamedef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 一个服务的配置类型
type ServiceDefine struct {
	MainConfig bool `protobuf:"varint,10,opt,name=MainConfig,json=mainConfig" json:"MainConfig,omitempty"`
	// Peer配置
	IP        string `protobuf:"bytes,11,opt,name=IP,json=iP" json:"IP,omitempty"`
	Port      uint32 `protobuf:"varint,12,opt,name=Port,json=port" json:"Port,omitempty"`
	PeerName  string `protobuf:"bytes,13,opt,name=PeerName,json=peerName" json:"PeerName,omitempty"`
	PeerIndex int32  `protobuf:"varint,14,opt,name=PeerIndex,json=peerIndex" json:"PeerIndex,omitempty"`
	Enable    bool   `protobuf:"varint,15,opt,name=Enable,json=enable" json:"Enable,omitempty"`
	// 基础服务
	Version     int32  `protobuf:"varint,201,opt,name=Version,json=version" json:"Version,omitempty"`
	Name        string `protobuf:"bytes,205,opt,name=Name,json=name" json:"Name,omitempty"`
	DisplayName string `protobuf:"bytes,206,opt,name=DisplayName,json=displayName" json:"DisplayName,omitempty"`
}

func (m *ServiceDefine) Reset()                    { *m = ServiceDefine{} }
func (m *ServiceDefine) String() string            { return proto.CompactTextString(m) }
func (*ServiceDefine) ProtoMessage()               {}
func (*ServiceDefine) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// 服务配置, 所有服务通用
type ServiceFile struct {
	Service []*ServiceDefine `protobuf:"bytes,1,rep,name=Service,json=service" json:"Service,omitempty"`
}

func (m *ServiceFile) Reset()                    { *m = ServiceFile{} }
func (m *ServiceFile) String() string            { return proto.CompactTextString(m) }
func (*ServiceFile) ProtoMessage()               {}
func (*ServiceFile) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ServiceFile) GetService() []*ServiceDefine {
	if m != nil {
		return m.Service
	}
	return nil
}

// 对ServiceDefine配置的选择
type LocalFile struct {
	ServiceConfig string `protobuf:"bytes,1,opt,name=ServiceConfig,json=serviceConfig" json:"ServiceConfig,omitempty"`
}

func (m *LocalFile) Reset()                    { *m = LocalFile{} }
func (m *LocalFile) String() string            { return proto.CompactTextString(m) }
func (*LocalFile) ProtoMessage()               {}
func (*LocalFile) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func init() {
	proto.RegisterType((*ServiceDefine)(nil), "gamedef.ServiceDefine")
	proto.RegisterType((*ServiceFile)(nil), "gamedef.ServiceFile")
	proto.RegisterType((*LocalFile)(nil), "gamedef.LocalFile")
}

var fileDescriptor1 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0x49, 0xbf, 0x7c, 0xf9, 0x39, 0x31, 0x15, 0x8e, 0x50, 0x46, 0x11, 0xa9, 0xc1, 0x45,
	0x57, 0xc1, 0x9f, 0x0b, 0x70, 0x61, 0x15, 0x0a, 0x2a, 0x21, 0x82, 0xfb, 0xb4, 0x39, 0x29, 0x03,
	0xe9, 0x4c, 0x98, 0x94, 0xa2, 0x97, 0xe8, 0x42, 0xaf, 0xc7, 0xa5, 0xd3, 0x99, 0x69, 0xc5, 0x5d,
	0xde, 0xf7, 0xcd, 0xe1, 0x61, 0x1e, 0x48, 0x7b, 0x52, 0x1b, 0xbe, 0xa0, 0xbc, 0x53, 0x72, 0x2d,
	0x31, 0x5c, 0x56, 0x2b, 0xaa, 0xa9, 0xc9, 0xbe, 0x3d, 0x48, 0x5f, 0xec, 0x34, 0xa5, 0x86, 0x0b,
	0xc2, 0x33, 0x80, 0xa7, 0x8a, 0x8b, 0x3b, 0x29, 0x1a, 0xbe, 0x64, 0x30, 0xf6, 0x26, 0x51, 0x09,
	0xab, 0x7d, 0x83, 0x43, 0x18, 0xcc, 0x0a, 0x96, 0xe8, 0x3e, 0x2e, 0x07, 0xbc, 0x40, 0x04, 0xbf,
	0x90, 0x6a, 0xcd, 0x0e, 0x74, 0x93, 0x96, 0x7e, 0xa7, 0xbf, 0xf1, 0x04, 0xa2, 0x82, 0x48, 0x3d,
	0x6b, 0x08, 0x4b, 0xcd, 0x9f, 0x51, 0xe7, 0x32, 0x9e, 0x42, 0xbc, 0xdd, 0x66, 0xa2, 0xa6, 0x37,
	0x36, 0xd4, 0xe3, 0xff, 0x32, 0xee, 0x76, 0x05, 0x8e, 0x20, 0xb8, 0x17, 0xd5, 0xbc, 0x25, 0x76,
	0x68, 0xc8, 0x01, 0x99, 0x84, 0xc7, 0x10, 0xbe, 0x92, 0xea, 0xb9, 0x14, 0xec, 0xc3, 0x33, 0x47,
	0xe1, 0xc6, 0x66, 0x3c, 0x02, 0xdf, 0x80, 0x3e, 0x3d, 0x43, 0xf2, 0xc5, 0x96, 0x72, 0x0e, 0xc9,
	0x94, 0xf7, 0x5d, 0x5b, 0xbd, 0x9b, 0xed, 0xcb, 0x6e, 0x49, 0xfd, 0xdb, 0x65, 0xb7, 0x90, 0xb8,
	0x97, 0x3f, 0x70, 0x4d, 0xb8, 0x84, 0xd0, 0x45, 0xe6, 0x8d, 0xff, 0x4d, 0x92, 0xeb, 0x51, 0xee,
	0x24, 0xe5, 0x7f, 0x04, 0x95, 0xa1, 0x53, 0x99, 0x5d, 0x41, 0xfc, 0x28, 0x17, 0x55, 0x6b, 0xce,
	0x2f, 0xf6, 0x1e, 0x9d, 0x39, 0x4b, 0xdc, 0x79, 0xb7, 0xe5, 0x3c, 0x30, 0xfa, 0x6f, 0x7e, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x26, 0x74, 0xd4, 0x91, 0x8f, 0x01, 0x00, 0x00,
}
