// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: backend/messages.proto

package backend

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PostLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // it means group_id
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *PostLoginRequest) Reset() {
	*x = PostLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostLoginRequest) ProtoMessage() {}

func (x *PostLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostLoginRequest.ProtoReflect.Descriptor instead.
func (*PostLoginRequest) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PostLoginRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type PostLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *Group `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *PostLoginResponse) Reset() {
	*x = PostLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostLoginResponse) ProtoMessage() {}

func (x *PostLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostLoginResponse.ProtoReflect.Descriptor instead.
func (*PostLoginResponse) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{1}
}

func (x *PostLoginResponse) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *PostLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type PostSubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddr    string `protobuf:"bytes,1,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	ContestId int32  `protobuf:"varint,2,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
}

func (x *PostSubmitRequest) Reset() {
	*x = PostSubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostSubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostSubmitRequest) ProtoMessage() {}

func (x *PostSubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostSubmitRequest.ProtoReflect.Descriptor instead.
func (*PostSubmitRequest) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PostSubmitRequest) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *PostSubmitRequest) GetContestId() int32 {
	if x != nil {
		return x.ContestId
	}
	return 0
}

type PostSubmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IpAddr     string                 `protobuf:"bytes,2,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	ContestId  int32                  `protobuf:"varint,3,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	SubmitedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=submited_at,json=submitedAt,proto3" json:"submited_at,omitempty"`
}

func (x *PostSubmitResponse) Reset() {
	*x = PostSubmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostSubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostSubmitResponse) ProtoMessage() {}

func (x *PostSubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostSubmitResponse.ProtoReflect.Descriptor instead.
func (*PostSubmitResponse) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{3}
}

func (x *PostSubmitResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PostSubmitResponse) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *PostSubmitResponse) GetContestId() int32 {
	if x != nil {
		return x.ContestId
	}
	return 0
}

func (x *PostSubmitResponse) GetSubmitedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SubmitedAt
	}
	return nil
}

type GetSubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubmitId int32 `protobuf:"varint,1,opt,name=submit_id,json=submitId,proto3" json:"submit_id,omitempty"`
}

func (x *GetSubmitRequest) Reset() {
	*x = GetSubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubmitRequest) ProtoMessage() {}

func (x *GetSubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubmitRequest.ProtoReflect.Descriptor instead.
func (*GetSubmitRequest) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GetSubmitRequest) GetSubmitId() int32 {
	if x != nil {
		return x.SubmitId
	}
	return 0
}

type GetSubmitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submit *Submit `protobuf:"bytes,1,opt,name=submit,proto3" json:"submit,omitempty"`
}

func (x *GetSubmitResponse) Reset() {
	*x = GetSubmitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubmitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubmitResponse) ProtoMessage() {}

func (x *GetSubmitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubmitResponse.ProtoReflect.Descriptor instead.
func (*GetSubmitResponse) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{5}
}

func (x *GetSubmitResponse) GetSubmit() *Submit {
	if x != nil {
		return x.Submit
	}
	return nil
}

type GetRankingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year         int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	ContainGuest bool  `protobuf:"varint,2,opt,name=contain_guest,json=containGuest,proto3" json:"contain_guest,omitempty"` // if it is true, return ranking which includes guests
}

func (x *GetRankingRequest) Reset() {
	*x = GetRankingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankingRequest) ProtoMessage() {}

func (x *GetRankingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankingRequest.ProtoReflect.Descriptor instead.
func (*GetRankingRequest) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{6}
}

func (x *GetRankingRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *GetRankingRequest) GetContainGuest() bool {
	if x != nil {
		return x.ContainGuest
	}
	return false
}

type GetRankingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*GetRankingResponse_Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *GetRankingResponse) Reset() {
	*x = GetRankingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankingResponse) ProtoMessage() {}

func (x *GetRankingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankingResponse.ProtoReflect.Descriptor instead.
func (*GetRankingResponse) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{7}
}

func (x *GetRankingResponse) GetRecords() []*GetRankingResponse_Record {
	if x != nil {
		return x.Records
	}
	return nil
}

type GetGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GetGroupRequest) Reset() {
	*x = GetGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupRequest) ProtoMessage() {}

func (x *GetGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupRequest.ProtoReflect.Descriptor instead.
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{8}
}

func (x *GetGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GetGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*GetGroupResponse_GroupInfo `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *GetGroupResponse) Reset() {
	*x = GetGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupResponse) ProtoMessage() {}

func (x *GetGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupResponse.ProtoReflect.Descriptor instead.
func (*GetGroupResponse) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{9}
}

func (x *GetGroupResponse) GetGroups() []*GetGroupResponse_GroupInfo {
	if x != nil {
		return x.Groups
	}
	return nil
}

type PingUnaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *PingUnaryRequest) Reset() {
	*x = PingUnaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingUnaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingUnaryRequest) ProtoMessage() {}

func (x *PingUnaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingUnaryRequest.ProtoReflect.Descriptor instead.
func (*PingUnaryRequest) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{10}
}

func (x *PingUnaryRequest) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type PingUnaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *PingUnaryResponse) Reset() {
	*x = PingUnaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingUnaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingUnaryResponse) ProtoMessage() {}

func (x *PingUnaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingUnaryResponse.ProtoReflect.Descriptor instead.
func (*PingUnaryResponse) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{11}
}

func (x *PingUnaryResponse) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type PingServerSideStreamingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,2,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *PingServerSideStreamingRequest) Reset() {
	*x = PingServerSideStreamingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingServerSideStreamingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingServerSideStreamingRequest) ProtoMessage() {}

func (x *PingServerSideStreamingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingServerSideStreamingRequest.ProtoReflect.Descriptor instead.
func (*PingServerSideStreamingRequest) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{12}
}

func (x *PingServerSideStreamingRequest) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type PingServerSideStreamingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,2,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *PingServerSideStreamingResponse) Reset() {
	*x = PingServerSideStreamingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingServerSideStreamingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingServerSideStreamingResponse) ProtoMessage() {}

func (x *PingServerSideStreamingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingServerSideStreamingResponse.ProtoReflect.Descriptor instead.
func (*PingServerSideStreamingResponse) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{13}
}

func (x *PingServerSideStreamingResponse) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type GetRankingResponse_Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank  int32  `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Group *Group `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *GetRankingResponse_Record) Reset() {
	*x = GetRankingResponse_Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRankingResponse_Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRankingResponse_Record) ProtoMessage() {}

func (x *GetRankingResponse_Record) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRankingResponse_Record.ProtoReflect.Descriptor instead.
func (*GetRankingResponse_Record) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{7, 0}
}

func (x *GetRankingResponse_Record) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *GetRankingResponse_Record) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

type GetGroupResponse_GroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group   *Group    `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Submits []*Submit `protobuf:"bytes,2,rep,name=submits,proto3" json:"submits,omitempty"`
}

func (x *GetGroupResponse_GroupInfo) Reset() {
	*x = GetGroupResponse_GroupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_messages_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupResponse_GroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupResponse_GroupInfo) ProtoMessage() {}

func (x *GetGroupResponse_GroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_backend_messages_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupResponse_GroupInfo.ProtoReflect.Descriptor instead.
func (*GetGroupResponse_GroupInfo) Descriptor() ([]byte, []int) {
	return file_backend_messages_proto_rawDescGZIP(), []int{9, 0}
}

func (x *GetGroupResponse_GroupInfo) GetGroup() *Group {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *GetGroupResponse_GroupInfo) GetSubmits() []*Submit {
	if x != nil {
		return x.Submits
	}
	return nil
}

var File_backend_messages_proto protoreflect.FileDescriptor

var file_backend_messages_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x47, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4b, 0x0a, 0x11, 0x50,
	0x6f, 0x73, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x99, 0x01, 0x0a, 0x12, 0x50, 0x6f, 0x73,
	0x74, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x22, 0x4c, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x5f,
	0x67, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x07, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x1a, 0x3a, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x22, 0x95, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x1a, 0x4c, 0x0a, 0x09, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x73, 0x22, 0x26, 0x0a, 0x10, 0x50, 0x69, 0x6e, 0x67,
	0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67,
	0x22, 0x27, 0x0a, 0x11, 0x50, 0x69, 0x6e, 0x67, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x34, 0x0a, 0x1e, 0x50, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22,
	0x35, 0x0a, 0x1f, 0x50, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_messages_proto_rawDescOnce sync.Once
	file_backend_messages_proto_rawDescData = file_backend_messages_proto_rawDesc
)

func file_backend_messages_proto_rawDescGZIP() []byte {
	file_backend_messages_proto_rawDescOnce.Do(func() {
		file_backend_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_messages_proto_rawDescData)
	})
	return file_backend_messages_proto_rawDescData
}

var file_backend_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_backend_messages_proto_goTypes = []interface{}{
	(*PostLoginRequest)(nil),                // 0: PostLoginRequest
	(*PostLoginResponse)(nil),               // 1: PostLoginResponse
	(*PostSubmitRequest)(nil),               // 2: PostSubmitRequest
	(*PostSubmitResponse)(nil),              // 3: PostSubmitResponse
	(*GetSubmitRequest)(nil),                // 4: GetSubmitRequest
	(*GetSubmitResponse)(nil),               // 5: GetSubmitResponse
	(*GetRankingRequest)(nil),               // 6: GetRankingRequest
	(*GetRankingResponse)(nil),              // 7: GetRankingResponse
	(*GetGroupRequest)(nil),                 // 8: GetGroupRequest
	(*GetGroupResponse)(nil),                // 9: GetGroupResponse
	(*PingUnaryRequest)(nil),                // 10: PingUnaryRequest
	(*PingUnaryResponse)(nil),               // 11: PingUnaryResponse
	(*PingServerSideStreamingRequest)(nil),  // 12: PingServerSideStreamingRequest
	(*PingServerSideStreamingResponse)(nil), // 13: PingServerSideStreamingResponse
	(*GetRankingResponse_Record)(nil),       // 14: GetRankingResponse.Record
	(*GetGroupResponse_GroupInfo)(nil),      // 15: GetGroupResponse.GroupInfo
	(*Group)(nil),                           // 16: Group
	(*timestamppb.Timestamp)(nil),           // 17: google.protobuf.Timestamp
	(*Submit)(nil),                          // 18: Submit
}
var file_backend_messages_proto_depIdxs = []int32{
	16, // 0: PostLoginResponse.group:type_name -> Group
	17, // 1: PostSubmitResponse.submited_at:type_name -> google.protobuf.Timestamp
	18, // 2: GetSubmitResponse.submit:type_name -> Submit
	14, // 3: GetRankingResponse.records:type_name -> GetRankingResponse.Record
	15, // 4: GetGroupResponse.groups:type_name -> GetGroupResponse.GroupInfo
	16, // 5: GetRankingResponse.Record.group:type_name -> Group
	16, // 6: GetGroupResponse.GroupInfo.group:type_name -> Group
	18, // 7: GetGroupResponse.GroupInfo.submits:type_name -> Submit
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_backend_messages_proto_init() }
func file_backend_messages_proto_init() {
	if File_backend_messages_proto != nil {
		return
	}
	file_backend_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_backend_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostLoginRequest); i {
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
		file_backend_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostLoginResponse); i {
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
		file_backend_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostSubmitRequest); i {
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
		file_backend_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostSubmitResponse); i {
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
		file_backend_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubmitRequest); i {
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
		file_backend_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubmitResponse); i {
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
		file_backend_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankingRequest); i {
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
		file_backend_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankingResponse); i {
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
		file_backend_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupRequest); i {
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
		file_backend_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupResponse); i {
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
		file_backend_messages_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingUnaryRequest); i {
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
		file_backend_messages_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingUnaryResponse); i {
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
		file_backend_messages_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingServerSideStreamingRequest); i {
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
		file_backend_messages_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingServerSideStreamingResponse); i {
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
		file_backend_messages_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRankingResponse_Record); i {
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
		file_backend_messages_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupResponse_GroupInfo); i {
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
			RawDescriptor: file_backend_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_backend_messages_proto_goTypes,
		DependencyIndexes: file_backend_messages_proto_depIdxs,
		MessageInfos:      file_backend_messages_proto_msgTypes,
	}.Build()
	File_backend_messages_proto = out.File
	file_backend_messages_proto_rawDesc = nil
	file_backend_messages_proto_goTypes = nil
	file_backend_messages_proto_depIdxs = nil
}
