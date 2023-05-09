// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: job.proto

package job

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

type IDsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *IDsReq) Reset() {
	*x = IDsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDsReq) ProtoMessage() {}

func (x *IDsReq) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDsReq.ProtoReflect.Descriptor instead.
func (*IDsReq) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{0}
}

func (x *IDsReq) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type BaseUUIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseUUIDResp) Reset() {
	*x = BaseUUIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseUUIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseUUIDResp) ProtoMessage() {}

func (x *BaseUUIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseUUIDResp.ProtoReflect.Descriptor instead.
func (*BaseUUIDResp) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{1}
}

func (x *BaseUUIDResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BaseUUIDResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type TaskInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CreatedAt      int64  `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt      int64  `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Status         uint32 `protobuf:"varint,4,opt,name=status,proto3" json:"status"`
	Name           string `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	TaskGroup      string `protobuf:"bytes,6,opt,name=task_group,json=taskGroup,proto3" json:"task_group"`
	CronExpression string `protobuf:"bytes,7,opt,name=cron_expression,json=cronExpression,proto3" json:"cron_expression"`
	Pattern        string `protobuf:"bytes,8,opt,name=pattern,proto3" json:"pattern"`
	Payload        string `protobuf:"bytes,9,opt,name=payload,proto3" json:"payload"`
}

func (x *TaskInfo) Reset() {
	*x = TaskInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInfo) ProtoMessage() {}

func (x *TaskInfo) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskInfo.ProtoReflect.Descriptor instead.
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{2}
}

func (x *TaskInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *TaskInfo) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *TaskInfo) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskInfo) GetTaskGroup() string {
	if x != nil {
		return x.TaskGroup
	}
	return ""
}

func (x *TaskInfo) GetCronExpression() string {
	if x != nil {
		return x.CronExpression
	}
	return ""
}

func (x *TaskInfo) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *TaskInfo) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type PageInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	PageSize uint64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
}

func (x *PageInfoReq) Reset() {
	*x = PageInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfoReq) ProtoMessage() {}

func (x *PageInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfoReq.ProtoReflect.Descriptor instead.
func (*PageInfoReq) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{3}
}

func (x *PageInfoReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageInfoReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type TaskListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	PageSize  uint64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	TaskGroup string `protobuf:"bytes,4,opt,name=task_group,json=taskGroup,proto3" json:"task_group"`
}

func (x *TaskListReq) Reset() {
	*x = TaskListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListReq) ProtoMessage() {}

func (x *TaskListReq) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListReq.ProtoReflect.Descriptor instead.
func (*TaskListReq) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{4}
}

func (x *TaskListReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *TaskListReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TaskListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskListReq) GetTaskGroup() string {
	if x != nil {
		return x.TaskGroup
	}
	return ""
}

type TaskLogListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint64         `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	Data  []*TaskLogInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
}

func (x *TaskLogListResp) Reset() {
	*x = TaskLogListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskLogListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskLogListResp) ProtoMessage() {}

func (x *TaskLogListResp) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskLogListResp.ProtoReflect.Descriptor instead.
func (*TaskLogListResp) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{5}
}

func (x *TaskLogListResp) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TaskLogListResp) GetData() []*TaskLogInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type TaskLogListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page"`
	PageSize uint64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	TaskId   uint64 `protobuf:"varint,3,opt,name=task_id,json=taskId,proto3" json:"task_id"`
	Result   uint32 `protobuf:"varint,4,opt,name=result,proto3" json:"result"`
}

func (x *TaskLogListReq) Reset() {
	*x = TaskLogListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskLogListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskLogListReq) ProtoMessage() {}

func (x *TaskLogListReq) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskLogListReq.ProtoReflect.Descriptor instead.
func (*TaskLogListReq) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{6}
}

func (x *TaskLogListReq) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *TaskLogListReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TaskLogListReq) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskLogListReq) GetResult() uint32 {
	if x != nil {
		return x.Result
	}
	return 0
}

// base message
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{7}
}

type UUIDsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids"`
}

func (x *UUIDsReq) Reset() {
	*x = UUIDsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDsReq) ProtoMessage() {}

func (x *UUIDsReq) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDsReq.ProtoReflect.Descriptor instead.
func (*UUIDsReq) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{8}
}

func (x *UUIDsReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UUIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *UUIDReq) Reset() {
	*x = UUIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUIDReq) ProtoMessage() {}

func (x *UUIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUIDReq.ProtoReflect.Descriptor instead.
func (*UUIDReq) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{9}
}

func (x *UUIDReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BaseIDResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseIDResp) Reset() {
	*x = BaseIDResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseIDResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseIDResp) ProtoMessage() {}

func (x *BaseIDResp) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseIDResp.ProtoReflect.Descriptor instead.
func (*BaseIDResp) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{10}
}

func (x *BaseIDResp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BaseIDResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type TaskListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint64      `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	Data  []*TaskInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data"`
}

func (x *TaskListResp) Reset() {
	*x = TaskListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskListResp) ProtoMessage() {}

func (x *TaskListResp) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskListResp.ProtoReflect.Descriptor instead.
func (*TaskListResp) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{11}
}

func (x *TaskListResp) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TaskListResp) GetData() []*TaskInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type IDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *IDReq) Reset() {
	*x = IDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDReq) ProtoMessage() {}

func (x *IDReq) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDReq.ProtoReflect.Descriptor instead.
func (*IDReq) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{12}
}

func (x *IDReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{13}
}

func (x *BaseResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type TaskLogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	StartedAt  int64  `protobuf:"varint,4,opt,name=started_at,json=startedAt,proto3" json:"started_at"`
	FinishedAt int64  `protobuf:"varint,5,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at"`
	Result     uint32 `protobuf:"varint,6,opt,name=result,proto3" json:"result"`
}

func (x *TaskLogInfo) Reset() {
	*x = TaskLogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskLogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskLogInfo) ProtoMessage() {}

func (x *TaskLogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskLogInfo.ProtoReflect.Descriptor instead.
func (*TaskLogInfo) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{14}
}

func (x *TaskLogInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskLogInfo) GetStartedAt() int64 {
	if x != nil {
		return x.StartedAt
	}
	return 0
}

func (x *TaskLogInfo) GetFinishedAt() int64 {
	if x != nil {
		return x.FinishedAt
	}
	return 0
}

func (x *TaskLogInfo) GetResult() uint32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_job_proto protoreflect.FileDescriptor

var file_job_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6a, 0x6f, 0x62,
	0x22, 0x1a, 0x0a, 0x06, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x30, 0x0a, 0x0c,
	0x42, 0x61, 0x73, 0x65, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x80,
	0x02, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x72, 0x6f, 0x6e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0x3e, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x71, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x4d, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x24, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x72, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x1c, 0x0a, 0x08, 0x55, 0x55, 0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x19,
	0x0a, 0x07, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x0a, 0x42, 0x61, 0x73,
	0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x47, 0x0a, 0x0c, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x08, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x75, 0x0a, 0x0b, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x92, 0x04, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x29, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x0a, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x0d, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x0f, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2a, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x0d, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0d,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a,
	0x0b, 0x67, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x28, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x0a, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x12, 0x10, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30, 0x0a, 0x0d, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x12, 0x10, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0d, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x0e, 0x67,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x13, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0a, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x12, 0x0b, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x6a, 0x6f, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_proto_rawDescOnce sync.Once
	file_job_proto_rawDescData = file_job_proto_rawDesc
)

func file_job_proto_rawDescGZIP() []byte {
	file_job_proto_rawDescOnce.Do(func() {
		file_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_proto_rawDescData)
	})
	return file_job_proto_rawDescData
}

var file_job_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_job_proto_goTypes = []interface{}{
	(*IDsReq)(nil),          // 0: job.IDsReq
	(*BaseUUIDResp)(nil),    // 1: job.BaseUUIDResp
	(*TaskInfo)(nil),        // 2: job.TaskInfo
	(*PageInfoReq)(nil),     // 3: job.PageInfoReq
	(*TaskListReq)(nil),     // 4: job.TaskListReq
	(*TaskLogListResp)(nil), // 5: job.TaskLogListResp
	(*TaskLogListReq)(nil),  // 6: job.TaskLogListReq
	(*Empty)(nil),           // 7: job.Empty
	(*UUIDsReq)(nil),        // 8: job.UUIDsReq
	(*UUIDReq)(nil),         // 9: job.UUIDReq
	(*BaseIDResp)(nil),      // 10: job.BaseIDResp
	(*TaskListResp)(nil),    // 11: job.TaskListResp
	(*IDReq)(nil),           // 12: job.IDReq
	(*BaseResp)(nil),        // 13: job.BaseResp
	(*TaskLogInfo)(nil),     // 14: job.TaskLogInfo
}
var file_job_proto_depIdxs = []int32{
	14, // 0: job.TaskLogListResp.data:type_name -> job.TaskLogInfo
	2,  // 1: job.TaskListResp.data:type_name -> job.TaskInfo
	7,  // 2: job.Job.initDatabase:input_type -> job.Empty
	2,  // 3: job.Job.createTask:input_type -> job.TaskInfo
	2,  // 4: job.Job.updateTask:input_type -> job.TaskInfo
	4,  // 5: job.Job.getTaskList:input_type -> job.TaskListReq
	12, // 6: job.Job.getTaskById:input_type -> job.IDReq
	0,  // 7: job.Job.deleteTask:input_type -> job.IDsReq
	14, // 8: job.Job.createTaskLog:input_type -> job.TaskLogInfo
	14, // 9: job.Job.updateTaskLog:input_type -> job.TaskLogInfo
	6,  // 10: job.Job.getTaskLogList:input_type -> job.TaskLogListReq
	12, // 11: job.Job.getTaskLogById:input_type -> job.IDReq
	0,  // 12: job.Job.deleteTaskLog:input_type -> job.IDsReq
	13, // 13: job.Job.initDatabase:output_type -> job.BaseResp
	10, // 14: job.Job.createTask:output_type -> job.BaseIDResp
	13, // 15: job.Job.updateTask:output_type -> job.BaseResp
	11, // 16: job.Job.getTaskList:output_type -> job.TaskListResp
	2,  // 17: job.Job.getTaskById:output_type -> job.TaskInfo
	13, // 18: job.Job.deleteTask:output_type -> job.BaseResp
	10, // 19: job.Job.createTaskLog:output_type -> job.BaseIDResp
	13, // 20: job.Job.updateTaskLog:output_type -> job.BaseResp
	5,  // 21: job.Job.getTaskLogList:output_type -> job.TaskLogListResp
	14, // 22: job.Job.getTaskLogById:output_type -> job.TaskLogInfo
	13, // 23: job.Job.deleteTaskLog:output_type -> job.BaseResp
	13, // [13:24] is the sub-list for method output_type
	2,  // [2:13] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_job_proto_init() }
func file_job_proto_init() {
	if File_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDsReq); i {
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
		file_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseUUIDResp); i {
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
		file_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskInfo); i {
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
		file_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfoReq); i {
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
		file_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListReq); i {
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
		file_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskLogListResp); i {
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
		file_job_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskLogListReq); i {
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
		file_job_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_job_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDsReq); i {
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
		file_job_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUIDReq); i {
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
		file_job_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseIDResp); i {
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
		file_job_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskListResp); i {
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
		file_job_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDReq); i {
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
		file_job_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_job_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskLogInfo); i {
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
			RawDescriptor: file_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_job_proto_goTypes,
		DependencyIndexes: file_job_proto_depIdxs,
		MessageInfos:      file_job_proto_msgTypes,
	}.Build()
	File_job_proto = out.File
	file_job_proto_rawDesc = nil
	file_job_proto_goTypes = nil
	file_job_proto_depIdxs = nil
}
