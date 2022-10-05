// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: internal/turbodprotocol/turbod.proto

package turbodprotocol

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	SessionId string `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_internal_turbodprotocol_turbod_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *HelloRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_internal_turbodprotocol_turbod_proto_rawDescGZIP(), []int{1}
}

type ShutdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShutdownRequest) Reset() {
	*x = ShutdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownRequest) ProtoMessage() {}

func (x *ShutdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownRequest.ProtoReflect.Descriptor instead.
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return file_internal_turbodprotocol_turbod_proto_rawDescGZIP(), []int{2}
}

type ShutdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShutdownResponse) Reset() {
	*x = ShutdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownResponse) ProtoMessage() {}

func (x *ShutdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownResponse.ProtoReflect.Descriptor instead.
func (*ShutdownResponse) Descriptor() ([]byte, []int) {
	return file_internal_turbodprotocol_turbod_proto_rawDescGZIP(), []int{3}
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_internal_turbodprotocol_turbod_proto_rawDescGZIP(), []int{4}
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DaemonStatus *DaemonStatus `protobuf:"bytes,1,opt,name=daemonStatus,proto3" json:"daemonStatus,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_internal_turbodprotocol_turbod_proto_rawDescGZIP(), []int{5}
}

func (x *StatusResponse) GetDaemonStatus() *DaemonStatus {
	if x != nil {
		return x.DaemonStatus
	}
	return nil
}

type NotifyOutputsWrittenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutputGlobs []string `protobuf:"bytes,1,rep,name=output_globs,json=outputGlobs,proto3" json:"output_globs,omitempty"`
	Hash        string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *NotifyOutputsWrittenRequest) Reset() {
	*x = NotifyOutputsWrittenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyOutputsWrittenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOutputsWrittenRequest) ProtoMessage() {}

func (x *NotifyOutputsWrittenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOutputsWrittenRequest.ProtoReflect.Descriptor instead.
func (*NotifyOutputsWrittenRequest) Descriptor() ([]byte, []int) {
	return file_internal_turbodprotocol_turbod_proto_rawDescGZIP(), []int{6}
}

func (x *NotifyOutputsWrittenRequest) GetOutputGlobs() []string {
	if x != nil {
		return x.OutputGlobs
	}
	return nil
}

func (x *NotifyOutputsWrittenRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type NotifyOutputsWrittenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyOutputsWrittenResponse) Reset() {
	*x = NotifyOutputsWrittenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyOutputsWrittenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyOutputsWrittenResponse) ProtoMessage() {}

func (x *NotifyOutputsWrittenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyOutputsWrittenResponse.ProtoReflect.Descriptor instead.
func (*NotifyOutputsWrittenResponse) Descriptor() ([]byte, []int) {
	return file_internal_turbodprotocol_turbod_proto_rawDescGZIP(), []int{7}
}

type GetChangedOutputsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutputGlobs []string `protobuf:"bytes,1,rep,name=output_globs,json=outputGlobs,proto3" json:"output_globs,omitempty"`
	Hash        string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *GetChangedOutputsRequest) Reset() {
	*x = GetChangedOutputsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChangedOutputsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChangedOutputsRequest) ProtoMessage() {}

func (x *GetChangedOutputsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChangedOutputsRequest.ProtoReflect.Descriptor instead.
func (*GetChangedOutputsRequest) Descriptor() ([]byte, []int) {
	return file_internal_turbodprotocol_turbod_proto_rawDescGZIP(), []int{8}
}

func (x *GetChangedOutputsRequest) GetOutputGlobs() []string {
	if x != nil {
		return x.OutputGlobs
	}
	return nil
}

func (x *GetChangedOutputsRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type GetChangedOutputsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChangedOutputGlobs []string `protobuf:"bytes,1,rep,name=changed_output_globs,json=changedOutputGlobs,proto3" json:"changed_output_globs,omitempty"`
}

func (x *GetChangedOutputsResponse) Reset() {
	*x = GetChangedOutputsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChangedOutputsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChangedOutputsResponse) ProtoMessage() {}

func (x *GetChangedOutputsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChangedOutputsResponse.ProtoReflect.Descriptor instead.
func (*GetChangedOutputsResponse) Descriptor() ([]byte, []int) {
	return file_internal_turbodprotocol_turbod_proto_rawDescGZIP(), []int{9}
}

func (x *GetChangedOutputsResponse) GetChangedOutputGlobs() []string {
	if x != nil {
		return x.ChangedOutputGlobs
	}
	return nil
}

type DaemonStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogFile    string `protobuf:"bytes,1,opt,name=log_file,json=logFile,proto3" json:"log_file,omitempty"`
	UptimeMsec uint64 `protobuf:"varint,2,opt,name=uptime_msec,json=uptimeMsec,proto3" json:"uptime_msec,omitempty"`
}

func (x *DaemonStatus) Reset() {
	*x = DaemonStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DaemonStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DaemonStatus) ProtoMessage() {}

func (x *DaemonStatus) ProtoReflect() protoreflect.Message {
	mi := &file_internal_turbodprotocol_turbod_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DaemonStatus.ProtoReflect.Descriptor instead.
func (*DaemonStatus) Descriptor() ([]byte, []int) {
	return file_internal_turbodprotocol_turbod_proto_rawDescGZIP(), []int{10}
}

func (x *DaemonStatus) GetLogFile() string {
	if x != nil {
		return x.LogFile
	}
	return ""
}

func (x *DaemonStatus) GetUptimeMsec() uint64 {
	if x != nil {
		return x.UptimeMsec
	}
	return 0
}

var File_internal_turbodprotocol_turbod_proto protoreflect.FileDescriptor

var file_internal_turbodprotocol_turbod_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x75, 0x72, 0x62, 0x6f,
	0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x64, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x47, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x0f, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x11, 0x0a, 0x0f, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x64, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0c,
	0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x54, 0x0a, 0x1b,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x57, 0x72, 0x69,
	0x74, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x22, 0x1e, 0x0a, 0x1c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x57, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x51, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x47, 0x6c, 0x6f, 0x62,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x4d, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x5f, 0x67, 0x6c, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x12, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x47,
	0x6c, 0x6f, 0x62, 0x73, 0x22, 0x4a, 0x0a, 0x0c, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x65, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x65, 0x63,
	0x32, 0xc3, 0x03, 0x0a, 0x06, 0x54, 0x75, 0x72, 0x62, 0x6f, 0x64, 0x12, 0x44, 0x0a, 0x05, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1c, 0x2e, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x64, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4d, 0x0a, 0x08, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x1f, 0x2e,
	0x74, 0x75, 0x72, 0x62, 0x6f, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53,
	0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x47, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x2e, 0x74, 0x75, 0x72,
	0x62, 0x6f, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x74, 0x75, 0x72, 0x62,
	0x6f, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x14, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x57, 0x72, 0x69, 0x74, 0x74, 0x65,
	0x6e, 0x12, 0x2b, 0x2e, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73,
	0x57, 0x72, 0x69, 0x74, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x57, 0x72, 0x69,
	0x74, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x12, 0x28, 0x2e, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x74, 0x75,
	0x72, 0x62, 0x6f, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x65, 0x72, 0x63, 0x65, 0x6c, 0x2f, 0x74, 0x75, 0x72, 0x62,
	0x6f, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x64, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_turbodprotocol_turbod_proto_rawDescOnce sync.Once
	file_internal_turbodprotocol_turbod_proto_rawDescData = file_internal_turbodprotocol_turbod_proto_rawDesc
)

func file_internal_turbodprotocol_turbod_proto_rawDescGZIP() []byte {
	file_internal_turbodprotocol_turbod_proto_rawDescOnce.Do(func() {
		file_internal_turbodprotocol_turbod_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_turbodprotocol_turbod_proto_rawDescData)
	})
	return file_internal_turbodprotocol_turbod_proto_rawDescData
}

var file_internal_turbodprotocol_turbod_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_internal_turbodprotocol_turbod_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),                 // 0: turbodprotocol.HelloRequest
	(*HelloResponse)(nil),                // 1: turbodprotocol.HelloResponse
	(*ShutdownRequest)(nil),              // 2: turbodprotocol.ShutdownRequest
	(*ShutdownResponse)(nil),             // 3: turbodprotocol.ShutdownResponse
	(*StatusRequest)(nil),                // 4: turbodprotocol.StatusRequest
	(*StatusResponse)(nil),               // 5: turbodprotocol.StatusResponse
	(*NotifyOutputsWrittenRequest)(nil),  // 6: turbodprotocol.NotifyOutputsWrittenRequest
	(*NotifyOutputsWrittenResponse)(nil), // 7: turbodprotocol.NotifyOutputsWrittenResponse
	(*GetChangedOutputsRequest)(nil),     // 8: turbodprotocol.GetChangedOutputsRequest
	(*GetChangedOutputsResponse)(nil),    // 9: turbodprotocol.GetChangedOutputsResponse
	(*DaemonStatus)(nil),                 // 10: turbodprotocol.DaemonStatus
}
var file_internal_turbodprotocol_turbod_proto_depIdxs = []int32{
	10, // 0: turbodprotocol.StatusResponse.daemonStatus:type_name -> turbodprotocol.DaemonStatus
	0,  // 1: turbodprotocol.Turbod.Hello:input_type -> turbodprotocol.HelloRequest
	2,  // 2: turbodprotocol.Turbod.Shutdown:input_type -> turbodprotocol.ShutdownRequest
	4,  // 3: turbodprotocol.Turbod.Status:input_type -> turbodprotocol.StatusRequest
	6,  // 4: turbodprotocol.Turbod.NotifyOutputsWritten:input_type -> turbodprotocol.NotifyOutputsWrittenRequest
	8,  // 5: turbodprotocol.Turbod.GetChangedOutputs:input_type -> turbodprotocol.GetChangedOutputsRequest
	1,  // 6: turbodprotocol.Turbod.Hello:output_type -> turbodprotocol.HelloResponse
	3,  // 7: turbodprotocol.Turbod.Shutdown:output_type -> turbodprotocol.ShutdownResponse
	5,  // 8: turbodprotocol.Turbod.Status:output_type -> turbodprotocol.StatusResponse
	7,  // 9: turbodprotocol.Turbod.NotifyOutputsWritten:output_type -> turbodprotocol.NotifyOutputsWrittenResponse
	9,  // 10: turbodprotocol.Turbod.GetChangedOutputs:output_type -> turbodprotocol.GetChangedOutputsResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_internal_turbodprotocol_turbod_proto_init() }
func file_internal_turbodprotocol_turbod_proto_init() {
	if File_internal_turbodprotocol_turbod_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_turbodprotocol_turbod_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_internal_turbodprotocol_turbod_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_internal_turbodprotocol_turbod_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownRequest); i {
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
		file_internal_turbodprotocol_turbod_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownResponse); i {
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
		file_internal_turbodprotocol_turbod_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_internal_turbodprotocol_turbod_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_internal_turbodprotocol_turbod_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyOutputsWrittenRequest); i {
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
		file_internal_turbodprotocol_turbod_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyOutputsWrittenResponse); i {
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
		file_internal_turbodprotocol_turbod_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChangedOutputsRequest); i {
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
		file_internal_turbodprotocol_turbod_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChangedOutputsResponse); i {
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
		file_internal_turbodprotocol_turbod_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DaemonStatus); i {
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
			RawDescriptor: file_internal_turbodprotocol_turbod_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_turbodprotocol_turbod_proto_goTypes,
		DependencyIndexes: file_internal_turbodprotocol_turbod_proto_depIdxs,
		MessageInfos:      file_internal_turbodprotocol_turbod_proto_msgTypes,
	}.Build()
	File_internal_turbodprotocol_turbod_proto = out.File
	file_internal_turbodprotocol_turbod_proto_rawDesc = nil
	file_internal_turbodprotocol_turbod_proto_goTypes = nil
	file_internal_turbodprotocol_turbod_proto_depIdxs = nil
}
