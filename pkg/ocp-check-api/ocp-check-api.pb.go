// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: api/ocp-check-api/ocp-check-api.proto

package ocp_check_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ListChecksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListChecksRequest) Reset() {
	*x = ListChecksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChecksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChecksRequest) ProtoMessage() {}

func (x *ListChecksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChecksRequest.ProtoReflect.Descriptor instead.
func (*ListChecksRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_check_api_ocp_check_api_proto_rawDescGZIP(), []int{0}
}

func (x *ListChecksRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListChecksRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListChecksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Checks []*Check `protobuf:"bytes,1,rep,name=checks,proto3" json:"checks,omitempty"`
}

func (x *ListChecksResponse) Reset() {
	*x = ListChecksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChecksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChecksResponse) ProtoMessage() {}

func (x *ListChecksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChecksResponse.ProtoReflect.Descriptor instead.
func (*ListChecksResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_check_api_ocp_check_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListChecksResponse) GetChecks() []*Check {
	if x != nil {
		return x.Checks
	}
	return nil
}

type CreateCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SolutionID uint64 `protobuf:"varint,1,opt,name=solutionID,proto3" json:"solutionID,omitempty"`
	TestID     uint64 `protobuf:"varint,2,opt,name=testID,proto3" json:"testID,omitempty"`
	RunnerID   uint64 `protobuf:"varint,3,opt,name=runnerID,proto3" json:"runnerID,omitempty"`
	Success    bool   `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateCheckRequest) Reset() {
	*x = CreateCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCheckRequest) ProtoMessage() {}

func (x *CreateCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCheckRequest.ProtoReflect.Descriptor instead.
func (*CreateCheckRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_check_api_ocp_check_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCheckRequest) GetSolutionID() uint64 {
	if x != nil {
		return x.SolutionID
	}
	return 0
}

func (x *CreateCheckRequest) GetTestID() uint64 {
	if x != nil {
		return x.TestID
	}
	return 0
}

func (x *CreateCheckRequest) GetRunnerID() uint64 {
	if x != nil {
		return x.RunnerID
	}
	return 0
}

func (x *CreateCheckRequest) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CreateCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckId uint64 `protobuf:"varint,1,opt,name=check_id,json=checkId,proto3" json:"check_id,omitempty"`
}

func (x *CreateCheckResponse) Reset() {
	*x = CreateCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCheckResponse) ProtoMessage() {}

func (x *CreateCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCheckResponse.ProtoReflect.Descriptor instead.
func (*CreateCheckResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_check_api_ocp_check_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCheckResponse) GetCheckId() uint64 {
	if x != nil {
		return x.CheckId
	}
	return 0
}

type RemoveCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckId uint64 `protobuf:"varint,1,opt,name=check_id,json=checkId,proto3" json:"check_id,omitempty"`
}

func (x *RemoveCheckRequest) Reset() {
	*x = RemoveCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCheckRequest) ProtoMessage() {}

func (x *RemoveCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCheckRequest.ProtoReflect.Descriptor instead.
func (*RemoveCheckRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_check_api_ocp_check_api_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveCheckRequest) GetCheckId() uint64 {
	if x != nil {
		return x.CheckId
	}
	return 0
}

type RemoveCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *RemoveCheckResponse) Reset() {
	*x = RemoveCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCheckResponse) ProtoMessage() {}

func (x *RemoveCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCheckResponse.ProtoReflect.Descriptor instead.
func (*RemoveCheckResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_check_api_ocp_check_api_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveCheckResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type DescribeCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckId uint64 `protobuf:"varint,1,opt,name=check_id,json=checkId,proto3" json:"check_id,omitempty"`
}

func (x *DescribeCheckRequest) Reset() {
	*x = DescribeCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCheckRequest) ProtoMessage() {}

func (x *DescribeCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCheckRequest.ProtoReflect.Descriptor instead.
func (*DescribeCheckRequest) Descriptor() ([]byte, []int) {
	return file_api_ocp_check_api_ocp_check_api_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeCheckRequest) GetCheckId() uint64 {
	if x != nil {
		return x.CheckId
	}
	return 0
}

type DescribeCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Check *Check `protobuf:"bytes,1,opt,name=check,proto3" json:"check,omitempty"`
}

func (x *DescribeCheckResponse) Reset() {
	*x = DescribeCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCheckResponse) ProtoMessage() {}

func (x *DescribeCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCheckResponse.ProtoReflect.Descriptor instead.
func (*DescribeCheckResponse) Descriptor() ([]byte, []int) {
	return file_api_ocp_check_api_ocp_check_api_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeCheckResponse) GetCheck() *Check {
	if x != nil {
		return x.Check
	}
	return nil
}

// Описание структуры "проверка"
type Check struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SolutionID uint64 `protobuf:"varint,2,opt,name=solutionID,proto3" json:"solutionID,omitempty"`
	TestID     uint64 `protobuf:"varint,3,opt,name=testID,proto3" json:"testID,omitempty"`
	RunnerID   uint64 `protobuf:"varint,4,opt,name=runnerID,proto3" json:"runnerID,omitempty"`
	Success    bool   `protobuf:"varint,5,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Check) Reset() {
	*x = Check{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Check) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Check) ProtoMessage() {}

func (x *Check) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_check_api_ocp_check_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Check.ProtoReflect.Descriptor instead.
func (*Check) Descriptor() ([]byte, []int) {
	return file_api_ocp_check_api_ocp_check_api_proto_rawDescGZIP(), []int{8}
}

func (x *Check) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Check) GetSolutionID() uint64 {
	if x != nil {
		return x.SolutionID
	}
	return 0
}

func (x *Check) GetTestID() uint64 {
	if x != nil {
		return x.TestID
	}
	return 0
}

func (x *Check) GetRunnerID() uint64 {
	if x != nil {
		return x.RunnerID
	}
	return 0
}

func (x *Check) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_ocp_check_api_ocp_check_api_proto protoreflect.FileDescriptor

var file_api_ocp_check_api_ocp_check_api_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x42, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x06, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x22, 0x82,
	0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x22,
	0x2b, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x3a, 0x0a, 0x14,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x85, 0x01,
	0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xc2, 0x03, 0x0a, 0x0b, 0x4f, 0x63, 0x70, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x41, 0x70, 0x69, 0x12, 0x62, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x12, 0x20, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09,
	0x12, 0x07, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x76, 0x0a, 0x0d, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x23, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2f, 0x7b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x65, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x22,
	0x07, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x12, 0x70, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2f,
	0x7b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f,
	0x6f, 0x63, 0x70, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x3b,
	0x6f, 0x63, 0x70, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ocp_check_api_ocp_check_api_proto_rawDescOnce sync.Once
	file_api_ocp_check_api_ocp_check_api_proto_rawDescData = file_api_ocp_check_api_ocp_check_api_proto_rawDesc
)

func file_api_ocp_check_api_ocp_check_api_proto_rawDescGZIP() []byte {
	file_api_ocp_check_api_ocp_check_api_proto_rawDescOnce.Do(func() {
		file_api_ocp_check_api_ocp_check_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ocp_check_api_ocp_check_api_proto_rawDescData)
	})
	return file_api_ocp_check_api_ocp_check_api_proto_rawDescData
}

var file_api_ocp_check_api_ocp_check_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_ocp_check_api_ocp_check_api_proto_goTypes = []interface{}{
	(*ListChecksRequest)(nil),     // 0: ocp.check.api.ListChecksRequest
	(*ListChecksResponse)(nil),    // 1: ocp.check.api.ListChecksResponse
	(*CreateCheckRequest)(nil),    // 2: ocp.check.api.CreateCheckRequest
	(*CreateCheckResponse)(nil),   // 3: ocp.check.api.CreateCheckResponse
	(*RemoveCheckRequest)(nil),    // 4: ocp.check.api.RemoveCheckRequest
	(*RemoveCheckResponse)(nil),   // 5: ocp.check.api.RemoveCheckResponse
	(*DescribeCheckRequest)(nil),  // 6: ocp.check.api.DescribeCheckRequest
	(*DescribeCheckResponse)(nil), // 7: ocp.check.api.DescribeCheckResponse
	(*Check)(nil),                 // 8: ocp.check.api.Check
}
var file_api_ocp_check_api_ocp_check_api_proto_depIdxs = []int32{
	8, // 0: ocp.check.api.ListChecksResponse.checks:type_name -> ocp.check.api.Check
	8, // 1: ocp.check.api.DescribeCheckResponse.check:type_name -> ocp.check.api.Check
	0, // 2: ocp.check.api.OcpCheckApi.ListChecks:input_type -> ocp.check.api.ListChecksRequest
	6, // 3: ocp.check.api.OcpCheckApi.DescribeCheck:input_type -> ocp.check.api.DescribeCheckRequest
	2, // 4: ocp.check.api.OcpCheckApi.CreateCheck:input_type -> ocp.check.api.CreateCheckRequest
	4, // 5: ocp.check.api.OcpCheckApi.RemoveCheck:input_type -> ocp.check.api.RemoveCheckRequest
	1, // 6: ocp.check.api.OcpCheckApi.ListChecks:output_type -> ocp.check.api.ListChecksResponse
	7, // 7: ocp.check.api.OcpCheckApi.DescribeCheck:output_type -> ocp.check.api.DescribeCheckResponse
	3, // 8: ocp.check.api.OcpCheckApi.CreateCheck:output_type -> ocp.check.api.CreateCheckResponse
	5, // 9: ocp.check.api.OcpCheckApi.RemoveCheck:output_type -> ocp.check.api.RemoveCheckResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_ocp_check_api_ocp_check_api_proto_init() }
func file_api_ocp_check_api_ocp_check_api_proto_init() {
	if File_api_ocp_check_api_ocp_check_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_check_api_ocp_check_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChecksRequest); i {
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
		file_api_ocp_check_api_ocp_check_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChecksResponse); i {
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
		file_api_ocp_check_api_ocp_check_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCheckRequest); i {
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
		file_api_ocp_check_api_ocp_check_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCheckResponse); i {
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
		file_api_ocp_check_api_ocp_check_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCheckRequest); i {
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
		file_api_ocp_check_api_ocp_check_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCheckResponse); i {
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
		file_api_ocp_check_api_ocp_check_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeCheckRequest); i {
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
		file_api_ocp_check_api_ocp_check_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeCheckResponse); i {
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
		file_api_ocp_check_api_ocp_check_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Check); i {
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
			RawDescriptor: file_api_ocp_check_api_ocp_check_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ocp_check_api_ocp_check_api_proto_goTypes,
		DependencyIndexes: file_api_ocp_check_api_ocp_check_api_proto_depIdxs,
		MessageInfos:      file_api_ocp_check_api_ocp_check_api_proto_msgTypes,
	}.Build()
	File_api_ocp_check_api_ocp_check_api_proto = out.File
	file_api_ocp_check_api_ocp_check_api_proto_rawDesc = nil
	file_api_ocp_check_api_ocp_check_api_proto_goTypes = nil
	file_api_ocp_check_api_ocp_check_api_proto_depIdxs = nil
}
