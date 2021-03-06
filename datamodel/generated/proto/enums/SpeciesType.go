// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: enums/SpeciesType.proto

package enums

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

// type:enum name:species_type
type SpeciesType int32

const (
	SpeciesType_AMPHIBIAN  SpeciesType = 0
	SpeciesType_ARTIFICIAL SpeciesType = 1
	SpeciesType_GASTROPOD  SpeciesType = 2
	SpeciesType_INSECTOID  SpeciesType = 3
	SpeciesType_MAMMAL     SpeciesType = 4
	SpeciesType_MAMMALS    SpeciesType = 5
	SpeciesType_OTHER      SpeciesType = 6
	SpeciesType_REPTILE    SpeciesType = 7
	SpeciesType_REPTILIAN  SpeciesType = 8
	SpeciesType_UNKOWN     SpeciesType = 9
)

// Enum value maps for SpeciesType.
var (
	SpeciesType_name = map[int32]string{
		0: "AMPHIBIAN",
		1: "ARTIFICIAL",
		2: "GASTROPOD",
		3: "INSECTOID",
		4: "MAMMAL",
		5: "MAMMALS",
		6: "OTHER",
		7: "REPTILE",
		8: "REPTILIAN",
		9: "UNKOWN",
	}
	SpeciesType_value = map[string]int32{
		"AMPHIBIAN":  0,
		"ARTIFICIAL": 1,
		"GASTROPOD":  2,
		"INSECTOID":  3,
		"MAMMAL":     4,
		"MAMMALS":    5,
		"OTHER":      6,
		"REPTILE":    7,
		"REPTILIAN":  8,
		"UNKOWN":     9,
	}
)

func (x SpeciesType) Enum() *SpeciesType {
	p := new(SpeciesType)
	*p = x
	return p
}

func (x SpeciesType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpeciesType) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_SpeciesType_proto_enumTypes[0].Descriptor()
}

func (SpeciesType) Type() protoreflect.EnumType {
	return &file_enums_SpeciesType_proto_enumTypes[0]
}

func (x SpeciesType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpeciesType.Descriptor instead.
func (SpeciesType) EnumDescriptor() ([]byte, []int) {
	return file_enums_SpeciesType_proto_rawDescGZIP(), []int{0}
}

var File_enums_SpeciesType_proto protoreflect.FileDescriptor

var file_enums_SpeciesType_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x96, 0x01, 0x0a, 0x0b, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x4d, 0x50,
	0x48, 0x49, 0x42, 0x49, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x52, 0x54, 0x49,
	0x46, 0x49, 0x43, 0x49, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x41, 0x53, 0x54,
	0x52, 0x4f, 0x50, 0x4f, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x53, 0x45, 0x43,
	0x54, 0x4f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x4d, 0x4d, 0x41, 0x4c,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x41, 0x4d, 0x4d, 0x41, 0x4c, 0x53, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45,
	0x50, 0x54, 0x49, 0x4c, 0x45, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x50, 0x54, 0x49,
	0x4c, 0x49, 0x41, 0x4e, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4f, 0x57, 0x4e,
	0x10, 0x09, 0x42, 0x2c, 0x5a, 0x2a, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_SpeciesType_proto_rawDescOnce sync.Once
	file_enums_SpeciesType_proto_rawDescData = file_enums_SpeciesType_proto_rawDesc
)

func file_enums_SpeciesType_proto_rawDescGZIP() []byte {
	file_enums_SpeciesType_proto_rawDescOnce.Do(func() {
		file_enums_SpeciesType_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_SpeciesType_proto_rawDescData)
	})
	return file_enums_SpeciesType_proto_rawDescData
}

var file_enums_SpeciesType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_SpeciesType_proto_goTypes = []interface{}{
	(SpeciesType)(0), // 0: SpeciesType
}
var file_enums_SpeciesType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_SpeciesType_proto_init() }
func file_enums_SpeciesType_proto_init() {
	if File_enums_SpeciesType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_SpeciesType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_SpeciesType_proto_goTypes,
		DependencyIndexes: file_enums_SpeciesType_proto_depIdxs,
		EnumInfos:         file_enums_SpeciesType_proto_enumTypes,
	}.Build()
	File_enums_SpeciesType_proto = out.File
	file_enums_SpeciesType_proto_rawDesc = nil
	file_enums_SpeciesType_proto_goTypes = nil
	file_enums_SpeciesType_proto_depIdxs = nil
}
