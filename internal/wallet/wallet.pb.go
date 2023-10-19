// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/wallet/wallet.proto

package wallet

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

type Wallet struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Balance              float32  `protobuf:"fixed32,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Wallet) Reset()         { *m = Wallet{} }
func (m *Wallet) String() string { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()    {}
func (*Wallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8911caa6983012d2, []int{0}
}

func (m *Wallet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wallet.Unmarshal(m, b)
}
func (m *Wallet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wallet.Marshal(b, m, deterministic)
}
func (m *Wallet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wallet.Merge(m, src)
}
func (m *Wallet) XXX_Size() int {
	return xxx_messageInfo_Wallet.Size(m)
}
func (m *Wallet) XXX_DiscardUnknown() {
	xxx_messageInfo_Wallet.DiscardUnknown(m)
}

var xxx_messageInfo_Wallet proto.InternalMessageInfo

func (m *Wallet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Wallet) GetBalance() float32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Wallet) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type CreateRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8911caa6983012d2, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type CreateResponse struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Wallet               *Wallet  `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8911caa6983012d2, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateResponse) GetWallet() *Wallet {
	if m != nil {
		return m.Wallet
	}
	return nil
}

func init() {
	proto.RegisterType((*Wallet)(nil), "Wallet")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
}

func init() {
	proto.RegisterFile("internal/wallet/wallet.proto", fileDescriptor_8911caa6983012d2)
}

var fileDescriptor_8911caa6983012d2 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0x69, 0x85, 0x9c, 0x3e, 0xb9, 0x0a, 0x19, 0xa4, 0x1c, 0x82, 0x47, 0x07, 0xad, 0xc3,
	0xa5, 0x70, 0x6e, 0x8e, 0x3a, 0xdd, 0xe2, 0x50, 0x07, 0xc1, 0x2d, 0x4d, 0x1f, 0xbd, 0x40, 0x4d,
	0xea, 0x6b, 0xa2, 0xf8, 0xdf, 0x0b, 0x4d, 0x2a, 0x74, 0xb8, 0x29, 0xfc, 0x1e, 0x1f, 0x5f, 0x7e,
	0x7c, 0x70, 0xa3, 0x8d, 0x43, 0x32, 0xb2, 0xaf, 0x7e, 0x64, 0xdf, 0xa3, 0x8b, 0x8f, 0x18, 0xc8,
	0x3a, 0x5b, 0xbc, 0x02, 0x7b, 0x9f, 0x98, 0x67, 0x90, 0xea, 0x36, 0x4f, 0xb6, 0x49, 0x79, 0x51,
	0xa7, 0xba, 0xe5, 0x39, 0xac, 0x1a, 0xd9, 0x4b, 0xa3, 0x30, 0x4f, 0xb7, 0x49, 0x99, 0xd6, 0x33,
	0xf2, 0x0d, 0x9c, 0x2b, 0x4f, 0x84, 0x46, 0xfd, 0xe6, 0x67, 0x53, 0xfe, 0x9f, 0x8b, 0x7b, 0x58,
	0xbf, 0x10, 0x4a, 0x87, 0x35, 0x7e, 0x79, 0x1c, 0x1d, 0xbf, 0x06, 0xe6, 0x47, 0xa4, 0xc3, 0x5c,
	0x1d, 0xa9, 0x38, 0x40, 0x36, 0x07, 0xc7, 0xc1, 0x9a, 0x11, 0x4f, 0x25, 0xf9, 0x2d, 0xb0, 0xa0,
	0x3c, 0x79, 0x5c, 0xee, 0x57, 0x22, 0x18, 0xd7, 0xf1, 0xbc, 0x7f, 0x82, 0x75, 0xb8, 0xbc, 0x21,
	0x7d, 0x6b, 0x85, 0xfc, 0x01, 0x58, 0xe8, 0xe6, 0x99, 0x58, 0xd8, 0x6c, 0xae, 0xc4, 0xf2, 0xd3,
	0xe7, 0xf2, 0xe3, 0xae, 0xd3, 0xee, 0xe8, 0x1b, 0xa1, 0xec, 0x67, 0x25, 0x8d, 0xd3, 0xbb, 0xd6,
	0x1f, 0x49, 0x9b, 0xae, 0xea, 0xec, 0xae, 0xa3, 0x41, 0xc5, 0xbd, 0x1a, 0x36, 0x0d, 0xf6, 0xf8,
	0x17, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x09, 0xb7, 0x9f, 0x50, 0x01, 0x00, 0x00,
}