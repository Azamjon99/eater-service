// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package eater

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

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0xdf, 0xf8, 0xe3, 0xb5, 0xb0, 0xb9, 0x5d, 0xb7, 0x65, 0x30, 0x1e, 0xe0, 0xf9, 0x4e,
	0x02, 0x21, 0x4d, 0x42, 0x42, 0xea, 0x5a, 0x14, 0x09, 0x81, 0xa8, 0x56, 0xf1, 0x00, 0x6f, 0x66,
	0xbe, 0x1b, 0x91, 0xd2, 0xa4, 0x38, 0xee, 0x50, 0x3f, 0x22, 0xdf, 0x0a, 0x25, 0xf6, 0xdd, 0xb5,
	0x93, 0x74, 0x6f, 0xbd, 0xbf, 0x93, 0x73, 0x7c, 0xed, 0xeb, 0x34, 0x62, 0x58, 0xa1, 0xb9, 0xcb,
	0xae, 0x11, 0xd6, 0xa6, 0xb4, 0x65, 0xb2, 0x87, 0xca, 0xa2, 0xf1, 0xc5, 0x50, 0x69, 0x6d, 0xb0,
	0xaa, 0x7c, 0x39, 0xf8, 0xab, 0xf2, 0x1c, 0x2d, 0x3d, 0x59, 0x1a, 0x7d, 0xff, 0xe4, 0xc0, 0x28,
	0x9b, 0x15, 0xb7, 0xae, 0x7a, 0xfb, 0x6f, 0x20, 0x06, 0x9f, 0xea, 0x9c, 0xa5, 0xcb, 0x96, 0x17,
	0x62, 0x6f, 0x99, 0xdd, 0x16, 0x9b, 0x75, 0x43, 0xe5, 0x08, 0x82, 0xea, 0x0a, 0xff, 0x6c, 0xb0,
	0xb2, 0xc9, 0x38, 0x86, 0xd5, 0xba, 0x2c, 0x2a, 0x94, 0x53, 0xf1, 0x6c, 0x56, 0x16, 0x37, 0x99,
	0x59, 0x2d, 0x57, 0xd5, 0xac, 0xd4, 0x28, 0x27, 0x10, 0x03, 0xf2, 0x1f, 0x75, 0xb8, 0x8f, 0xf8,
	0x2a, 0xe4, 0xf7, 0xb5, 0x56, 0x16, 0x9b, 0xe4, 0x85, 0x29, 0x6f, 0xb2, 0x1c, 0x65, 0x02, 0x5d,
	0x48, 0x51, 0xa7, 0xbd, 0x9a, 0x8f, 0x9b, 0x8b, 0xe7, 0x29, 0xda, 0x28, 0xeb, 0x08, 0x5a, 0x84,
	0x82, 0x8e, 0xbb, 0x82, 0x4f, 0x79, 0x2f, 0xc4, 0x54, 0xeb, 0xa9, 0x3b, 0x5f, 0x29, 0x81, 0x0b,
	0xf2, 0x8e, 0x22, 0xe6, 0x6d, 0x1f, 0xc5, 0xd0, 0xb5, 0x46, 0xce, 0x43, 0x88, 0x6a, 0x32, 0x4f,
	0xda, 0x98, 0xfd, 0x73, 0xcc, 0x31, 0xf4, 0x47, 0x35, 0xfb, 0x5b, 0x98, 0xdb, 0x4e, 0xd1, 0x72,
	0xdb, 0x5c, 0x70, 0xdb, 0x21, 0xe3, 0x11, 0x7c, 0xc9, 0x2a, 0xc2, 0x97, 0x5b, 0x77, 0x0d, 0x12,
	0xe8, 0x42, 0x1e, 0x41, 0x9f, 0xc6, 0x97, 0x62, 0xaa, 0xf5, 0x42, 0x6d, 0x57, 0x58, 0xd8, 0x99,
	0x32, 0x5a, 0x4e, 0x20, 0x06, 0x7c, 0x29, 0xda, 0xdc, 0x47, 0x7c, 0x16, 0x07, 0x6e, 0x87, 0x61,
	0xca, 0x09, 0x74, 0x18, 0x05, 0x25, 0x7d, 0x12, 0xb7, 0x93, 0xa2, 0x8d, 0xdb, 0x89, 0x01, 0xb7,
	0xd3, 0xe6, 0x3e, 0xe2, 0x87, 0x98, 0xd4, 0xfb, 0x0d, 0x24, 0x3a, 0xa4, 0x33, 0xe8, 0x17, 0x28,
	0xf2, 0xd5, 0x4e, 0x9d, 0x47, 0xb6, 0xc8, 0xd5, 0x35, 0x7e, 0xab, 0x5f, 0x57, 0x29, 0x81, 0x0b,
	0x1e, 0x59, 0xc8, 0xbc, 0xed, 0x42, 0xec, 0xb9, 0x2b, 0xe4, 0x7c, 0x23, 0x08, 0x2a, 0x7e, 0x65,
	0x23, 0xe8, 0x9d, 0xe7, 0xe2, 0x49, 0x8a, 0xd6, 0xd9, 0xf6, 0x81, 0x7e, 0x92, 0xe7, 0x20, 0x20,
	0xbc, 0x94, 0x3b, 0x5c, 0x5a, 0x2a, 0xa8, 0x78, 0xa9, 0x08, 0x7a, 0x67, 0x2a, 0xf6, 0xeb, 0xdd,
	0x37, 0x90, 0x0e, 0xec, 0x18, 0xda, 0x88, 0x32, 0x4e, 0x7a, 0x14, 0x1e, 0xe1, 0x95, 0xb2, 0xf5,
	0xeb, 0x69, 0xd5, 0xc6, 0xa8, 0xc2, 0xca, 0x09, 0xc4, 0x80, 0x47, 0xd8, 0xe6, 0x3c, 0x42, 0x77,
	0x1a, 0x81, 0xd6, 0xfc, 0x29, 0xca, 0x33, 0xe8, 0x17, 0x78, 0x84, 0xbb, 0x74, 0x1f, 0xfd, 0x5b,
	0x9c, 0xd6, 0x9d, 0xb7, 0x75, 0xda, 0xf1, 0x6b, 0x78, 0x40, 0xa5, 0x45, 0xde, 0x3c, 0xfc, 0x90,
	0x5f, 0xe9, 0x83, 0x18, 0xd4, 0xdb, 0x9b, 0x63, 0x9e, 0xdd, 0xa1, 0xd9, 0xca, 0x31, 0x84, 0x25,
	0x65, 0x1d, 0xb6, 0xa8, 0x37, 0x2f, 0xc5, 0xd8, 0x6d, 0xe4, 0x5e, 0x71, 0xfb, 0x7f, 0x01, 0x7d,
	0x98, 0xc2, 0x5e, 0xee, 0x50, 0x5d, 0xe8, 0xe5, 0xd3, 0x9f, 0x8f, 0xe1, 0xbc, 0xf9, 0x28, 0xfd,
	0x7a, 0xd4, 0x7c, 0x5d, 0xde, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xfd, 0x93, 0xb3, 0xb3,
	0x06, 0x00, 0x00,
}
