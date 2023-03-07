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
// source: sudo/mpc_virtual_service/platform/user.proto

package user

import (
	_ "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/options"
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

type UserBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedTs  int64  `protobuf:"varint,2,opt,name=created_ts,json=createdTs,proto3" json:"created_ts,omitempty"`
	UpdatedTs  int64  `protobuf:"varint,3,opt,name=updated_ts,json=updatedTs,proto3" json:"updated_ts,omitempty"`
	Account    string `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	Unlocktime int64  `protobuf:"varint,5,opt,name=unlocktime,proto3" json:"unlocktime,omitempty"`
	Username   string `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	Email      string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Gender     uint64 `protobuf:"varint,8,opt,name=gender,proto3" json:"gender,omitempty"`
	Avatar     string `protobuf:"bytes,9,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Phone      string `protobuf:"bytes,10,opt,name=phone,proto3" json:"phone,omitempty"`
	JobDesc    string `protobuf:"bytes,11,opt,name=job_desc,json=jobDesc,proto3" json:"job_desc,omitempty"`
	Status     string `protobuf:"bytes,12,opt,name=status,proto3" json:"status,omitempty"`
	Banned     bool   `protobuf:"varint,13,opt,name=banned,proto3" json:"banned,omitempty"`
}

func (x *UserBase) Reset() {
	*x = UserBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBase) ProtoMessage() {}

func (x *UserBase) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBase.ProtoReflect.Descriptor instead.
func (*UserBase) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserBase) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserBase) GetCreatedTs() int64 {
	if x != nil {
		return x.CreatedTs
	}
	return 0
}

func (x *UserBase) GetUpdatedTs() int64 {
	if x != nil {
		return x.UpdatedTs
	}
	return 0
}

func (x *UserBase) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserBase) GetUnlocktime() int64 {
	if x != nil {
		return x.Unlocktime
	}
	return 0
}

func (x *UserBase) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserBase) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserBase) GetGender() uint64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserBase) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UserBase) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserBase) GetJobDesc() string {
	if x != nil {
		return x.JobDesc
	}
	return ""
}

func (x *UserBase) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UserBase) GetBanned() bool {
	if x != nil {
		return x.Banned
	}
	return false
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base   *UserBase    `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Groups []*UserGroup `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_user_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetBase() *UserBase {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *User) GetGroups() []*UserGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

type UserGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Users       []*User `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UserGroup) Reset() {
	*x = UserGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGroup) ProtoMessage() {}

func (x *UserGroup) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGroup.ProtoReflect.Descriptor instead.
func (*UserGroup) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UserGroup) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_sudo_mpc_virtual_service_platform_user_proto protoreflect.FileDescriptor

var file_sudo_mpc_virtual_service_platform_user_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x2f, 0x73, 0x75, 0x64,
	0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x02, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62,
	0x44, 0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x62, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x64, 0x22, 0x86, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x3e, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x75,
	0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61,
	0x73, 0x65, 0x42, 0x03, 0xc8, 0x3e, 0x01, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x7a, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x37, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sudo_mpc_virtual_service_platform_user_proto_rawDescOnce sync.Once
	file_sudo_mpc_virtual_service_platform_user_proto_rawDescData = file_sudo_mpc_virtual_service_platform_user_proto_rawDesc
)

func file_sudo_mpc_virtual_service_platform_user_proto_rawDescGZIP() []byte {
	file_sudo_mpc_virtual_service_platform_user_proto_rawDescOnce.Do(func() {
		file_sudo_mpc_virtual_service_platform_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_sudo_mpc_virtual_service_platform_user_proto_rawDescData)
	})
	return file_sudo_mpc_virtual_service_platform_user_proto_rawDescData
}

var file_sudo_mpc_virtual_service_platform_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sudo_mpc_virtual_service_platform_user_proto_goTypes = []interface{}{
	(*UserBase)(nil),  // 0: sudo.protobuf.platform.user.UserBase
	(*User)(nil),      // 1: sudo.protobuf.platform.user.User
	(*UserGroup)(nil), // 2: sudo.protobuf.platform.user.UserGroup
}
var file_sudo_mpc_virtual_service_platform_user_proto_depIdxs = []int32{
	0, // 0: sudo.protobuf.platform.user.User.base:type_name -> sudo.protobuf.platform.user.UserBase
	2, // 1: sudo.protobuf.platform.user.User.groups:type_name -> sudo.protobuf.platform.user.UserGroup
	1, // 2: sudo.protobuf.platform.user.UserGroup.users:type_name -> sudo.protobuf.platform.user.User
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_sudo_mpc_virtual_service_platform_user_proto_init() }
func file_sudo_mpc_virtual_service_platform_user_proto_init() {
	if File_sudo_mpc_virtual_service_platform_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sudo_mpc_virtual_service_platform_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBase); i {
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
		file_sudo_mpc_virtual_service_platform_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_sudo_mpc_virtual_service_platform_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGroup); i {
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
			RawDescriptor: file_sudo_mpc_virtual_service_platform_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sudo_mpc_virtual_service_platform_user_proto_goTypes,
		DependencyIndexes: file_sudo_mpc_virtual_service_platform_user_proto_depIdxs,
		MessageInfos:      file_sudo_mpc_virtual_service_platform_user_proto_msgTypes,
	}.Build()
	File_sudo_mpc_virtual_service_platform_user_proto = out.File
	file_sudo_mpc_virtual_service_platform_user_proto_rawDesc = nil
	file_sudo_mpc_virtual_service_platform_user_proto_goTypes = nil
	file_sudo_mpc_virtual_service_platform_user_proto_depIdxs = nil
}
