// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.25.3
// source: proto/whatsapp.proto

package whatsapp

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

type SendOrderCustomerMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CompanyId     int64                  `protobuf:"varint,1,opt,name=companyId,proto3" json:"companyId,omitempty"`
	ToPhone       string                 `protobuf:"bytes,2,opt,name=toPhone,proto3" json:"toPhone,omitempty"`
	CustomerName  string                 `protobuf:"bytes,3,opt,name=customerName,proto3" json:"customerName,omitempty"`
	Products      []*OrderProductRequest `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendOrderCustomerMessageRequest) Reset() {
	*x = SendOrderCustomerMessageRequest{}
	mi := &file_proto_whatsapp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendOrderCustomerMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOrderCustomerMessageRequest) ProtoMessage() {}

func (x *SendOrderCustomerMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatsapp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOrderCustomerMessageRequest.ProtoReflect.Descriptor instead.
func (*SendOrderCustomerMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{0}
}

func (x *SendOrderCustomerMessageRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *SendOrderCustomerMessageRequest) GetToPhone() string {
	if x != nil {
		return x.ToPhone
	}
	return ""
}

func (x *SendOrderCustomerMessageRequest) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *SendOrderCustomerMessageRequest) GetProducts() []*OrderProductRequest {
	if x != nil {
		return x.Products
	}
	return nil
}

type SendOrderCustomerMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendOrderCustomerMessageResponse) Reset() {
	*x = SendOrderCustomerMessageResponse{}
	mi := &file_proto_whatsapp_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendOrderCustomerMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOrderCustomerMessageResponse) ProtoMessage() {}

func (x *SendOrderCustomerMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatsapp_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOrderCustomerMessageResponse.ProtoReflect.Descriptor instead.
func (*SendOrderCustomerMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{1}
}

func (x *SendOrderCustomerMessageResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SendOrderCustomerMessageResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SendOrderBusinessMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CompanyId     int64                  `protobuf:"varint,1,opt,name=companyId,proto3" json:"companyId,omitempty"`
	CustomerName  string                 `protobuf:"bytes,2,opt,name=customerName,proto3" json:"customerName,omitempty"`
	ContactNumber string                 `protobuf:"bytes,3,opt,name=contactNumber,proto3" json:"contactNumber,omitempty"`
	Products      []*OrderProductRequest `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendOrderBusinessMessageRequest) Reset() {
	*x = SendOrderBusinessMessageRequest{}
	mi := &file_proto_whatsapp_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendOrderBusinessMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOrderBusinessMessageRequest) ProtoMessage() {}

func (x *SendOrderBusinessMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatsapp_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOrderBusinessMessageRequest.ProtoReflect.Descriptor instead.
func (*SendOrderBusinessMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{2}
}

func (x *SendOrderBusinessMessageRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *SendOrderBusinessMessageRequest) GetCustomerName() string {
	if x != nil {
		return x.CustomerName
	}
	return ""
}

func (x *SendOrderBusinessMessageRequest) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
	}
	return ""
}

func (x *SendOrderBusinessMessageRequest) GetProducts() []*OrderProductRequest {
	if x != nil {
		return x.Products
	}
	return nil
}

type SendOrderBusinessMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendOrderBusinessMessageResponse) Reset() {
	*x = SendOrderBusinessMessageResponse{}
	mi := &file_proto_whatsapp_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendOrderBusinessMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendOrderBusinessMessageResponse) ProtoMessage() {}

func (x *SendOrderBusinessMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatsapp_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendOrderBusinessMessageResponse.ProtoReflect.Descriptor instead.
func (*SendOrderBusinessMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{3}
}

func (x *SendOrderBusinessMessageResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *SendOrderBusinessMessageResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type OrderProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Quantity      uint64                 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price         uint64                 `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderProductRequest) Reset() {
	*x = OrderProductRequest{}
	mi := &file_proto_whatsapp_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProductRequest) ProtoMessage() {}

func (x *OrderProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatsapp_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProductRequest.ProtoReflect.Descriptor instead.
func (*OrderProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{4}
}

func (x *OrderProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderProductRequest) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderProductRequest) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateBusinessPhoneRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CompanyId     int64                  `protobuf:"varint,1,opt,name=companyId,proto3" json:"companyId,omitempty"`
	Phone         string                 `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBusinessPhoneRequest) Reset() {
	*x = CreateBusinessPhoneRequest{}
	mi := &file_proto_whatsapp_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBusinessPhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBusinessPhoneRequest) ProtoMessage() {}

func (x *CreateBusinessPhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatsapp_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBusinessPhoneRequest.ProtoReflect.Descriptor instead.
func (*CreateBusinessPhoneRequest) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{5}
}

func (x *CreateBusinessPhoneRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *CreateBusinessPhoneRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type CreateBusinessPhoneResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBusinessPhoneResponse) Reset() {
	*x = CreateBusinessPhoneResponse{}
	mi := &file_proto_whatsapp_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBusinessPhoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBusinessPhoneResponse) ProtoMessage() {}

func (x *CreateBusinessPhoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatsapp_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBusinessPhoneResponse.ProtoReflect.Descriptor instead.
func (*CreateBusinessPhoneResponse) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{6}
}

func (x *CreateBusinessPhoneResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateBusinessPhoneResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateBusinessPhoneRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CompanyId     int64                  `protobuf:"varint,1,opt,name=companyId,proto3" json:"companyId,omitempty"`
	Phone         string                 `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBusinessPhoneRequest) Reset() {
	*x = UpdateBusinessPhoneRequest{}
	mi := &file_proto_whatsapp_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBusinessPhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBusinessPhoneRequest) ProtoMessage() {}

func (x *UpdateBusinessPhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatsapp_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBusinessPhoneRequest.ProtoReflect.Descriptor instead.
func (*UpdateBusinessPhoneRequest) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBusinessPhoneRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *UpdateBusinessPhoneRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type UpdateBusinessPhoneResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        int64                  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBusinessPhoneResponse) Reset() {
	*x = UpdateBusinessPhoneResponse{}
	mi := &file_proto_whatsapp_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBusinessPhoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBusinessPhoneResponse) ProtoMessage() {}

func (x *UpdateBusinessPhoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatsapp_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBusinessPhoneResponse.ProtoReflect.Descriptor instead.
func (*UpdateBusinessPhoneResponse) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateBusinessPhoneResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateBusinessPhoneResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetBusinessPhoneByCompanyIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CompanyId     int64                  `protobuf:"varint,1,opt,name=companyId,proto3" json:"companyId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBusinessPhoneByCompanyIdRequest) Reset() {
	*x = GetBusinessPhoneByCompanyIdRequest{}
	mi := &file_proto_whatsapp_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBusinessPhoneByCompanyIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessPhoneByCompanyIdRequest) ProtoMessage() {}

func (x *GetBusinessPhoneByCompanyIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatsapp_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessPhoneByCompanyIdRequest.ProtoReflect.Descriptor instead.
func (*GetBusinessPhoneByCompanyIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{9}
}

func (x *GetBusinessPhoneByCompanyIdRequest) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type GetBusinessPhoneByCompanyIdResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyId     int64                  `protobuf:"varint,2,opt,name=companyId,proto3" json:"companyId,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Status        int64                  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBusinessPhoneByCompanyIdResponse) Reset() {
	*x = GetBusinessPhoneByCompanyIdResponse{}
	mi := &file_proto_whatsapp_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBusinessPhoneByCompanyIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessPhoneByCompanyIdResponse) ProtoMessage() {}

func (x *GetBusinessPhoneByCompanyIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_whatsapp_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessPhoneByCompanyIdResponse.ProtoReflect.Descriptor instead.
func (*GetBusinessPhoneByCompanyIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{10}
}

func (x *GetBusinessPhoneByCompanyIdResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetBusinessPhoneByCompanyIdResponse) GetCompanyId() int64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *GetBusinessPhoneByCompanyIdResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetBusinessPhoneByCompanyIdResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetBusinessPhoneByCompanyIdResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_whatsapp_proto protoreflect.FileDescriptor

var file_proto_whatsapp_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70,
	0x22, 0xb8, 0x01, 0x0a, 0x1f, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x39, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x50, 0x0a, 0x20, 0x53,
	0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xc4, 0x01,
	0x0a, 0x1f, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x68,
	0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x22, 0x50, 0x0a, 0x20, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5b, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x22, 0x50, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x4b, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x50, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x22, 0x4b, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x42, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32,
	0xbb, 0x04, 0x0a, 0x0f, 0x57, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x29, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x77, 0x68, 0x61,
	0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x29, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x24, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70,
	0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x24, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x77, 0x68, 0x61,
	0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x7a, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64,
	0x12, 0x2c, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x79, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_whatsapp_proto_rawDescOnce sync.Once
	file_proto_whatsapp_proto_rawDescData = file_proto_whatsapp_proto_rawDesc
)

func file_proto_whatsapp_proto_rawDescGZIP() []byte {
	file_proto_whatsapp_proto_rawDescOnce.Do(func() {
		file_proto_whatsapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_whatsapp_proto_rawDescData)
	})
	return file_proto_whatsapp_proto_rawDescData
}

var file_proto_whatsapp_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_whatsapp_proto_goTypes = []any{
	(*SendOrderCustomerMessageRequest)(nil),     // 0: whatsapp.SendOrderCustomerMessageRequest
	(*SendOrderCustomerMessageResponse)(nil),    // 1: whatsapp.SendOrderCustomerMessageResponse
	(*SendOrderBusinessMessageRequest)(nil),     // 2: whatsapp.SendOrderBusinessMessageRequest
	(*SendOrderBusinessMessageResponse)(nil),    // 3: whatsapp.SendOrderBusinessMessageResponse
	(*OrderProductRequest)(nil),                 // 4: whatsapp.OrderProductRequest
	(*CreateBusinessPhoneRequest)(nil),          // 5: whatsapp.CreateBusinessPhoneRequest
	(*CreateBusinessPhoneResponse)(nil),         // 6: whatsapp.CreateBusinessPhoneResponse
	(*UpdateBusinessPhoneRequest)(nil),          // 7: whatsapp.UpdateBusinessPhoneRequest
	(*UpdateBusinessPhoneResponse)(nil),         // 8: whatsapp.UpdateBusinessPhoneResponse
	(*GetBusinessPhoneByCompanyIdRequest)(nil),  // 9: whatsapp.GetBusinessPhoneByCompanyIdRequest
	(*GetBusinessPhoneByCompanyIdResponse)(nil), // 10: whatsapp.GetBusinessPhoneByCompanyIdResponse
}
var file_proto_whatsapp_proto_depIdxs = []int32{
	4,  // 0: whatsapp.SendOrderCustomerMessageRequest.products:type_name -> whatsapp.OrderProductRequest
	4,  // 1: whatsapp.SendOrderBusinessMessageRequest.products:type_name -> whatsapp.OrderProductRequest
	0,  // 2: whatsapp.WhatsappService.SendOrderCustomerMessage:input_type -> whatsapp.SendOrderCustomerMessageRequest
	2,  // 3: whatsapp.WhatsappService.SendOrderBusinessMessage:input_type -> whatsapp.SendOrderBusinessMessageRequest
	5,  // 4: whatsapp.WhatsappService.CreateBusinessPhone:input_type -> whatsapp.CreateBusinessPhoneRequest
	7,  // 5: whatsapp.WhatsappService.UpdateBusinessPhone:input_type -> whatsapp.UpdateBusinessPhoneRequest
	9,  // 6: whatsapp.WhatsappService.GetBusinessPhoneByCompanyId:input_type -> whatsapp.GetBusinessPhoneByCompanyIdRequest
	1,  // 7: whatsapp.WhatsappService.SendOrderCustomerMessage:output_type -> whatsapp.SendOrderCustomerMessageResponse
	3,  // 8: whatsapp.WhatsappService.SendOrderBusinessMessage:output_type -> whatsapp.SendOrderBusinessMessageResponse
	6,  // 9: whatsapp.WhatsappService.CreateBusinessPhone:output_type -> whatsapp.CreateBusinessPhoneResponse
	8,  // 10: whatsapp.WhatsappService.UpdateBusinessPhone:output_type -> whatsapp.UpdateBusinessPhoneResponse
	10, // 11: whatsapp.WhatsappService.GetBusinessPhoneByCompanyId:output_type -> whatsapp.GetBusinessPhoneByCompanyIdResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_whatsapp_proto_init() }
func file_proto_whatsapp_proto_init() {
	if File_proto_whatsapp_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_whatsapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_whatsapp_proto_goTypes,
		DependencyIndexes: file_proto_whatsapp_proto_depIdxs,
		MessageInfos:      file_proto_whatsapp_proto_msgTypes,
	}.Build()
	File_proto_whatsapp_proto = out.File
	file_proto_whatsapp_proto_rawDesc = nil
	file_proto_whatsapp_proto_goTypes = nil
	file_proto_whatsapp_proto_depIdxs = nil
}
