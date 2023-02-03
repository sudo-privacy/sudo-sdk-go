// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: protobuf/mpc_virtual_service/platform/service/common.proto

package platformpb

import (
	context "context"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	jwt1 "github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/jwt"
	jwt "github.com/sudo-privacy/sudo-sdk-go/protobuf/virtualservice/platformpb/jwt"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

var File_protobuf_mpc_virtual_service_platform_service_common_proto protoreflect.FileDescriptor

var file_protobuf_mpc_virtual_service_platform_service_common_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6a, 0x77, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x94, 0x01, 0x0a,
	0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x89, 0x01, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x77, 0x74, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x77, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x77, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6a, 0x77, 0x74,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x06, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x30, 0x5a, 0x2e, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_protobuf_mpc_virtual_service_platform_service_common_proto_goTypes = []interface{}{
	(*jwt.CreateJwtRequest)(nil),   // 0: protobuf.platform.jwt.CreateJwtRequest
	(*jwt1.CreateJwtResponse)(nil), // 1: sudo.protobuf.platform.jwt.CreateJwtResponse
}
var file_protobuf_mpc_virtual_service_platform_service_common_proto_depIdxs = []int32{
	0, // 0: protobuf.platform.service.Common.CreateJwt:input_type -> protobuf.platform.jwt.CreateJwtRequest
	1, // 1: protobuf.platform.service.Common.CreateJwt:output_type -> sudo.protobuf.platform.jwt.CreateJwtResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_mpc_virtual_service_platform_service_common_proto_init() }
func file_protobuf_mpc_virtual_service_platform_service_common_proto_init() {
	if File_protobuf_mpc_virtual_service_platform_service_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_mpc_virtual_service_platform_service_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_mpc_virtual_service_platform_service_common_proto_goTypes,
		DependencyIndexes: file_protobuf_mpc_virtual_service_platform_service_common_proto_depIdxs,
	}.Build()
	File_protobuf_mpc_virtual_service_platform_service_common_proto = out.File
	file_protobuf_mpc_virtual_service_platform_service_common_proto_rawDesc = nil
	file_protobuf_mpc_virtual_service_platform_service_common_proto_goTypes = nil
	file_protobuf_mpc_virtual_service_platform_service_common_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CommonClient is the client API for Common service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommonClient interface {
	CreateJwt(ctx context.Context, in *jwt.CreateJwtRequest, opts ...grpc.CallOption) (*jwt1.CreateJwtResponse, error)
}

type commonClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonClient(cc grpc.ClientConnInterface) CommonClient {
	return &commonClient{cc}
}

func (c *commonClient) CreateJwt(ctx context.Context, in *jwt.CreateJwtRequest, opts ...grpc.CallOption) (*jwt1.CreateJwtResponse, error) {
	out := new(jwt1.CreateJwtResponse)
	err := c.cc.Invoke(ctx, "/protobuf.platform.service.Common/CreateJwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServer is the server API for Common service.
type CommonServer interface {
	CreateJwt(context.Context, *jwt.CreateJwtRequest) (*jwt1.CreateJwtResponse, error)
}

// UnimplementedCommonServer can be embedded to have forward compatible implementations.
type UnimplementedCommonServer struct {
}

func (*UnimplementedCommonServer) CreateJwt(context.Context, *jwt.CreateJwtRequest) (*jwt1.CreateJwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJwt not implemented")
}

func RegisterCommonServer(s *grpc.Server, srv CommonServer) {
	s.RegisterService(&_Common_serviceDesc, srv)
}

func _Common_CreateJwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(jwt.CreateJwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).CreateJwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.platform.service.Common/CreateJwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).CreateJwt(ctx, req.(*jwt.CreateJwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Common_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.platform.service.Common",
	HandlerType: (*CommonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJwt",
			Handler:    _Common_CreateJwt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/mpc_virtual_service/platform/service/common.proto",
}