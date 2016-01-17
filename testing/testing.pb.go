// Code generated by protoc-gen-go.
// source: testing/testing.proto
// DO NOT EDIT!

package lion_testing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Foo struct {
	One         string `protobuf:"bytes,1,opt,name=one" json:"one,omitempty"`
	Two         int32  `protobuf:"varint,2,opt,name=two" json:"two,omitempty"`
	StringField string `protobuf:"bytes,3,opt,name=string_field" json:"string_field,omitempty"`
	Int32Field  int32  `protobuf:"varint,4,opt,name=int32_field" json:"int32_field,omitempty"`
	Bar         *Bar   `protobuf:"bytes,5,opt,name=bar" json:"bar,omitempty"`
}

func (m *Foo) Reset()                    { *m = Foo{} }
func (m *Foo) String() string            { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()               {}
func (*Foo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Foo) GetBar() *Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

type Bar struct {
	One         string `protobuf:"bytes,1,opt,name=one" json:"one,omitempty"`
	Two         string `protobuf:"bytes,2,opt,name=two" json:"two,omitempty"`
	StringField string `protobuf:"bytes,3,opt,name=string_field" json:"string_field,omitempty"`
	Int32Field  int32  `protobuf:"varint,4,opt,name=int32_field" json:"int32_field,omitempty"`
}

func (m *Bar) Reset()                    { *m = Bar{} }
func (m *Bar) String() string            { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()               {}
func (*Bar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Baz struct {
	Bat *Baz_Bat `protobuf:"bytes,1,opt,name=bat" json:"bat,omitempty"`
}

func (m *Baz) Reset()                    { *m = Baz{} }
func (m *Baz) String() string            { return proto.CompactTextString(m) }
func (*Baz) ProtoMessage()               {}
func (*Baz) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Baz) GetBat() *Baz_Bat {
	if m != nil {
		return m.Bat
	}
	return nil
}

type Baz_Bat struct {
	Ban *Baz_Bat_Ban `protobuf:"bytes,1,opt,name=ban" json:"ban,omitempty"`
}

func (m *Baz_Bat) Reset()                    { *m = Baz_Bat{} }
func (m *Baz_Bat) String() string            { return proto.CompactTextString(m) }
func (*Baz_Bat) ProtoMessage()               {}
func (*Baz_Bat) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *Baz_Bat) GetBan() *Baz_Bat_Ban {
	if m != nil {
		return m.Ban
	}
	return nil
}

type Baz_Bat_Ban struct {
	StringField string `protobuf:"bytes,1,opt,name=string_field" json:"string_field,omitempty"`
	Int32Field  int32  `protobuf:"varint,2,opt,name=int32_field" json:"int32_field,omitempty"`
}

func (m *Baz_Bat_Ban) Reset()                    { *m = Baz_Bat_Ban{} }
func (m *Baz_Bat_Ban) String() string            { return proto.CompactTextString(m) }
func (*Baz_Bat_Ban) ProtoMessage()               {}
func (*Baz_Bat_Ban) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0, 0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Foo)(nil), "lion.testing.Foo")
	proto.RegisterType((*Bar)(nil), "lion.testing.Bar")
	proto.RegisterType((*Baz)(nil), "lion.testing.Baz")
	proto.RegisterType((*Baz_Bat)(nil), "lion.testing.Baz.Bat")
	proto.RegisterType((*Baz_Bat_Ban)(nil), "lion.testing.Baz.Bat.Ban")
	proto.RegisterType((*Empty)(nil), "lion.testing.Empty")
}

var fileDescriptor0 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x49, 0x2d, 0x2e,
	0xc9, 0xcc, 0x4b, 0xd7, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x3c, 0x39, 0x99,
	0xf9, 0x79, 0x7a, 0x50, 0x31, 0xa5, 0x2c, 0x2e, 0x66, 0xb7, 0xfc, 0x7c, 0x21, 0x6e, 0x2e, 0xe6,
	0xfc, 0xbc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x10, 0xa7, 0xa4, 0x3c, 0x5f, 0x82, 0x09,
	0xc8, 0x61, 0x15, 0x12, 0xe1, 0xe2, 0x29, 0x2e, 0x29, 0x02, 0x2a, 0x8d, 0x4f, 0xcb, 0x4c, 0xcd,
	0x49, 0x91, 0x60, 0x06, 0x2b, 0x11, 0xe6, 0xe2, 0xce, 0xcc, 0x2b, 0x31, 0x36, 0x82, 0x0a, 0xb2,
	0x80, 0x95, 0xca, 0x71, 0x31, 0x27, 0x25, 0x16, 0x49, 0xb0, 0x02, 0x39, 0xdc, 0x46, 0x82, 0x7a,
	0xc8, 0xf6, 0xe8, 0x39, 0x25, 0x16, 0x29, 0x79, 0x71, 0x31, 0x03, 0x29, 0x9c, 0x76, 0x71, 0x92,
	0x60, 0x97, 0x52, 0x17, 0x23, 0xc8, 0xb0, 0x2a, 0x21, 0x25, 0x90, 0x9d, 0x25, 0x60, 0xc3, 0xb8,
	0x8d, 0x44, 0xd1, 0xed, 0xac, 0x02, 0xe2, 0x12, 0xa9, 0x78, 0x90, 0xd2, 0x12, 0x21, 0x35, 0x90,
	0xd2, 0x3c, 0xa8, 0x52, 0x49, 0xac, 0x4a, 0x81, 0x38, 0x4f, 0xca, 0x00, 0xa4, 0x3c, 0x0f, 0xc3,
	0x31, 0x8c, 0xd8, 0x1c, 0x03, 0x0e, 0x23, 0x25, 0x76, 0x2e, 0x56, 0xd7, 0xdc, 0x82, 0x92, 0xca,
	0x24, 0x36, 0x70, 0x10, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x43, 0xc3, 0xa7, 0x7b,
	0x01, 0x00, 0x00,
}
