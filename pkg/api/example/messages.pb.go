// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/example/messages.proto

package example

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatePostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	AuthorId      string                 `protobuf:"bytes,2,opt,name=author_id,proto3" json:"author_id,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	mi := &file_api_example_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_example_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_api_example_messages_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePostRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *CreatePostRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreatePostResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PostId        uint64                 `protobuf:"varint,1,opt,name=post_id,proto3" json:"post_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePostResponse) Reset() {
	*x = CreatePostResponse{}
	mi := &file_api_example_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostResponse) ProtoMessage() {}

func (x *CreatePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_example_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostResponse.ProtoReflect.Descriptor instead.
func (*CreatePostResponse) Descriptor() ([]byte, []int) {
	return file_api_example_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePostResponse) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type ListPostsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPostsRequest) Reset() {
	*x = ListPostsRequest{}
	mi := &file_api_example_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsRequest) ProtoMessage() {}

func (x *ListPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_example_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsRequest.ProtoReflect.Descriptor instead.
func (*ListPostsRequest) Descriptor() ([]byte, []int) {
	return file_api_example_messages_proto_rawDescGZIP(), []int{2}
}

type ListPostsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPostsResponse) Reset() {
	*x = ListPostsResponse{}
	mi := &file_api_example_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsResponse) ProtoMessage() {}

func (x *ListPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_example_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsResponse.ProtoReflect.Descriptor instead.
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return file_api_example_messages_proto_rawDescGZIP(), []int{3}
}

var File_api_example_messages_proto protoreflect.FileDescriptor

const file_api_example_messages_proto_rawDesc = "" +
	"\n" +
	"\x1aapi/example/messages.proto\x12$github.com.PavelNen.grpc.api.example\x1a\x1bbuf/validate/validate.proto\"\x82\x01\n" +
	"\x11CreatePostRequest\x12 \n" +
	"\x05title\x18\x01 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x01\x18\x80\x02R\x05title\x12%\n" +
	"\tauthor_id\x18\x02 \x01(\tB\a\xbaH\x04r\x02\x10\x01R\tauthor_id\x12$\n" +
	"\acontent\x18\x03 \x01(\tB\n" +
	"\xbaH\ar\x05\x10\x01\x18\x80\x10R\acontent\".\n" +
	"\x12CreatePostResponse\x12\x18\n" +
	"\apost_id\x18\x01 \x01(\x04R\apost_id\"\x12\n" +
	"\x10ListPostsRequest\"\x13\n" +
	"\x11ListPostsResponseB\x82\x02\n" +
	"(com.github.com.PavelNen.grpc.api.exampleB\rMessagesProtoP\x01Z\x0fpkg/api/example\xa2\x02\x06GCPGAE\xaa\x02$Github.Com.PavelNen.Grpc.Api.Example\xca\x02$Github\\Com\\PavelNen\\Grpc\\Api\\Example\xe2\x020Github\\Com\\PavelNen\\Grpc\\Api\\Example\\GPBMetadata\xea\x02)Github::Com::PavelNen::Grpc::Api::Exampleb\x06proto3"

var (
	file_api_example_messages_proto_rawDescOnce sync.Once
	file_api_example_messages_proto_rawDescData []byte
)

func file_api_example_messages_proto_rawDescGZIP() []byte {
	file_api_example_messages_proto_rawDescOnce.Do(func() {
		file_api_example_messages_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_example_messages_proto_rawDesc), len(file_api_example_messages_proto_rawDesc)))
	})
	return file_api_example_messages_proto_rawDescData
}

var file_api_example_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_example_messages_proto_goTypes = []any{
	(*CreatePostRequest)(nil),  // 0: github.com.PavelNen.grpc.api.example.CreatePostRequest
	(*CreatePostResponse)(nil), // 1: github.com.PavelNen.grpc.api.example.CreatePostResponse
	(*ListPostsRequest)(nil),   // 2: github.com.PavelNen.grpc.api.example.ListPostsRequest
	(*ListPostsResponse)(nil),  // 3: github.com.PavelNen.grpc.api.example.ListPostsResponse
}
var file_api_example_messages_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_example_messages_proto_init() }
func file_api_example_messages_proto_init() {
	if File_api_example_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_example_messages_proto_rawDesc), len(file_api_example_messages_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_example_messages_proto_goTypes,
		DependencyIndexes: file_api_example_messages_proto_depIdxs,
		MessageInfos:      file_api_example_messages_proto_msgTypes,
	}.Build()
	File_api_example_messages_proto = out.File
	file_api_example_messages_proto_goTypes = nil
	file_api_example_messages_proto_depIdxs = nil
}
