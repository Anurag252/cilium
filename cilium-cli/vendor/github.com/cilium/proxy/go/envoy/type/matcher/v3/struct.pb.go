// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.19.4
// source: envoy/type/matcher/v3/struct.proto

package matcherv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// StructMatcher provides a general interface to check if a given value is matched in
// google.protobuf.Struct. It uses `path` to retrieve the value
// from the struct and then check if it's matched to the specified value.
//
// For example, for the following Struct:
//
// .. code-block:: yaml
//
//	fields:
//	  a:
//	    struct_value:
//	      fields:
//	        b:
//	          struct_value:
//	            fields:
//	              c:
//	                string_value: pro
//	        t:
//	          list_value:
//	            values:
//	              - string_value: m
//	              - string_value: n
//
// The following MetadataMatcher is matched as the path [a, b, c] will retrieve a string value "pro"
// from the Metadata which is matched to the specified prefix match.
//
// .. code-block:: yaml
//
//	path:
//	- key: a
//	- key: b
//	- key: c
//	value:
//	  string_match:
//	    prefix: pr
//
// The following StructMatcher is matched as the code will match one of the string values in the
// list at the path [a, t].
//
// .. code-block:: yaml
//
//	path:
//	- key: a
//	- key: t
//	value:
//	  list_match:
//	    one_of:
//	      string_match:
//	        exact: m
//
// An example use of StructMatcher is to match metadata in envoy.v*.core.Node.
type StructMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path to retrieve the Value from the Struct.
	Path []*StructMatcher_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	// The StructMatcher is matched if the value retrieved by path is matched to this value.
	Value *ValueMatcher `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *StructMatcher) Reset() {
	*x = StructMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v3_struct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructMatcher) ProtoMessage() {}

func (x *StructMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v3_struct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructMatcher.ProtoReflect.Descriptor instead.
func (*StructMatcher) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v3_struct_proto_rawDescGZIP(), []int{0}
}

func (x *StructMatcher) GetPath() []*StructMatcher_PathSegment {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *StructMatcher) GetValue() *ValueMatcher {
	if x != nil {
		return x.Value
	}
	return nil
}

// Specifies the segment in a path to retrieve value from Struct.
type StructMatcher_PathSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Segment:
	//
	//	*StructMatcher_PathSegment_Key
	Segment isStructMatcher_PathSegment_Segment `protobuf_oneof:"segment"`
}

func (x *StructMatcher_PathSegment) Reset() {
	*x = StructMatcher_PathSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_v3_struct_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructMatcher_PathSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructMatcher_PathSegment) ProtoMessage() {}

func (x *StructMatcher_PathSegment) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_v3_struct_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructMatcher_PathSegment.ProtoReflect.Descriptor instead.
func (*StructMatcher_PathSegment) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_v3_struct_proto_rawDescGZIP(), []int{0, 0}
}

func (m *StructMatcher_PathSegment) GetSegment() isStructMatcher_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (x *StructMatcher_PathSegment) GetKey() string {
	if x, ok := x.GetSegment().(*StructMatcher_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

type isStructMatcher_PathSegment_Segment interface {
	isStructMatcher_PathSegment_Segment()
}

type StructMatcher_PathSegment_Key struct {
	// If specified, use the key to retrieve the value in a Struct.
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*StructMatcher_PathSegment_Key) isStructMatcher_PathSegment_Segment() {}

var File_envoy_type_matcher_v3_struct_proto protoreflect.FileDescriptor

var file_envoy_type_matcher_v3_struct_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x21, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f,
	0x76, 0x33, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x43, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x33, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x6f, 0x0a, 0x0b, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x3a, 0x33, 0x9a, 0xc5,
	0x88, 0x1e, 0x2e, 0x0a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x0e, 0x0a, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x03, 0xf8, 0x42,
	0x01, 0x3a, 0x27, 0x9a, 0xc5, 0x88, 0x1e, 0x22, 0x0a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x84, 0x01, 0x0a, 0x23, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x76, 0x33, 0x42, 0x0b, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_type_matcher_v3_struct_proto_rawDescOnce sync.Once
	file_envoy_type_matcher_v3_struct_proto_rawDescData = file_envoy_type_matcher_v3_struct_proto_rawDesc
)

func file_envoy_type_matcher_v3_struct_proto_rawDescGZIP() []byte {
	file_envoy_type_matcher_v3_struct_proto_rawDescOnce.Do(func() {
		file_envoy_type_matcher_v3_struct_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_matcher_v3_struct_proto_rawDescData)
	})
	return file_envoy_type_matcher_v3_struct_proto_rawDescData
}

var file_envoy_type_matcher_v3_struct_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_type_matcher_v3_struct_proto_goTypes = []interface{}{
	(*StructMatcher)(nil),             // 0: envoy.type.matcher.v3.StructMatcher
	(*StructMatcher_PathSegment)(nil), // 1: envoy.type.matcher.v3.StructMatcher.PathSegment
	(*ValueMatcher)(nil),              // 2: envoy.type.matcher.v3.ValueMatcher
}
var file_envoy_type_matcher_v3_struct_proto_depIdxs = []int32{
	1, // 0: envoy.type.matcher.v3.StructMatcher.path:type_name -> envoy.type.matcher.v3.StructMatcher.PathSegment
	2, // 1: envoy.type.matcher.v3.StructMatcher.value:type_name -> envoy.type.matcher.v3.ValueMatcher
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_type_matcher_v3_struct_proto_init() }
func file_envoy_type_matcher_v3_struct_proto_init() {
	if File_envoy_type_matcher_v3_struct_proto != nil {
		return
	}
	file_envoy_type_matcher_v3_value_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_matcher_v3_struct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructMatcher); i {
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
		file_envoy_type_matcher_v3_struct_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructMatcher_PathSegment); i {
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
	file_envoy_type_matcher_v3_struct_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StructMatcher_PathSegment_Key)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_type_matcher_v3_struct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_matcher_v3_struct_proto_goTypes,
		DependencyIndexes: file_envoy_type_matcher_v3_struct_proto_depIdxs,
		MessageInfos:      file_envoy_type_matcher_v3_struct_proto_msgTypes,
	}.Build()
	File_envoy_type_matcher_v3_struct_proto = out.File
	file_envoy_type_matcher_v3_struct_proto_rawDesc = nil
	file_envoy_type_matcher_v3_struct_proto_goTypes = nil
	file_envoy_type_matcher_v3_struct_proto_depIdxs = nil
}
