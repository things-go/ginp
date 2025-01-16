// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.1
// source: non_nested.proto

package enums

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

// NonNestedStatus 状态值
// #[enum]
type NonNestedStatus int32

const (
	// 未定义
	// #[enum(mapping="unspecified")]
	NonNestedStatus_NonNestedStatusAnnote_Unspecified NonNestedStatus = 0
	// 上
	// #[enum(mapping="up")]
	NonNestedStatus_NonNestedStatusAnnote_Up NonNestedStatus = 1
	// 下
	// #[enum(mapping="down")]
	NonNestedStatus_NonNestedStatusAnnote_Down NonNestedStatus = 2
	// 左
	// #[enum(mapping="left")]
	NonNestedStatus_NonNestedStatusAnnote_Left NonNestedStatus = 3
	// 右
	// #[enum(mapping="right")]
	NonNestedStatus_NonNestedStatusAnnote_Right NonNestedStatus = 4
)

// Enum value maps for NonNestedStatus.
var (
	NonNestedStatus_name = map[int32]string{
		0: "NonNestedStatusAnnote_Unspecified",
		1: "NonNestedStatusAnnote_Up",
		2: "NonNestedStatusAnnote_Down",
		3: "NonNestedStatusAnnote_Left",
		4: "NonNestedStatusAnnote_Right",
	}
	NonNestedStatus_value = map[string]int32{
		"NonNestedStatusAnnote_Unspecified": 0,
		"NonNestedStatusAnnote_Up":          1,
		"NonNestedStatusAnnote_Down":        2,
		"NonNestedStatusAnnote_Left":        3,
		"NonNestedStatusAnnote_Right":       4,
	}
)

func (x NonNestedStatus) Enum() *NonNestedStatus {
	p := new(NonNestedStatus)
	*p = x
	return p
}

func (x NonNestedStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NonNestedStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_non_nested_proto_enumTypes[0].Descriptor()
}

func (NonNestedStatus) Type() protoreflect.EnumType {
	return &file_non_nested_proto_enumTypes[0]
}

func (x NonNestedStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NonNestedStatus.Descriptor instead.
func (NonNestedStatus) EnumDescriptor() ([]byte, []int) {
	return file_non_nested_proto_rawDescGZIP(), []int{0}
}

var File_non_nested_proto protoreflect.FileDescriptor

var file_non_nested_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6e, 0x6f, 0x6e, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x5f, 0x73, 0x61, 0x62,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2a, 0xb7, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x6e,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x21,
	0x4e, 0x6f, 0x6e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x4e, 0x6f, 0x6e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x55, 0x70, 0x10,
	0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x6f, 0x6e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x44, 0x6f, 0x77, 0x6e, 0x10,
	0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x4e, 0x6f, 0x6e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x4c, 0x65, 0x66, 0x74, 0x10,
	0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x4e, 0x6f, 0x6e, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x52, 0x69, 0x67, 0x68, 0x74,
	0x10, 0x04, 0x42, 0x50, 0x0a, 0x1a, 0x63, 0x6e, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2d,
	0x67, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x42, 0x05, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2d, 0x67, 0x6f, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x3b, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_non_nested_proto_rawDescOnce sync.Once
	file_non_nested_proto_rawDescData = file_non_nested_proto_rawDesc
)

func file_non_nested_proto_rawDescGZIP() []byte {
	file_non_nested_proto_rawDescOnce.Do(func() {
		file_non_nested_proto_rawDescData = protoimpl.X.CompressGZIP(file_non_nested_proto_rawDescData)
	})
	return file_non_nested_proto_rawDescData
}

var file_non_nested_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_non_nested_proto_goTypes = []any{
	(NonNestedStatus)(0), // 0: protogen_saber.enums.NonNestedStatus
}
var file_non_nested_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_non_nested_proto_init() }
func file_non_nested_proto_init() {
	if File_non_nested_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_non_nested_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_non_nested_proto_goTypes,
		DependencyIndexes: file_non_nested_proto_depIdxs,
		EnumInfos:         file_non_nested_proto_enumTypes,
	}.Build()
	File_non_nested_proto = out.File
	file_non_nested_proto_rawDesc = nil
	file_non_nested_proto_goTypes = nil
	file_non_nested_proto_depIdxs = nil
}
