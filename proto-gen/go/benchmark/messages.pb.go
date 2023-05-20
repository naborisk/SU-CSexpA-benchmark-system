// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: benchmark/messages.proto

package benchmark

import (
	backend "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HttpMethod int32

const (
	HttpMethod_GET    HttpMethod = 0
	HttpMethod_POST   HttpMethod = 1
	HttpMethod_PUT    HttpMethod = 2
	HttpMethod_DELETE HttpMethod = 3
)

// Enum value maps for HttpMethod.
var (
	HttpMethod_name = map[int32]string{
		0: "GET",
		1: "POST",
		2: "PUT",
		3: "DELETE",
	}
	HttpMethod_value = map[string]int32{
		"GET":    0,
		"POST":   1,
		"PUT":    2,
		"DELETE": 3,
	}
)

func (x HttpMethod) Enum() *HttpMethod {
	p := new(HttpMethod)
	*p = x
	return p
}

func (x HttpMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HttpMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_benchmark_messages_proto_enumTypes[0].Descriptor()
}

func (HttpMethod) Type() protoreflect.EnumType {
	return &file_benchmark_messages_proto_enumTypes[0]
}

func (x HttpMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HttpMethod.Descriptor instead.
func (HttpMethod) EnumDescriptor() ([]byte, []int) {
	return file_benchmark_messages_proto_rawDescGZIP(), []int{0}
}

type ExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks   []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	GroupId string  `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"` // for logging
	Year    int32   `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *ExecuteRequest) Reset() {
	*x = ExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRequest) ProtoMessage() {}

func (x *ExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRequest.ProtoReflect.Descriptor instead.
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return file_benchmark_messages_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteRequest) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *ExecuteRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ExecuteRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request      *HttpRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	ThreadNum    int32        `protobuf:"varint,6,opt,name=thread_num,json=threadNum,proto3" json:"thread_num,omitempty"`          // the number of threads for a task
	AttemptCount int32        `protobuf:"varint,7,opt,name=attempt_count,json=attemptCount,proto3" json:"attempt_count,omitempty"` // the count of attempting for a task
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_benchmark_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Task) GetRequest() *HttpRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Task) GetThreadNum() int32 {
	if x != nil {
		return x.ThreadNum
	}
	return 0
}

func (x *Task) GetAttemptCount() int32 {
	if x != nil {
		return x.AttemptCount
	}
	return 0
}

type ExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok                bool           `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	ErrorMessage      *string        `protobuf:"bytes,2,opt,name=error_message,json=errorMessage,proto3,oneof" json:"error_message,omitempty"` // if ok is false, this field is set
	TimeElapsed       int64          `protobuf:"varint,3,opt,name=time_elapsed,json=timeElapsed,proto3" json:"time_elapsed,omitempty"`         // in milliseconds
	TotalRequests     int32          `protobuf:"varint,4,opt,name=total_requests,json=totalRequests,proto3" json:"total_requests,omitempty"`
	RequestsPerSecond int32          `protobuf:"varint,5,opt,name=requests_per_second,json=requestsPerSecond,proto3" json:"requests_per_second,omitempty"`
	Task              *Task          `protobuf:"bytes,6,opt,name=task,proto3" json:"task,omitempty"`
	Status            backend.Status `protobuf:"varint,7,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
}

func (x *ExecuteResponse) Reset() {
	*x = ExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResponse) ProtoMessage() {}

func (x *ExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResponse.ProtoReflect.Descriptor instead.
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return file_benchmark_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ExecuteResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ExecuteResponse) GetErrorMessage() string {
	if x != nil && x.ErrorMessage != nil {
		return *x.ErrorMessage
	}
	return ""
}

func (x *ExecuteResponse) GetTimeElapsed() int64 {
	if x != nil {
		return x.TimeElapsed
	}
	return 0
}

func (x *ExecuteResponse) GetTotalRequests() int32 {
	if x != nil {
		return x.TotalRequests
	}
	return 0
}

func (x *ExecuteResponse) GetRequestsPerSecond() int32 {
	if x != nil {
		return x.RequestsPerSecond
	}
	return 0
}

func (x *ExecuteResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *ExecuteResponse) GetStatus() backend.Status {
	if x != nil {
		return x.Status
	}
	return backend.Status(0)
}

type HttpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string     `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"` // e.g.) http://10.255.255.255/endpoint
	Method      HttpMethod `protobuf:"varint,2,opt,name=method,proto3,enum=HttpMethod" json:"method,omitempty"`
	ContentType string     `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	Body        string     `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *HttpRequest) Reset() {
	*x = HttpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_benchmark_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRequest) ProtoMessage() {}

func (x *HttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_benchmark_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRequest.ProtoReflect.Descriptor instead.
func (*HttpRequest) Descriptor() ([]byte, []int) {
	return file_benchmark_messages_proto_rawDescGZIP(), []int{3}
}

func (x *HttpRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HttpRequest) GetMethod() HttpMethod {
	if x != nil {
		return x.Method
	}
	return HttpMethod_GET
}

func (x *HttpRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *HttpRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_benchmark_messages_proto protoreflect.FileDescriptor

var file_benchmark_messages_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x22, 0x72, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x26, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x4e, 0x75,
	0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x93, 0x02, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x6c, 0x61,
	0x70, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65,
	0x45, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x2e,
	0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x19,
	0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x7b, 0x0a, 0x0b,
	0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x23, 0x0a,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x2a, 0x34, 0x0a, 0x0a, 0x48, 0x74, 0x74,
	0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55,
	0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x42,
	0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x68,
	0x6b, 0x69, 0x6c, 0x61, 0x62, 0x2f, 0x53, 0x55, 0x2d, 0x43, 0x53, 0x65, 0x78, 0x70, 0x41, 0x2d,
	0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x65,
	0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_benchmark_messages_proto_rawDescOnce sync.Once
	file_benchmark_messages_proto_rawDescData = file_benchmark_messages_proto_rawDesc
)

func file_benchmark_messages_proto_rawDescGZIP() []byte {
	file_benchmark_messages_proto_rawDescOnce.Do(func() {
		file_benchmark_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_benchmark_messages_proto_rawDescData)
	})
	return file_benchmark_messages_proto_rawDescData
}

var file_benchmark_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_benchmark_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_benchmark_messages_proto_goTypes = []interface{}{
	(HttpMethod)(0),         // 0: HttpMethod
	(*ExecuteRequest)(nil),  // 1: ExecuteRequest
	(*Task)(nil),            // 2: Task
	(*ExecuteResponse)(nil), // 3: ExecuteResponse
	(*HttpRequest)(nil),     // 4: HttpRequest
	(backend.Status)(0),     // 5: Status
}
var file_benchmark_messages_proto_depIdxs = []int32{
	2, // 0: ExecuteRequest.tasks:type_name -> Task
	4, // 1: Task.request:type_name -> HttpRequest
	2, // 2: ExecuteResponse.task:type_name -> Task
	5, // 3: ExecuteResponse.status:type_name -> Status
	0, // 4: HttpRequest.method:type_name -> HttpMethod
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_benchmark_messages_proto_init() }
func file_benchmark_messages_proto_init() {
	if File_benchmark_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_benchmark_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteRequest); i {
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
		file_benchmark_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_benchmark_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteResponse); i {
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
		file_benchmark_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpRequest); i {
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
	file_benchmark_messages_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_benchmark_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_benchmark_messages_proto_goTypes,
		DependencyIndexes: file_benchmark_messages_proto_depIdxs,
		EnumInfos:         file_benchmark_messages_proto_enumTypes,
		MessageInfos:      file_benchmark_messages_proto_msgTypes,
	}.Build()
	File_benchmark_messages_proto = out.File
	file_benchmark_messages_proto_rawDesc = nil
	file_benchmark_messages_proto_goTypes = nil
	file_benchmark_messages_proto_depIdxs = nil
}
