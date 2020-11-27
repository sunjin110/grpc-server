// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: book/book.proto

package book

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	comm "grpc-server/src/infra/grpc/proto/comm"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	User  string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_book_book_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_book_book_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_book_book_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeleteRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type ListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookInfoList []*BookInfo `protobuf:"bytes,1,rep,name=book_info_list,json=bookInfoList,proto3" json:"book_info_list,omitempty"`
}

func (x *ListReply) Reset() {
	*x = ListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReply) ProtoMessage() {}

func (x *ListReply) ProtoReflect() protoreflect.Message {
	mi := &file_book_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReply.ProtoReflect.Descriptor instead.
func (*ListReply) Descriptor() ([]byte, []int) {
	return file_book_book_proto_rawDescGZIP(), []int{3}
}

func (x *ListReply) GetBookInfoList() []*BookInfo {
	if x != nil {
		return x.BookInfoList
	}
	return nil
}

type BookInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	User  string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *BookInfo) Reset() {
	*x = BookInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookInfo) ProtoMessage() {}

func (x *BookInfo) ProtoReflect() protoreflect.Message {
	mi := &file_book_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookInfo.ProtoReflect.Descriptor instead.
func (*BookInfo) Descriptor() ([]byte, []int) {
	return file_book_book_proto_rawDescGZIP(), []int{4}
}

func (x *BookInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BookInfo) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BookInfo) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

var File_book_book_proto protoreflect.FileDescriptor

var file_book_book_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x1a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x37, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x6b,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x48,
	0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xc0, 0x01, 0x0a, 0x07, 0x42, 0x6f, 0x6f,
	0x6b, 0x52, 0x70, 0x63, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x31, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x13, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_book_proto_rawDescOnce sync.Once
	file_book_book_proto_rawDescData = file_book_book_proto_rawDesc
)

func file_book_book_proto_rawDescGZIP() []byte {
	file_book_book_proto_rawDescOnce.Do(func() {
		file_book_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_book_proto_rawDescData)
	})
	return file_book_book_proto_rawDescData
}

var file_book_book_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_book_book_proto_goTypes = []interface{}{
	(*CreateRequest)(nil), // 0: book.CreateRequest
	(*UpdateRequest)(nil), // 1: book.UpdateRequest
	(*DeleteRequest)(nil), // 2: book.DeleteRequest
	(*ListReply)(nil),     // 3: book.ListReply
	(*BookInfo)(nil),      // 4: book.BookInfo
	(*comm.Empty)(nil),    // 5: comm.Empty
}
var file_book_book_proto_depIdxs = []int32{
	4, // 0: book.ListReply.book_info_list:type_name -> book.BookInfo
	0, // 1: book.BookRpc.Create:input_type -> book.CreateRequest
	1, // 2: book.BookRpc.UpdatePrice:input_type -> book.UpdateRequest
	2, // 3: book.BookRpc.Delete:input_type -> book.DeleteRequest
	5, // 4: book.BookRpc.List:input_type -> comm.Empty
	5, // 5: book.BookRpc.Create:output_type -> comm.Empty
	5, // 6: book.BookRpc.UpdatePrice:output_type -> comm.Empty
	5, // 7: book.BookRpc.Delete:output_type -> comm.Empty
	3, // 8: book.BookRpc.List:output_type -> book.ListReply
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_book_book_proto_init() }
func file_book_book_proto_init() {
	if File_book_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_book_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_book_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_book_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_book_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReply); i {
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
		file_book_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookInfo); i {
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
			RawDescriptor: file_book_book_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_book_book_proto_goTypes,
		DependencyIndexes: file_book_book_proto_depIdxs,
		MessageInfos:      file_book_book_proto_msgTypes,
	}.Build()
	File_book_book_proto = out.File
	file_book_book_proto_rawDesc = nil
	file_book_book_proto_goTypes = nil
	file_book_book_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BookRpcClient is the client API for BookRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookRpcClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*comm.Empty, error)
	UpdatePrice(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*comm.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*comm.Empty, error)
	List(ctx context.Context, in *comm.Empty, opts ...grpc.CallOption) (*ListReply, error)
}

type bookRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewBookRpcClient(cc grpc.ClientConnInterface) BookRpcClient {
	return &bookRpcClient{cc}
}

func (c *bookRpcClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*comm.Empty, error) {
	out := new(comm.Empty)
	err := c.cc.Invoke(ctx, "/book.BookRpc/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookRpcClient) UpdatePrice(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*comm.Empty, error) {
	out := new(comm.Empty)
	err := c.cc.Invoke(ctx, "/book.BookRpc/UpdatePrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookRpcClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*comm.Empty, error) {
	out := new(comm.Empty)
	err := c.cc.Invoke(ctx, "/book.BookRpc/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookRpcClient) List(ctx context.Context, in *comm.Empty, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, "/book.BookRpc/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookRpcServer is the server API for BookRpc service.
type BookRpcServer interface {
	Create(context.Context, *CreateRequest) (*comm.Empty, error)
	UpdatePrice(context.Context, *UpdateRequest) (*comm.Empty, error)
	Delete(context.Context, *DeleteRequest) (*comm.Empty, error)
	List(context.Context, *comm.Empty) (*ListReply, error)
}

// UnimplementedBookRpcServer can be embedded to have forward compatible implementations.
type UnimplementedBookRpcServer struct {
}

func (*UnimplementedBookRpcServer) Create(context.Context, *CreateRequest) (*comm.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedBookRpcServer) UpdatePrice(context.Context, *UpdateRequest) (*comm.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePrice not implemented")
}
func (*UnimplementedBookRpcServer) Delete(context.Context, *DeleteRequest) (*comm.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedBookRpcServer) List(context.Context, *comm.Empty) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterBookRpcServer(s *grpc.Server, srv BookRpcServer) {
	s.RegisterService(&_BookRpc_serviceDesc, srv)
}

func _BookRpc_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookRpcServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookRpc/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookRpcServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookRpc_UpdatePrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookRpcServer).UpdatePrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookRpc/UpdatePrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookRpcServer).UpdatePrice(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookRpc_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookRpcServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookRpc/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookRpcServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookRpc_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(comm.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookRpcServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookRpc/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookRpcServer).List(ctx, req.(*comm.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookRpc",
	HandlerType: (*BookRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BookRpc_Create_Handler,
		},
		{
			MethodName: "UpdatePrice",
			Handler:    _BookRpc_UpdatePrice_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BookRpc_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _BookRpc_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book/book.proto",
}
