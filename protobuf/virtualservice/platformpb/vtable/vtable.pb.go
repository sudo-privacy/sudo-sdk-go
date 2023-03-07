// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: protobuf/mpc_virtual_service/platform/vtable.proto

package vtable

import (
	datasource "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/datasource"
	_ "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/options"
	vtable "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/basic/protobuf/virtualservice/platformpb/vtable"
	common "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/virtualservice/platformpb/common"
	user "github.com/sudo-privacy/sudo-sdk-go/v2/protobuf/virtualservice/platformpb/user"
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

type Vtable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base      *vtable.VtableBase     `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	OwnerUser *user.User             `protobuf:"bytes,2,opt,name=owner_user,json=ownerUser,proto3" json:"owner_user,omitempty"`
	Datasrc   *datasource.DataSource `protobuf:"bytes,3,opt,name=datasrc,proto3" json:"datasrc,omitempty"`
	Asset     *common.Asset          `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
	Visible   bool                   `protobuf:"varint,5,opt,name=visible,proto3" json:"visible,omitempty"`
}

func (x *Vtable) Reset() {
	*x = Vtable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vtable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vtable) ProtoMessage() {}

func (x *Vtable) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vtable.ProtoReflect.Descriptor instead.
func (*Vtable) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescGZIP(), []int{0}
}

func (x *Vtable) GetBase() *vtable.VtableBase {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *Vtable) GetOwnerUser() *user.User {
	if x != nil {
		return x.OwnerUser
	}
	return nil
}

func (x *Vtable) GetDatasrc() *datasource.DataSource {
	if x != nil {
		return x.Datasrc
	}
	return nil
}

func (x *Vtable) GetAsset() *common.Asset {
	if x != nil {
		return x.Asset
	}
	return nil
}

func (x *Vtable) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

type ListVtablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryOptions             *vtable.VtableQueryOptions `protobuf:"bytes,1,opt,name=query_options,json=queryOptions,proto3" json:"query_options,omitempty"`
	WithGbaInfo              bool                       `protobuf:"varint,2,opt,name=with_gba_info,json=withGbaInfo,proto3" json:"with_gba_info,omitempty"`
	WithDepartmentVisibility bool                       `protobuf:"varint,3,opt,name=with_department_visibility,json=withDepartmentVisibility,proto3" json:"with_department_visibility,omitempty"`
}

func (x *ListVtablesRequest) Reset() {
	*x = ListVtablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVtablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVtablesRequest) ProtoMessage() {}

func (x *ListVtablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVtablesRequest.ProtoReflect.Descriptor instead.
func (*ListVtablesRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescGZIP(), []int{1}
}

func (x *ListVtablesRequest) GetQueryOptions() *vtable.VtableQueryOptions {
	if x != nil {
		return x.QueryOptions
	}
	return nil
}

func (x *ListVtablesRequest) GetWithGbaInfo() bool {
	if x != nil {
		return x.WithGbaInfo
	}
	return false
}

func (x *ListVtablesRequest) GetWithDepartmentVisibility() bool {
	if x != nil {
		return x.WithDepartmentVisibility
	}
	return false
}

type ListVtablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*Vtable `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int64     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListVtablesResponse) Reset() {
	*x = ListVtablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVtablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVtablesResponse) ProtoMessage() {}

func (x *ListVtablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVtablesResponse.ProtoReflect.Descriptor instead.
func (*ListVtablesResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescGZIP(), []int{2}
}

func (x *ListVtablesResponse) GetData() []*Vtable {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListVtablesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateVtablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vtables []*vtable.VtableBase `protobuf:"bytes,1,rep,name=vtables,proto3" json:"vtables,omitempty"`
}

func (x *CreateVtablesRequest) Reset() {
	*x = CreateVtablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVtablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVtablesRequest) ProtoMessage() {}

func (x *CreateVtablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVtablesRequest.ProtoReflect.Descriptor instead.
func (*CreateVtablesRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescGZIP(), []int{3}
}

func (x *CreateVtablesRequest) GetVtables() []*vtable.VtableBase {
	if x != nil {
		return x.Vtables
	}
	return nil
}

type CreateVtablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*vtable.VtableBase `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateVtablesResponse) Reset() {
	*x = CreateVtablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVtablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVtablesResponse) ProtoMessage() {}

func (x *CreateVtablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVtablesResponse.ProtoReflect.Descriptor instead.
func (*CreateVtablesResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescGZIP(), []int{4}
}

func (x *CreateVtablesResponse) GetData() []*vtable.VtableBase {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_protobuf_mpc_virtual_service_platform_vtable_proto protoreflect.FileDescriptor

var file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDesc = []byte{
	0x0a, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x32,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63,
	0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3c, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x37, 0x73, 0x75, 0x64, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x73, 0x75, 0x64,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x70, 0x63, 0x5f, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a, 0x06, 0x56, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x42, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x56,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x42, 0x03, 0xc8, 0x3e, 0x01, 0x52, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x48, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x72, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x72, 0x63, 0x12, 0x35, 0x0a, 0x05, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x05, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0xd3, 0x01, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x5b, 0x0a, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x75, 0x64,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x56, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x03, 0xc0,
	0x3e, 0x01, 0x52, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x22, 0x0a, 0x0d, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x67, 0x62, 0x61, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x77, 0x69, 0x74, 0x68, 0x47, 0x62, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x1a, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x77, 0x69, 0x74, 0x68, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x22, 0x61, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x56, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x5b, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a,
	0x07, 0x76, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x73, 0x75, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x56,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x52, 0x07, 0x76, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x22, 0x56, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x75, 0x64, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x56, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x37, 0x5a, 0x35, 0x70, 0x61,
	0x61, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x2f, 0x76, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescOnce sync.Once
	file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescData = file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDesc
)

func file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescGZIP() []byte {
	file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescOnce.Do(func() {
		file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescData)
	})
	return file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDescData
}

var file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protobuf_mpc_virtual_service_platform_vtable_proto_goTypes = []interface{}{
	(*Vtable)(nil),                    // 0: protobuf.platform.vtable.Vtable
	(*ListVtablesRequest)(nil),        // 1: protobuf.platform.vtable.ListVtablesRequest
	(*ListVtablesResponse)(nil),       // 2: protobuf.platform.vtable.ListVtablesResponse
	(*CreateVtablesRequest)(nil),      // 3: protobuf.platform.vtable.CreateVtablesRequest
	(*CreateVtablesResponse)(nil),     // 4: protobuf.platform.vtable.CreateVtablesResponse
	(*vtable.VtableBase)(nil),         // 5: sudo.protobuf.platform.vtable.VtableBase
	(*user.User)(nil),                 // 6: protobuf.platform.user.User
	(*datasource.DataSource)(nil),     // 7: sudo.protobuf.platform.data_source.DataSource
	(*common.Asset)(nil),              // 8: protobuf.platform.common.Asset
	(*vtable.VtableQueryOptions)(nil), // 9: sudo.protobuf.platform.vtable.VtableQueryOptions
}
var file_protobuf_mpc_virtual_service_platform_vtable_proto_depIdxs = []int32{
	5, // 0: protobuf.platform.vtable.Vtable.base:type_name -> sudo.protobuf.platform.vtable.VtableBase
	6, // 1: protobuf.platform.vtable.Vtable.owner_user:type_name -> protobuf.platform.user.User
	7, // 2: protobuf.platform.vtable.Vtable.datasrc:type_name -> sudo.protobuf.platform.data_source.DataSource
	8, // 3: protobuf.platform.vtable.Vtable.asset:type_name -> protobuf.platform.common.Asset
	9, // 4: protobuf.platform.vtable.ListVtablesRequest.query_options:type_name -> sudo.protobuf.platform.vtable.VtableQueryOptions
	0, // 5: protobuf.platform.vtable.ListVtablesResponse.data:type_name -> protobuf.platform.vtable.Vtable
	5, // 6: protobuf.platform.vtable.CreateVtablesRequest.vtables:type_name -> sudo.protobuf.platform.vtable.VtableBase
	5, // 7: protobuf.platform.vtable.CreateVtablesResponse.data:type_name -> sudo.protobuf.platform.vtable.VtableBase
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_protobuf_mpc_virtual_service_platform_vtable_proto_init() }
func file_protobuf_mpc_virtual_service_platform_vtable_proto_init() {
	if File_protobuf_mpc_virtual_service_platform_vtable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vtable); i {
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
		file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVtablesRequest); i {
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
		file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVtablesResponse); i {
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
		file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVtablesRequest); i {
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
		file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVtablesResponse); i {
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
			RawDescriptor: file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_mpc_virtual_service_platform_vtable_proto_goTypes,
		DependencyIndexes: file_protobuf_mpc_virtual_service_platform_vtable_proto_depIdxs,
		MessageInfos:      file_protobuf_mpc_virtual_service_platform_vtable_proto_msgTypes,
	}.Build()
	File_protobuf_mpc_virtual_service_platform_vtable_proto = out.File
	file_protobuf_mpc_virtual_service_platform_vtable_proto_rawDesc = nil
	file_protobuf_mpc_virtual_service_platform_vtable_proto_goTypes = nil
	file_protobuf_mpc_virtual_service_platform_vtable_proto_depIdxs = nil
}
