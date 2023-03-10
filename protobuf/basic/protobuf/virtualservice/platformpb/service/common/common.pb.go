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
// source: sudo/mpc_virtual_service/platform/service/common.proto

package platformpb

import (
	context "context"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	jwt "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/jwt"
	party "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/party"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_sudo_mpc_virtual_service_platform_service_common_proto protoreflect.FileDescriptor

var file_sudo_mpc_virtual_service_platform_service_common_proto_rawDesc = []byte{
	0x0a, 0x36, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2d, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xa1, 0x05, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x8e, 0x01, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x77, 0x74, 0x12, 0x2c, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x77, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x6a, 0x77, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x77, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x6a, 0x77, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x92, 0x41, 0x06, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x92, 0x01, 0x0a, 0x0a,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4a, 0x77, 0x74, 0x12, 0x2d, 0x2e, 0x73, 0x75, 0x64,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4a,
	0x77, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x75, 0x64, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4a, 0x77,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x22, 0x11, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6a, 0x77, 0x74, 0x2f, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x06, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x5c, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4a, 0x77, 0x74, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6a, 0x77, 0x74, 0x2f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x92, 0x41, 0x06, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x7f,
	0x0a, 0x0b, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4a, 0x77, 0x74, 0x12, 0x2e, 0x2e,
	0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x44, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x4a, 0x77, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x6a, 0x77, 0x74, 0x2f, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x06, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x91, 0x01, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x2e, 0x2e,
	0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e,
	0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x92, 0x41, 0x06, 0x0a, 0x04, 0x4d,
	0x69, 0x73, 0x63, 0x42, 0x2b, 0x5a, 0x29, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_sudo_mpc_virtual_service_platform_service_common_proto_goTypes = []interface{}{
	(*jwt.CreateJwtRequest)(nil),    // 0: sudo.protobuf.platform.jwt.CreateJwtRequest
	(*jwt.RefreshJwtRequest)(nil),   // 1: sudo.protobuf.platform.jwt.RefreshJwtRequest
	(*emptypb.Empty)(nil),           // 2: google.protobuf.Empty
	(*jwt.DeactiveJwtRequest)(nil),  // 3: sudo.protobuf.platform.jwt.DeactiveJwtRequest
	(*party.PingPartyRequest)(nil),  // 4: sudo.protobuf.platform.party.PingPartyRequest
	(*jwt.CreateJwtResponse)(nil),   // 5: sudo.protobuf.platform.jwt.CreateJwtResponse
	(*jwt.RefreshJwtResponse)(nil),  // 6: sudo.protobuf.platform.jwt.RefreshJwtResponse
	(*party.PingPartyResponse)(nil), // 7: sudo.protobuf.platform.party.PingPartyResponse
}
var file_sudo_mpc_virtual_service_platform_service_common_proto_depIdxs = []int32{
	0, // 0: sudo.protobuf.platform.service.Common.CreateJwt:input_type -> sudo.protobuf.platform.jwt.CreateJwtRequest
	1, // 1: sudo.protobuf.platform.service.Common.RefreshJwt:input_type -> sudo.protobuf.platform.jwt.RefreshJwtRequest
	2, // 2: sudo.protobuf.platform.service.Common.CheckJwt:input_type -> google.protobuf.Empty
	3, // 3: sudo.protobuf.platform.service.Common.DeactiveJwt:input_type -> sudo.protobuf.platform.jwt.DeactiveJwtRequest
	4, // 4: sudo.protobuf.platform.service.Common.PingParty:input_type -> sudo.protobuf.platform.party.PingPartyRequest
	5, // 5: sudo.protobuf.platform.service.Common.CreateJwt:output_type -> sudo.protobuf.platform.jwt.CreateJwtResponse
	6, // 6: sudo.protobuf.platform.service.Common.RefreshJwt:output_type -> sudo.protobuf.platform.jwt.RefreshJwtResponse
	2, // 7: sudo.protobuf.platform.service.Common.CheckJwt:output_type -> google.protobuf.Empty
	2, // 8: sudo.protobuf.platform.service.Common.DeactiveJwt:output_type -> google.protobuf.Empty
	7, // 9: sudo.protobuf.platform.service.Common.PingParty:output_type -> sudo.protobuf.platform.party.PingPartyResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sudo_mpc_virtual_service_platform_service_common_proto_init() }
func file_sudo_mpc_virtual_service_platform_service_common_proto_init() {
	if File_sudo_mpc_virtual_service_platform_service_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sudo_mpc_virtual_service_platform_service_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sudo_mpc_virtual_service_platform_service_common_proto_goTypes,
		DependencyIndexes: file_sudo_mpc_virtual_service_platform_service_common_proto_depIdxs,
	}.Build()
	File_sudo_mpc_virtual_service_platform_service_common_proto = out.File
	file_sudo_mpc_virtual_service_platform_service_common_proto_rawDesc = nil
	file_sudo_mpc_virtual_service_platform_service_common_proto_goTypes = nil
	file_sudo_mpc_virtual_service_platform_service_common_proto_depIdxs = nil
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
	CreateJwt(ctx context.Context, in *jwt.CreateJwtRequest, opts ...grpc.CallOption) (*jwt.CreateJwtResponse, error)
	RefreshJwt(ctx context.Context, in *jwt.RefreshJwtRequest, opts ...grpc.CallOption) (*jwt.RefreshJwtResponse, error)
	CheckJwt(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeactiveJwt(ctx context.Context, in *jwt.DeactiveJwtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PingParty(ctx context.Context, in *party.PingPartyRequest, opts ...grpc.CallOption) (*party.PingPartyResponse, error)
}

type commonClient struct {
	cc grpc.ClientConnInterface
}

func NewCommonClient(cc grpc.ClientConnInterface) CommonClient {
	return &commonClient{cc}
}

func (c *commonClient) CreateJwt(ctx context.Context, in *jwt.CreateJwtRequest, opts ...grpc.CallOption) (*jwt.CreateJwtResponse, error) {
	out := new(jwt.CreateJwtResponse)
	err := c.cc.Invoke(ctx, "/sudo.protobuf.platform.service.Common/CreateJwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) RefreshJwt(ctx context.Context, in *jwt.RefreshJwtRequest, opts ...grpc.CallOption) (*jwt.RefreshJwtResponse, error) {
	out := new(jwt.RefreshJwtResponse)
	err := c.cc.Invoke(ctx, "/sudo.protobuf.platform.service.Common/RefreshJwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) CheckJwt(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sudo.protobuf.platform.service.Common/CheckJwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) DeactiveJwt(ctx context.Context, in *jwt.DeactiveJwtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/sudo.protobuf.platform.service.Common/DeactiveJwt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) PingParty(ctx context.Context, in *party.PingPartyRequest, opts ...grpc.CallOption) (*party.PingPartyResponse, error) {
	out := new(party.PingPartyResponse)
	err := c.cc.Invoke(ctx, "/sudo.protobuf.platform.service.Common/PingParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServer is the server API for Common service.
type CommonServer interface {
	CreateJwt(context.Context, *jwt.CreateJwtRequest) (*jwt.CreateJwtResponse, error)
	RefreshJwt(context.Context, *jwt.RefreshJwtRequest) (*jwt.RefreshJwtResponse, error)
	CheckJwt(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	DeactiveJwt(context.Context, *jwt.DeactiveJwtRequest) (*emptypb.Empty, error)
	PingParty(context.Context, *party.PingPartyRequest) (*party.PingPartyResponse, error)
}

// UnimplementedCommonServer can be embedded to have forward compatible implementations.
type UnimplementedCommonServer struct {
}

func (*UnimplementedCommonServer) CreateJwt(context.Context, *jwt.CreateJwtRequest) (*jwt.CreateJwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJwt not implemented")
}
func (*UnimplementedCommonServer) RefreshJwt(context.Context, *jwt.RefreshJwtRequest) (*jwt.RefreshJwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshJwt not implemented")
}
func (*UnimplementedCommonServer) CheckJwt(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckJwt not implemented")
}
func (*UnimplementedCommonServer) DeactiveJwt(context.Context, *jwt.DeactiveJwtRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactiveJwt not implemented")
}
func (*UnimplementedCommonServer) PingParty(context.Context, *party.PingPartyRequest) (*party.PingPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingParty not implemented")
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
		FullMethod: "/sudo.protobuf.platform.service.Common/CreateJwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).CreateJwt(ctx, req.(*jwt.CreateJwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_RefreshJwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(jwt.RefreshJwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).RefreshJwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sudo.protobuf.platform.service.Common/RefreshJwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).RefreshJwt(ctx, req.(*jwt.RefreshJwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_CheckJwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).CheckJwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sudo.protobuf.platform.service.Common/CheckJwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).CheckJwt(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_DeactiveJwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(jwt.DeactiveJwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).DeactiveJwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sudo.protobuf.platform.service.Common/DeactiveJwt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).DeactiveJwt(ctx, req.(*jwt.DeactiveJwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_PingParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(party.PingPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).PingParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sudo.protobuf.platform.service.Common/PingParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).PingParty(ctx, req.(*party.PingPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Common_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sudo.protobuf.platform.service.Common",
	HandlerType: (*CommonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJwt",
			Handler:    _Common_CreateJwt_Handler,
		},
		{
			MethodName: "RefreshJwt",
			Handler:    _Common_RefreshJwt_Handler,
		},
		{
			MethodName: "CheckJwt",
			Handler:    _Common_CheckJwt_Handler,
		},
		{
			MethodName: "DeactiveJwt",
			Handler:    _Common_DeactiveJwt_Handler,
		},
		{
			MethodName: "PingParty",
			Handler:    _Common_PingParty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sudo/mpc_virtual_service/platform/service/common.proto",
}
