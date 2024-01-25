// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: idl/hello.proto

package hello

import (
	context "context"
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

type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_idl_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RespBody string `protobuf:"bytes,1,opt,name=RespBody,proto3" json:"RespBody,omitempty"`
}

func (x *HelloResp) Reset() {
	*x = HelloResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResp) ProtoMessage() {}

func (x *HelloResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResp.ProtoReflect.Descriptor instead.
func (*HelloResp) Descriptor() ([]byte, []int) {
	return file_idl_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResp) GetRespBody() string {
	if x != nil {
		return x.RespBody
	}
	return ""
}

var File_idl_hello_proto protoreflect.FileDescriptor

var file_idl_hello_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x64, 0x6c, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x1e, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x09, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x42, 0x6f, 0x64,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x42, 0x6f, 0x64,
	0x79, 0x32, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2a, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0f, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x42, 0x2f, 0x5a,
	0x2d, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f,
	0x6b, 0x69, 0x74, 0x65, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6b,
	0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_hello_proto_rawDescOnce sync.Once
	file_idl_hello_proto_rawDescData = file_idl_hello_proto_rawDesc
)

func file_idl_hello_proto_rawDescGZIP() []byte {
	file_idl_hello_proto_rawDescOnce.Do(func() {
		file_idl_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_hello_proto_rawDescData)
	})
	return file_idl_hello_proto_rawDescData
}

var file_idl_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_idl_hello_proto_goTypes = []interface{}{
	(*HelloReq)(nil),  // 0: hello.HelloReq
	(*HelloResp)(nil), // 1: hello.HelloResp
}
var file_idl_hello_proto_depIdxs = []int32{
	0, // 0: hello.HelloService.Hello:input_type -> hello.HelloReq
	1, // 1: hello.HelloService.Hello:output_type -> hello.HelloResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_idl_hello_proto_init() }
func file_idl_hello_proto_init() {
	if File_idl_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq); i {
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
		file_idl_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResp); i {
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
			RawDescriptor: file_idl_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_hello_proto_goTypes,
		DependencyIndexes: file_idl_hello_proto_depIdxs,
		MessageInfos:      file_idl_hello_proto_msgTypes,
	}.Build()
	File_idl_hello_proto = out.File
	file_idl_hello_proto_rawDesc = nil
	file_idl_hello_proto_goTypes = nil
	file_idl_hello_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.7.3. DO NOT EDIT.

type HelloService interface {
	Hello(ctx context.Context, req *HelloReq) (res *HelloResp, err error)
}