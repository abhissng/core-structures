// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/discovery/discovery.proto

package discoveryv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DiscoveryMessage_Status int32

const (
	DiscoveryMessage_STATUS_UNSPECIFIED DiscoveryMessage_Status = 0
	DiscoveryMessage_STATUS_PENDING     DiscoveryMessage_Status = 1
	DiscoveryMessage_STATUS_COMPLETED   DiscoveryMessage_Status = 2
	DiscoveryMessage_STATUS_FAILED      DiscoveryMessage_Status = 3
)

// Enum value maps for DiscoveryMessage_Status.
var (
	DiscoveryMessage_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_PENDING",
		2: "STATUS_COMPLETED",
		3: "STATUS_FAILED",
	}
	DiscoveryMessage_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_PENDING":     1,
		"STATUS_COMPLETED":   2,
		"STATUS_FAILED":      3,
	}
)

func (x DiscoveryMessage_Status) Enum() *DiscoveryMessage_Status {
	p := new(DiscoveryMessage_Status)
	*p = x
	return p
}

func (x DiscoveryMessage_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiscoveryMessage_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_discovery_discovery_proto_enumTypes[0].Descriptor()
}

func (DiscoveryMessage_Status) Type() protoreflect.EnumType {
	return &file_proto_discovery_discovery_proto_enumTypes[0]
}

func (x DiscoveryMessage_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiscoveryMessage_Status.Descriptor instead.
func (DiscoveryMessage_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_discovery_discovery_proto_rawDescGZIP(), []int{1, 0}
}

type DiscoveryMessage_Action int32

const (
	DiscoveryMessage_ACTION_UNSPECIFIED DiscoveryMessage_Action = 0
	DiscoveryMessage_ACTION_EXECUTE     DiscoveryMessage_Action = 1
	DiscoveryMessage_ACTION_ROLLBACK    DiscoveryMessage_Action = 2
)

// Enum value maps for DiscoveryMessage_Action.
var (
	DiscoveryMessage_Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "ACTION_EXECUTE",
		2: "ACTION_ROLLBACK",
	}
	DiscoveryMessage_Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"ACTION_EXECUTE":     1,
		"ACTION_ROLLBACK":    2,
	}
)

func (x DiscoveryMessage_Action) Enum() *DiscoveryMessage_Action {
	p := new(DiscoveryMessage_Action)
	*p = x
	return p
}

func (x DiscoveryMessage_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiscoveryMessage_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_discovery_discovery_proto_enumTypes[1].Descriptor()
}

func (DiscoveryMessage_Action) Type() protoreflect.EnumType {
	return &file_proto_discovery_discovery_proto_enumTypes[1]
}

func (x DiscoveryMessage_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiscoveryMessage_Action.Descriptor instead.
func (DiscoveryMessage_Action) EnumDescriptor() ([]byte, []int) {
	return file_proto_discovery_discovery_proto_rawDescGZIP(), []int{1, 1}
}

// Core message structure
type Core struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *Core_UserInformation  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Primary       *Core_PrimaryInfo      `protobuf:"bytes,2,opt,name=primary,proto3" json:"primary,omitempty"`
	Input         *Core_InputInfo        `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	Error         *anypb.Any             `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Core) Reset() {
	*x = Core{}
	mi := &file_proto_discovery_discovery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Core) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Core) ProtoMessage() {}

func (x *Core) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discovery_discovery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Core.ProtoReflect.Descriptor instead.
func (*Core) Descriptor() ([]byte, []int) {
	return file_proto_discovery_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *Core) GetUser() *Core_UserInformation {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Core) GetPrimary() *Core_PrimaryInfo {
	if x != nil {
		return x.Primary
	}
	return nil
}

func (x *Core) GetInput() *Core_InputInfo {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Core) GetError() *anypb.Any {
	if x != nil {
		return x.Error
	}
	return nil
}

// Generic Message structure
type DiscoveryMessage struct {
	state          protoimpl.MessageState  `protogen:"open.v1"`
	CorrelationId  string                  `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	RequestId      string                  `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Payload        *anypb.Any              `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Status         DiscoveryMessage_Status `protobuf:"varint,4,opt,name=status,proto3,enum=discovery.v1.DiscoveryMessage_Status" json:"status,omitempty"`
	Action         DiscoveryMessage_Action `protobuf:"varint,5,opt,name=action,proto3,enum=discovery.v1.DiscoveryMessage_Action" json:"action,omitempty"`
	Error          *anypb.Any              `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	Timestamp      *timestamppb.Timestamp  `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CurrentService string                  `protobuf:"bytes,8,opt,name=current_service,json=currentService,proto3" json:"current_service,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DiscoveryMessage) Reset() {
	*x = DiscoveryMessage{}
	mi := &file_proto_discovery_discovery_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoveryMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryMessage) ProtoMessage() {}

func (x *DiscoveryMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discovery_discovery_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryMessage.ProtoReflect.Descriptor instead.
func (*DiscoveryMessage) Descriptor() ([]byte, []int) {
	return file_proto_discovery_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *DiscoveryMessage) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *DiscoveryMessage) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *DiscoveryMessage) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *DiscoveryMessage) GetStatus() DiscoveryMessage_Status {
	if x != nil {
		return x.Status
	}
	return DiscoveryMessage_STATUS_UNSPECIFIED
}

func (x *DiscoveryMessage) GetAction() DiscoveryMessage_Action {
	if x != nil {
		return x.Action
	}
	return DiscoveryMessage_ACTION_UNSPECIFIED
}

func (x *DiscoveryMessage) GetError() *anypb.Any {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *DiscoveryMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *DiscoveryMessage) GetCurrentService() string {
	if x != nil {
		return x.CurrentService
	}
	return ""
}

type Core_UserInformation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"` // Add other user fields as needed
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Core_UserInformation) Reset() {
	*x = Core_UserInformation{}
	mi := &file_proto_discovery_discovery_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Core_UserInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Core_UserInformation) ProtoMessage() {}

func (x *Core_UserInformation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discovery_discovery_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Core_UserInformation.ProtoReflect.Descriptor instead.
func (*Core_UserInformation) Descriptor() ([]byte, []int) {
	return file_proto_discovery_discovery_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Core_UserInformation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Core_UserInformation) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Core_PrimaryInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PrimaryId     string                 `protobuf:"bytes,1,opt,name=primary_id,json=primaryId,proto3" json:"primary_id,omitempty"`
	PrimaryType   string                 `protobuf:"bytes,2,opt,name=primary_type,json=primaryType,proto3" json:"primary_type,omitempty"` // Add primary-specific fields
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Core_PrimaryInfo) Reset() {
	*x = Core_PrimaryInfo{}
	mi := &file_proto_discovery_discovery_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Core_PrimaryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Core_PrimaryInfo) ProtoMessage() {}

func (x *Core_PrimaryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discovery_discovery_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Core_PrimaryInfo.ProtoReflect.Descriptor instead.
func (*Core_PrimaryInfo) Descriptor() ([]byte, []int) {
	return file_proto_discovery_discovery_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Core_PrimaryInfo) GetPrimaryId() string {
	if x != nil {
		return x.PrimaryId
	}
	return ""
}

func (x *Core_PrimaryInfo) GetPrimaryType() string {
	if x != nil {
		return x.PrimaryType
	}
	return ""
}

type Core_InputInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InputType     string                 `protobuf:"bytes,1,opt,name=input_type,json=inputType,proto3" json:"input_type,omitempty"`
	InputData     []byte                 `protobuf:"bytes,2,opt,name=input_data,json=inputData,proto3" json:"input_data,omitempty"` // Add input-specific fields
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Core_InputInfo) Reset() {
	*x = Core_InputInfo{}
	mi := &file_proto_discovery_discovery_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Core_InputInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Core_InputInfo) ProtoMessage() {}

func (x *Core_InputInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_discovery_discovery_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Core_InputInfo.ProtoReflect.Descriptor instead.
func (*Core_InputInfo) Descriptor() ([]byte, []int) {
	return file_proto_discovery_discovery_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Core_InputInfo) GetInputType() string {
	if x != nil {
		return x.InputType
	}
	return ""
}

func (x *Core_InputInfo) GetInputData() []byte {
	if x != nil {
		return x.InputData
	}
	return nil
}

var File_proto_discovery_discovery_proto protoreflect.FileDescriptor

var file_proto_discovery_discovery_proto_rawDesc = string([]byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x03, 0x0a, 0x04,
	0x43, 0x6f, 0x72, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x07,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x72,
	0x65, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x40, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x4f, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x49, 0x0a, 0x09, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x22, 0xbf, 0x04, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x25, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3d,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5d, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x22, 0x49, 0x0a, 0x06, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x4f, 0x4c, 0x4c,
	0x42, 0x41, 0x43, 0x4b, 0x10, 0x02, 0x32, 0x66, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x10, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x1e,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1e,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x1e,
	0x5a, 0x1c, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_discovery_discovery_proto_rawDescOnce sync.Once
	file_proto_discovery_discovery_proto_rawDescData []byte
)

func file_proto_discovery_discovery_proto_rawDescGZIP() []byte {
	file_proto_discovery_discovery_proto_rawDescOnce.Do(func() {
		file_proto_discovery_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_discovery_discovery_proto_rawDesc), len(file_proto_discovery_discovery_proto_rawDesc)))
	})
	return file_proto_discovery_discovery_proto_rawDescData
}

var file_proto_discovery_discovery_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_discovery_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_discovery_discovery_proto_goTypes = []any{
	(DiscoveryMessage_Status)(0),  // 0: discovery.v1.DiscoveryMessage.Status
	(DiscoveryMessage_Action)(0),  // 1: discovery.v1.DiscoveryMessage.Action
	(*Core)(nil),                  // 2: discovery.v1.Core
	(*DiscoveryMessage)(nil),      // 3: discovery.v1.DiscoveryMessage
	(*Core_UserInformation)(nil),  // 4: discovery.v1.Core.UserInformation
	(*Core_PrimaryInfo)(nil),      // 5: discovery.v1.Core.PrimaryInfo
	(*Core_InputInfo)(nil),        // 6: discovery.v1.Core.InputInfo
	(*anypb.Any)(nil),             // 7: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_proto_discovery_discovery_proto_depIdxs = []int32{
	4,  // 0: discovery.v1.Core.user:type_name -> discovery.v1.Core.UserInformation
	5,  // 1: discovery.v1.Core.primary:type_name -> discovery.v1.Core.PrimaryInfo
	6,  // 2: discovery.v1.Core.input:type_name -> discovery.v1.Core.InputInfo
	7,  // 3: discovery.v1.Core.error:type_name -> google.protobuf.Any
	7,  // 4: discovery.v1.DiscoveryMessage.payload:type_name -> google.protobuf.Any
	0,  // 5: discovery.v1.DiscoveryMessage.status:type_name -> discovery.v1.DiscoveryMessage.Status
	1,  // 6: discovery.v1.DiscoveryMessage.action:type_name -> discovery.v1.DiscoveryMessage.Action
	7,  // 7: discovery.v1.DiscoveryMessage.error:type_name -> google.protobuf.Any
	8,  // 8: discovery.v1.DiscoveryMessage.timestamp:type_name -> google.protobuf.Timestamp
	3,  // 9: discovery.v1.DiscoveryService.ProcessDiscovery:input_type -> discovery.v1.DiscoveryMessage
	3,  // 10: discovery.v1.DiscoveryService.ProcessDiscovery:output_type -> discovery.v1.DiscoveryMessage
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_discovery_discovery_proto_init() }
func file_proto_discovery_discovery_proto_init() {
	if File_proto_discovery_discovery_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_discovery_discovery_proto_rawDesc), len(file_proto_discovery_discovery_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_discovery_discovery_proto_goTypes,
		DependencyIndexes: file_proto_discovery_discovery_proto_depIdxs,
		EnumInfos:         file_proto_discovery_discovery_proto_enumTypes,
		MessageInfos:      file_proto_discovery_discovery_proto_msgTypes,
	}.Build()
	File_proto_discovery_discovery_proto = out.File
	file_proto_discovery_discovery_proto_goTypes = nil
	file_proto_discovery_discovery_proto_depIdxs = nil
}
