// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.6.1
// source: data/v1/data.proto

package data_v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ExportDBV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExportDBV1Request) Reset() {
	*x = ExportDBV1Request{}
	mi := &file_data_v1_data_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportDBV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportDBV1Request) ProtoMessage() {}

func (x *ExportDBV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_data_v1_data_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportDBV1Request.ProtoReflect.Descriptor instead.
func (*ExportDBV1Request) Descriptor() ([]byte, []int) {
	return file_data_v1_data_proto_rawDescGZIP(), []int{0}
}

type ExportDBV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Db []byte `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
}

func (x *ExportDBV1Response) Reset() {
	*x = ExportDBV1Response{}
	mi := &file_data_v1_data_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportDBV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportDBV1Response) ProtoMessage() {}

func (x *ExportDBV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_data_v1_data_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportDBV1Response.ProtoReflect.Descriptor instead.
func (*ExportDBV1Response) Descriptor() ([]byte, []int) {
	return file_data_v1_data_proto_rawDescGZIP(), []int{1}
}

func (x *ExportDBV1Response) GetDb() []byte {
	if x != nil {
		return x.Db
	}
	return nil
}

type ImportDBV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Db []byte `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
}

func (x *ImportDBV1Request) Reset() {
	*x = ImportDBV1Request{}
	mi := &file_data_v1_data_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportDBV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportDBV1Request) ProtoMessage() {}

func (x *ImportDBV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_data_v1_data_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportDBV1Request.ProtoReflect.Descriptor instead.
func (*ImportDBV1Request) Descriptor() ([]byte, []int) {
	return file_data_v1_data_proto_rawDescGZIP(), []int{2}
}

func (x *ImportDBV1Request) GetDb() []byte {
	if x != nil {
		return x.Db
	}
	return nil
}

type ImportDBV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ImportDBV1Response) Reset() {
	*x = ImportDBV1Response{}
	mi := &file_data_v1_data_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportDBV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportDBV1Response) ProtoMessage() {}

func (x *ImportDBV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_data_v1_data_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportDBV1Response.ProtoReflect.Descriptor instead.
func (*ImportDBV1Response) Descriptor() ([]byte, []int) {
	return file_data_v1_data_proto_rawDescGZIP(), []int{3}
}

var File_data_v1_data_proto protoreflect.FileDescriptor

var file_data_v1_data_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x42, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x24, 0x0a, 0x12, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x42, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x02, 0x64, 0x62, 0x22, 0x23, 0x0a, 0x11, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x44, 0x42, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x64,
	0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x64, 0x62, 0x22, 0x14, 0x0a, 0x12, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x42, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xc0, 0x01, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x41, 0x50, 0x49, 0x12, 0x58, 0x0a,
	0x0a, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x42, 0x56, 0x31, 0x12, 0x1a, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x42, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x42, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5b, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x44, 0x42, 0x56, 0x31, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x42, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x44, 0x42, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x11, 0x5a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x3b,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_v1_data_proto_rawDescOnce sync.Once
	file_data_v1_data_proto_rawDescData = file_data_v1_data_proto_rawDesc
)

func file_data_v1_data_proto_rawDescGZIP() []byte {
	file_data_v1_data_proto_rawDescOnce.Do(func() {
		file_data_v1_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_v1_data_proto_rawDescData)
	})
	return file_data_v1_data_proto_rawDescData
}

var file_data_v1_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_data_v1_data_proto_goTypes = []any{
	(*ExportDBV1Request)(nil),  // 0: data.v1.ExportDBV1Request
	(*ExportDBV1Response)(nil), // 1: data.v1.ExportDBV1Response
	(*ImportDBV1Request)(nil),  // 2: data.v1.ImportDBV1Request
	(*ImportDBV1Response)(nil), // 3: data.v1.ImportDBV1Response
}
var file_data_v1_data_proto_depIdxs = []int32{
	0, // 0: data.v1.DataAPI.ExportDBV1:input_type -> data.v1.ExportDBV1Request
	2, // 1: data.v1.DataAPI.ImportDBV1:input_type -> data.v1.ImportDBV1Request
	1, // 2: data.v1.DataAPI.ExportDBV1:output_type -> data.v1.ExportDBV1Response
	3, // 3: data.v1.DataAPI.ImportDBV1:output_type -> data.v1.ImportDBV1Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_data_v1_data_proto_init() }
func file_data_v1_data_proto_init() {
	if File_data_v1_data_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_data_v1_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_data_v1_data_proto_goTypes,
		DependencyIndexes: file_data_v1_data_proto_depIdxs,
		MessageInfos:      file_data_v1_data_proto_msgTypes,
	}.Build()
	File_data_v1_data_proto = out.File
	file_data_v1_data_proto_rawDesc = nil
	file_data_v1_data_proto_goTypes = nil
	file_data_v1_data_proto_depIdxs = nil
}
