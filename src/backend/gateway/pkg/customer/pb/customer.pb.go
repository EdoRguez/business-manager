// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: pkg/customer/pb/customer.proto

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

type CreateCustomerRequest struct {
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

func (x *CreateCustomerRequest) Reset() {
	*x = CreateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_customer_pb_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerRequest) ProtoMessage() {}

func (x *CreateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_customer_pb_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_customer_pb_customer_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCustomerRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *CreateCustomerRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateCustomerRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateCustomerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateCustomerRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateCustomerRequest) GetIdentificationNumber() string {
	if x != nil {
		return x.IdentificationNumber
	}
	return ""
}

func (x *CreateCustomerRequest) GetIdentificationType() string {
	if x != nil {
		return x.IdentificationType
	}
	return ""
}

type CreateCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateCustomerResponse) Reset() {
	*x = CreateCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_customer_pb_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerResponse) ProtoMessage() {}

func (x *CreateCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_customer_pb_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerResponse.ProtoReflect.Descriptor instead.
func (*CreateCustomerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_customer_pb_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCustomerResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateCustomerResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateCustomerResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCustomerRequest) Reset() {
	*x = GetCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_customer_pb_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRequest) ProtoMessage() {}

func (x *GetCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_customer_pb_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_customer_pb_customer_proto_rawDescGZIP(), []int{2}
}

func (x *GetCustomerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCustomerResponse struct {
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

func (x *GetCustomerResponse) Reset() {
	*x = GetCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_customer_pb_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerResponse) ProtoMessage() {}

func (x *GetCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_customer_pb_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_customer_pb_customer_proto_rawDescGZIP(), []int{3}
}

func (x *GetCustomerResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetCustomerResponse) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *GetCustomerResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetCustomerResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetCustomerResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetCustomerResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetCustomerResponse) GetIdentificationNumber() string {
	if x != nil {
		return x.IdentificationNumber
	}
	return ""
}

func (x *GetCustomerResponse) GetIdentificationType() string {
	if x != nil {
		return x.IdentificationType
	}
	return ""
}

func (x *GetCustomerResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetCustomerResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetCustomersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyId int64 `protobuf:"varint,1,opt,name=companyId,proto3" json:"companyId,omitempty"`
	Limit     int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset    int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetCustomersRequest) Reset() {
	*x = GetCustomersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_customer_pb_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomersRequest) ProtoMessage() {}

func (x *GetCustomersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_customer_pb_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomersRequest.ProtoReflect.Descriptor instead.
func (*GetCustomersRequest) Descriptor() ([]byte, []int) {
	return file_pkg_customer_pb_customer_proto_rawDescGZIP(), []int{4}
}

func (x *GetCustomersRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *GetCustomersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetCustomersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetCustomersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customers []*GetCustomerResponse `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
	Status    int64                  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error     string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetCustomersResponse) Reset() {
	*x = GetCustomersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_customer_pb_customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomersResponse) ProtoMessage() {}

func (x *GetCustomersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_customer_pb_customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomersResponse.ProtoReflect.Descriptor instead.
func (*GetCustomersResponse) Descriptor() ([]byte, []int) {
	return file_pkg_customer_pb_customer_proto_rawDescGZIP(), []int{5}
}

func (x *GetCustomersResponse) GetCustomers() []*GetCustomerResponse {
	if x != nil {
		return x.Customers
	}
	return nil
}

func (x *GetCustomersResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetCustomersResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateCustomerRequest struct {
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

func (x *UpdateCustomerRequest) Reset() {
	*x = UpdateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_customer_pb_customer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerRequest) ProtoMessage() {}

func (x *UpdateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_customer_pb_customer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerRequest.ProtoReflect.Descriptor instead.
func (*UpdateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_customer_pb_customer_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCustomerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCustomerRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateCustomerRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateCustomerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateCustomerRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateCustomerRequest) GetIdentificationNumber() string {
	if x != nil {
		return x.IdentificationNumber
	}
	return ""
}

func (x *UpdateCustomerRequest) GetIdentificationType() string {
	if x != nil {
		return x.IdentificationType
	}
	return ""
}

type UpdateCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UpdateCustomerResponse) Reset() {
	*x = UpdateCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_customer_pb_customer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCustomerResponse) ProtoMessage() {}

func (x *UpdateCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_customer_pb_customer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCustomerResponse.ProtoReflect.Descriptor instead.
func (*UpdateCustomerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_customer_pb_customer_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCustomerResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateCustomerResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type DeleteCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCustomerRequest) Reset() {
	*x = DeleteCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_customer_pb_customer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCustomerRequest) ProtoMessage() {}

func (x *DeleteCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_customer_pb_customer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCustomerRequest.ProtoReflect.Descriptor instead.
func (*DeleteCustomerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_customer_pb_customer_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCustomerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int64  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteCustomerResponse) Reset() {
	*x = DeleteCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_customer_pb_customer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCustomerResponse) ProtoMessage() {}

func (x *DeleteCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_customer_pb_customer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCustomerResponse.ProtoReflect.Descriptor instead.
func (*DeleteCustomerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_customer_pb_customer_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCustomerResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteCustomerResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_customer_pb_customer_proto protoreflect.FileDescriptor

var file_pkg_customer_pb_customer_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0xff, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x56, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x24,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xbb, 0x02, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x32, 0x0a, 0x14, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x61, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x7b, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a,
	0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0xf1, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x46, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x27,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32,
	0xef, 0x02, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x47, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x45, 0x64, 0x6f, 0x52, 0x67, 0x75, 0x65, 0x7a, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_customer_pb_customer_proto_rawDescOnce sync.Once
	file_pkg_customer_pb_customer_proto_rawDescData = file_pkg_customer_pb_customer_proto_rawDesc
)

func file_pkg_customer_pb_customer_proto_rawDescGZIP() []byte {
	file_pkg_customer_pb_customer_proto_rawDescOnce.Do(func() {
		file_pkg_customer_pb_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_customer_pb_customer_proto_rawDescData)
	})
	return file_pkg_customer_pb_customer_proto_rawDescData
}

var file_pkg_customer_pb_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_customer_pb_customer_proto_goTypes = []interface{}{
	(*CreateCustomerRequest)(nil),  // 0: pb.CreateCustomerRequest
	(*CreateCustomerResponse)(nil), // 1: pb.CreateCustomerResponse
	(*GetCustomerRequest)(nil),     // 2: pb.GetCustomerRequest
	(*GetCustomerResponse)(nil),    // 3: pb.GetCustomerResponse
	(*GetCustomersRequest)(nil),    // 4: pb.GetCustomersRequest
	(*GetCustomersResponse)(nil),   // 5: pb.GetCustomersResponse
	(*UpdateCustomerRequest)(nil),  // 6: pb.UpdateCustomerRequest
	(*UpdateCustomerResponse)(nil), // 7: pb.UpdateCustomerResponse
	(*DeleteCustomerRequest)(nil),  // 8: pb.DeleteCustomerRequest
	(*DeleteCustomerResponse)(nil), // 9: pb.DeleteCustomerResponse
}
var file_pkg_customer_pb_customer_proto_depIdxs = []int32{
	3, // 0: pb.GetCustomersResponse.customers:type_name -> pb.GetCustomerResponse
	0, // 1: pb.CustomerService.CreateCustomer:input_type -> pb.CreateCustomerRequest
	2, // 2: pb.CustomerService.GetCustomer:input_type -> pb.GetCustomerRequest
	4, // 3: pb.CustomerService.GetCustomers:input_type -> pb.GetCustomersRequest
	6, // 4: pb.CustomerService.UpdateCustomer:input_type -> pb.UpdateCustomerRequest
	8, // 5: pb.CustomerService.DeleteCustomer:input_type -> pb.DeleteCustomerRequest
	1, // 6: pb.CustomerService.CreateCustomer:output_type -> pb.CreateCustomerResponse
	3, // 7: pb.CustomerService.GetCustomer:output_type -> pb.GetCustomerResponse
	5, // 8: pb.CustomerService.GetCustomers:output_type -> pb.GetCustomersResponse
	7, // 9: pb.CustomerService.UpdateCustomer:output_type -> pb.UpdateCustomerResponse
	9, // 10: pb.CustomerService.DeleteCustomer:output_type -> pb.DeleteCustomerResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_customer_pb_customer_proto_init() }
func file_pkg_customer_pb_customer_proto_init() {
	if File_pkg_customer_pb_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_customer_pb_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerRequest); i {
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
		file_pkg_customer_pb_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCustomerResponse); i {
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
		file_pkg_customer_pb_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerRequest); i {
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
		file_pkg_customer_pb_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerResponse); i {
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
		file_pkg_customer_pb_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomersRequest); i {
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
		file_pkg_customer_pb_customer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomersResponse); i {
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
		file_pkg_customer_pb_customer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCustomerRequest); i {
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
		file_pkg_customer_pb_customer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCustomerResponse); i {
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
		file_pkg_customer_pb_customer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCustomerRequest); i {
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
		file_pkg_customer_pb_customer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCustomerResponse); i {
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
			RawDescriptor: file_pkg_customer_pb_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_customer_pb_customer_proto_goTypes,
		DependencyIndexes: file_pkg_customer_pb_customer_proto_depIdxs,
		MessageInfos:      file_pkg_customer_pb_customer_proto_msgTypes,
	}.Build()
	File_pkg_customer_pb_customer_proto = out.File
	file_pkg_customer_pb_customer_proto_rawDesc = nil
	file_pkg_customer_pb_customer_proto_goTypes = nil
	file_pkg_customer_pb_customer_proto_depIdxs = nil
}
