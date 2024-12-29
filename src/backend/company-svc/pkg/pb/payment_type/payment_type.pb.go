// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.27.1
// source: proto/payment_type.proto

package payment_type

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

type GetPaymentTypeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPaymentTypeRequest) Reset() {
	*x = GetPaymentTypeRequest{}
	mi := &file_proto_payment_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPaymentTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentTypeRequest) ProtoMessage() {}

func (x *GetPaymentTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentTypeRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentTypeRequest) Descriptor() ([]byte, []int) {
	return file_proto_payment_type_proto_rawDescGZIP(), []int{0}
}

func (x *GetPaymentTypeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPaymentTypeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImagePath     *string                `protobuf:"bytes,3,opt,name=imagePath,proto3,oneof" json:"imagePath,omitempty"`
	Status        int64                  `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                 `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPaymentTypeResponse) Reset() {
	*x = GetPaymentTypeResponse{}
	mi := &file_proto_payment_type_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPaymentTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentTypeResponse) ProtoMessage() {}

func (x *GetPaymentTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_type_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentTypeResponse.ProtoReflect.Descriptor instead.
func (*GetPaymentTypeResponse) Descriptor() ([]byte, []int) {
	return file_proto_payment_type_proto_rawDescGZIP(), []int{1}
}

func (x *GetPaymentTypeResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetPaymentTypeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPaymentTypeResponse) GetImagePath() string {
	if x != nil && x.ImagePath != nil {
		return *x.ImagePath
	}
	return ""
}

func (x *GetPaymentTypeResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetPaymentTypeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetPaymentTypesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPaymentTypesRequest) Reset() {
	*x = GetPaymentTypesRequest{}
	mi := &file_proto_payment_type_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPaymentTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentTypesRequest) ProtoMessage() {}

func (x *GetPaymentTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_type_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentTypesRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentTypesRequest) Descriptor() ([]byte, []int) {
	return file_proto_payment_type_proto_rawDescGZIP(), []int{2}
}

func (x *GetPaymentTypesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetPaymentTypesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetPaymentTypesResponse struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	PaymentTypes  []*GetPaymentTypeResponse `protobuf:"bytes,1,rep,name=payment_types,json=paymentTypes,proto3" json:"payment_types,omitempty"`
	Status        int64                     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Error         string                    `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPaymentTypesResponse) Reset() {
	*x = GetPaymentTypesResponse{}
	mi := &file_proto_payment_type_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPaymentTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentTypesResponse) ProtoMessage() {}

func (x *GetPaymentTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_payment_type_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentTypesResponse.ProtoReflect.Descriptor instead.
func (*GetPaymentTypesResponse) Descriptor() ([]byte, []int) {
	return file_proto_payment_type_proto_rawDescGZIP(), []int{3}
}

func (x *GetPaymentTypesResponse) GetPaymentTypes() []*GetPaymentTypeResponse {
	if x != nil {
		return x.PaymentTypes
	}
	return nil
}

func (x *GetPaymentTypesResponse) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetPaymentTypesResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_proto_payment_type_proto protoreflect.FileDescriptor

var file_proto_payment_type_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x9b, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68,
	0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22,
	0x46, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xd1, 0x01, 0x0a,
	0x12, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x5e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_payment_type_proto_rawDescOnce sync.Once
	file_proto_payment_type_proto_rawDescData = file_proto_payment_type_proto_rawDesc
)

func file_proto_payment_type_proto_rawDescGZIP() []byte {
	file_proto_payment_type_proto_rawDescOnce.Do(func() {
		file_proto_payment_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_payment_type_proto_rawDescData)
	})
	return file_proto_payment_type_proto_rawDescData
}

var file_proto_payment_type_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_payment_type_proto_goTypes = []any{
	(*GetPaymentTypeRequest)(nil),   // 0: payment_type.GetPaymentTypeRequest
	(*GetPaymentTypeResponse)(nil),  // 1: payment_type.GetPaymentTypeResponse
	(*GetPaymentTypesRequest)(nil),  // 2: payment_type.GetPaymentTypesRequest
	(*GetPaymentTypesResponse)(nil), // 3: payment_type.GetPaymentTypesResponse
}
var file_proto_payment_type_proto_depIdxs = []int32{
	1, // 0: payment_type.GetPaymentTypesResponse.payment_types:type_name -> payment_type.GetPaymentTypeResponse
	0, // 1: payment_type.PaymentTypeService.GetPaymentType:input_type -> payment_type.GetPaymentTypeRequest
	2, // 2: payment_type.PaymentTypeService.GetPaymentTypes:input_type -> payment_type.GetPaymentTypesRequest
	1, // 3: payment_type.PaymentTypeService.GetPaymentType:output_type -> payment_type.GetPaymentTypeResponse
	3, // 4: payment_type.PaymentTypeService.GetPaymentTypes:output_type -> payment_type.GetPaymentTypesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_payment_type_proto_init() }
func file_proto_payment_type_proto_init() {
	if File_proto_payment_type_proto != nil {
		return
	}
	file_proto_payment_type_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_payment_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_payment_type_proto_goTypes,
		DependencyIndexes: file_proto_payment_type_proto_depIdxs,
		MessageInfos:      file_proto_payment_type_proto_msgTypes,
	}.Build()
	File_proto_payment_type_proto = out.File
	file_proto_payment_type_proto_rawDesc = nil
	file_proto_payment_type_proto_goTypes = nil
	file_proto_payment_type_proto_depIdxs = nil
}
