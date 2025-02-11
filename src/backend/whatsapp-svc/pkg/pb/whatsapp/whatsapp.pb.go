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
	ToPhone       string                 `protobuf:"bytes,1,opt,name=toPhone,proto3" json:"toPhone,omitempty"`
	CustomerName  string                 `protobuf:"bytes,2,opt,name=customerName,proto3" json:"customerName,omitempty"`
	ContactNumber string                 `protobuf:"bytes,3,opt,name=contactNumber,proto3" json:"contactNumber,omitempty"`
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

func (x *SendOrderCustomerMessageRequest) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
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
	mi := &file_proto_whatsapp_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProductRequest) ProtoMessage() {}

func (x *OrderProductRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use OrderProductRequest.ProtoReflect.Descriptor instead.
func (*OrderProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_whatsapp_proto_rawDescGZIP(), []int{2}
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

var File_proto_whatsapp_proto protoreflect.FileDescriptor

var file_proto_whatsapp_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70,
	0x22, 0xc0, 0x01, 0x0a, 0x1f, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x68, 0x61,
	0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x22, 0x50, 0x0a, 0x20, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5b, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x32, 0x84, 0x01, 0x0a, 0x0f, 0x57, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x29, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x77,
	0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_whatsapp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_whatsapp_proto_goTypes = []any{
	(*SendOrderCustomerMessageRequest)(nil),  // 0: whatsapp.SendOrderCustomerMessageRequest
	(*SendOrderCustomerMessageResponse)(nil), // 1: whatsapp.SendOrderCustomerMessageResponse
	(*OrderProductRequest)(nil),              // 2: whatsapp.OrderProductRequest
}
var file_proto_whatsapp_proto_depIdxs = []int32{
	2, // 0: whatsapp.SendOrderCustomerMessageRequest.products:type_name -> whatsapp.OrderProductRequest
	0, // 1: whatsapp.WhatsappService.SendOrderCustomerMessage:input_type -> whatsapp.SendOrderCustomerMessageRequest
	1, // 2: whatsapp.WhatsappService.SendOrderCustomerMessage:output_type -> whatsapp.SendOrderCustomerMessageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
			NumMessages:   3,
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
