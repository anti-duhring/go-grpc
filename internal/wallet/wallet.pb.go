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

type TransferRequest struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount               float32  `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency             string   `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferRequest) Reset()         { *m = TransferRequest{} }
func (m *TransferRequest) String() string { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()    {}
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8911caa6983012d2, []int{3}
}

func (m *TransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRequest.Unmarshal(m, b)
}
func (m *TransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRequest.Marshal(b, m, deterministic)
}
func (m *TransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRequest.Merge(m, src)
}
func (m *TransferRequest) XXX_Size() int {
	return xxx_messageInfo_TransferRequest.Size(m)
}
func (m *TransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRequest proto.InternalMessageInfo

func (m *TransferRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TransferRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TransferRequest) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransferRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type TransferResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferResponse) Reset()         { *m = TransferResponse{} }
func (m *TransferResponse) String() string { return proto.CompactTextString(m) }
func (*TransferResponse) ProtoMessage()    {}
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8911caa6983012d2, []int{4}
}

func (m *TransferResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferResponse.Unmarshal(m, b)
}
func (m *TransferResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferResponse.Marshal(b, m, deterministic)
}
func (m *TransferResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferResponse.Merge(m, src)
}
func (m *TransferResponse) XXX_Size() int {
	return xxx_messageInfo_TransferResponse.Size(m)
}
func (m *TransferResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferResponse proto.InternalMessageInfo

func (m *TransferResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Wallet)(nil), "Wallet")
	proto.RegisterType((*CreateRequest)(nil), "CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "CreateResponse")
	proto.RegisterType((*TransferRequest)(nil), "TransferRequest")
	proto.RegisterType((*TransferResponse)(nil), "TransferResponse")
}

func init() {
	proto.RegisterFile("internal/wallet/wallet.proto", fileDescriptor_8911caa6983012d2)
}

var fileDescriptor_8911caa6983012d2 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x69, 0x2c, 0x69, 0x3b, 0xd2, 0x3f, 0xee, 0x41, 0x42, 0x11, 0x2c, 0x39, 0x68, 0x15,
	0xba, 0x81, 0xfa, 0x0d, 0xf4, 0xd4, 0x8b, 0x87, 0x28, 0x08, 0xde, 0xb6, 0xe9, 0x34, 0x5d, 0x4c,
	0x77, 0xe3, 0xee, 0xac, 0xe2, 0xb7, 0x97, 0x6e, 0x12, 0x4b, 0x03, 0x9e, 0x92, 0xb7, 0x0c, 0xef,
	0xfd, 0xde, 0x0c, 0x5c, 0x49, 0x45, 0x68, 0x94, 0x28, 0x92, 0x6f, 0x51, 0x14, 0x48, 0xf5, 0x87,
	0x97, 0x46, 0x93, 0x8e, 0x9f, 0x21, 0x7c, 0xf3, 0x9a, 0x8d, 0x20, 0x90, 0x9b, 0xa8, 0x33, 0xeb,
	0xcc, 0x07, 0x69, 0x20, 0x37, 0x2c, 0x82, 0xde, 0x5a, 0x14, 0x42, 0x65, 0x18, 0x05, 0xb3, 0xce,
	0x3c, 0x48, 0x1b, 0xc9, 0xa6, 0xd0, 0xcf, 0x9c, 0x31, 0xa8, 0xb2, 0x9f, 0xe8, 0xcc, 0xcf, 0xff,
	0xe9, 0xf8, 0x16, 0x86, 0x4f, 0x06, 0x05, 0x61, 0x8a, 0x9f, 0x0e, 0x2d, 0xb1, 0x4b, 0x08, 0x9d,
	0x45, 0xb3, 0x6a, 0xac, 0x6b, 0x15, 0xaf, 0x60, 0xd4, 0x0c, 0xda, 0x52, 0x2b, 0x8b, 0xff, 0x4d,
	0xb2, 0x6b, 0x08, 0x2b, 0x64, 0xcf, 0x71, 0xbe, 0xec, 0xf1, 0x8a, 0x38, 0xad, 0x9f, 0x63, 0x09,
	0xe3, 0x57, 0x23, 0x94, 0xdd, 0xa2, 0x69, 0x52, 0x19, 0x74, 0xb7, 0x46, 0xef, 0x6b, 0x27, 0xff,
	0x7f, 0x28, 0x48, 0xda, 0x7b, 0x0c, 0xd2, 0x80, 0xf4, 0x21, 0x4f, 0xec, 0xb5, 0x53, 0xe4, 0x4b,
	0x04, 0x69, 0xad, 0x4e, 0xea, 0x75, 0x5b, 0xf5, 0xee, 0x61, 0x72, 0x8c, 0x3a, 0x72, 0x5b, 0x12,
	0xe4, 0x6c, 0xc3, 0x5d, 0xa9, 0xe5, 0x07, 0x0c, 0x2b, 0xd0, 0x17, 0x34, 0x5f, 0x32, 0x43, 0x76,
	0x07, 0x61, 0x55, 0x99, 0x8d, 0xf8, 0xc9, 0x92, 0xa6, 0x63, 0xde, 0xda, 0x45, 0x02, 0xfd, 0x26,
	0x87, 0x4d, 0x78, 0xab, 0xdd, 0xf4, 0x82, 0xb7, 0x21, 0x1e, 0xe7, 0xef, 0x37, 0xb9, 0xa4, 0x9d,
	0x5b, 0xf3, 0x4c, 0xef, 0x13, 0xa1, 0x48, 0x2e, 0x36, 0x6e, 0x67, 0xa4, 0xca, 0x93, 0x5c, 0x2f,
	0x72, 0x53, 0x66, 0xf5, 0xdd, 0xd7, 0xa1, 0x3f, 0xfc, 0xc3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xef, 0x68, 0xb7, 0x74, 0x18, 0x02, 0x00, 0x00,
}
