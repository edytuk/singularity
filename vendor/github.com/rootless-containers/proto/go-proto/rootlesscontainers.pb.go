// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rootlesscontainers.proto

/*
Package rootlesscontainers is a generated protocol buffer package.

The rootlesscontainers package is maintained at https://rootlesscontaine.rs/ .
If you want to extend the resource definition, please open a PR.

It is generated from these files:
	rootlesscontainers.proto

It has these top-level messages:
	Resource
*/
package rootlesscontainers

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Resource defines the schema for "user.rootlesscontainers" xattr values.
// The resource can be used as a persistent storage for emulated `chown(2)` syscall.
// Syscall emulators SHOULD try to hide this xattr from the emulated environment.
type Resource struct {
	// Zero-value MUST be parsed as a literally zero-value, not "unset".
	// To keep both uid and gid unchaged, the entire xattr value SHOULD be removed.
	// To keep either one of uid or gid unchaged, 0xFFFFFFFF (in other words,
	// `(uint32_t) -1`, see also chown(2)) value SHOULD be set.
	// (Because some protobuf bindings cannot distinguish "unset" from zero-value.)
	Uid uint32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Gid uint32 `protobuf:"varint,2,opt,name=gid" json:"gid,omitempty"`
}

func (m *Resource) Reset()                    { *m = Resource{} }
func (m *Resource) String() string            { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()               {}
func (*Resource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Resource) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Resource) GetGid() uint32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func init() {
	proto.RegisterType((*Resource)(nil), "rootlesscontainers.Resource")
}

func init() { proto.RegisterFile("rootlesscontainers.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xca, 0xcf, 0x2f,
	0xc9, 0x49, 0x2d, 0x2e, 0x4e, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d, 0x2a, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc2, 0x94, 0x51, 0xd2, 0xe3, 0xe2, 0x08, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x15, 0x12, 0xe0, 0x62, 0x2e, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0d, 0x02, 0x31, 0x41, 0x22, 0xe9, 0x99, 0x29, 0x12, 0x4c, 0x10, 0x91, 0xf4, 0xcc, 0x94,
	0x24, 0x36, 0xb0, 0x51, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xac, 0x07, 0x53, 0x66,
	0x00, 0x00, 0x00,
}