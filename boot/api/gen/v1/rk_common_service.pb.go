// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: v1/rk_common_service.proto

package rk_grpc_common_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
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

type HealthyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthyRequest) Reset() {
	*x = HealthyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthyRequest) ProtoMessage() {}

func (x *HealthyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthyRequest.ProtoReflect.Descriptor instead.
func (*HealthyRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{0}
}

type GcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GcRequest) Reset() {
	*x = GcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcRequest) ProtoMessage() {}

func (x *GcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcRequest.ProtoReflect.Descriptor instead.
func (*GcRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{1}
}

type InfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InfoRequest) Reset() {
	*x = InfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoRequest) ProtoMessage() {}

func (x *InfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoRequest.ProtoReflect.Descriptor instead.
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{2}
}

type ConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigsRequest) Reset() {
	*x = ConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigsRequest) ProtoMessage() {}

func (x *ConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigsRequest.ProtoReflect.Descriptor instead.
func (*ConfigsRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{3}
}

type ApisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApisRequest) Reset() {
	*x = ApisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApisRequest) ProtoMessage() {}

func (x *ApisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApisRequest.ProtoReflect.Descriptor instead.
func (*ApisRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{4}
}

type SysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysRequest) Reset() {
	*x = SysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysRequest) ProtoMessage() {}

func (x *SysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysRequest.ProtoReflect.Descriptor instead.
func (*SysRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{5}
}

type ReqRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqRequest) Reset() {
	*x = ReqRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqRequest) ProtoMessage() {}

func (x *ReqRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqRequest.ProtoReflect.Descriptor instead.
func (*ReqRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{6}
}

type EntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EntriesRequest) Reset() {
	*x = EntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntriesRequest) ProtoMessage() {}

func (x *EntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntriesRequest.ProtoReflect.Descriptor instead.
func (*EntriesRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{7}
}

type CertsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CertsRequest) Reset() {
	*x = CertsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertsRequest) ProtoMessage() {}

func (x *CertsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertsRequest.ProtoReflect.Descriptor instead.
func (*CertsRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{8}
}

type LogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogsRequest) Reset() {
	*x = LogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsRequest) ProtoMessage() {}

func (x *LogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsRequest.ProtoReflect.Descriptor instead.
func (*LogsRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{9}
}

type DepsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DepsRequest) Reset() {
	*x = DepsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DepsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepsRequest) ProtoMessage() {}

func (x *DepsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepsRequest.ProtoReflect.Descriptor instead.
func (*DepsRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{10}
}

type LicenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LicenseRequest) Reset() {
	*x = LicenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LicenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LicenseRequest) ProtoMessage() {}

func (x *LicenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LicenseRequest.ProtoReflect.Descriptor instead.
func (*LicenseRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{11}
}

type ReadmeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadmeRequest) Reset() {
	*x = ReadmeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadmeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadmeRequest) ProtoMessage() {}

func (x *ReadmeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadmeRequest.ProtoReflect.Descriptor instead.
func (*ReadmeRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{12}
}

type GitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GitRequest) Reset() {
	*x = GitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitRequest) ProtoMessage() {}

func (x *GitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitRequest.ProtoReflect.Descriptor instead.
func (*GitRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{13}
}

type GwErrorMappingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GwErrorMappingRequest) Reset() {
	*x = GwErrorMappingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_rk_common_service_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GwErrorMappingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GwErrorMappingRequest) ProtoMessage() {}

func (x *GwErrorMappingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_rk_common_service_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GwErrorMappingRequest.ProtoReflect.Descriptor instead.
func (*GwErrorMappingRequest) Descriptor() ([]byte, []int) {
	return file_v1_rk_common_service_proto_rawDescGZIP(), []int{14}
}

var File_v1_rk_common_service_proto protoreflect.FileDescriptor

var file_v1_rk_common_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x31, 0x2f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x72, 0x6b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x10, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x0b, 0x0a, 0x09, 0x47, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x0d, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x0d, 0x0a, 0x0b, 0x41, 0x70, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x0c, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0c, 0x0a,
	0x0a, 0x52, 0x65, 0x71, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a,
	0x0c, 0x43, 0x65, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a,
	0x0b, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b,
	0x44, 0x65, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a,
	0x0d, 0x52, 0x65, 0x61, 0x64, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0c,
	0x0a, 0x0a, 0x47, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x17, 0x0a, 0x15,
	0x47, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xae, 0x07, 0x0a, 0x0f, 0x52, 0x6b, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x79, 0x12, 0x19, 0x2e, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x02, 0x47, 0x63,
	0x12, 0x14, 0x2e, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x72, 0x6b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x19, 0x2e, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a,
	0x04, 0x41, 0x70, 0x69, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x03, 0x53, 0x79, 0x73, 0x12,
	0x15, 0x2e, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x15, 0x2e, 0x72, 0x6b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x03, 0x47, 0x69,
	0x74, 0x12, 0x15, 0x2e, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x19,
	0x2e, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x05, 0x43, 0x65, 0x72, 0x74, 0x73, 0x12, 0x17, 0x2e,
	0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x39, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x6b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04,
	0x44, 0x65, 0x70, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x73, 0x65, 0x12, 0x19, 0x2e, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x64,
	0x6d, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x47, 0x77, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x2e, 0x72, 0x6b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6b, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_rk_common_service_proto_rawDescOnce sync.Once
	file_v1_rk_common_service_proto_rawDescData = file_v1_rk_common_service_proto_rawDesc
)

func file_v1_rk_common_service_proto_rawDescGZIP() []byte {
	file_v1_rk_common_service_proto_rawDescOnce.Do(func() {
		file_v1_rk_common_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_rk_common_service_proto_rawDescData)
	})
	return file_v1_rk_common_service_proto_rawDescData
}

var file_v1_rk_common_service_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_v1_rk_common_service_proto_goTypes = []interface{}{
	(*HealthyRequest)(nil),        // 0: rk.api.v1.HealthyRequest
	(*GcRequest)(nil),             // 1: rk.api.v1.GcRequest
	(*InfoRequest)(nil),           // 2: rk.api.v1.InfoRequest
	(*ConfigsRequest)(nil),        // 3: rk.api.v1.ConfigsRequest
	(*ApisRequest)(nil),           // 4: rk.api.v1.ApisRequest
	(*SysRequest)(nil),            // 5: rk.api.v1.SysRequest
	(*ReqRequest)(nil),            // 6: rk.api.v1.ReqRequest
	(*EntriesRequest)(nil),        // 7: rk.api.v1.EntriesRequest
	(*CertsRequest)(nil),          // 8: rk.api.v1.CertsRequest
	(*LogsRequest)(nil),           // 9: rk.api.v1.LogsRequest
	(*DepsRequest)(nil),           // 10: rk.api.v1.DepsRequest
	(*LicenseRequest)(nil),        // 11: rk.api.v1.LicenseRequest
	(*ReadmeRequest)(nil),         // 12: rk.api.v1.ReadmeRequest
	(*GitRequest)(nil),            // 13: rk.api.v1.GitRequest
	(*GwErrorMappingRequest)(nil), // 14: rk.api.v1.GwErrorMappingRequest
	(*structpb.Struct)(nil),       // 15: google.protobuf.Struct
}
var file_v1_rk_common_service_proto_depIdxs = []int32{
	0,  // 0: rk.api.v1.RkCommonService.Healthy:input_type -> rk.api.v1.HealthyRequest
	1,  // 1: rk.api.v1.RkCommonService.Gc:input_type -> rk.api.v1.GcRequest
	2,  // 2: rk.api.v1.RkCommonService.Info:input_type -> rk.api.v1.InfoRequest
	3,  // 3: rk.api.v1.RkCommonService.Configs:input_type -> rk.api.v1.ConfigsRequest
	4,  // 4: rk.api.v1.RkCommonService.Apis:input_type -> rk.api.v1.ApisRequest
	5,  // 5: rk.api.v1.RkCommonService.Sys:input_type -> rk.api.v1.SysRequest
	6,  // 6: rk.api.v1.RkCommonService.Req:input_type -> rk.api.v1.ReqRequest
	13, // 7: rk.api.v1.RkCommonService.Git:input_type -> rk.api.v1.GitRequest
	7,  // 8: rk.api.v1.RkCommonService.Entries:input_type -> rk.api.v1.EntriesRequest
	8,  // 9: rk.api.v1.RkCommonService.Certs:input_type -> rk.api.v1.CertsRequest
	9,  // 10: rk.api.v1.RkCommonService.Logs:input_type -> rk.api.v1.LogsRequest
	10, // 11: rk.api.v1.RkCommonService.Deps:input_type -> rk.api.v1.DepsRequest
	11, // 12: rk.api.v1.RkCommonService.License:input_type -> rk.api.v1.LicenseRequest
	12, // 13: rk.api.v1.RkCommonService.Readme:input_type -> rk.api.v1.ReadmeRequest
	14, // 14: rk.api.v1.RkCommonService.GwErrorMapping:input_type -> rk.api.v1.GwErrorMappingRequest
	15, // 15: rk.api.v1.RkCommonService.Healthy:output_type -> google.protobuf.Struct
	15, // 16: rk.api.v1.RkCommonService.Gc:output_type -> google.protobuf.Struct
	15, // 17: rk.api.v1.RkCommonService.Info:output_type -> google.protobuf.Struct
	15, // 18: rk.api.v1.RkCommonService.Configs:output_type -> google.protobuf.Struct
	15, // 19: rk.api.v1.RkCommonService.Apis:output_type -> google.protobuf.Struct
	15, // 20: rk.api.v1.RkCommonService.Sys:output_type -> google.protobuf.Struct
	15, // 21: rk.api.v1.RkCommonService.Req:output_type -> google.protobuf.Struct
	15, // 22: rk.api.v1.RkCommonService.Git:output_type -> google.protobuf.Struct
	15, // 23: rk.api.v1.RkCommonService.Entries:output_type -> google.protobuf.Struct
	15, // 24: rk.api.v1.RkCommonService.Certs:output_type -> google.protobuf.Struct
	15, // 25: rk.api.v1.RkCommonService.Logs:output_type -> google.protobuf.Struct
	15, // 26: rk.api.v1.RkCommonService.Deps:output_type -> google.protobuf.Struct
	15, // 27: rk.api.v1.RkCommonService.License:output_type -> google.protobuf.Struct
	15, // 28: rk.api.v1.RkCommonService.Readme:output_type -> google.protobuf.Struct
	15, // 29: rk.api.v1.RkCommonService.GwErrorMapping:output_type -> google.protobuf.Struct
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_v1_rk_common_service_proto_init() }
func file_v1_rk_common_service_proto_init() {
	if File_v1_rk_common_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_rk_common_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthyRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigsRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApisRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntriesRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertsRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DepsRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LicenseRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadmeRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitRequest); i {
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
		file_v1_rk_common_service_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GwErrorMappingRequest); i {
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
			RawDescriptor: file_v1_rk_common_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_rk_common_service_proto_goTypes,
		DependencyIndexes: file_v1_rk_common_service_proto_depIdxs,
		MessageInfos:      file_v1_rk_common_service_proto_msgTypes,
	}.Build()
	File_v1_rk_common_service_proto = out.File
	file_v1_rk_common_service_proto_rawDesc = nil
	file_v1_rk_common_service_proto_goTypes = nil
	file_v1_rk_common_service_proto_depIdxs = nil
}
