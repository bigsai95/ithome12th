// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: users.proto

package mypb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserPwd  string `protobuf:"bytes,2,opt,name=user_pwd,json=userPwd,proto3" json:"user_pwd,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserRequest) GetUserPwd() string {
	if x != nil {
		return x.UserPwd
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type UsersListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UsersListRequest) Reset() {
	*x = UsersListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersListRequest) ProtoMessage() {}

func (x *UsersListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersListRequest.ProtoReflect.Descriptor instead.
func (*UsersListRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{2}
}

type UsersListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	UserName []string `protobuf:"bytes,2,rep,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *UsersListResponse) Reset() {
	*x = UsersListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersListResponse) ProtoMessage() {}

func (x *UsersListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersListResponse.ProtoReflect.Descriptor instead.
func (*UsersListResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{3}
}

func (x *UsersListResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *UsersListResponse) GetUserName() []string {
	if x != nil {
		return x.UserName
	}
	return nil
}

var File_users_proto protoreflect.FileDescriptor

var file_users_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d,
	0x79, 0x70, 0x62, 0x22, 0x45, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x77, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x50, 0x77, 0x64, 0x22, 0x26, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x32, 0xba, 0x01, 0x0a, 0x0e, 0x4d, 0x79, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11,
	0x2e, 0x6d, 0x79, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x79, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x6d, 0x79, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x79, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x6d, 0x79, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x79, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x3b, 0x6d, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_users_proto_rawDescOnce sync.Once
	file_users_proto_rawDescData = file_users_proto_rawDesc
)

func file_users_proto_rawDescGZIP() []byte {
	file_users_proto_rawDescOnce.Do(func() {
		file_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_users_proto_rawDescData)
	})
	return file_users_proto_rawDescData
}

var file_users_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_users_proto_goTypes = []interface{}{
	(*UserRequest)(nil),       // 0: mypb.UserRequest
	(*UserResponse)(nil),      // 1: mypb.UserResponse
	(*UsersListRequest)(nil),  // 2: mypb.UsersListRequest
	(*UsersListResponse)(nil), // 3: mypb.UsersListResponse
}
var file_users_proto_depIdxs = []int32{
	0, // 0: mypb.MyprotoService.AddUser:input_type -> mypb.UserRequest
	0, // 1: mypb.MyprotoService.LoginUser:input_type -> mypb.UserRequest
	2, // 2: mypb.MyprotoService.UsersList:input_type -> mypb.UsersListRequest
	1, // 3: mypb.MyprotoService.AddUser:output_type -> mypb.UserResponse
	1, // 4: mypb.MyprotoService.LoginUser:output_type -> mypb.UserResponse
	3, // 5: mypb.MyprotoService.UsersList:output_type -> mypb.UsersListResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_users_proto_init() }
func file_users_proto_init() {
	if File_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersListRequest); i {
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
		file_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersListResponse); i {
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
			RawDescriptor: file_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_proto_goTypes,
		DependencyIndexes: file_users_proto_depIdxs,
		MessageInfos:      file_users_proto_msgTypes,
	}.Build()
	File_users_proto = out.File
	file_users_proto_rawDesc = nil
	file_users_proto_goTypes = nil
	file_users_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MyprotoServiceClient is the client API for MyprotoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MyprotoServiceClient interface {
	AddUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	LoginUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	UsersList(ctx context.Context, in *UsersListRequest, opts ...grpc.CallOption) (*UsersListResponse, error)
}

type myprotoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyprotoServiceClient(cc grpc.ClientConnInterface) MyprotoServiceClient {
	return &myprotoServiceClient{cc}
}

func (c *myprotoServiceClient) AddUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/mypb.MyprotoService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myprotoServiceClient) LoginUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/mypb.MyprotoService/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myprotoServiceClient) UsersList(ctx context.Context, in *UsersListRequest, opts ...grpc.CallOption) (*UsersListResponse, error) {
	out := new(UsersListResponse)
	err := c.cc.Invoke(ctx, "/mypb.MyprotoService/UsersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyprotoServiceServer is the server API for MyprotoService service.
type MyprotoServiceServer interface {
	AddUser(context.Context, *UserRequest) (*UserResponse, error)
	LoginUser(context.Context, *UserRequest) (*UserResponse, error)
	UsersList(context.Context, *UsersListRequest) (*UsersListResponse, error)
}

// UnimplementedMyprotoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMyprotoServiceServer struct {
}

func (*UnimplementedMyprotoServiceServer) AddUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (*UnimplementedMyprotoServiceServer) LoginUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (*UnimplementedMyprotoServiceServer) UsersList(context.Context, *UsersListRequest) (*UsersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsersList not implemented")
}

func RegisterMyprotoServiceServer(s *grpc.Server, srv MyprotoServiceServer) {
	s.RegisterService(&_MyprotoService_serviceDesc, srv)
}

func _MyprotoService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyprotoServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mypb.MyprotoService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyprotoServiceServer).AddUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyprotoService_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyprotoServiceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mypb.MyprotoService/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyprotoServiceServer).LoginUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyprotoService_UsersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyprotoServiceServer).UsersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mypb.MyprotoService/UsersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyprotoServiceServer).UsersList(ctx, req.(*UsersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MyprotoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mypb.MyprotoService",
	HandlerType: (*MyprotoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _MyprotoService_AddUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _MyprotoService_LoginUser_Handler,
		},
		{
			MethodName: "UsersList",
			Handler:    _MyprotoService_UsersList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
