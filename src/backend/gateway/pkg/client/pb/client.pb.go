// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: pkg/client/pb/client.proto

package pb

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

type CreateClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId            int64  `protobuf:"varint,1,opt,name=companyId,proto3" json:"companyId,omitempty"`
	FirstName            string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	IdentificationNumber string `protobuf:"bytes,6,opt,name=identificationNumber,proto3" json:"identificationNumber,omitempty"`
	IdentificationType   string `protobuf:"bytes,7,opt,name=identificationType,proto3" json:"identificationType,omitempty"`
}

func (x *CreateClientRequest) Reset() {
	*x = CreateClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_pb_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientRequest) ProtoMessage() {}

func (x *CreateClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_pb_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientRequest.ProtoReflect.Descriptor instead.
func (*CreateClientRequest) Descriptor() ([]byte, []int) {
	return file_pkg_client_pb_client_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClientRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *CreateClientRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateClientRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateClientRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateClientRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateClientRequest) GetIdentificationNumber() string {
	if x != nil {
		return x.IdentificationNumber
	}
	return ""
}

func (x *CreateClientRequest) GetIdentificationType() string {
	if x != nil {
		return x.IdentificationType
	}
	return ""
}

type CreateClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateClientResponse) Reset() {
	*x = CreateClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_pb_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientResponse) ProtoMessage() {}

func (x *CreateClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_pb_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientResponse.ProtoReflect.Descriptor instead.
func (*CreateClientResponse) Descriptor() ([]byte, []int) {
	return file_pkg_client_pb_client_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClientResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateClientResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateClientResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetClientRequest) Reset() {
	*x = GetClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_pb_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientRequest) ProtoMessage() {}

func (x *GetClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_pb_client_proto_msgTypes[2]
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
	return file_pkg_client_pb_client_proto_rawDescGZIP(), []int{2}
}

func (x *GetClientRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyId            int64  `protobuf:"varint,2,opt,name=companyId,proto3" json:"companyId,omitempty"`
	FirstName            string `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	IdentificationNumber string `protobuf:"bytes,7,opt,name=identificationNumber,proto3" json:"identificationNumber,omitempty"`
	IdentificationType   string `protobuf:"bytes,8,opt,name=identificationType,proto3" json:"identificationType,omitempty"`
	Status               int64  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	Error                string `protobuf:"bytes,10,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetClientResponse) Reset() {
	*x = GetClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_pb_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientResponse) ProtoMessage() {}

func (x *GetClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_pb_client_proto_msgTypes[3]
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
	return file_pkg_client_pb_client_proto_rawDescGZIP(), []int{3}
}

func (x *GetClientResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetClientResponse) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *GetClientResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetClientResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetClientResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetClientResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetClientResponse) GetIdentificationNumber() string {
	if x != nil {
		return x.IdentificationNumber
	}
	return ""
}

func (x *GetClientResponse) GetIdentificationType() string {
	if x != nil {
		return x.IdentificationType
	}
	return ""
}

func (x *GetClientResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetClientResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetClientsRequest) Reset() {
	*x = GetClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_pb_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsRequest) ProtoMessage() {}

func (x *GetClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_pb_client_proto_msgTypes[4]
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
	return file_pkg_client_pb_client_proto_rawDescGZIP(), []int{4}
}

func (x *GetClientsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetClientsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetClientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*GetClientResponse `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
	Status  int64                `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error   string               `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetClientsResponse) Reset() {
	*x = GetClientsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_pb_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsResponse) ProtoMessage() {}

func (x *GetClientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_pb_client_proto_msgTypes[5]
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
	return file_pkg_client_pb_client_proto_rawDescGZIP(), []int{5}
}

func (x *GetClientsResponse) GetClients() []*GetClientResponse {
	if x != nil {
		return x.Clients
	}
	return nil
}

func (x *GetClientsResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetClientsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetClientsByCompanyIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetClientsByCompanyIdRequest) Reset() {
	*x = GetClientsByCompanyIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_pb_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsByCompanyIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsByCompanyIdRequest) ProtoMessage() {}

func (x *GetClientsByCompanyIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_pb_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsByCompanyIdRequest.ProtoReflect.Descriptor instead.
func (*GetClientsByCompanyIdRequest) Descriptor() ([]byte, []int) {
	return file_pkg_client_pb_client_proto_rawDescGZIP(), []int{6}
}

func (x *GetClientsByCompanyIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetClientsByCompanyIdRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetClientsByCompanyIdRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type UpdateClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	IdentificationNumber string `protobuf:"bytes,6,opt,name=identificationNumber,proto3" json:"identificationNumber,omitempty"`
	IdentificationType   string `protobuf:"bytes,7,opt,name=identificationType,proto3" json:"identificationType,omitempty"`
}

func (x *UpdateClientRequest) Reset() {
	*x = UpdateClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_pb_client_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClientRequest) ProtoMessage() {}

func (x *UpdateClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_pb_client_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClientRequest.ProtoReflect.Descriptor instead.
func (*UpdateClientRequest) Descriptor() ([]byte, []int) {
	return file_pkg_client_pb_client_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateClientRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateClientRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateClientRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateClientRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateClientRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateClientRequest) GetIdentificationNumber() string {
	if x != nil {
		return x.IdentificationNumber
	}
	return ""
}

func (x *UpdateClientRequest) GetIdentificationType() string {
	if x != nil {
		return x.IdentificationType
	}
	return ""
}

type UpdateClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpdateClientResponse) Reset() {
	*x = UpdateClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_pb_client_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClientResponse) ProtoMessage() {}

func (x *UpdateClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_pb_client_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClientResponse.ProtoReflect.Descriptor instead.
func (*UpdateClientResponse) Descriptor() ([]byte, []int) {
	return file_pkg_client_pb_client_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateClientResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateClientResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteClientRequest) Reset() {
	*x = DeleteClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_pb_client_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientRequest) ProtoMessage() {}

func (x *DeleteClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_pb_client_proto_msgTypes[9]
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
	return file_pkg_client_pb_client_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteClientRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteClientResponse) Reset() {
	*x = DeleteClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_client_pb_client_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientResponse) ProtoMessage() {}

func (x *DeleteClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_client_pb_client_proto_msgTypes[10]
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
	return file_pkg_client_pb_client_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteClientResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteClientResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_client_pb_client_proto protoreflect.FileDescriptor

var file_pkg_client_pb_client_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0xfd, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x32, 0x0a, 0x14,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x2e, 0x0a, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x54, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb9, 0x02, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x14, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x41, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x73, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5c,
	0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xef, 0x01, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x44,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0xa2, 0x03, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x41, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x64, 0x6f, 0x52, 0x67, 0x75, 0x65, 0x7a, 0x2f, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_client_pb_client_proto_rawDescOnce sync.Once
	file_pkg_client_pb_client_proto_rawDescData = file_pkg_client_pb_client_proto_rawDesc
)

func file_pkg_client_pb_client_proto_rawDescGZIP() []byte {
	file_pkg_client_pb_client_proto_rawDescOnce.Do(func() {
		file_pkg_client_pb_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_client_pb_client_proto_rawDescData)
	})
	return file_pkg_client_pb_client_proto_rawDescData
}

var file_pkg_client_pb_client_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_client_pb_client_proto_goTypes = []interface{}{
	(*CreateClientRequest)(nil),          // 0: pb.CreateClientRequest
	(*CreateClientResponse)(nil),         // 1: pb.CreateClientResponse
	(*GetClientRequest)(nil),             // 2: pb.GetClientRequest
	(*GetClientResponse)(nil),            // 3: pb.GetClientResponse
	(*GetClientsRequest)(nil),            // 4: pb.GetClientsRequest
	(*GetClientsResponse)(nil),           // 5: pb.GetClientsResponse
	(*GetClientsByCompanyIdRequest)(nil), // 6: pb.GetClientsByCompanyIdRequest
	(*UpdateClientRequest)(nil),          // 7: pb.UpdateClientRequest
	(*UpdateClientResponse)(nil),         // 8: pb.UpdateClientResponse
	(*DeleteClientRequest)(nil),          // 9: pb.DeleteClientRequest
	(*DeleteClientResponse)(nil),         // 10: pb.DeleteClientResponse
}
var file_pkg_client_pb_client_proto_depIdxs = []int32{
	3,  // 0: pb.GetClientsResponse.clients:type_name -> pb.GetClientResponse
	0,  // 1: pb.ClientService.CreateClient:input_type -> pb.CreateClientRequest
	2,  // 2: pb.ClientService.GetClient:input_type -> pb.GetClientRequest
	4,  // 3: pb.ClientService.GetClients:input_type -> pb.GetClientsRequest
	6,  // 4: pb.ClientService.GetClientsByCompanyId:input_type -> pb.GetClientsByCompanyIdRequest
	7,  // 5: pb.ClientService.UpdateClient:input_type -> pb.UpdateClientRequest
	9,  // 6: pb.ClientService.DeleteClient:input_type -> pb.DeleteClientRequest
	1,  // 7: pb.ClientService.CreateClient:output_type -> pb.CreateClientResponse
	3,  // 8: pb.ClientService.GetClient:output_type -> pb.GetClientResponse
	5,  // 9: pb.ClientService.GetClients:output_type -> pb.GetClientsResponse
	5,  // 10: pb.ClientService.GetClientsByCompanyId:output_type -> pb.GetClientsResponse
	8,  // 11: pb.ClientService.UpdateClient:output_type -> pb.UpdateClientResponse
	10, // 12: pb.ClientService.DeleteClient:output_type -> pb.DeleteClientResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_client_pb_client_proto_init() }
func file_pkg_client_pb_client_proto_init() {
	if File_pkg_client_pb_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_client_pb_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientRequest); i {
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
		file_pkg_client_pb_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientResponse); i {
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
		file_pkg_client_pb_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_client_pb_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_client_pb_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_client_pb_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_client_pb_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsByCompanyIdRequest); i {
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
		file_pkg_client_pb_client_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClientRequest); i {
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
		file_pkg_client_pb_client_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClientResponse); i {
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
		file_pkg_client_pb_client_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_client_pb_client_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_pkg_client_pb_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_client_pb_client_proto_goTypes,
		DependencyIndexes: file_pkg_client_pb_client_proto_depIdxs,
		MessageInfos:      file_pkg_client_pb_client_proto_msgTypes,
	}.Build()
	File_pkg_client_pb_client_proto = out.File
	file_pkg_client_pb_client_proto_rawDesc = nil
	file_pkg_client_pb_client_proto_goTypes = nil
	file_pkg_client_pb_client_proto_depIdxs = nil
}
