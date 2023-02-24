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
// source: sudo/mpc_virtual_service/platform/party.proto

package party

import (
	enums "github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/enums"
	_ "github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/options"
	paginator "github.com/sudo-privacy/sudo-sdk-go/protobuf/basic/protobuf/virtualservice/platformpb/paginator"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PingPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Echo string `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
}

func (x *PingPartyRequest) Reset() {
	*x = PingPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingPartyRequest) ProtoMessage() {}

func (x *PingPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingPartyRequest.ProtoReflect.Descriptor instead.
func (*PingPartyRequest) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_party_proto_rawDescGZIP(), []int{0}
}

func (x *PingPartyRequest) GetEcho() string {
	if x != nil {
		return x.Echo
	}
	return ""
}

type PingPartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Echo string `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
	// 参与方的名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 参与方的party id
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// 参与方的角色
	Role enums.Party_Role `protobuf:"varint,4,opt,name=role,proto3,enum=sudo.protobuf.enums.Party_Role" json:"role,omitempty"`
}

func (x *PingPartyResponse) Reset() {
	*x = PingPartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingPartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingPartyResponse) ProtoMessage() {}

func (x *PingPartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingPartyResponse.ProtoReflect.Descriptor instead.
func (*PingPartyResponse) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_party_proto_rawDescGZIP(), []int{1}
}

func (x *PingPartyResponse) GetEcho() string {
	if x != nil {
		return x.Echo
	}
	return ""
}

func (x *PingPartyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PingPartyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PingPartyResponse) GetRole() enums.Party_Role {
	if x != nil {
		return x.Role
	}
	return enums.Party_TUSITA
}

type PingFurnaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Echo string `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
	// 连通测试的Furnace地址
	FurnaceUrl string `protobuf:"bytes,2,opt,name=furnace_url,json=furnaceUrl,proto3" json:"furnace_url,omitempty"`
}

func (x *PingFurnaceRequest) Reset() {
	*x = PingFurnaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingFurnaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingFurnaceRequest) ProtoMessage() {}

func (x *PingFurnaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingFurnaceRequest.ProtoReflect.Descriptor instead.
func (*PingFurnaceRequest) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_party_proto_rawDescGZIP(), []int{2}
}

func (x *PingFurnaceRequest) GetEcho() string {
	if x != nil {
		return x.Echo
	}
	return ""
}

func (x *PingFurnaceRequest) GetFurnaceUrl() string {
	if x != nil {
		return x.FurnaceUrl
	}
	return ""
}

type PartyQueryOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pager          *paginator.Paginator `protobuf:"bytes,1,opt,name=pager,proto3" json:"pager,omitempty"`
	ExcludePartyId string               `protobuf:"bytes,2,opt,name=exclude_party_id,json=excludePartyId,proto3" json:"exclude_party_id,omitempty"`
	PartyIds       []string             `protobuf:"bytes,3,rep,name=party_ids,json=partyIds,proto3" json:"party_ids,omitempty"`
	PartyId        string               `protobuf:"bytes,4,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	PartyAlias     string               `protobuf:"bytes,5,opt,name=party_alias,json=partyAlias,proto3" json:"party_alias,omitempty"`
	Url            string               `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Role           enums.Party_Role     `protobuf:"varint,7,opt,name=role,proto3,enum=sudo.protobuf.enums.Party_Role" json:"role,omitempty"`
	Status         enums.Party_Status   `protobuf:"varint,8,opt,name=status,proto3,enum=sudo.protobuf.enums.Party_Status" json:"status,omitempty"`
	QueryTusita    bool                 `protobuf:"varint,9,opt,name=query_tusita,json=queryTusita,proto3" json:"query_tusita,omitempty"`
}

func (x *PartyQueryOption) Reset() {
	*x = PartyQueryOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyQueryOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyQueryOption) ProtoMessage() {}

func (x *PartyQueryOption) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyQueryOption.ProtoReflect.Descriptor instead.
func (*PartyQueryOption) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_party_proto_rawDescGZIP(), []int{3}
}

func (x *PartyQueryOption) GetPager() *paginator.Paginator {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *PartyQueryOption) GetExcludePartyId() string {
	if x != nil {
		return x.ExcludePartyId
	}
	return ""
}

func (x *PartyQueryOption) GetPartyIds() []string {
	if x != nil {
		return x.PartyIds
	}
	return nil
}

func (x *PartyQueryOption) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PartyQueryOption) GetPartyAlias() string {
	if x != nil {
		return x.PartyAlias
	}
	return ""
}

func (x *PartyQueryOption) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PartyQueryOption) GetRole() enums.Party_Role {
	if x != nil {
		return x.Role
	}
	return enums.Party_TUSITA
}

func (x *PartyQueryOption) GetStatus() enums.Party_Status {
	if x != nil {
		return x.Status
	}
	return enums.Party_UNKNOWN_STATUS
}

func (x *PartyQueryOption) GetQueryTusita() bool {
	if x != nil {
		return x.QueryTusita
	}
	return false
}

type Party struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedTs   int64              `protobuf:"varint,2,opt,name=created_ts,json=createdTs,proto3" json:"created_ts,omitempty"`
	UpdatedTs   int64              `protobuf:"varint,3,opt,name=updated_ts,json=updatedTs,proto3" json:"updated_ts,omitempty"`
	BlockedTs   int64              `protobuf:"varint,4,opt,name=blocked_ts,json=blockedTs,proto3" json:"blocked_ts,omitempty"`
	Description string             `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	PartyAlias  string             `protobuf:"bytes,6,opt,name=party_alias,json=partyAlias,proto3" json:"party_alias,omitempty"`
	PartyId     string             `protobuf:"bytes,7,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Role        enums.Party_Role   `protobuf:"varint,8,opt,name=role,proto3,enum=sudo.protobuf.enums.Party_Role" json:"role,omitempty"`
	Status      enums.Party_Status `protobuf:"varint,9,opt,name=status,proto3,enum=sudo.protobuf.enums.Party_Status" json:"status,omitempty"`
	Url         string             `protobuf:"bytes,10,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Party) Reset() {
	*x = Party{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Party) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Party) ProtoMessage() {}

func (x *Party) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Party.ProtoReflect.Descriptor instead.
func (*Party) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_party_proto_rawDescGZIP(), []int{4}
}

func (x *Party) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Party) GetCreatedTs() int64 {
	if x != nil {
		return x.CreatedTs
	}
	return 0
}

func (x *Party) GetUpdatedTs() int64 {
	if x != nil {
		return x.UpdatedTs
	}
	return 0
}

func (x *Party) GetBlockedTs() int64 {
	if x != nil {
		return x.BlockedTs
	}
	return 0
}

func (x *Party) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Party) GetPartyAlias() string {
	if x != nil {
		return x.PartyAlias
	}
	return ""
}

func (x *Party) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *Party) GetRole() enums.Party_Role {
	if x != nil {
		return x.Role
	}
	return enums.Party_TUSITA
}

func (x *Party) GetStatus() enums.Party_Status {
	if x != nil {
		return x.Status
	}
	return enums.Party_UNKNOWN_STATUS
}

func (x *Party) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetPartiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tusitas []*Party `protobuf:"bytes,1,rep,name=tusitas,proto3" json:"tusitas,omitempty"`
	Total   int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetPartiesResponse) Reset() {
	*x = GetPartiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartiesResponse) ProtoMessage() {}

func (x *GetPartiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartiesResponse.ProtoReflect.Descriptor instead.
func (*GetPartiesResponse) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_party_proto_rawDescGZIP(), []int{5}
}

func (x *GetPartiesResponse) GetTusitas() []*Party {
	if x != nil {
		return x.Tusitas
	}
	return nil
}

func (x *GetPartiesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_sudo_mpc_virtual_service_platform_party_proto protoreflect.FileDescriptor

var file_sudo_mpc_virtual_service_platform_party_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f,
	0x73, 0x75, 0x64, 0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x6d, 0x70,
	0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x50, 0x69,
	0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x63,
	0x68, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x11, 0x50, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x33, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x4e, 0x0a, 0x12, 0x50, 0x69, 0x6e, 0x67, 0x46, 0x75, 0x72,
	0x6e, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x65,
	0x63, 0x68, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x12,
	0x24, 0x0a, 0x0b, 0x66, 0x75, 0x72, 0x6e, 0x61, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x66, 0x75, 0x72, 0x6e, 0x61,
	0x63, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x82, 0x03, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x74, 0x79, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x75, 0x64, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x42, 0x03, 0xc0, 0x3e, 0x01, 0x52, 0x05, 0x70, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x33, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73,
	0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x74, 0x75, 0x73, 0x69, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x54, 0x75, 0x73, 0x69, 0x74, 0x61, 0x22, 0xd4, 0x02, 0x0a, 0x05, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x54, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x54, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x54,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x79, 0x41,
	0x6c, 0x69, 0x61, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x33, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x69, 0x0a, 0x12, 0x67, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x74, 0x75, 0x73, 0x69, 0x74,
	0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x07, 0x74,
	0x75, 0x73, 0x69, 0x74, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sudo_mpc_virtual_service_platform_party_proto_rawDescOnce sync.Once
	file_sudo_mpc_virtual_service_platform_party_proto_rawDescData = file_sudo_mpc_virtual_service_platform_party_proto_rawDesc
)

func file_sudo_mpc_virtual_service_platform_party_proto_rawDescGZIP() []byte {
	file_sudo_mpc_virtual_service_platform_party_proto_rawDescOnce.Do(func() {
		file_sudo_mpc_virtual_service_platform_party_proto_rawDescData = protoimpl.X.CompressGZIP(file_sudo_mpc_virtual_service_platform_party_proto_rawDescData)
	})
	return file_sudo_mpc_virtual_service_platform_party_proto_rawDescData
}

var file_sudo_mpc_virtual_service_platform_party_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sudo_mpc_virtual_service_platform_party_proto_goTypes = []interface{}{
	(*PingPartyRequest)(nil),    // 0: sudo.protobuf.platform.party.PingPartyRequest
	(*PingPartyResponse)(nil),   // 1: sudo.protobuf.platform.party.PingPartyResponse
	(*PingFurnaceRequest)(nil),  // 2: sudo.protobuf.platform.party.PingFurnaceRequest
	(*PartyQueryOption)(nil),    // 3: sudo.protobuf.platform.party.PartyQueryOption
	(*Party)(nil),               // 4: sudo.protobuf.platform.party.Party
	(*GetPartiesResponse)(nil),  // 5: sudo.protobuf.platform.party.getPartiesResponse
	(enums.Party_Role)(0),       // 6: sudo.protobuf.enums.Party.Role
	(*paginator.Paginator)(nil), // 7: sudo.protobuf.platform.paginator.Paginator
	(enums.Party_Status)(0),     // 8: sudo.protobuf.enums.Party.Status
}
var file_sudo_mpc_virtual_service_platform_party_proto_depIdxs = []int32{
	6, // 0: sudo.protobuf.platform.party.PingPartyResponse.role:type_name -> sudo.protobuf.enums.Party.Role
	7, // 1: sudo.protobuf.platform.party.PartyQueryOption.pager:type_name -> sudo.protobuf.platform.paginator.Paginator
	6, // 2: sudo.protobuf.platform.party.PartyQueryOption.role:type_name -> sudo.protobuf.enums.Party.Role
	8, // 3: sudo.protobuf.platform.party.PartyQueryOption.status:type_name -> sudo.protobuf.enums.Party.Status
	6, // 4: sudo.protobuf.platform.party.Party.role:type_name -> sudo.protobuf.enums.Party.Role
	8, // 5: sudo.protobuf.platform.party.Party.status:type_name -> sudo.protobuf.enums.Party.Status
	4, // 6: sudo.protobuf.platform.party.getPartiesResponse.tusitas:type_name -> sudo.protobuf.platform.party.Party
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_sudo_mpc_virtual_service_platform_party_proto_init() }
func file_sudo_mpc_virtual_service_platform_party_proto_init() {
	if File_sudo_mpc_virtual_service_platform_party_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingPartyRequest); i {
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
		file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingPartyResponse); i {
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
		file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingFurnaceRequest); i {
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
		file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartyQueryOption); i {
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
		file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Party); i {
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
		file_sudo_mpc_virtual_service_platform_party_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartiesResponse); i {
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
			RawDescriptor: file_sudo_mpc_virtual_service_platform_party_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sudo_mpc_virtual_service_platform_party_proto_goTypes,
		DependencyIndexes: file_sudo_mpc_virtual_service_platform_party_proto_depIdxs,
		MessageInfos:      file_sudo_mpc_virtual_service_platform_party_proto_msgTypes,
	}.Build()
	File_sudo_mpc_virtual_service_platform_party_proto = out.File
	file_sudo_mpc_virtual_service_platform_party_proto_rawDesc = nil
	file_sudo_mpc_virtual_service_platform_party_proto_goTypes = nil
	file_sudo_mpc_virtual_service_platform_party_proto_depIdxs = nil
}
