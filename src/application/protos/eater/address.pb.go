// Code generated by protoc-gen-go. DO NOT EDIT.
// source: address.proto

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

type Location struct {
	Longitude            float64  `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             float64  `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{0}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

type Address struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EaterId              string    `protobuf:"bytes,2,opt,name=eater_id,proto3" json:"eater_id,omitempty"`
	Name                 string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Location             *Location `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	CreatedAt            string    `protobuf:"bytes,5,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt            string    `protobuf:"bytes,6,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{1}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Address) GetEaterId() string {
	if m != nil {
		return m.EaterId
	}
	return ""
}

func (m *Address) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Address) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Address) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Address) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type AddAddressRequest struct {
	EaterId              string   `protobuf:"bytes,1,opt,name=eater_id,json=eaterId,proto3" json:"eater_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Longitude            float64  `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             float64  `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAddressRequest) Reset()         { *m = AddAddressRequest{} }
func (m *AddAddressRequest) String() string { return proto.CompactTextString(m) }
func (*AddAddressRequest) ProtoMessage()    {}
func (*AddAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{2}
}

func (m *AddAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAddressRequest.Unmarshal(m, b)
}
func (m *AddAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAddressRequest.Marshal(b, m, deterministic)
}
func (m *AddAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAddressRequest.Merge(m, src)
}
func (m *AddAddressRequest) XXX_Size() int {
	return xxx_messageInfo_AddAddressRequest.Size(m)
}
func (m *AddAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddAddressRequest proto.InternalMessageInfo

func (m *AddAddressRequest) GetEaterId() string {
	if m != nil {
		return m.EaterId
	}
	return ""
}

func (m *AddAddressRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddAddressRequest) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *AddAddressRequest) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

type AddAddressResponse struct {
	Address              *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddAddressResponse) Reset()         { *m = AddAddressResponse{} }
func (m *AddAddressResponse) String() string { return proto.CompactTextString(m) }
func (*AddAddressResponse) ProtoMessage()    {}
func (*AddAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{3}
}

func (m *AddAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddAddressResponse.Unmarshal(m, b)
}
func (m *AddAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddAddressResponse.Marshal(b, m, deterministic)
}
func (m *AddAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddAddressResponse.Merge(m, src)
}
func (m *AddAddressResponse) XXX_Size() int {
	return xxx_messageInfo_AddAddressResponse.Size(m)
}
func (m *AddAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddAddressResponse proto.InternalMessageInfo

func (m *AddAddressResponse) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type UpdateAddressRequest struct {
	AddressId            string   `protobuf:"bytes,1,opt,name=address_id,proto3" json:"address_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Longitude            float64  `protobuf:"fixed64,3,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             float64  `protobuf:"fixed64,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAddressRequest) Reset()         { *m = UpdateAddressRequest{} }
func (m *UpdateAddressRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAddressRequest) ProtoMessage()    {}
func (*UpdateAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{4}
}

func (m *UpdateAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAddressRequest.Unmarshal(m, b)
}
func (m *UpdateAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAddressRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAddressRequest.Merge(m, src)
}
func (m *UpdateAddressRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAddressRequest.Size(m)
}
func (m *UpdateAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAddressRequest proto.InternalMessageInfo

func (m *UpdateAddressRequest) GetAddressId() string {
	if m != nil {
		return m.AddressId
	}
	return ""
}

func (m *UpdateAddressRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateAddressRequest) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *UpdateAddressRequest) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

type UpdateAddressResponse struct {
	Address              *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAddressResponse) Reset()         { *m = UpdateAddressResponse{} }
func (m *UpdateAddressResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateAddressResponse) ProtoMessage()    {}
func (*UpdateAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{5}
}

func (m *UpdateAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAddressResponse.Unmarshal(m, b)
}
func (m *UpdateAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAddressResponse.Marshal(b, m, deterministic)
}
func (m *UpdateAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAddressResponse.Merge(m, src)
}
func (m *UpdateAddressResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateAddressResponse.Size(m)
}
func (m *UpdateAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAddressResponse proto.InternalMessageInfo

func (m *UpdateAddressResponse) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type DeleteAddressRequest struct {
	AddressId            string   `protobuf:"bytes,1,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAddressRequest) Reset()         { *m = DeleteAddressRequest{} }
func (m *DeleteAddressRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAddressRequest) ProtoMessage()    {}
func (*DeleteAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{6}
}

func (m *DeleteAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAddressRequest.Unmarshal(m, b)
}
func (m *DeleteAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAddressRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAddressRequest.Merge(m, src)
}
func (m *DeleteAddressRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAddressRequest.Size(m)
}
func (m *DeleteAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAddressRequest proto.InternalMessageInfo

func (m *DeleteAddressRequest) GetAddressId() string {
	if m != nil {
		return m.AddressId
	}
	return ""
}

type DeleteAddressResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAddressResponse) Reset()         { *m = DeleteAddressResponse{} }
func (m *DeleteAddressResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteAddressResponse) ProtoMessage()    {}
func (*DeleteAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{7}
}

func (m *DeleteAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAddressResponse.Unmarshal(m, b)
}
func (m *DeleteAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAddressResponse.Marshal(b, m, deterministic)
}
func (m *DeleteAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAddressResponse.Merge(m, src)
}
func (m *DeleteAddressResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteAddressResponse.Size(m)
}
func (m *DeleteAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAddressResponse proto.InternalMessageInfo

type GetAddressRequest struct {
	AddressId            string   `protobuf:"bytes,1,opt,name=address_id,json=addressId,proto3" json:"address_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAddressRequest) Reset()         { *m = GetAddressRequest{} }
func (m *GetAddressRequest) String() string { return proto.CompactTextString(m) }
func (*GetAddressRequest) ProtoMessage()    {}
func (*GetAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{8}
}

func (m *GetAddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAddressRequest.Unmarshal(m, b)
}
func (m *GetAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAddressRequest.Marshal(b, m, deterministic)
}
func (m *GetAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAddressRequest.Merge(m, src)
}
func (m *GetAddressRequest) XXX_Size() int {
	return xxx_messageInfo_GetAddressRequest.Size(m)
}
func (m *GetAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAddressRequest proto.InternalMessageInfo

func (m *GetAddressRequest) GetAddressId() string {
	if m != nil {
		return m.AddressId
	}
	return ""
}

type GetAddressResponse struct {
	Address              *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAddressResponse) Reset()         { *m = GetAddressResponse{} }
func (m *GetAddressResponse) String() string { return proto.CompactTextString(m) }
func (*GetAddressResponse) ProtoMessage()    {}
func (*GetAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{9}
}

func (m *GetAddressResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAddressResponse.Unmarshal(m, b)
}
func (m *GetAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAddressResponse.Marshal(b, m, deterministic)
}
func (m *GetAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAddressResponse.Merge(m, src)
}
func (m *GetAddressResponse) XXX_Size() int {
	return xxx_messageInfo_GetAddressResponse.Size(m)
}
func (m *GetAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAddressResponse proto.InternalMessageInfo

func (m *GetAddressResponse) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

type ListAddressByEaterRequest struct {
	EaterId              string   `protobuf:"bytes,1,opt,name=eater_id,json=eaterId,proto3" json:"eater_id,omitempty"`
	Sort                 string   `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
	Page                 int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32    `protobuf:"varint,4,opt,name=page_size,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAddressByEaterRequest) Reset()         { *m = ListAddressByEaterRequest{} }
func (m *ListAddressByEaterRequest) String() string { return proto.CompactTextString(m) }
func (*ListAddressByEaterRequest) ProtoMessage()    {}
func (*ListAddressByEaterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{10}
}

func (m *ListAddressByEaterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAddressByEaterRequest.Unmarshal(m, b)
}
func (m *ListAddressByEaterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAddressByEaterRequest.Marshal(b, m, deterministic)
}
func (m *ListAddressByEaterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAddressByEaterRequest.Merge(m, src)
}
func (m *ListAddressByEaterRequest) XXX_Size() int {
	return xxx_messageInfo_ListAddressByEaterRequest.Size(m)
}
func (m *ListAddressByEaterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAddressByEaterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAddressByEaterRequest proto.InternalMessageInfo

func (m *ListAddressByEaterRequest) GetEaterId() string {
	if m != nil {
		return m.EaterId
	}
	return ""
}

func (m *ListAddressByEaterRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *ListAddressByEaterRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListAddressByEaterRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type ListAddressByEaterResponse struct {
	Addresses            []*Address `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListAddressByEaterResponse) Reset()         { *m = ListAddressByEaterResponse{} }
func (m *ListAddressByEaterResponse) String() string { return proto.CompactTextString(m) }
func (*ListAddressByEaterResponse) ProtoMessage()    {}
func (*ListAddressByEaterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_982c640dad8fe78e, []int{11}
}

func (m *ListAddressByEaterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAddressByEaterResponse.Unmarshal(m, b)
}
func (m *ListAddressByEaterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAddressByEaterResponse.Marshal(b, m, deterministic)
}
func (m *ListAddressByEaterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAddressByEaterResponse.Merge(m, src)
}
func (m *ListAddressByEaterResponse) XXX_Size() int {
	return xxx_messageInfo_ListAddressByEaterResponse.Size(m)
}
func (m *ListAddressByEaterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAddressByEaterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAddressByEaterResponse proto.InternalMessageInfo

func (m *ListAddressByEaterResponse) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func init() {
	proto.RegisterType((*Location)(nil), "Location")
	proto.RegisterType((*Address)(nil), "Address")
	proto.RegisterType((*AddAddressRequest)(nil), "AddAddressRequest")
	proto.RegisterType((*AddAddressResponse)(nil), "AddAddressResponse")
	proto.RegisterType((*UpdateAddressRequest)(nil), "UpdateAddressRequest")
	proto.RegisterType((*UpdateAddressResponse)(nil), "UpdateAddressResponse")
	proto.RegisterType((*DeleteAddressRequest)(nil), "DeleteAddressRequest")
	proto.RegisterType((*DeleteAddressResponse)(nil), "DeleteAddressResponse")
	proto.RegisterType((*GetAddressRequest)(nil), "GetAddressRequest")
	proto.RegisterType((*GetAddressResponse)(nil), "GetAddressResponse")
	proto.RegisterType((*ListAddressByEaterRequest)(nil), "ListAddressByEaterRequest")
	proto.RegisterType((*ListAddressByEaterResponse)(nil), "ListAddressByEaterResponse")
}

func init() { proto.RegisterFile("address.proto", fileDescriptor_982c640dad8fe78e) }

var fileDescriptor_982c640dad8fe78e = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5f, 0x6b, 0xfa, 0x30,
	0x14, 0xa5, 0xf5, 0x4f, 0xdb, 0x2b, 0xbf, 0x1f, 0x18, 0x94, 0x55, 0xd9, 0x44, 0x02, 0x1b, 0x3e,
	0x75, 0xe0, 0x18, 0x08, 0x7b, 0x52, 0x1c, 0x43, 0xf0, 0xa9, 0xb0, 0x97, 0xbd, 0x48, 0x67, 0x82,
	0x14, 0xba, 0xb6, 0x6b, 0xe2, 0xc3, 0x86, 0x8f, 0xfb, 0x40, 0xfb, 0x88, 0x23, 0x49, 0x63, 0x6b,
	0x27, 0x63, 0xc2, 0x9e, 0x9a, 0x9c, 0x93, 0x73, 0xef, 0x39, 0x37, 0xa1, 0xf0, 0x2f, 0x20, 0x24,
	0xa3, 0x8c, 0x79, 0x69, 0x96, 0xf0, 0x04, 0xcf, 0xc1, 0x5e, 0x26, 0xeb, 0x80, 0x87, 0x49, 0x8c,
	0xce, 0xc1, 0x89, 0x92, 0x78, 0x13, 0xf2, 0x2d, 0xa1, 0xae, 0x31, 0x34, 0x46, 0x86, 0x5f, 0x00,
	0xa8, 0x0f, 0x76, 0x14, 0x70, 0x45, 0x9a, 0x92, 0xdc, 0xef, 0xf1, 0xa7, 0x01, 0xd6, 0x54, 0xd5,
	0x45, 0xff, 0xc1, 0x0c, 0x89, 0x94, 0x3b, 0xbe, 0x19, 0x12, 0xa1, 0xa3, 0x01, 0xa7, 0xd9, 0x2a,
	0x24, 0x52, 0xe7, 0xf8, 0xfb, 0x3d, 0x42, 0x50, 0x8f, 0x83, 0x17, 0xea, 0xd6, 0x24, 0x2e, 0xd7,
	0xe8, 0x12, 0xec, 0x28, 0x77, 0xe4, 0xd6, 0x87, 0xc6, 0xa8, 0x35, 0x76, 0x3c, 0x6d, 0xd1, 0xdf,
	0x53, 0x68, 0x00, 0xb0, 0xce, 0x44, 0x21, 0xb2, 0x0a, 0xb8, 0xdb, 0x90, 0x05, 0x4a, 0x88, 0xe0,
	0xb7, 0x29, 0xd1, 0x7c, 0x53, 0xf1, 0x05, 0x82, 0x77, 0xd0, 0x9e, 0x12, 0x92, 0x9b, 0xf6, 0xe9,
	0xeb, 0x96, 0x32, 0x8e, 0x7a, 0x25, 0xaf, 0x2a, 0x81, 0x25, 0xf7, 0x8b, 0xc2, 0xaa, 0x59, 0xb2,
	0x7a, 0x30, 0xb0, 0xda, 0x4f, 0x03, 0xab, 0x57, 0x06, 0x36, 0x01, 0x54, 0xee, 0xce, 0xd2, 0x24,
	0x66, 0x14, 0x61, 0xb0, 0xf2, 0xdb, 0x91, 0xdd, 0x5b, 0x63, 0xdb, 0xd3, 0x47, 0x34, 0x81, 0x3f,
	0x0c, 0xe8, 0x3c, 0xca, 0x18, 0x15, 0xef, 0x03, 0x80, 0xfc, 0x4c, 0xe1, 0xbe, 0x84, 0xfc, 0x71,
	0x80, 0x3b, 0xe8, 0x56, 0x5c, 0x9c, 0x90, 0xe1, 0x16, 0x3a, 0x73, 0x1a, 0xd1, 0x6f, 0x11, 0x2e,
	0x8e, 0x44, 0x70, 0x72, 0x64, 0x41, 0xf0, 0x19, 0x74, 0x2b, 0x32, 0xd5, 0x13, 0x8f, 0xa1, 0xfd,
	0x40, 0xf9, 0x69, 0xc5, 0x26, 0x80, 0xca, 0x9a, 0x13, 0xdc, 0xef, 0xa0, 0xb7, 0x0c, 0x99, 0x96,
	0xce, 0xde, 0xee, 0xc5, 0x0b, 0xf9, 0xdd, 0x0b, 0x62, 0x49, 0xc6, 0xf5, 0x05, 0x88, 0xb5, 0xc0,
	0xd2, 0x60, 0xa3, 0x66, 0xdf, 0xf0, 0xe5, 0x5a, 0x5c, 0x8a, 0xf8, 0xae, 0x58, 0xf8, 0xae, 0xe6,
	0xde, 0xf0, 0x0b, 0x00, 0xcf, 0xa1, 0x7f, 0xac, 0x7b, 0xee, 0xff, 0x0a, 0x74, 0x44, 0x2a, 0x12,
	0xd4, 0x0e, 0x12, 0x14, 0xd4, 0xcc, 0x79, 0xb2, 0xbc, 0x6b, 0x69, 0xec, 0xb9, 0x29, 0x7f, 0x04,
	0x37, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x11, 0x83, 0x3f, 0x84, 0x19, 0x04, 0x00, 0x00,
}
