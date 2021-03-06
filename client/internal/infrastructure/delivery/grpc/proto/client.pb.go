// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: internal/infrastructure/delivery/grpc/proto/client.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Email        string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Names        string `protobuf:"bytes,3,opt,name=Names,proto3" json:"Names,omitempty"`
	LastNames    string `protobuf:"bytes,4,opt,name=LastNames,proto3" json:"LastNames,omitempty"`
	BirthDate    string `protobuf:"bytes,5,opt,name=BirthDate,proto3" json:"BirthDate,omitempty"`
	Gender       string `protobuf:"bytes,6,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Address      string `protobuf:"bytes,7,opt,name=Address,proto3" json:"Address,omitempty"`
	Reference    string `protobuf:"bytes,8,opt,name=Reference,proto3" json:"Reference,omitempty"`
	CI           string `protobuf:"bytes,9,opt,name=CI,proto3" json:"CI,omitempty"`
	Telephone    string `protobuf:"bytes,10,opt,name=Telephone,proto3" json:"Telephone,omitempty"`
	LastIP       string `protobuf:"bytes,11,opt,name=LastIP,proto3" json:"LastIP,omitempty"`
	Locked       bool   `protobuf:"varint,12,opt,name=Locked,proto3" json:"Locked,omitempty"`
	BusinessName string `protobuf:"bytes,13,opt,name=BusinessName,proto3" json:"BusinessName,omitempty"`
	Nit          string `protobuf:"bytes,14,opt,name=Nit,proto3" json:"Nit,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Client) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Client) GetNames() string {
	if x != nil {
		return x.Names
	}
	return ""
}

func (x *Client) GetLastNames() string {
	if x != nil {
		return x.LastNames
	}
	return ""
}

func (x *Client) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *Client) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Client) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Client) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *Client) GetCI() string {
	if x != nil {
		return x.CI
	}
	return ""
}

func (x *Client) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

func (x *Client) GetLastIP() string {
	if x != nil {
		return x.LastIP
	}
	return ""
}

func (x *Client) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

func (x *Client) GetBusinessName() string {
	if x != nil {
		return x.BusinessName
	}
	return ""
}

func (x *Client) GetNit() string {
	if x != nil {
		return x.Nit
	}
	return ""
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP(), []int{1}
}

func (x *IdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdsRequest) Reset() {
	*x = IdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsRequest) ProtoMessage() {}

func (x *IdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdsRequest.ProtoReflect.Descriptor instead.
func (*IdsRequest) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP(), []int{2}
}

func (x *IdsRequest) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *Client `protobuf:"bytes,1,opt,name=Client,proto3" json:"Client,omitempty"`
}

func (x *ClientResponse) Reset() {
	*x = ClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResponse) ProtoMessage() {}

func (x *ClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResponse.ProtoReflect.Descriptor instead.
func (*ClientResponse) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP(), []int{3}
}

func (x *ClientResponse) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client   *Client `protobuf:"bytes,1,opt,name=Client,proto3" json:"Client,omitempty"`
	Password string  `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRequest) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *CreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP(), []int{5}
}

func (x *MessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ClientNilRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientNilRequest) Reset() {
	*x = ClientNilRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientNilRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientNilRequest) ProtoMessage() {}

func (x *ClientNilRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientNilRequest.ProtoReflect.Descriptor instead.
func (*ClientNilRequest) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP(), []int{6}
}

type ClientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*Client `protobuf:"bytes,1,rep,name=Clients,proto3" json:"Clients,omitempty"`
}

func (x *ClientsResponse) Reset() {
	*x = ClientsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientsResponse) ProtoMessage() {}

func (x *ClientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientsResponse.ProtoReflect.Descriptor instead.
func (*ClientsResponse) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP(), []int{7}
}

func (x *ClientsResponse) GetClients() []*Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Client *Client `protobuf:"bytes,2,opt,name=Client,proto3" json:"Client,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP(), []int{9}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok     bool    `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Client *Client `protobuf:"bytes,2,opt,name=Client,proto3" json:"Client,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP(), []int{10}
}

func (x *LoginResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *LoginResponse) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

var File_internal_infrastructure_delivery_grpc_proto_client_proto protoreflect.FileDescriptor

var file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDesc = []byte{
	0x0a, 0x38, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0xe4, 0x02, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x43, 0x49, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x43, 0x49, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x50, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x61, 0x73, 0x74, 0x49, 0x50, 0x12, 0x16, 0x0a, 0x06, 0x4c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x4c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x69, 0x74, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4e, 0x69, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0a, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x38, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x22, 0x53, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x47, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x47,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12,
	0x26, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x32, 0x8e, 0x04, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x11, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x18, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4e, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x03,
	0x42, 0x61, 0x6e, 0x12, 0x11, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x05, 0x55, 0x6e, 0x62, 0x61, 0x6e, 0x12, 0x11, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescOnce sync.Once
	file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescData = file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDesc
)

func file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescGZIP() []byte {
	file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescOnce.Do(func() {
		file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescData)
	})
	return file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDescData
}

var file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_internal_infrastructure_delivery_grpc_proto_client_proto_goTypes = []interface{}{
	(*Client)(nil),           // 0: Client.Client
	(*IdRequest)(nil),        // 1: Client.IdRequest
	(*IdsRequest)(nil),       // 2: Client.IdsRequest
	(*ClientResponse)(nil),   // 3: Client.ClientResponse
	(*CreateRequest)(nil),    // 4: Client.CreateRequest
	(*MessageResponse)(nil),  // 5: Client.MessageResponse
	(*ClientNilRequest)(nil), // 6: Client.ClientNilRequest
	(*ClientsResponse)(nil),  // 7: Client.ClientsResponse
	(*UpdateRequest)(nil),    // 8: Client.UpdateRequest
	(*LoginRequest)(nil),     // 9: Client.LoginRequest
	(*LoginResponse)(nil),    // 10: Client.LoginResponse
}
var file_internal_infrastructure_delivery_grpc_proto_client_proto_depIdxs = []int32{
	0,  // 0: Client.ClientResponse.Client:type_name -> Client.Client
	0,  // 1: Client.CreateRequest.Client:type_name -> Client.Client
	0,  // 2: Client.ClientsResponse.Clients:type_name -> Client.Client
	0,  // 3: Client.UpdateRequest.Client:type_name -> Client.Client
	0,  // 4: Client.LoginResponse.Client:type_name -> Client.Client
	1,  // 5: Client.ClientService.Get:input_type -> Client.IdRequest
	4,  // 6: Client.ClientService.Create:input_type -> Client.CreateRequest
	6,  // 7: Client.ClientService.List:input_type -> Client.ClientNilRequest
	1,  // 8: Client.ClientService.Delete:input_type -> Client.IdRequest
	2,  // 9: Client.ClientService.Deletes:input_type -> Client.IdsRequest
	8,  // 10: Client.ClientService.Update:input_type -> Client.UpdateRequest
	9,  // 11: Client.ClientService.Login:input_type -> Client.LoginRequest
	1,  // 12: Client.ClientService.Ban:input_type -> Client.IdRequest
	1,  // 13: Client.ClientService.Unban:input_type -> Client.IdRequest
	3,  // 14: Client.ClientService.Get:output_type -> Client.ClientResponse
	5,  // 15: Client.ClientService.Create:output_type -> Client.MessageResponse
	7,  // 16: Client.ClientService.List:output_type -> Client.ClientsResponse
	5,  // 17: Client.ClientService.Delete:output_type -> Client.MessageResponse
	5,  // 18: Client.ClientService.Deletes:output_type -> Client.MessageResponse
	5,  // 19: Client.ClientService.Update:output_type -> Client.MessageResponse
	10, // 20: Client.ClientService.Login:output_type -> Client.LoginResponse
	5,  // 21: Client.ClientService.Ban:output_type -> Client.MessageResponse
	5,  // 22: Client.ClientService.Unban:output_type -> Client.MessageResponse
	14, // [14:23] is the sub-list for method output_type
	5,  // [5:14] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_internal_infrastructure_delivery_grpc_proto_client_proto_init() }
func file_internal_infrastructure_delivery_grpc_proto_client_proto_init() {
	if File_internal_infrastructure_delivery_grpc_proto_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientNilRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_infrastructure_delivery_grpc_proto_client_proto_goTypes,
		DependencyIndexes: file_internal_infrastructure_delivery_grpc_proto_client_proto_depIdxs,
		MessageInfos:      file_internal_infrastructure_delivery_grpc_proto_client_proto_msgTypes,
	}.Build()
	File_internal_infrastructure_delivery_grpc_proto_client_proto = out.File
	file_internal_infrastructure_delivery_grpc_proto_client_proto_rawDesc = nil
	file_internal_infrastructure_delivery_grpc_proto_client_proto_goTypes = nil
	file_internal_infrastructure_delivery_grpc_proto_client_proto_depIdxs = nil
}
