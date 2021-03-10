// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: entities/Character.proto

package entities

import (
	enums "assessment/datamodel/generated/proto/enums"
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

// table:characters
type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                             // column:id pk generated returned cached
	Name            string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                          // column:name unique index cached
	Height          int32             `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`                                     // column:height
	Mass            int32             `protobuf:"varint,4,opt,name=mass,proto3" json:"mass,omitempty"`                                         // column:mass
	HairColor       enums.Color       `protobuf:"varint,5,opt,name=hairColor,proto3,enum=Color" json:"hairColor,omitempty"`                    // column:hair_color type:enum
	SkinColor       enums.Color       `protobuf:"varint,6,opt,name=skinColor,proto3,enum=Color" json:"skinColor,omitempty"`                    // column:skin_color type:enum
	EyeColor        enums.Color       `protobuf:"varint,7,opt,name=eyeColor,proto3,enum=Color" json:"eyeColor,omitempty"`                      // column:eye_color type:enum
	BirthYear       float32           `protobuf:"fixed32,8,opt,name=birthYear,proto3" json:"birthYear,omitempty"`                              // column:birth_year
	Gender          enums.Gender      `protobuf:"varint,9,opt,name=gender,proto3,enum=Gender" json:"gender,omitempty"`                         // column:gender type:enum
	Homeworld       *Planet           `protobuf:"bytes,10,opt,name=homeworld,proto3" json:"homeworld,omitempty"`                               // column:homeworld entity table:planets nullable
	Species         *Species          `protobuf:"bytes,11,opt,name=species,proto3" json:"species,omitempty"`                                   // column:species entity table:species nullable
	MainAffiliation enums.Affiliation `protobuf:"varint,12,opt,name=mainAffiliation,proto3,enum=Affiliation" json:"mainAffiliation,omitempty"` // column:main_affiliation type:enum
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_Character_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_entities_Character_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_entities_Character_proto_rawDescGZIP(), []int{0}
}

func (x *Character) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Character) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Character) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Character) GetMass() int32 {
	if x != nil {
		return x.Mass
	}
	return 0
}

func (x *Character) GetHairColor() enums.Color {
	if x != nil {
		return x.HairColor
	}
	return enums.Color_BLACK
}

func (x *Character) GetSkinColor() enums.Color {
	if x != nil {
		return x.SkinColor
	}
	return enums.Color_BLACK
}

func (x *Character) GetEyeColor() enums.Color {
	if x != nil {
		return x.EyeColor
	}
	return enums.Color_BLACK
}

func (x *Character) GetBirthYear() float32 {
	if x != nil {
		return x.BirthYear
	}
	return 0
}

func (x *Character) GetGender() enums.Gender {
	if x != nil {
		return x.Gender
	}
	return enums.Gender_MALE
}

func (x *Character) GetHomeworld() *Planet {
	if x != nil {
		return x.Homeworld
	}
	return nil
}

func (x *Character) GetSpecies() *Species {
	if x != nil {
		return x.Species
	}
	return nil
}

func (x *Character) GetMainAffiliation() enums.Affiliation {
	if x != nil {
		return x.MainAffiliation
	}
	return enums.Affiliation_JEDI_ORDER
}

type Characters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Character `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Characters) Reset() {
	*x = Characters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entities_Character_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Characters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Characters) ProtoMessage() {}

func (x *Characters) ProtoReflect() protoreflect.Message {
	mi := &file_entities_Character_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Characters.ProtoReflect.Descriptor instead.
func (*Characters) Descriptor() ([]byte, []int) {
	return file_entities_Character_proto_rawDescGZIP(), []int{1}
}

func (x *Characters) GetValues() []*Character {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_entities_Character_proto protoreflect.FileDescriptor

var file_entities_Character_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2f, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x03, 0x0a, 0x09, 0x43, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x6d, 0x61, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x09, 0x68, 0x61, 0x69, 0x72, 0x43,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x06, 0x2e, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x52, 0x09, 0x68, 0x61, 0x69, 0x72, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x24, 0x0a,
	0x09, 0x73, 0x6b, 0x69, 0x6e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x06, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x09, 0x73, 0x6b, 0x69, 0x6e, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x08, 0x65, 0x79, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x06, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x08, 0x65,
	0x79, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x59, 0x65, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x59, 0x65, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x52, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x22, 0x0a,
	0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x36, 0x0a, 0x0f, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x41, 0x66, 0x66,
	0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x66,
	0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x0a, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entities_Character_proto_rawDescOnce sync.Once
	file_entities_Character_proto_rawDescData = file_entities_Character_proto_rawDesc
)

func file_entities_Character_proto_rawDescGZIP() []byte {
	file_entities_Character_proto_rawDescOnce.Do(func() {
		file_entities_Character_proto_rawDescData = protoimpl.X.CompressGZIP(file_entities_Character_proto_rawDescData)
	})
	return file_entities_Character_proto_rawDescData
}

var file_entities_Character_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_entities_Character_proto_goTypes = []interface{}{
	(*Character)(nil),      // 0: Character
	(*Characters)(nil),     // 1: Characters
	(enums.Color)(0),       // 2: Color
	(enums.Gender)(0),      // 3: Gender
	(*Planet)(nil),         // 4: Planet
	(*Species)(nil),        // 5: Species
	(enums.Affiliation)(0), // 6: Affiliation
}
var file_entities_Character_proto_depIdxs = []int32{
	2, // 0: Character.hairColor:type_name -> Color
	2, // 1: Character.skinColor:type_name -> Color
	2, // 2: Character.eyeColor:type_name -> Color
	3, // 3: Character.gender:type_name -> Gender
	4, // 4: Character.homeworld:type_name -> Planet
	5, // 5: Character.species:type_name -> Species
	6, // 6: Character.mainAffiliation:type_name -> Affiliation
	0, // 7: Characters.values:type_name -> Character
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_entities_Character_proto_init() }
func file_entities_Character_proto_init() {
	if File_entities_Character_proto != nil {
		return
	}
	file_entities_Planet_proto_init()
	file_entities_Species_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entities_Character_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_entities_Character_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Characters); i {
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
			RawDescriptor: file_entities_Character_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entities_Character_proto_goTypes,
		DependencyIndexes: file_entities_Character_proto_depIdxs,
		MessageInfos:      file_entities_Character_proto_msgTypes,
	}.Build()
	File_entities_Character_proto = out.File
	file_entities_Character_proto_rawDesc = nil
	file_entities_Character_proto_goTypes = nil
	file_entities_Character_proto_depIdxs = nil
}