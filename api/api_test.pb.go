// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.2
// source: api/api_test.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_api_api_test_proto protoreflect.FileDescriptor

var file_api_api_test_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd6, 0x02, 0x0a, 0x0b, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x55, 0x0a, 0x0d, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x69, 0x61, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x69, 0x61, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x69, 0x61, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x69, 0x61, 0x6b, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x12, 0x55, 0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x52, 0x69, 0x61, 0x6b, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x69, 0x61, 0x6b, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x69, 0x61, 0x6b,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x72,
	0x69, 0x61, 0x6b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x59, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x69, 0x61, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x69, 0x61, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x69, 0x61, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x69, 0x61, 0x6b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x65, 0x74, 0x65, 0x72, 0x6c, 0x65, 0x61, 0x72,
	0x6e, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_api_test_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: api.Empty
	(*RiakStoreReq)(nil),   // 1: api.RiakStoreReq
	(*RiakFetchReq)(nil),   // 2: api.RiakFetchReq
	(*RiakDeleteReq)(nil),  // 3: api.RiakDeleteReq
	(*HelloResp)(nil),      // 4: api.HelloResp
	(*RiakStoreResp)(nil),  // 5: api.RiakStoreResp
	(*RiakFetchResp)(nil),  // 6: api.RiakFetchResp
	(*RiakDeleteResp)(nil), // 7: api.RiakDeleteResp
}
var file_api_api_test_proto_depIdxs = []int32{
	0, // 0: api.TestService.Hello:input_type -> api.Empty
	1, // 1: api.TestService.TestRiakStore:input_type -> api.RiakStoreReq
	2, // 2: api.TestService.TestRiakFetch:input_type -> api.RiakFetchReq
	3, // 3: api.TestService.TestRiakDelete:input_type -> api.RiakDeleteReq
	4, // 4: api.TestService.Hello:output_type -> api.HelloResp
	5, // 5: api.TestService.TestRiakStore:output_type -> api.RiakStoreResp
	6, // 6: api.TestService.TestRiakFetch:output_type -> api.RiakFetchResp
	7, // 7: api.TestService.TestRiakDelete:output_type -> api.RiakDeleteResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_api_test_proto_init() }
func file_api_api_test_proto_init() {
	if File_api_api_test_proto != nil {
		return
	}
	file_api_default_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_api_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_test_proto_goTypes,
		DependencyIndexes: file_api_api_test_proto_depIdxs,
	}.Build()
	File_api_api_test_proto = out.File
	file_api_api_test_proto_rawDesc = nil
	file_api_api_test_proto_goTypes = nil
	file_api_api_test_proto_depIdxs = nil
}
