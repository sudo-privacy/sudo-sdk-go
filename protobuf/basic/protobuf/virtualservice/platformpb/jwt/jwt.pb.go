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
// source: sudo/mpc_virtual_service/platform/jwt.proto

package jwt

import (
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

type CreateJwtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account  string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Cipher   string `protobuf:"bytes,3,opt,name=cipher,proto3" json:"cipher,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Tfa      string `protobuf:"bytes,5,opt,name=tfa,proto3" json:"tfa,omitempty"`
}

func (x *CreateJwtRequest) Reset() {
	*x = CreateJwtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJwtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJwtRequest) ProtoMessage() {}

func (x *CreateJwtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJwtRequest.ProtoReflect.Descriptor instead.
func (*CreateJwtRequest) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescGZIP(), []int{0}
}

func (x *CreateJwtRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CreateJwtRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateJwtRequest) GetCipher() string {
	if x != nil {
		return x.Cipher
	}
	return ""
}

func (x *CreateJwtRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateJwtRequest) GetTfa() string {
	if x != nil {
		return x.Tfa
	}
	return ""
}

type CreateJwtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access  string `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Refresh string `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

func (x *CreateJwtResponse) Reset() {
	*x = CreateJwtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJwtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJwtResponse) ProtoMessage() {}

func (x *CreateJwtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJwtResponse.ProtoReflect.Descriptor instead.
func (*CreateJwtResponse) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescGZIP(), []int{1}
}

func (x *CreateJwtResponse) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

func (x *CreateJwtResponse) GetRefresh() string {
	if x != nil {
		return x.Refresh
	}
	return ""
}

type RefreshJwtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Refresh string `protobuf:"bytes,1,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

func (x *RefreshJwtRequest) Reset() {
	*x = RefreshJwtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshJwtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshJwtRequest) ProtoMessage() {}

func (x *RefreshJwtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshJwtRequest.ProtoReflect.Descriptor instead.
func (*RefreshJwtRequest) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshJwtRequest) GetRefresh() string {
	if x != nil {
		return x.Refresh
	}
	return ""
}

type RefreshJwtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access  string `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Refresh string `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

func (x *RefreshJwtResponse) Reset() {
	*x = RefreshJwtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshJwtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshJwtResponse) ProtoMessage() {}

func (x *RefreshJwtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshJwtResponse.ProtoReflect.Descriptor instead.
func (*RefreshJwtResponse) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshJwtResponse) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

func (x *RefreshJwtResponse) GetRefresh() string {
	if x != nil {
		return x.Refresh
	}
	return ""
}

type DeactiveJwtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access  string `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	Refresh string `protobuf:"bytes,2,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

func (x *DeactiveJwtRequest) Reset() {
	*x = DeactiveJwtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactiveJwtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactiveJwtRequest) ProtoMessage() {}

func (x *DeactiveJwtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactiveJwtRequest.ProtoReflect.Descriptor instead.
func (*DeactiveJwtRequest) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescGZIP(), []int{4}
}

func (x *DeactiveJwtRequest) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

func (x *DeactiveJwtRequest) GetRefresh() string {
	if x != nil {
		return x.Refresh
	}
	return ""
}

var File_sudo_mpc_virtual_service_platform_jwt_proto protoreflect.FileDescriptor

var file_sudo_mpc_virtual_service_platform_jwt_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73,
	0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x6a, 0x77, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x77, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x66, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x66, 0x61, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4a, 0x77, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x22, 0x2d, 0x0a, 0x11, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4a, 0x77, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x22, 0x46, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4a, 0x77, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x22, 0x46, 0x0a, 0x12, 0x44, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x4a, 0x77, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescOnce sync.Once
	file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescData = file_sudo_mpc_virtual_service_platform_jwt_proto_rawDesc
)

func file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescGZIP() []byte {
	file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescOnce.Do(func() {
		file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescData = protoimpl.X.CompressGZIP(file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescData)
	})
	return file_sudo_mpc_virtual_service_platform_jwt_proto_rawDescData
}

var file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sudo_mpc_virtual_service_platform_jwt_proto_goTypes = []interface{}{
	(*CreateJwtRequest)(nil),   // 0: sudo.protobuf.platform.jwt.CreateJwtRequest
	(*CreateJwtResponse)(nil),  // 1: sudo.protobuf.platform.jwt.CreateJwtResponse
	(*RefreshJwtRequest)(nil),  // 2: sudo.protobuf.platform.jwt.RefreshJwtRequest
	(*RefreshJwtResponse)(nil), // 3: sudo.protobuf.platform.jwt.RefreshJwtResponse
	(*DeactiveJwtRequest)(nil), // 4: sudo.protobuf.platform.jwt.DeactiveJwtRequest
}
var file_sudo_mpc_virtual_service_platform_jwt_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sudo_mpc_virtual_service_platform_jwt_proto_init() }
func file_sudo_mpc_virtual_service_platform_jwt_proto_init() {
	if File_sudo_mpc_virtual_service_platform_jwt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJwtRequest); i {
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
		file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJwtResponse); i {
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
		file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshJwtRequest); i {
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
		file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshJwtResponse); i {
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
		file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeactiveJwtRequest); i {
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
			RawDescriptor: file_sudo_mpc_virtual_service_platform_jwt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sudo_mpc_virtual_service_platform_jwt_proto_goTypes,
		DependencyIndexes: file_sudo_mpc_virtual_service_platform_jwt_proto_depIdxs,
		MessageInfos:      file_sudo_mpc_virtual_service_platform_jwt_proto_msgTypes,
	}.Build()
	File_sudo_mpc_virtual_service_platform_jwt_proto = out.File
	file_sudo_mpc_virtual_service_platform_jwt_proto_rawDesc = nil
	file_sudo_mpc_virtual_service_platform_jwt_proto_goTypes = nil
	file_sudo_mpc_virtual_service_platform_jwt_proto_depIdxs = nil
}
