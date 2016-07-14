// Code generated by protoc-gen-go.
// source: example/current/main.proto
// DO NOT EDIT!

package main

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

type Bar struct {
	One string `protobuf:"bytes,1,opt,name=one" json:"one,omitempty"`
}

func (m *Bar) Reset()                    { *m = Bar{} }
func (m *Bar) String() string            { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()               {}
func (*Bar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Foo struct {
	Bar   *Bar   `protobuf:"bytes,1,opt,name=bar" json:"bar,omitempty"`
	Two   string `protobuf:"bytes,2,opt,name=two" json:"two,omitempty"`
	Three uint64 `protobuf:"varint,3,opt,name=three" json:"three,omitempty"`
}

func (m *Foo) Reset()                    { *m = Foo{} }
func (m *Foo) String() string            { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()               {}
func (*Foo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Foo) GetBar() *Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

func init() {
	proto.RegisterType((*Bar)(nil), "main.Bar")
	proto.RegisterType((*Foo)(nil), "main.Foo")
}

func init() { proto.RegisterFile("example/current/main.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x4f, 0x2e, 0x2d, 0x2a, 0x4a, 0xcd, 0x2b, 0xd1, 0xcf, 0x4d, 0xcc, 0xcc,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xc4, 0xb9, 0x98, 0x9d, 0x12,
	0x8b, 0x84, 0x04, 0xb8, 0x98, 0xf3, 0xf3, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40,
	0x4c, 0x25, 0x2f, 0x2e, 0x66, 0xb7, 0xfc, 0x7c, 0x21, 0x69, 0x2e, 0xe6, 0xa4, 0xc4, 0x22, 0xb0,
	0x04, 0xb7, 0x11, 0xa7, 0x1e, 0x58, 0x3f, 0x50, 0x43, 0x10, 0x48, 0x14, 0xa4, 0xab, 0xa4, 0x3c,
	0x5f, 0x82, 0x09, 0xa2, 0x0b, 0xc8, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x28, 0x4a, 0x4d, 0x95,
	0x60, 0x06, 0x8a, 0xb1, 0x04, 0x41, 0x38, 0x49, 0x6c, 0x60, 0x1b, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x18, 0xbe, 0x99, 0xe0, 0x8f, 0x00, 0x00, 0x00,
}
