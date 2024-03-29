// Copyright 2022 Sudo Technology Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: sudo/mpc_virtual_service/platform/service/coordinator.proto

package platformpb

import (
	context "context"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	party "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/party"
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

var File_sudo_mpc_virtual_service_platform_service_coordinator_proto protoreflect.FileDescriptor

var file_sudo_mpc_virtual_service_platform_service_coordinator_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x73,
	0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x73, 0x75, 0x64,
	0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc0, 0x01, 0x0a, 0x0b, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0xb0, 0x01, 0x0a, 0x0b, 0x50,
	0x69, 0x6e, 0x67, 0x46, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x12, 0x30, 0x2e, 0x73, 0x75, 0x64,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x46, 0x75,
	0x72, 0x6e, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73,
	0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x69, 0x6e, 0x67,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x76, 0x32,
	0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x66, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x92, 0x41, 0x0d,
	0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x2b, 0x5a,
	0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_sudo_mpc_virtual_service_platform_service_coordinator_proto_goTypes = []interface{}{
	(*party.PingFurnaceRequest)(nil), // 0: sudo.protobuf.platform.party.PingFurnaceRequest
	(*party.PingPartyResponse)(nil),  // 1: sudo.protobuf.platform.party.PingPartyResponse
}
var file_sudo_mpc_virtual_service_platform_service_coordinator_proto_depIdxs = []int32{
	0, // 0: sudo.protobuf.platform.service.Coordinator.PingFurnace:input_type -> sudo.protobuf.platform.party.PingFurnaceRequest
	1, // 1: sudo.protobuf.platform.service.Coordinator.PingFurnace:output_type -> sudo.protobuf.platform.party.PingPartyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sudo_mpc_virtual_service_platform_service_coordinator_proto_init() }
func file_sudo_mpc_virtual_service_platform_service_coordinator_proto_init() {
	if File_sudo_mpc_virtual_service_platform_service_coordinator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sudo_mpc_virtual_service_platform_service_coordinator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sudo_mpc_virtual_service_platform_service_coordinator_proto_goTypes,
		DependencyIndexes: file_sudo_mpc_virtual_service_platform_service_coordinator_proto_depIdxs,
	}.Build()
	File_sudo_mpc_virtual_service_platform_service_coordinator_proto = out.File
	file_sudo_mpc_virtual_service_platform_service_coordinator_proto_rawDesc = nil
	file_sudo_mpc_virtual_service_platform_service_coordinator_proto_goTypes = nil
	file_sudo_mpc_virtual_service_platform_service_coordinator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoordinatorClient interface {
	// 检查与Furnace是否连通
	//
	// 测试与Furnace的连通情况，返回Furnace的基本信息，如角色、party id等。
	PingFurnace(ctx context.Context, in *party.PingFurnaceRequest, opts ...grpc.CallOption) (*party.PingPartyResponse, error)
}

type coordinatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinatorClient(cc grpc.ClientConnInterface) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) PingFurnace(ctx context.Context, in *party.PingFurnaceRequest, opts ...grpc.CallOption) (*party.PingPartyResponse, error) {
	out := new(party.PingPartyResponse)
	err := c.cc.Invoke(ctx, "/sudo.protobuf.platform.service.Coordinator/PingFurnace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
type CoordinatorServer interface {
	// 检查与Furnace是否连通
	//
	// 测试与Furnace的连通情况，返回Furnace的基本信息，如角色、party id等。
	PingFurnace(context.Context, *party.PingFurnaceRequest) (*party.PingPartyResponse, error)
}

// UnimplementedCoordinatorServer can be embedded to have forward compatible implementations.
type UnimplementedCoordinatorServer struct {
}

func (*UnimplementedCoordinatorServer) PingFurnace(context.Context, *party.PingFurnaceRequest) (*party.PingPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingFurnace not implemented")
}

func RegisterCoordinatorServer(s *grpc.Server, srv CoordinatorServer) {
	s.RegisterService(&_Coordinator_serviceDesc, srv)
}

func _Coordinator_PingFurnace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(party.PingFurnaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).PingFurnace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sudo.protobuf.platform.service.Coordinator/PingFurnace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).PingFurnace(ctx, req.(*party.PingFurnaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coordinator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sudo.protobuf.platform.service.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingFurnace",
			Handler:    _Coordinator_PingFurnace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sudo/mpc_virtual_service/platform/service/coordinator.proto",
}
