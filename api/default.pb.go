// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/default.proto

package v1

import (
	common "github.com/peterlearn/spgs_proto/common"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_default_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_default_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_default_proto_rawDescGZIP(), []int{0}
}

type HelloResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string         `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Country common.Country `protobuf:"varint,2,opt,name=Country,proto3,enum=common.Country" json:"Country,omitempty"`
}

func (x *HelloResp) Reset() {
	*x = HelloResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_default_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResp) ProtoMessage() {}

func (x *HelloResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_default_proto_msgTypes[1]
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
	return file_api_default_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResp) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *HelloResp) GetCountry() common.Country {
	if x != nil {
		return x.Country
	}
	return common.Country(0)
}

// riak存储请求
type RiakStoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket string `protobuf:"bytes,1,opt,name=Bucket,proto3" json:"Bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	Data   string `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *RiakStoreReq) Reset() {
	*x = RiakStoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_default_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiakStoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiakStoreReq) ProtoMessage() {}

func (x *RiakStoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_default_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiakStoreReq.ProtoReflect.Descriptor instead.
func (*RiakStoreReq) Descriptor() ([]byte, []int) {
	return file_api_default_proto_rawDescGZIP(), []int{2}
}

func (x *RiakStoreReq) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *RiakStoreReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *RiakStoreReq) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type RiakStoreResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RiakStoreResp) Reset() {
	*x = RiakStoreResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_default_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiakStoreResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiakStoreResp) ProtoMessage() {}

func (x *RiakStoreResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_default_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiakStoreResp.ProtoReflect.Descriptor instead.
func (*RiakStoreResp) Descriptor() ([]byte, []int) {
	return file_api_default_proto_rawDescGZIP(), []int{3}
}

// 获取riak中数据
type RiakFetchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket string   `protobuf:"bytes,1,opt,name=Bucket,proto3" json:"Bucket,omitempty"`
	Key    []string `protobuf:"bytes,2,rep,name=Key,proto3" json:"Key,omitempty"`
}

func (x *RiakFetchReq) Reset() {
	*x = RiakFetchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_default_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiakFetchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiakFetchReq) ProtoMessage() {}

func (x *RiakFetchReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_default_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiakFetchReq.ProtoReflect.Descriptor instead.
func (*RiakFetchReq) Descriptor() ([]byte, []int) {
	return file_api_default_proto_rawDescGZIP(), []int{4}
}

func (x *RiakFetchReq) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *RiakFetchReq) GetKey() []string {
	if x != nil {
		return x.Key
	}
	return nil
}

type KV struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Val string `protobuf:"bytes,2,opt,name=Val,proto3" json:"Val,omitempty"`
}

func (x *KV) Reset() {
	*x = KV{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_default_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KV) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KV) ProtoMessage() {}

func (x *KV) ProtoReflect() protoreflect.Message {
	mi := &file_api_default_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KV.ProtoReflect.Descriptor instead.
func (*KV) Descriptor() ([]byte, []int) {
	return file_api_default_proto_rawDescGZIP(), []int{5}
}

func (x *KV) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KV) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type RiakFetchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*KV `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (x *RiakFetchResp) Reset() {
	*x = RiakFetchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_default_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiakFetchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiakFetchResp) ProtoMessage() {}

func (x *RiakFetchResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_default_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiakFetchResp.ProtoReflect.Descriptor instead.
func (*RiakFetchResp) Descriptor() ([]byte, []int) {
	return file_api_default_proto_rawDescGZIP(), []int{6}
}

func (x *RiakFetchResp) GetData() []*KV {
	if x != nil {
		return x.Data
	}
	return nil
}

// 删除riak中数据
type RiakDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket string `protobuf:"bytes,1,opt,name=Bucket,proto3" json:"Bucket,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *RiakDeleteReq) Reset() {
	*x = RiakDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_default_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiakDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiakDeleteReq) ProtoMessage() {}

func (x *RiakDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_default_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiakDeleteReq.ProtoReflect.Descriptor instead.
func (*RiakDeleteReq) Descriptor() ([]byte, []int) {
	return file_api_default_proto_rawDescGZIP(), []int{7}
}

func (x *RiakDeleteReq) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *RiakDeleteReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type RiakDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RiakDeleteResp) Reset() {
	*x = RiakDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_default_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RiakDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RiakDeleteResp) ProtoMessage() {}

func (x *RiakDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_default_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RiakDeleteResp.ProtoReflect.Descriptor instead.
func (*RiakDeleteResp) Descriptor() ([]byte, []int) {
	return file_api_default_proto_rawDescGZIP(), []int{8}
}

var File_api_default_proto protoreflect.FileDescriptor

var file_api_default_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x50, 0x0a, 0x09, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x4c, 0x0a, 0x0c, 0x52, 0x69, 0x61, 0x6b, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x69, 0x61, 0x6b, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x38, 0x0a, 0x0c, 0x52, 0x69, 0x61, 0x6b, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79,
	0x22, 0x28, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x56, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x56, 0x61, 0x6c, 0x22, 0x2c, 0x0a, 0x0d, 0x52, 0x69,
	0x61, 0x6b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4b, 0x56, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x39, 0x0a, 0x0d, 0x52, 0x69, 0x61, 0x6b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4b, 0x65, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x69, 0x61, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x65, 0x74, 0x65, 0x72, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_default_proto_rawDescOnce sync.Once
	file_api_default_proto_rawDescData = file_api_default_proto_rawDesc
)

func file_api_default_proto_rawDescGZIP() []byte {
	file_api_default_proto_rawDescOnce.Do(func() {
		file_api_default_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_default_proto_rawDescData)
	})
	return file_api_default_proto_rawDescData
}

var file_api_default_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_default_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: api.Empty
	(*HelloResp)(nil),      // 1: api.HelloResp
	(*RiakStoreReq)(nil),   // 2: api.RiakStoreReq
	(*RiakStoreResp)(nil),  // 3: api.RiakStoreResp
	(*RiakFetchReq)(nil),   // 4: api.RiakFetchReq
	(*KV)(nil),             // 5: api.KV
	(*RiakFetchResp)(nil),  // 6: api.RiakFetchResp
	(*RiakDeleteReq)(nil),  // 7: api.RiakDeleteReq
	(*RiakDeleteResp)(nil), // 8: api.RiakDeleteResp
	(common.Country)(0),    // 9: common.Country
}
var file_api_default_proto_depIdxs = []int32{
	9, // 0: api.HelloResp.Country:type_name -> common.Country
	5, // 1: api.RiakFetchResp.Data:type_name -> api.KV
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_default_proto_init() }
func file_api_default_proto_init() {
	if File_api_default_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_default_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_default_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_api_default_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiakStoreReq); i {
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
		file_api_default_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiakStoreResp); i {
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
		file_api_default_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiakFetchReq); i {
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
		file_api_default_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KV); i {
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
		file_api_default_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiakFetchResp); i {
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
		file_api_default_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiakDeleteReq); i {
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
		file_api_default_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RiakDeleteResp); i {
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
			RawDescriptor: file_api_default_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_default_proto_goTypes,
		DependencyIndexes: file_api_default_proto_depIdxs,
		MessageInfos:      file_api_default_proto_msgTypes,
	}.Build()
	File_api_default_proto = out.File
	file_api_default_proto_rawDesc = nil
	file_api_default_proto_goTypes = nil
	file_api_default_proto_depIdxs = nil
}
