// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        v5.26.1
// source: requestmgmt/requestmgmt.proto

package project

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

type NewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wid string `protobuf:"bytes,1,opt,name=wid,proto3" json:"wid,omitempty"`
	Rid string `protobuf:"bytes,2,opt,name=rid,proto3" json:"rid,omitempty"`
	Id  int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewRequest) Reset() {
	*x = NewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestmgmt_requestmgmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRequest) ProtoMessage() {}

func (x *NewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_requestmgmt_requestmgmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRequest.ProtoReflect.Descriptor instead.
func (*NewRequest) Descriptor() ([]byte, []int) {
	return file_requestmgmt_requestmgmt_proto_rawDescGZIP(), []int{0}
}

func (x *NewRequest) GetWid() string {
	if x != nil {
		return x.Wid
	}
	return ""
}

func (x *NewRequest) GetRid() string {
	if x != nil {
		return x.Rid
	}
	return ""
}

func (x *NewRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wid string `protobuf:"bytes,1,opt,name=wid,proto3" json:"wid,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_requestmgmt_requestmgmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_requestmgmt_requestmgmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_requestmgmt_requestmgmt_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetWid() string {
	if x != nil {
		return x.Wid
	}
	return ""
}

var File_requestmgmt_requestmgmt_proto protoreflect.FileDescriptor

var file_requestmgmt_requestmgmt_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x22, 0x40, 0x0a, 0x0a,
	0x4e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x69, 0x64, 0x32, 0x55, 0x0a, 0x11, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e,
	0x4e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x68, 0x75, 0x62, 0x68, 0x61, 0x6d, 0x67, 0x6f, 0x79, 0x61, 0x6c, 0x31, 0x34, 0x30,
	0x32, 0x2f, 0x68, 0x70, 0x65, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_requestmgmt_requestmgmt_proto_rawDescOnce sync.Once
	file_requestmgmt_requestmgmt_proto_rawDescData = file_requestmgmt_requestmgmt_proto_rawDesc
)

func file_requestmgmt_requestmgmt_proto_rawDescGZIP() []byte {
	file_requestmgmt_requestmgmt_proto_rawDescOnce.Do(func() {
		file_requestmgmt_requestmgmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_requestmgmt_requestmgmt_proto_rawDescData)
	})
	return file_requestmgmt_requestmgmt_proto_rawDescData
}

var file_requestmgmt_requestmgmt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_requestmgmt_requestmgmt_proto_goTypes = []interface{}{
	(*NewRequest)(nil), // 0: requestmgmt.NewRequest
	(*Request)(nil),    // 1: requestmgmt.Request
}
var file_requestmgmt_requestmgmt_proto_depIdxs = []int32{
	0, // 0: requestmgmt.RequestManagement.CreateRequest:input_type -> requestmgmt.NewRequest
	1, // 1: requestmgmt.RequestManagement.CreateRequest:output_type -> requestmgmt.Request
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_requestmgmt_requestmgmt_proto_init() }
func file_requestmgmt_requestmgmt_proto_init() {
	if File_requestmgmt_requestmgmt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_requestmgmt_requestmgmt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRequest); i {
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
		file_requestmgmt_requestmgmt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			RawDescriptor: file_requestmgmt_requestmgmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_requestmgmt_requestmgmt_proto_goTypes,
		DependencyIndexes: file_requestmgmt_requestmgmt_proto_depIdxs,
		MessageInfos:      file_requestmgmt_requestmgmt_proto_msgTypes,
	}.Build()
	File_requestmgmt_requestmgmt_proto = out.File
	file_requestmgmt_requestmgmt_proto_rawDesc = nil
	file_requestmgmt_requestmgmt_proto_goTypes = nil
	file_requestmgmt_requestmgmt_proto_depIdxs = nil
}
