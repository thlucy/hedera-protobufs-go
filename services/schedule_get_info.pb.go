// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: schedule_get_info.proto

package services

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

//*
// Gets information about a schedule in the network's action queue.
//
// Responds with <tt>INVALID_SCHEDULE_ID</tt> if the requested schedule doesn't exist.
type ScheduleGetInfoQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// standard info sent from client to node including the signed payment, and what kind of response
	// is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The id of the schedule to interrogate
	ScheduleID *ScheduleID `protobuf:"bytes,2,opt,name=scheduleID,proto3" json:"scheduleID,omitempty"`
}

func (x *ScheduleGetInfoQuery) Reset() {
	*x = ScheduleGetInfoQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_get_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleGetInfoQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleGetInfoQuery) ProtoMessage() {}

func (x *ScheduleGetInfoQuery) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_get_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleGetInfoQuery.ProtoReflect.Descriptor instead.
func (*ScheduleGetInfoQuery) Descriptor() ([]byte, []int) {
	return file_schedule_get_info_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleGetInfoQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ScheduleGetInfoQuery) GetScheduleID() *ScheduleID {
	if x != nil {
		return x.ScheduleID
	}
	return nil
}

//*
// Information summarizing schedule state
type ScheduleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The id of the schedule
	ScheduleID *ScheduleID `protobuf:"bytes,1,opt,name=scheduleID,proto3" json:"scheduleID,omitempty"`
	// Types that are assignable to Data:
	//	*ScheduleInfo_DeletionTime
	//	*ScheduleInfo_ExecutionTime
	Data isScheduleInfo_Data `protobuf_oneof:"data"`
	//*
	// The time at which the schedule will be evaluated for execution and then expire.
	//
	// Note: Before Long Term Scheduled Transactions are enabled, Scheduled Transactions will _never_ execute at expiration - they
	//       will _only_ execute during the initial ScheduleCreate transaction or via ScheduleSign transactions and will _always_
	//       expire at expirationTime.
	ExpirationTime *Timestamp `protobuf:"bytes,4,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	//*
	// The scheduled transaction
	ScheduledTransactionBody *SchedulableTransactionBody `protobuf:"bytes,5,opt,name=scheduledTransactionBody,proto3" json:"scheduledTransactionBody,omitempty"`
	//*
	// The publicly visible memo of the schedule
	Memo string `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
	//*
	// The key used to delete the schedule from state
	AdminKey *Key `protobuf:"bytes,7,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	//*
	// The Ed25519 keys the network deems to have signed the scheduled transaction
	Signers *KeyList `protobuf:"bytes,8,opt,name=signers,proto3" json:"signers,omitempty"`
	//*
	// The id of the account that created the schedule
	CreatorAccountID *AccountID `protobuf:"bytes,9,opt,name=creatorAccountID,proto3" json:"creatorAccountID,omitempty"`
	//*
	// The id of the account responsible for the service fee of the scheduled transaction
	PayerAccountID *AccountID `protobuf:"bytes,10,opt,name=payerAccountID,proto3" json:"payerAccountID,omitempty"`
	//*
	// The transaction id that will be used in the record of the scheduled transaction (if it
	// executes)
	ScheduledTransactionID *TransactionID `protobuf:"bytes,11,opt,name=scheduledTransactionID,proto3" json:"scheduledTransactionID,omitempty"`
	//*
	// The ledger ID the response was returned from; please see <a href="https://github.com/hashgraph/hedera-improvement-proposal/blob/master/HIP/hip-198.md">HIP-198</a> for the network-specific IDs.
	LedgerId []byte `protobuf:"bytes,12,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
	//*
	// When set to true, the transaction will be evaluated for execution at expiration_time instead
	// of when all required signatures are received.
	// When set to false, the transaction will execute immediately after sufficient signatures are received
	// to sign the contained transaction. During the initial ScheduleCreate transaction or via ScheduleSign transactions.
	//
	// Note: this field is unused until Long Term Scheduled Transactions are enabled.
	WaitForExpiry bool `protobuf:"varint,13,opt,name=wait_for_expiry,json=waitForExpiry,proto3" json:"wait_for_expiry,omitempty"`
}

func (x *ScheduleInfo) Reset() {
	*x = ScheduleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_get_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleInfo) ProtoMessage() {}

func (x *ScheduleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_get_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleInfo.ProtoReflect.Descriptor instead.
func (*ScheduleInfo) Descriptor() ([]byte, []int) {
	return file_schedule_get_info_proto_rawDescGZIP(), []int{1}
}

func (x *ScheduleInfo) GetScheduleID() *ScheduleID {
	if x != nil {
		return x.ScheduleID
	}
	return nil
}

func (m *ScheduleInfo) GetData() isScheduleInfo_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ScheduleInfo) GetDeletionTime() *Timestamp {
	if x, ok := x.GetData().(*ScheduleInfo_DeletionTime); ok {
		return x.DeletionTime
	}
	return nil
}

func (x *ScheduleInfo) GetExecutionTime() *Timestamp {
	if x, ok := x.GetData().(*ScheduleInfo_ExecutionTime); ok {
		return x.ExecutionTime
	}
	return nil
}

func (x *ScheduleInfo) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *ScheduleInfo) GetScheduledTransactionBody() *SchedulableTransactionBody {
	if x != nil {
		return x.ScheduledTransactionBody
	}
	return nil
}

func (x *ScheduleInfo) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *ScheduleInfo) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *ScheduleInfo) GetSigners() *KeyList {
	if x != nil {
		return x.Signers
	}
	return nil
}

func (x *ScheduleInfo) GetCreatorAccountID() *AccountID {
	if x != nil {
		return x.CreatorAccountID
	}
	return nil
}

func (x *ScheduleInfo) GetPayerAccountID() *AccountID {
	if x != nil {
		return x.PayerAccountID
	}
	return nil
}

func (x *ScheduleInfo) GetScheduledTransactionID() *TransactionID {
	if x != nil {
		return x.ScheduledTransactionID
	}
	return nil
}

func (x *ScheduleInfo) GetLedgerId() []byte {
	if x != nil {
		return x.LedgerId
	}
	return nil
}

func (x *ScheduleInfo) GetWaitForExpiry() bool {
	if x != nil {
		return x.WaitForExpiry
	}
	return false
}

type isScheduleInfo_Data interface {
	isScheduleInfo_Data()
}

type ScheduleInfo_DeletionTime struct {
	//*
	// If the schedule has been deleted, the consensus time when this occurred
	DeletionTime *Timestamp `protobuf:"bytes,2,opt,name=deletion_time,json=deletionTime,proto3,oneof"`
}

type ScheduleInfo_ExecutionTime struct {
	//*
	// If the schedule has been executed, the consensus time when this occurred
	ExecutionTime *Timestamp `protobuf:"bytes,3,opt,name=execution_time,json=executionTime,proto3,oneof"`
}

func (*ScheduleInfo_DeletionTime) isScheduleInfo_Data() {}

func (*ScheduleInfo_ExecutionTime) isScheduleInfo_Data() {}

//*
// Response wrapper for the <tt>ScheduleInfo</tt>
type ScheduleGetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard response from node to client, including the requested fields: cost, or state proof, or
	// both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The information requested about this schedule instance
	ScheduleInfo *ScheduleInfo `protobuf:"bytes,2,opt,name=scheduleInfo,proto3" json:"scheduleInfo,omitempty"`
}

func (x *ScheduleGetInfoResponse) Reset() {
	*x = ScheduleGetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_get_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleGetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleGetInfoResponse) ProtoMessage() {}

func (x *ScheduleGetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_get_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleGetInfoResponse.ProtoReflect.Descriptor instead.
func (*ScheduleGetInfoResponse) Descriptor() ([]byte, []int) {
	return file_schedule_get_info_proto_rawDescGZIP(), []int{2}
}

func (x *ScheduleGetInfoResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ScheduleGetInfoResponse) GetScheduleInfo() *ScheduleInfo {
	if x != nil {
		return x.ScheduleInfo
	}
	return nil
}

var File_schedule_get_info_proto protoreflect.FileDescriptor

var file_schedule_get_info_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x14, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x0a,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x22, 0xc7, 0x05, 0x0a, 0x0c, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x0a, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x49, 0x44, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x37,
	0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x38, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5d, 0x0a, 0x18,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x18, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12,
	0x26, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x73, 0x12, 0x3c, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x10, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x38, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x4c, 0x0a, 0x16, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52,
	0x16, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x66, 0x6f, 0x72,
	0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x77,
	0x61, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x17, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x37, 0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schedule_get_info_proto_rawDescOnce sync.Once
	file_schedule_get_info_proto_rawDescData = file_schedule_get_info_proto_rawDesc
)

func file_schedule_get_info_proto_rawDescGZIP() []byte {
	file_schedule_get_info_proto_rawDescOnce.Do(func() {
		file_schedule_get_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_schedule_get_info_proto_rawDescData)
	})
	return file_schedule_get_info_proto_rawDescData
}

var file_schedule_get_info_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schedule_get_info_proto_goTypes = []interface{}{
	(*ScheduleGetInfoQuery)(nil),       // 0: proto.ScheduleGetInfoQuery
	(*ScheduleInfo)(nil),               // 1: proto.ScheduleInfo
	(*ScheduleGetInfoResponse)(nil),    // 2: proto.ScheduleGetInfoResponse
	(*QueryHeader)(nil),                // 3: proto.QueryHeader
	(*ScheduleID)(nil),                 // 4: proto.ScheduleID
	(*Timestamp)(nil),                  // 5: proto.Timestamp
	(*SchedulableTransactionBody)(nil), // 6: proto.SchedulableTransactionBody
	(*Key)(nil),                        // 7: proto.Key
	(*KeyList)(nil),                    // 8: proto.KeyList
	(*AccountID)(nil),                  // 9: proto.AccountID
	(*TransactionID)(nil),              // 10: proto.TransactionID
	(*ResponseHeader)(nil),             // 11: proto.ResponseHeader
}
var file_schedule_get_info_proto_depIdxs = []int32{
	3,  // 0: proto.ScheduleGetInfoQuery.header:type_name -> proto.QueryHeader
	4,  // 1: proto.ScheduleGetInfoQuery.scheduleID:type_name -> proto.ScheduleID
	4,  // 2: proto.ScheduleInfo.scheduleID:type_name -> proto.ScheduleID
	5,  // 3: proto.ScheduleInfo.deletion_time:type_name -> proto.Timestamp
	5,  // 4: proto.ScheduleInfo.execution_time:type_name -> proto.Timestamp
	5,  // 5: proto.ScheduleInfo.expirationTime:type_name -> proto.Timestamp
	6,  // 6: proto.ScheduleInfo.scheduledTransactionBody:type_name -> proto.SchedulableTransactionBody
	7,  // 7: proto.ScheduleInfo.adminKey:type_name -> proto.Key
	8,  // 8: proto.ScheduleInfo.signers:type_name -> proto.KeyList
	9,  // 9: proto.ScheduleInfo.creatorAccountID:type_name -> proto.AccountID
	9,  // 10: proto.ScheduleInfo.payerAccountID:type_name -> proto.AccountID
	10, // 11: proto.ScheduleInfo.scheduledTransactionID:type_name -> proto.TransactionID
	11, // 12: proto.ScheduleGetInfoResponse.header:type_name -> proto.ResponseHeader
	1,  // 13: proto.ScheduleGetInfoResponse.scheduleInfo:type_name -> proto.ScheduleInfo
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_schedule_get_info_proto_init() }
func file_schedule_get_info_proto_init() {
	if File_schedule_get_info_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_timestamp_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	file_schedulable_transaction_body_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_schedule_get_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleGetInfoQuery); i {
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
		file_schedule_get_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleInfo); i {
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
		file_schedule_get_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleGetInfoResponse); i {
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
	file_schedule_get_info_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ScheduleInfo_DeletionTime)(nil),
		(*ScheduleInfo_ExecutionTime)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schedule_get_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schedule_get_info_proto_goTypes,
		DependencyIndexes: file_schedule_get_info_proto_depIdxs,
		MessageInfos:      file_schedule_get_info_proto_msgTypes,
	}.Build()
	File_schedule_get_info_proto = out.File
	file_schedule_get_info_proto_rawDesc = nil
	file_schedule_get_info_proto_goTypes = nil
	file_schedule_get_info_proto_depIdxs = nil
}
