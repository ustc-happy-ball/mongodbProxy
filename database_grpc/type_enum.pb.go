// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: type_enum.proto

package databaseGrpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RESPONSE_STATUS int32

const (
	RESPONSE_STATUS_SUCCESS RESPONSE_STATUS = 0
	RESPONSE_STATUS_FAIL    RESPONSE_STATUS = 1
)

// Enum value maps for RESPONSE_STATUS.
var (
	RESPONSE_STATUS_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
	}
	RESPONSE_STATUS_value = map[string]int32{
		"SUCCESS": 0,
		"FAIL":    1,
	}
)

func (x RESPONSE_STATUS) Enum() *RESPONSE_STATUS {
	p := new(RESPONSE_STATUS)
	*p = x
	return p
}

func (x RESPONSE_STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RESPONSE_STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_type_enum_proto_enumTypes[0].Descriptor()
}

func (RESPONSE_STATUS) Type() protoreflect.EnumType {
	return &file_type_enum_proto_enumTypes[0]
}

func (x RESPONSE_STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RESPONSE_STATUS.Descriptor instead.
func (RESPONSE_STATUS) EnumDescriptor() ([]byte, []int) {
	return file_type_enum_proto_rawDescGZIP(), []int{0}
}

type MESSAGE_TYPE int32

const (
	MESSAGE_TYPE_REQUEST  MESSAGE_TYPE = 0
	MESSAGE_TYPE_RESPONSE MESSAGE_TYPE = 1
)

// Enum value maps for MESSAGE_TYPE.
var (
	MESSAGE_TYPE_name = map[int32]string{
		0: "REQUEST",
		1: "RESPONSE",
	}
	MESSAGE_TYPE_value = map[string]int32{
		"REQUEST":  0,
		"RESPONSE": 1,
	}
)

func (x MESSAGE_TYPE) Enum() *MESSAGE_TYPE {
	p := new(MESSAGE_TYPE)
	*p = x
	return p
}

func (x MESSAGE_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MESSAGE_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_type_enum_proto_enumTypes[1].Descriptor()
}

func (MESSAGE_TYPE) Type() protoreflect.EnumType {
	return &file_type_enum_proto_enumTypes[1]
}

func (x MESSAGE_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MESSAGE_TYPE.Descriptor instead.
func (MESSAGE_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_type_enum_proto_rawDescGZIP(), []int{1}
}

type ITEM int32

const (
	ITEM_PLAYER ITEM = 0
)

// Enum value maps for ITEM.
var (
	ITEM_name = map[int32]string{
		0: "PLAYER",
	}
	ITEM_value = map[string]int32{
		"PLAYER": 0,
	}
)

func (x ITEM) Enum() *ITEM {
	p := new(ITEM)
	*p = x
	return p
}

func (x ITEM) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ITEM) Descriptor() protoreflect.EnumDescriptor {
	return file_type_enum_proto_enumTypes[2].Descriptor()
}

func (ITEM) Type() protoreflect.EnumType {
	return &file_type_enum_proto_enumTypes[2]
}

func (x ITEM) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ITEM.Descriptor instead.
func (ITEM) EnumDescriptor() ([]byte, []int) {
	return file_type_enum_proto_rawDescGZIP(), []int{2}
}

type MESSAGE_CODE int32

const (
	MESSAGE_CODE_FIND_REQUEST    MESSAGE_CODE = 0
	MESSAGE_CODE_ADD_REQUEST     MESSAGE_CODE = 1
	MESSAGE_CODE_DELETE_REQUEST  MESSAGE_CODE = 2
	MESSAGE_CODE_UPDATE_REQUEST  MESSAGE_CODE = 3
	MESSAGE_CODE_FIND_RESPONSE   MESSAGE_CODE = 4
	MESSAGE_CODE_ADD_RESPONSE    MESSAGE_CODE = 5
	MESSAGE_CODE_DELETE_RESPONSE MESSAGE_CODE = 6
	MESSAGE_CODE_UPDATE_RESPONSE MESSAGE_CODE = 7
)

// Enum value maps for MESSAGE_CODE.
var (
	MESSAGE_CODE_name = map[int32]string{
		0: "FIND_REQUEST",
		1: "ADD_REQUEST",
		2: "DELETE_REQUEST",
		3: "UPDATE_REQUEST",
		4: "FIND_RESPONSE",
		5: "ADD_RESPONSE",
		6: "DELETE_RESPONSE",
		7: "UPDATE_RESPONSE",
	}
	MESSAGE_CODE_value = map[string]int32{
		"FIND_REQUEST":    0,
		"ADD_REQUEST":     1,
		"DELETE_REQUEST":  2,
		"UPDATE_REQUEST":  3,
		"FIND_RESPONSE":   4,
		"ADD_RESPONSE":    5,
		"DELETE_RESPONSE": 6,
		"UPDATE_RESPONSE": 7,
	}
)

func (x MESSAGE_CODE) Enum() *MESSAGE_CODE {
	p := new(MESSAGE_CODE)
	*p = x
	return p
}

func (x MESSAGE_CODE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MESSAGE_CODE) Descriptor() protoreflect.EnumDescriptor {
	return file_type_enum_proto_enumTypes[3].Descriptor()
}

func (MESSAGE_CODE) Type() protoreflect.EnumType {
	return &file_type_enum_proto_enumTypes[3]
}

func (x MESSAGE_CODE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MESSAGE_CODE.Descriptor instead.
func (MESSAGE_CODE) EnumDescriptor() ([]byte, []int) {
	return file_type_enum_proto_rawDescGZIP(), []int{3}
}

var File_type_enum_proto protoreflect.FileDescriptor

var file_type_enum_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2a,
	0x28, 0x0a, 0x0f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x2a, 0x29, 0x0a, 0x0c, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e,
	0x53, 0x45, 0x10, 0x01, 0x2a, 0x12, 0x0a, 0x04, 0x49, 0x54, 0x45, 0x4d, 0x12, 0x0a, 0x0a, 0x06,
	0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x10, 0x00, 0x2a, 0xa8, 0x01, 0x0a, 0x0c, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x49, 0x4e,
	0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41,
	0x44, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x02,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45,
	0x53, 0x54, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x49, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x53,
	0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x44, 0x44, 0x5f, 0x52,
	0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x06, 0x12, 0x13,
	0x0a, 0x0f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53,
	0x45, 0x10, 0x07, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_type_enum_proto_rawDescOnce sync.Once
	file_type_enum_proto_rawDescData = file_type_enum_proto_rawDesc
)

func file_type_enum_proto_rawDescGZIP() []byte {
	file_type_enum_proto_rawDescOnce.Do(func() {
		file_type_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_type_enum_proto_rawDescData)
	})
	return file_type_enum_proto_rawDescData
}

var file_type_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_type_enum_proto_goTypes = []interface{}{
	(RESPONSE_STATUS)(0), // 0: databaseGrpc.RESPONSE_STATUS
	(MESSAGE_TYPE)(0),    // 1: databaseGrpc.MESSAGE_TYPE
	(ITEM)(0),            // 2: databaseGrpc.ITEM
	(MESSAGE_CODE)(0),    // 3: databaseGrpc.MESSAGE_CODE
}
var file_type_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_type_enum_proto_init() }
func file_type_enum_proto_init() {
	if File_type_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_type_enum_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_type_enum_proto_goTypes,
		DependencyIndexes: file_type_enum_proto_depIdxs,
		EnumInfos:         file_type_enum_proto_enumTypes,
	}.Build()
	File_type_enum_proto = out.File
	file_type_enum_proto_rawDesc = nil
	file_type_enum_proto_goTypes = nil
	file_type_enum_proto_depIdxs = nil
}
