// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: notice/api.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_notice_api_proto protoreflect.FileDescriptor

var file_notice_api_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x1a, 0x1f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x60, 0x0a, 0x10, 0x54, 0x68, 0x69, 0x72, 0x64,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x53,
	0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x2e, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x19,
	0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x6e, 0x6f, 0x74,
	0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x22, 0x5a, 0x20, 0x70, 0x65, 0x74,
	0x65, 0x72, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2d, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_notice_api_proto_goTypes = []interface{}{
	(*SendMessageReq)(nil), // 0: third_party_notice.SendMessageReq
	(*Empty)(nil),          // 1: third_party_notice.Empty
}
var file_notice_api_proto_depIdxs = []int32{
	0, // 0: third_party_notice.ThirdPartyNotice.SendMessage:input_type -> third_party_notice.SendMessageReq
	1, // 1: third_party_notice.ThirdPartyNotice.SendMessage:output_type -> third_party_notice.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notice_api_proto_init() }
func file_notice_api_proto_init() {
	if File_notice_api_proto != nil {
		return
	}
	file_notice_third_party_notice_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_notice_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notice_api_proto_goTypes,
		DependencyIndexes: file_notice_api_proto_depIdxs,
	}.Build()
	File_notice_api_proto = out.File
	file_notice_api_proto_rawDesc = nil
	file_notice_api_proto_goTypes = nil
	file_notice_api_proto_depIdxs = nil
}
