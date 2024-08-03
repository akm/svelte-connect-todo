// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: task/v1/task.proto

package taskv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type TaskStatus int32

const (
	TaskStatus_UNKNOWN_UNSPECIFIED TaskStatus = 0
	TaskStatus_TODO                TaskStatus = 1
	TaskStatus_DONE                TaskStatus = 2
)

// Enum value maps for TaskStatus.
var (
	TaskStatus_name = map[int32]string{
		0: "UNKNOWN_UNSPECIFIED",
		1: "TODO",
		2: "DONE",
	}
	TaskStatus_value = map[string]int32{
		"UNKNOWN_UNSPECIFIED": 0,
		"TODO":                1,
		"DONE":                2,
	}
)

func (x TaskStatus) Enum() *TaskStatus {
	p := new(TaskStatus)
	*p = x
	return p
}

func (x TaskStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_task_v1_task_proto_enumTypes[0].Descriptor()
}

func (TaskStatus) Type() protoreflect.EnumType {
	return &file_task_v1_task_proto_enumTypes[0]
}

func (x TaskStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskStatus.Descriptor instead.
func (TaskStatus) EnumDescriptor() ([]byte, []int) {
	return file_task_v1_task_proto_rawDescGZIP(), []int{0}
}

type ShowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ShowRequest) Reset() {
	*x = ShowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_v1_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowRequest) ProtoMessage() {}

func (x *ShowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_v1_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowRequest.ProtoReflect.Descriptor instead.
func (*ShowRequest) Descriptor() ([]byte, []int) {
	return file_task_v1_task_proto_rawDescGZIP(), []int{0}
}

func (x *ShowRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_v1_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_v1_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_task_v1_task_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status TaskStatus `protobuf:"varint,3,opt,name=status,proto3,enum=task.v1.TaskStatus" json:"status,omitempty"`
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_v1_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_v1_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_task_v1_task_proto_rawDescGZIP(), []int{2}
}

func (x *TaskResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskResponse) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_UNKNOWN_UNSPECIFIED
}

type TaskServiceCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// タスク名は1文字以上である必要があります。
	Name   string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status TaskStatus `protobuf:"varint,3,opt,name=status,proto3,enum=task.v1.TaskStatus" json:"status,omitempty"`
}

func (x *TaskServiceCreateRequest) Reset() {
	*x = TaskServiceCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_v1_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceCreateRequest) ProtoMessage() {}

func (x *TaskServiceCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_v1_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceCreateRequest.ProtoReflect.Descriptor instead.
func (*TaskServiceCreateRequest) Descriptor() ([]byte, []int) {
	return file_task_v1_task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskServiceCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskServiceCreateRequest) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_UNKNOWN_UNSPECIFIED
}

type TaskServiceUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// タスク名は1文字以上である必要があります。
	Name   string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status TaskStatus `protobuf:"varint,3,opt,name=status,proto3,enum=task.v1.TaskStatus" json:"status,omitempty"`
}

func (x *TaskServiceUpdateRequest) Reset() {
	*x = TaskServiceUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_v1_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceUpdateRequest) ProtoMessage() {}

func (x *TaskServiceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_v1_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceUpdateRequest.ProtoReflect.Descriptor instead.
func (*TaskServiceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_task_v1_task_proto_rawDescGZIP(), []int{4}
}

func (x *TaskServiceUpdateRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskServiceUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskServiceUpdateRequest) GetStatus() TaskStatus {
	if x != nil {
		return x.Status
	}
	return TaskStatus_UNKNOWN_UNSPECIFIED
}

type TaskServiceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *TaskServiceListRequest) Reset() {
	*x = TaskServiceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_v1_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceListRequest) ProtoMessage() {}

func (x *TaskServiceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_v1_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceListRequest.ProtoReflect.Descriptor instead.
func (*TaskServiceListRequest) Descriptor() ([]byte, []int) {
	return file_task_v1_task_proto_rawDescGZIP(), []int{5}
}

func (x *TaskServiceListRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *TaskServiceListRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type TaskServiceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint64          `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Items []*TaskResponse `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *TaskServiceListResponse) Reset() {
	*x = TaskServiceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_task_v1_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskServiceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskServiceListResponse) ProtoMessage() {}

func (x *TaskServiceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_v1_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskServiceListResponse.ProtoReflect.Descriptor instead.
func (*TaskServiceListResponse) Descriptor() ([]byte, []int) {
	return file_task_v1_task_proto_rawDescGZIP(), []int{6}
}

func (x *TaskServiceListResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TaskServiceListResponse) GetItems() []*TaskResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_task_v1_task_proto protoreflect.FileDescriptor

var file_task_v1_task_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0b, 0x53, 0x68,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5f, 0x0a, 0x0c, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x64, 0x0a, 0x18, 0x54,
	0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x74, 0x0a, 0x18, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x46, 0x0a, 0x16, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x5c, 0x0a, 0x17, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x39, 0x0a,
	0x0a, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x4f, 0x44, 0x4f, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x32, 0xd8, 0x02, 0x0a, 0x0b, 0x54, 0x61, 0x73,
	0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1f, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x14, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x1b, 0x5a, 0x19, 0x61, 0x70, 0x69, 0x73, 0x76, 0x72, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_v1_task_proto_rawDescOnce sync.Once
	file_task_v1_task_proto_rawDescData = file_task_v1_task_proto_rawDesc
)

func file_task_v1_task_proto_rawDescGZIP() []byte {
	file_task_v1_task_proto_rawDescOnce.Do(func() {
		file_task_v1_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_v1_task_proto_rawDescData)
	})
	return file_task_v1_task_proto_rawDescData
}

var file_task_v1_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_task_v1_task_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_task_v1_task_proto_goTypes = []any{
	(TaskStatus)(0),                  // 0: task.v1.TaskStatus
	(*ShowRequest)(nil),              // 1: task.v1.ShowRequest
	(*DeleteRequest)(nil),            // 2: task.v1.DeleteRequest
	(*TaskResponse)(nil),             // 3: task.v1.TaskResponse
	(*TaskServiceCreateRequest)(nil), // 4: task.v1.TaskServiceCreateRequest
	(*TaskServiceUpdateRequest)(nil), // 5: task.v1.TaskServiceUpdateRequest
	(*TaskServiceListRequest)(nil),   // 6: task.v1.TaskServiceListRequest
	(*TaskServiceListResponse)(nil),  // 7: task.v1.TaskServiceListResponse
}
var file_task_v1_task_proto_depIdxs = []int32{
	0, // 0: task.v1.TaskResponse.status:type_name -> task.v1.TaskStatus
	0, // 1: task.v1.TaskServiceCreateRequest.status:type_name -> task.v1.TaskStatus
	0, // 2: task.v1.TaskServiceUpdateRequest.status:type_name -> task.v1.TaskStatus
	3, // 3: task.v1.TaskServiceListResponse.items:type_name -> task.v1.TaskResponse
	6, // 4: task.v1.TaskService.List:input_type -> task.v1.TaskServiceListRequest
	1, // 5: task.v1.TaskService.Show:input_type -> task.v1.ShowRequest
	4, // 6: task.v1.TaskService.Create:input_type -> task.v1.TaskServiceCreateRequest
	5, // 7: task.v1.TaskService.Update:input_type -> task.v1.TaskServiceUpdateRequest
	2, // 8: task.v1.TaskService.Delete:input_type -> task.v1.DeleteRequest
	7, // 9: task.v1.TaskService.List:output_type -> task.v1.TaskServiceListResponse
	3, // 10: task.v1.TaskService.Show:output_type -> task.v1.TaskResponse
	3, // 11: task.v1.TaskService.Create:output_type -> task.v1.TaskResponse
	3, // 12: task.v1.TaskService.Update:output_type -> task.v1.TaskResponse
	3, // 13: task.v1.TaskService.Delete:output_type -> task.v1.TaskResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_task_v1_task_proto_init() }
func file_task_v1_task_proto_init() {
	if File_task_v1_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_task_v1_task_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ShowRequest); i {
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
		file_task_v1_task_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRequest); i {
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
		file_task_v1_task_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TaskResponse); i {
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
		file_task_v1_task_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TaskServiceCreateRequest); i {
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
		file_task_v1_task_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*TaskServiceUpdateRequest); i {
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
		file_task_v1_task_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*TaskServiceListRequest); i {
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
		file_task_v1_task_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*TaskServiceListResponse); i {
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
			RawDescriptor: file_task_v1_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_v1_task_proto_goTypes,
		DependencyIndexes: file_task_v1_task_proto_depIdxs,
		EnumInfos:         file_task_v1_task_proto_enumTypes,
		MessageInfos:      file_task_v1_task_proto_msgTypes,
	}.Build()
	File_task_v1_task_proto = out.File
	file_task_v1_task_proto_rawDesc = nil
	file_task_v1_task_proto_goTypes = nil
	file_task_v1_task_proto_depIdxs = nil
}
