// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: encoding/encoding.proto

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

type TestModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hobby     []string          `protobuf:"bytes,3,rep,name=hobby,proto3" json:"hobby,omitempty"`
	SnakeCase map[string]string `protobuf:"bytes,4,rep,name=snake_case,json=snakeCase,proto3" json:"snake_case,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestModel) Reset() {
	*x = TestModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encoding_encoding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestModel) ProtoMessage() {}

func (x *TestModel) ProtoReflect() protoreflect.Message {
	mi := &file_encoding_encoding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestModel.ProtoReflect.Descriptor instead.
func (*TestModel) Descriptor() ([]byte, []int) {
	return file_encoding_encoding_proto_rawDescGZIP(), []int{0}
}

func (x *TestModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestModel) GetHobby() []string {
	if x != nil {
		return x.Hobby
	}
	return nil
}

func (x *TestModel) GetSnakeCase() map[string]string {
	if x != nil {
		return x.SnakeCase
	}
	return nil
}

var File_encoding_encoding_proto protoreflect.FileDescriptor

var file_encoding_encoding_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x22,
	0xc4, 0x01, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x68, 0x6f, 0x62, 0x62, 0x79, 0x12, 0x3e, 0x0a, 0x0a, 0x73, 0x6e, 0x61, 0x6b,
	0x65, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53,
	0x6e, 0x61, 0x6b, 0x65, 0x43, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x73,
	0x6e, 0x61, 0x6b, 0x65, 0x43, 0x61, 0x73, 0x65, 0x1a, 0x3c, 0x0a, 0x0e, 0x53, 0x6e, 0x61, 0x6b,
	0x65, 0x43, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x79, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_encoding_encoding_proto_rawDescOnce sync.Once
	file_encoding_encoding_proto_rawDescData = file_encoding_encoding_proto_rawDesc
)

func file_encoding_encoding_proto_rawDescGZIP() []byte {
	file_encoding_encoding_proto_rawDescOnce.Do(func() {
		file_encoding_encoding_proto_rawDescData = protoimpl.X.CompressGZIP(file_encoding_encoding_proto_rawDescData)
	})
	return file_encoding_encoding_proto_rawDescData
}

var file_encoding_encoding_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_encoding_encoding_proto_goTypes = []interface{}{
	(*TestModel)(nil), // 0: test.test_model
	nil,               // 1: test.test_model.SnakeCaseEntry
}
var file_encoding_encoding_proto_depIdxs = []int32{
	1, // 0: test.test_model.snake_case:type_name -> test.test_model.SnakeCaseEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_encoding_encoding_proto_init() }
func file_encoding_encoding_proto_init() {
	if File_encoding_encoding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encoding_encoding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestModel); i {
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
			RawDescriptor: file_encoding_encoding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_encoding_encoding_proto_goTypes,
		DependencyIndexes: file_encoding_encoding_proto_depIdxs,
		MessageInfos:      file_encoding_encoding_proto_msgTypes,
	}.Build()
	File_encoding_encoding_proto = out.File
	file_encoding_encoding_proto_rawDesc = nil
	file_encoding_encoding_proto_goTypes = nil
	file_encoding_encoding_proto_depIdxs = nil
}
