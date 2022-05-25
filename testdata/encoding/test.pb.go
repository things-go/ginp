// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: encoding/test.proto

package encoding

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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sub  *Sub   `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encoding_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encoding_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_encoding_test_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetSub() *Sub {
	if x != nil {
		return x.Sub
	}
	return nil
}

type Sub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,json=naming,proto3" json:"name,omitempty"`
}

func (x *Sub) Reset() {
	*x = Sub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encoding_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sub) ProtoMessage() {}

func (x *Sub) ProtoReflect() protoreflect.Message {
	mi := &file_encoding_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sub.ProtoReflect.Descriptor instead.
func (*Sub) Descriptor() ([]byte, []int) {
	return file_encoding_test_proto_rawDescGZIP(), []int{1}
}

func (x *Sub) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_encoding_test_proto protoreflect.FileDescriptor

var file_encoding_test_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x42,
	0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x03, 0x73,
	0x75, 0x62, 0x22, 0x1b, 0x0a, 0x03, 0x53, 0x75, 0x62, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x42,
	0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x79,
	0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_encoding_test_proto_rawDescOnce sync.Once
	file_encoding_test_proto_rawDescData = file_encoding_test_proto_rawDesc
)

func file_encoding_test_proto_rawDescGZIP() []byte {
	file_encoding_test_proto_rawDescOnce.Do(func() {
		file_encoding_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_encoding_test_proto_rawDescData)
	})
	return file_encoding_test_proto_rawDescData
}

var file_encoding_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_encoding_test_proto_goTypes = []interface{}{
	(*HelloRequest)(nil), // 0: binding.HelloRequest
	(*Sub)(nil),          // 1: binding.Sub
}
var file_encoding_test_proto_depIdxs = []int32{
	1, // 0: binding.HelloRequest.sub:type_name -> binding.Sub
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_encoding_test_proto_init() }
func file_encoding_test_proto_init() {
	if File_encoding_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encoding_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_encoding_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sub); i {
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
			RawDescriptor: file_encoding_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_encoding_test_proto_goTypes,
		DependencyIndexes: file_encoding_test_proto_depIdxs,
		MessageInfos:      file_encoding_test_proto_msgTypes,
	}.Build()
	File_encoding_test_proto = out.File
	file_encoding_test_proto_rawDesc = nil
	file_encoding_test_proto_goTypes = nil
	file_encoding_test_proto_depIdxs = nil
}
