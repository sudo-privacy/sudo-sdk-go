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
// source: sudo/mpc_virtual_service/platform/apl.proto

package apl

import (
	enums "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/enums"
	_ "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Reviewer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AplId        uint64                   `protobuf:"varint,1,opt,name=apl_id,json=aplId,proto3" json:"apl_id,omitempty"`
	Reviewer     string                   `protobuf:"bytes,2,opt,name=reviewer,proto3" json:"reviewer,omitempty"`
	Action       enums.Application_Action `protobuf:"varint,3,opt,name=action,proto3,enum=sudo.protobuf.enums.Application_Action" json:"action,omitempty"`
	UpdatedAt    int64                    `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Reason       string                   `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	ActionUserId uint64                   `protobuf:"varint,6,opt,name=action_user_id,json=actionUserId,proto3" json:"action_user_id,omitempty"`
}

func (x *Reviewer) Reset() {
	*x = Reviewer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reviewer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reviewer) ProtoMessage() {}

func (x *Reviewer) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reviewer.ProtoReflect.Descriptor instead.
func (*Reviewer) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_apl_proto_rawDescGZIP(), []int{0}
}

func (x *Reviewer) GetAplId() uint64 {
	if x != nil {
		return x.AplId
	}
	return 0
}

func (x *Reviewer) GetReviewer() string {
	if x != nil {
		return x.Reviewer
	}
	return ""
}

func (x *Reviewer) GetAction() enums.Application_Action {
	if x != nil {
		return x.Action
	}
	return enums.Application_UNKNOWN_ACTION
}

func (x *Reviewer) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Reviewer) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Reviewer) GetActionUserId() uint64 {
	if x != nil {
		return x.ActionUserId
	}
	return 0
}

type CallbackLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AplId        uint64                   `protobuf:"varint,1,opt,name=apl_id,json=aplId,proto3" json:"apl_id,omitempty"`
	RefId        uint64                   `protobuf:"varint,2,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	Reviewer     string                   `protobuf:"bytes,3,opt,name=reviewer,proto3" json:"reviewer,omitempty"`
	Action       enums.Application_Action `protobuf:"varint,4,opt,name=action,proto3,enum=sudo.protobuf.enums.Application_Action" json:"action,omitempty"`
	ActionUserId uint64                   `protobuf:"varint,5,opt,name=action_user_id,json=actionUserId,proto3" json:"action_user_id,omitempty"`
	Successed    bool                     `protobuf:"varint,6,opt,name=successed,proto3" json:"successed,omitempty"`
	ErrMsg       string                   `protobuf:"bytes,7,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	RetryNo      int64                    `protobuf:"varint,8,opt,name=retry_no,json=retryNo,proto3" json:"retry_no,omitempty"`
	CreatedTs    int64                    `protobuf:"varint,9,opt,name=created_ts,json=createdTs,proto3" json:"created_ts,omitempty"`
	UpdatedTs    int64                    `protobuf:"varint,10,opt,name=updated_ts,json=updatedTs,proto3" json:"updated_ts,omitempty"`
}

func (x *CallbackLog) Reset() {
	*x = CallbackLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackLog) ProtoMessage() {}

func (x *CallbackLog) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackLog.ProtoReflect.Descriptor instead.
func (*CallbackLog) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_apl_proto_rawDescGZIP(), []int{1}
}

func (x *CallbackLog) GetAplId() uint64 {
	if x != nil {
		return x.AplId
	}
	return 0
}

func (x *CallbackLog) GetRefId() uint64 {
	if x != nil {
		return x.RefId
	}
	return 0
}

func (x *CallbackLog) GetReviewer() string {
	if x != nil {
		return x.Reviewer
	}
	return ""
}

func (x *CallbackLog) GetAction() enums.Application_Action {
	if x != nil {
		return x.Action
	}
	return enums.Application_UNKNOWN_ACTION
}

func (x *CallbackLog) GetActionUserId() uint64 {
	if x != nil {
		return x.ActionUserId
	}
	return 0
}

func (x *CallbackLog) GetSuccessed() bool {
	if x != nil {
		return x.Successed
	}
	return false
}

func (x *CallbackLog) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

func (x *CallbackLog) GetRetryNo() int64 {
	if x != nil {
		return x.RetryNo
	}
	return 0
}

func (x *CallbackLog) GetCreatedTs() int64 {
	if x != nil {
		return x.CreatedTs
	}
	return 0
}

func (x *CallbackLog) GetUpdatedTs() int64 {
	if x != nil {
		return x.UpdatedTs
	}
	return 0
}

type AplTakeActionParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action               enums.Application_Action   `protobuf:"varint,1,opt,name=action,proto3,enum=sudo.protobuf.enums.Application_Action" json:"action,omitempty"`
	AdditionalParameters map[string]*structpb.Value `protobuf:"bytes,2,rep,name=additional_parameters,json=additionalParameters,proto3" json:"additional_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Reason               string                     `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	RefId                uint64                     `protobuf:"varint,4,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	Type                 string                     `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *AplTakeActionParam) Reset() {
	*x = AplTakeActionParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AplTakeActionParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AplTakeActionParam) ProtoMessage() {}

func (x *AplTakeActionParam) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AplTakeActionParam.ProtoReflect.Descriptor instead.
func (*AplTakeActionParam) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_apl_proto_rawDescGZIP(), []int{2}
}

func (x *AplTakeActionParam) GetAction() enums.Application_Action {
	if x != nil {
		return x.Action
	}
	return enums.Application_UNKNOWN_ACTION
}

func (x *AplTakeActionParam) GetAdditionalParameters() map[string]*structpb.Value {
	if x != nil {
		return x.AdditionalParameters
	}
	return nil
}

func (x *AplTakeActionParam) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *AplTakeActionParam) GetRefId() uint64 {
	if x != nil {
		return x.RefId
	}
	return 0
}

func (x *AplTakeActionParam) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AplTakeActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param *AplTakeActionParam `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"`
	// 请求会转发到的中心节点ID
	// 若非空，则覆盖grpc metadata中设置的to-tusita
	ToTusita string `protobuf:"bytes,2,opt,name=to_tusita,json=toTusita,proto3" json:"to_tusita,omitempty"`
}

func (x *AplTakeActionRequest) Reset() {
	*x = AplTakeActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AplTakeActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AplTakeActionRequest) ProtoMessage() {}

func (x *AplTakeActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AplTakeActionRequest.ProtoReflect.Descriptor instead.
func (*AplTakeActionRequest) Descriptor() ([]byte, []int) {
	return file_sudo_mpc_virtual_service_platform_apl_proto_rawDescGZIP(), []int{3}
}

func (x *AplTakeActionRequest) GetParam() *AplTakeActionParam {
	if x != nil {
		return x.Param
	}
	return nil
}

func (x *AplTakeActionRequest) GetToTusita() string {
	if x != nil {
		return x.ToTusita
	}
	return ""
}

var File_sudo_mpc_virtual_service_platform_apl_proto protoreflect.FileDescriptor

var file_sudo_mpc_virtual_service_platform_apl_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73,
	0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2f, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdb, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x12, 0x15, 0x0a,
	0x06, 0x61, 0x70, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x61,
	0x70, 0x6c, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72,
	0x12, 0x3f, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x27, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xce,
	0x02, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x4c, 0x6f, 0x67, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x61, 0x70, 0x6c, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x72, 0x65, 0x66, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f,
	0x6e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4e,
	0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x73, 0x22,
	0xf8, 0x02, 0x0a, 0x12, 0x41, 0x70, 0x6c, 0x54, 0x61, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x3f, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7d, 0x0a, 0x15, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x61, 0x70, 0x6c, 0x2e, 0x41, 0x70, 0x6c, 0x54, 0x61, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x14, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x15,
	0x0a, 0x06, 0x72, 0x65, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x72, 0x65, 0x66, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x5f, 0x0a, 0x19, 0x41, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7e, 0x0a, 0x14, 0x41, 0x70,
	0x6c, 0x54, 0x61, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x49, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x70, 0x6c, 0x2e, 0x41,
	0x70, 0x6c, 0x54, 0x61, 0x6b, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x42, 0x03, 0xc8, 0x3e, 0x01, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x6f, 0x5f, 0x74, 0x75, 0x73, 0x69, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x6f, 0x54, 0x75, 0x73, 0x69, 0x74, 0x61, 0x42, 0x31, 0x5a, 0x2f, 0x70, 0x72,
	0x69, 0x76, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sudo_mpc_virtual_service_platform_apl_proto_rawDescOnce sync.Once
	file_sudo_mpc_virtual_service_platform_apl_proto_rawDescData = file_sudo_mpc_virtual_service_platform_apl_proto_rawDesc
)

func file_sudo_mpc_virtual_service_platform_apl_proto_rawDescGZIP() []byte {
	file_sudo_mpc_virtual_service_platform_apl_proto_rawDescOnce.Do(func() {
		file_sudo_mpc_virtual_service_platform_apl_proto_rawDescData = protoimpl.X.CompressGZIP(file_sudo_mpc_virtual_service_platform_apl_proto_rawDescData)
	})
	return file_sudo_mpc_virtual_service_platform_apl_proto_rawDescData
}

var file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sudo_mpc_virtual_service_platform_apl_proto_goTypes = []interface{}{
	(*Reviewer)(nil),              // 0: sudo.protobuf.platform.apl.Reviewer
	(*CallbackLog)(nil),           // 1: sudo.protobuf.platform.apl.CallbackLog
	(*AplTakeActionParam)(nil),    // 2: sudo.protobuf.platform.apl.AplTakeActionParam
	(*AplTakeActionRequest)(nil),  // 3: sudo.protobuf.platform.apl.AplTakeActionRequest
	nil,                           // 4: sudo.protobuf.platform.apl.AplTakeActionParam.AdditionalParametersEntry
	(enums.Application_Action)(0), // 5: sudo.protobuf.enums.Application.Action
	(*structpb.Value)(nil),        // 6: google.protobuf.Value
}
var file_sudo_mpc_virtual_service_platform_apl_proto_depIdxs = []int32{
	5, // 0: sudo.protobuf.platform.apl.Reviewer.action:type_name -> sudo.protobuf.enums.Application.Action
	5, // 1: sudo.protobuf.platform.apl.CallbackLog.action:type_name -> sudo.protobuf.enums.Application.Action
	5, // 2: sudo.protobuf.platform.apl.AplTakeActionParam.action:type_name -> sudo.protobuf.enums.Application.Action
	4, // 3: sudo.protobuf.platform.apl.AplTakeActionParam.additional_parameters:type_name -> sudo.protobuf.platform.apl.AplTakeActionParam.AdditionalParametersEntry
	2, // 4: sudo.protobuf.platform.apl.AplTakeActionRequest.param:type_name -> sudo.protobuf.platform.apl.AplTakeActionParam
	6, // 5: sudo.protobuf.platform.apl.AplTakeActionParam.AdditionalParametersEntry.value:type_name -> google.protobuf.Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_sudo_mpc_virtual_service_platform_apl_proto_init() }
func file_sudo_mpc_virtual_service_platform_apl_proto_init() {
	if File_sudo_mpc_virtual_service_platform_apl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reviewer); i {
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
		file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackLog); i {
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
		file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AplTakeActionParam); i {
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
		file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AplTakeActionRequest); i {
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
			RawDescriptor: file_sudo_mpc_virtual_service_platform_apl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sudo_mpc_virtual_service_platform_apl_proto_goTypes,
		DependencyIndexes: file_sudo_mpc_virtual_service_platform_apl_proto_depIdxs,
		MessageInfos:      file_sudo_mpc_virtual_service_platform_apl_proto_msgTypes,
	}.Build()
	File_sudo_mpc_virtual_service_platform_apl_proto = out.File
	file_sudo_mpc_virtual_service_platform_apl_proto_rawDesc = nil
	file_sudo_mpc_virtual_service_platform_apl_proto_goTypes = nil
	file_sudo_mpc_virtual_service_platform_apl_proto_depIdxs = nil
}
