// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: proto/clients/client.proto

package clients

import (
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

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Fio   string `protobuf:"bytes,2,opt,name=fio,proto3" json:"fio,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[0]
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
	return file_proto_clients_client_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Client) GetFio() string {
	if x != nil {
		return x.Fio
	}
	return ""
}

func (x *Client) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Client) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type ListClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListClientRequest) Reset() {
	*x = ListClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClientRequest) ProtoMessage() {}

func (x *ListClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClientRequest.ProtoReflect.Descriptor instead.
func (*ListClientRequest) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{1}
}

type ListClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*Client `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *ListClientResponse) Reset() {
	*x = ListClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClientResponse) ProtoMessage() {}

func (x *ListClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClientResponse.ProtoReflect.Descriptor instead.
func (*ListClientResponse) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{2}
}

func (x *ListClientResponse) GetClients() []*Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

type GetClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetClientsRequest) Reset() {
	*x = GetClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsRequest) ProtoMessage() {}

func (x *GetClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsRequest.ProtoReflect.Descriptor instead.
func (*GetClientsRequest) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{3}
}

func (x *GetClientsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *GetClientResponse) Reset() {
	*x = GetClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientResponse) ProtoMessage() {}

func (x *GetClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientResponse.ProtoReflect.Descriptor instead.
func (*GetClientResponse) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{4}
}

func (x *GetClientResponse) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type GetClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetClientRequest) Reset() {
	*x = GetClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientRequest) ProtoMessage() {}

func (x *GetClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientRequest.ProtoReflect.Descriptor instead.
func (*GetClientRequest) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{5}
}

func (x *GetClientRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetClientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*Client `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *GetClientsResponse) Reset() {
	*x = GetClientsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsResponse) ProtoMessage() {}

func (x *GetClientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsResponse.ProtoReflect.Descriptor instead.
func (*GetClientsResponse) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{6}
}

func (x *GetClientsResponse) GetClients() []*Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

type GetClientsByTYPERequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *GetClientsByTYPERequest) Reset() {
	*x = GetClientsByTYPERequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsByTYPERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsByTYPERequest) ProtoMessage() {}

func (x *GetClientsByTYPERequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsByTYPERequest.ProtoReflect.Descriptor instead.
func (*GetClientsByTYPERequest) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{7}
}

func (x *GetClientsByTYPERequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type GetClientsByTYPEResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*Client `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *GetClientsByTYPEResponse) Reset() {
	*x = GetClientsByTYPEResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsByTYPEResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsByTYPEResponse) ProtoMessage() {}

func (x *GetClientsByTYPEResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsByTYPEResponse.ProtoReflect.Descriptor instead.
func (*GetClientsByTYPEResponse) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{8}
}

func (x *GetClientsByTYPEResponse) GetClients() []*Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

type GetClientsByMAKERequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Make string `protobuf:"bytes,1,opt,name=make,proto3" json:"make,omitempty"`
}

func (x *GetClientsByMAKERequest) Reset() {
	*x = GetClientsByMAKERequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsByMAKERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsByMAKERequest) ProtoMessage() {}

func (x *GetClientsByMAKERequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsByMAKERequest.ProtoReflect.Descriptor instead.
func (*GetClientsByMAKERequest) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{9}
}

func (x *GetClientsByMAKERequest) GetMake() string {
	if x != nil {
		return x.Make
	}
	return ""
}

type GetClientsByMAKEResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*Client `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *GetClientsByMAKEResponse) Reset() {
	*x = GetClientsByMAKEResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsByMAKEResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsByMAKEResponse) ProtoMessage() {}

func (x *GetClientsByMAKEResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsByMAKEResponse.ProtoReflect.Descriptor instead.
func (*GetClientsByMAKEResponse) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{10}
}

func (x *GetClientsByMAKEResponse) GetClients() []*Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

type AddClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *AddClientRequest) Reset() {
	*x = AddClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClientRequest) ProtoMessage() {}

func (x *AddClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClientRequest.ProtoReflect.Descriptor instead.
func (*AddClientRequest) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{11}
}

func (x *AddClientRequest) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type AddClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *AddClientResponse) Reset() {
	*x = AddClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClientResponse) ProtoMessage() {}

func (x *AddClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClientResponse.ProtoReflect.Descriptor instead.
func (*AddClientResponse) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{12}
}

func (x *AddClientResponse) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type DeleteClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteClientRequest) Reset() {
	*x = DeleteClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientRequest) ProtoMessage() {}

func (x *DeleteClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClientRequest.ProtoReflect.Descriptor instead.
func (*DeleteClientRequest) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteClientRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteClientResponse) Reset() {
	*x = DeleteClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_clients_client_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientResponse) ProtoMessage() {}

func (x *DeleteClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_clients_client_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClientResponse.ProtoReflect.Descriptor instead.
func (*DeleteClientResponse) Descriptor() ([]byte, []int) {
	return file_proto_clients_client_proto_rawDescGZIP(), []int{14}
}

var File_proto_clients_client_proto protoreflect.FileDescriptor

var file_proto_clients_client_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x22, 0x56, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x66, 0x69, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x13, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x4e, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63,
	0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x4e, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x2d, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x54, 0x59, 0x50, 0x45, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x54, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x54, 0x59, 0x50, 0x45, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x2d,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4d, 0x41,
	0x4b, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x6b,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x22, 0x54, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4d, 0x41, 0x4b,
	0x45, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x4a, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72,
	0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22,
	0x4b, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf3, 0x05, 0x0a, 0x0a,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x70, 0x69, 0x12, 0x63, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63,
	0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x60, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x63, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x29, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x54, 0x59, 0x50, 0x45, 0x12, 0x2f, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79,
	0x54, 0x59, 0x50, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x79, 0x54, 0x59, 0x50, 0x45, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4d, 0x41, 0x4b,
	0x45, 0x12, 0x2f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4d, 0x41, 0x4b, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x4d, 0x41, 0x4b, 0x45, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x12, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x63, 0x72, 0x6d, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_clients_client_proto_rawDescOnce sync.Once
	file_proto_clients_client_proto_rawDescData = file_proto_clients_client_proto_rawDesc
)

func file_proto_clients_client_proto_rawDescGZIP() []byte {
	file_proto_clients_client_proto_rawDescOnce.Do(func() {
		file_proto_clients_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_clients_client_proto_rawDescData)
	})
	return file_proto_clients_client_proto_rawDescData
}

var file_proto_clients_client_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_clients_client_proto_goTypes = []interface{}{
	(*Client)(nil),                   // 0: grpccrm.grpc.client.v1.Client
	(*ListClientRequest)(nil),        // 1: grpccrm.grpc.client.v1.ListClientRequest
	(*ListClientResponse)(nil),       // 2: grpccrm.grpc.client.v1.ListClientResponse
	(*GetClientsRequest)(nil),        // 3: grpccrm.grpc.client.v1.GetClientsRequest
	(*GetClientResponse)(nil),        // 4: grpccrm.grpc.client.v1.GetClientResponse
	(*GetClientRequest)(nil),         // 5: grpccrm.grpc.client.v1.GetClientRequest
	(*GetClientsResponse)(nil),       // 6: grpccrm.grpc.client.v1.GetClientsResponse
	(*GetClientsByTYPERequest)(nil),  // 7: grpccrm.grpc.client.v1.GetClientsByTYPERequest
	(*GetClientsByTYPEResponse)(nil), // 8: grpccrm.grpc.client.v1.GetClientsByTYPEResponse
	(*GetClientsByMAKERequest)(nil),  // 9: grpccrm.grpc.client.v1.GetClientsByMAKERequest
	(*GetClientsByMAKEResponse)(nil), // 10: grpccrm.grpc.client.v1.GetClientsByMAKEResponse
	(*AddClientRequest)(nil),         // 11: grpccrm.grpc.client.v1.AddClientRequest
	(*AddClientResponse)(nil),        // 12: grpccrm.grpc.client.v1.AddClientResponse
	(*DeleteClientRequest)(nil),      // 13: grpccrm.grpc.client.v1.DeleteClientRequest
	(*DeleteClientResponse)(nil),     // 14: grpccrm.grpc.client.v1.DeleteClientResponse
}
var file_proto_clients_client_proto_depIdxs = []int32{
	0,  // 0: grpccrm.grpc.client.v1.ListClientResponse.clients:type_name -> grpccrm.grpc.client.v1.Client
	0,  // 1: grpccrm.grpc.client.v1.GetClientResponse.client:type_name -> grpccrm.grpc.client.v1.Client
	0,  // 2: grpccrm.grpc.client.v1.GetClientsResponse.clients:type_name -> grpccrm.grpc.client.v1.Client
	0,  // 3: grpccrm.grpc.client.v1.GetClientsByTYPEResponse.clients:type_name -> grpccrm.grpc.client.v1.Client
	0,  // 4: grpccrm.grpc.client.v1.GetClientsByMAKEResponse.clients:type_name -> grpccrm.grpc.client.v1.Client
	0,  // 5: grpccrm.grpc.client.v1.AddClientRequest.client:type_name -> grpccrm.grpc.client.v1.Client
	0,  // 6: grpccrm.grpc.client.v1.AddClientResponse.client:type_name -> grpccrm.grpc.client.v1.Client
	1,  // 7: grpccrm.grpc.client.v1.ClientsApi.ListClient:input_type -> grpccrm.grpc.client.v1.ListClientRequest
	5,  // 8: grpccrm.grpc.client.v1.ClientsApi.GetClient:input_type -> grpccrm.grpc.client.v1.GetClientRequest
	3,  // 9: grpccrm.grpc.client.v1.ClientsApi.GetClients:input_type -> grpccrm.grpc.client.v1.GetClientsRequest
	7,  // 10: grpccrm.grpc.client.v1.ClientsApi.GetClientsByTYPE:input_type -> grpccrm.grpc.client.v1.GetClientsByTYPERequest
	9,  // 11: grpccrm.grpc.client.v1.ClientsApi.GetClientsByMAKE:input_type -> grpccrm.grpc.client.v1.GetClientsByMAKERequest
	11, // 12: grpccrm.grpc.client.v1.ClientsApi.AddClient:input_type -> grpccrm.grpc.client.v1.AddClientRequest
	13, // 13: grpccrm.grpc.client.v1.ClientsApi.DeleteClient:input_type -> grpccrm.grpc.client.v1.DeleteClientRequest
	2,  // 14: grpccrm.grpc.client.v1.ClientsApi.ListClient:output_type -> grpccrm.grpc.client.v1.ListClientResponse
	4,  // 15: grpccrm.grpc.client.v1.ClientsApi.GetClient:output_type -> grpccrm.grpc.client.v1.GetClientResponse
	6,  // 16: grpccrm.grpc.client.v1.ClientsApi.GetClients:output_type -> grpccrm.grpc.client.v1.GetClientsResponse
	8,  // 17: grpccrm.grpc.client.v1.ClientsApi.GetClientsByTYPE:output_type -> grpccrm.grpc.client.v1.GetClientsByTYPEResponse
	10, // 18: grpccrm.grpc.client.v1.ClientsApi.GetClientsByMAKE:output_type -> grpccrm.grpc.client.v1.GetClientsByMAKEResponse
	12, // 19: grpccrm.grpc.client.v1.ClientsApi.AddClient:output_type -> grpccrm.grpc.client.v1.AddClientResponse
	14, // 20: grpccrm.grpc.client.v1.ClientsApi.DeleteClient:output_type -> grpccrm.grpc.client.v1.DeleteClientResponse
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_clients_client_proto_init() }
func file_proto_clients_client_proto_init() {
	if File_proto_clients_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_clients_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_clients_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClientRequest); i {
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
		file_proto_clients_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClientResponse); i {
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
		file_proto_clients_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsRequest); i {
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
		file_proto_clients_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientResponse); i {
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
		file_proto_clients_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientRequest); i {
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
		file_proto_clients_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsResponse); i {
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
		file_proto_clients_client_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsByTYPERequest); i {
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
		file_proto_clients_client_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsByTYPEResponse); i {
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
		file_proto_clients_client_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsByMAKERequest); i {
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
		file_proto_clients_client_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsByMAKEResponse); i {
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
		file_proto_clients_client_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClientRequest); i {
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
		file_proto_clients_client_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClientResponse); i {
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
		file_proto_clients_client_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClientRequest); i {
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
		file_proto_clients_client_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClientResponse); i {
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
			RawDescriptor: file_proto_clients_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_clients_client_proto_goTypes,
		DependencyIndexes: file_proto_clients_client_proto_depIdxs,
		MessageInfos:      file_proto_clients_client_proto_msgTypes,
	}.Build()
	File_proto_clients_client_proto = out.File
	file_proto_clients_client_proto_rawDesc = nil
	file_proto_clients_client_proto_goTypes = nil
	file_proto_clients_client_proto_depIdxs = nil
}
