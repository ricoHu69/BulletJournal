// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: daemon/grpc/types/join_group_response.proto

package types

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

type ActionEnum int32

const (
	// ACCEPT means 0
	ActionEnum_ACCEPT ActionEnum = 0
	// DECLINE means 1
	ActionEnum_DECLINE ActionEnum = 1
)

// Enum value maps for ActionEnum.
var (
	ActionEnum_name = map[int32]string{
		0: "ACCEPT",
		1: "DECLINE",
	}
	ActionEnum_value = map[string]int32{
		"ACCEPT":  0,
		"DECLINE": 1,
	}
)

func (x ActionEnum) Enum() *ActionEnum {
	p := new(ActionEnum)
	*p = x
	return p
}

func (x ActionEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ActionEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_daemon_grpc_types_join_group_response_proto_enumTypes[0].Descriptor()
}

func (ActionEnum) Type() protoreflect.EnumType {
	return &file_daemon_grpc_types_join_group_response_proto_enumTypes[0]
}

func (x ActionEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ActionEnum.Descriptor instead.
func (ActionEnum) EnumDescriptor() ([]byte, []int) {
	return file_daemon_grpc_types_join_group_response_proto_rawDescGZIP(), []int{0}
}

type JoinGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Action ActionEnum `protobuf:"varint,2,opt,name=action,proto3,enum=types.ActionEnum" json:"action,omitempty"`
}

func (x *JoinGroupResponse) Reset() {
	*x = JoinGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_daemon_grpc_types_join_group_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGroupResponse) ProtoMessage() {}

func (x *JoinGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_daemon_grpc_types_join_group_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGroupResponse.ProtoReflect.Descriptor instead.
func (*JoinGroupResponse) Descriptor() ([]byte, []int) {
	return file_daemon_grpc_types_join_group_response_proto_rawDescGZIP(), []int{0}
}

func (x *JoinGroupResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *JoinGroupResponse) GetAction() ActionEnum {
	if x != nil {
		return x.Action
	}
	return ActionEnum_ACCEPT
}

var File_daemon_grpc_types_join_group_response_proto protoreflect.FileDescriptor

var file_daemon_grpc_types_join_group_response_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x22, 0x50, 0x0a, 0x11, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x25, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x43, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x64, 0x6d, 0x78, 0x2f, 0x42, 0x75, 0x6c, 0x6c, 0x65, 0x74, 0x4a, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_daemon_grpc_types_join_group_response_proto_rawDescOnce sync.Once
	file_daemon_grpc_types_join_group_response_proto_rawDescData = file_daemon_grpc_types_join_group_response_proto_rawDesc
)

func file_daemon_grpc_types_join_group_response_proto_rawDescGZIP() []byte {
	file_daemon_grpc_types_join_group_response_proto_rawDescOnce.Do(func() {
		file_daemon_grpc_types_join_group_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_daemon_grpc_types_join_group_response_proto_rawDescData)
	})
	return file_daemon_grpc_types_join_group_response_proto_rawDescData
}

var file_daemon_grpc_types_join_group_response_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_daemon_grpc_types_join_group_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_daemon_grpc_types_join_group_response_proto_goTypes = []interface{}{
	(ActionEnum)(0),           // 0: types.ActionEnum
	(*JoinGroupResponse)(nil), // 1: types.JoinGroupResponse
}
var file_daemon_grpc_types_join_group_response_proto_depIdxs = []int32{
	0, // 0: types.JoinGroupResponse.action:type_name -> types.ActionEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_daemon_grpc_types_join_group_response_proto_init() }
func file_daemon_grpc_types_join_group_response_proto_init() {
	if File_daemon_grpc_types_join_group_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_daemon_grpc_types_join_group_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGroupResponse); i {
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
			RawDescriptor: file_daemon_grpc_types_join_group_response_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_daemon_grpc_types_join_group_response_proto_goTypes,
		DependencyIndexes: file_daemon_grpc_types_join_group_response_proto_depIdxs,
		EnumInfos:         file_daemon_grpc_types_join_group_response_proto_enumTypes,
		MessageInfos:      file_daemon_grpc_types_join_group_response_proto_msgTypes,
	}.Build()
	File_daemon_grpc_types_join_group_response_proto = out.File
	file_daemon_grpc_types_join_group_response_proto_rawDesc = nil
	file_daemon_grpc_types_join_group_response_proto_goTypes = nil
	file_daemon_grpc_types_join_group_response_proto_depIdxs = nil
}
