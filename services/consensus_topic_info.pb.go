// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: consensus_topic_info.proto

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
// Current state of a topic.
type ConsensusTopicInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The memo associated with the topic (UTF-8 encoding max 100 bytes)
	Memo string `protobuf:"bytes,1,opt,name=memo,proto3" json:"memo,omitempty"`
	//*
	// When a topic is created, its running hash is initialized to 48 bytes of binary zeros.
	// For each submitted message, the topic's running hash is then updated to the output
	// of a particular SHA-384 digest whose input data include the previous running hash.
	//
	// See the TransactionReceipt.proto documentation for an exact description of the
	// data included in the SHA-384 digest used for the update.
	RunningHash []byte `protobuf:"bytes,2,opt,name=runningHash,proto3" json:"runningHash,omitempty"`
	//*
	// Sequence number (starting at 1 for the first submitMessage) of messages on the topic.
	SequenceNumber uint64 `protobuf:"varint,3,opt,name=sequenceNumber,proto3" json:"sequenceNumber,omitempty"`
	//*
	// Effective consensus timestamp at (and after) which submitMessage calls will no longer succeed on the topic
	// and the topic will expire and after AUTORENEW_GRACE_PERIOD be automatically deleted.
	ExpirationTime *Timestamp `protobuf:"bytes,4,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	//*
	// Access control for update/delete of the topic. Null if there is no key.
	AdminKey *Key `protobuf:"bytes,5,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	//*
	// Access control for ConsensusService.submitMessage. Null if there is no key.
	SubmitKey *Key `protobuf:"bytes,6,opt,name=submitKey,proto3" json:"submitKey,omitempty"`
	//*
	// If an auto-renew account is specified, when the topic expires, its lifetime will be extended
	// by up to this duration (depending on the solvency of the auto-renew account). If the
	// auto-renew account has no funds at all, the topic will be deleted instead.
	AutoRenewPeriod *Duration `protobuf:"bytes,7,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	//*
	// The account, if any, to charge for automatic renewal of the topic's lifetime upon expiry.
	AutoRenewAccount *AccountID `protobuf:"bytes,8,opt,name=autoRenewAccount,proto3" json:"autoRenewAccount,omitempty"`
	//*
	// The ledger ID the response was returned from; please see <a href="https://github.com/hashgraph/hedera-improvement-proposal/blob/master/HIP/hip-198.md">HIP-198</a> for the network-specific IDs.
	LedgerId []byte `protobuf:"bytes,9,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
}

func (x *ConsensusTopicInfo) Reset() {
	*x = ConsensusTopicInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consensus_topic_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusTopicInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusTopicInfo) ProtoMessage() {}

func (x *ConsensusTopicInfo) ProtoReflect() protoreflect.Message {
	mi := &file_consensus_topic_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusTopicInfo.ProtoReflect.Descriptor instead.
func (*ConsensusTopicInfo) Descriptor() ([]byte, []int) {
	return file_consensus_topic_info_proto_rawDescGZIP(), []int{0}
}

func (x *ConsensusTopicInfo) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *ConsensusTopicInfo) GetRunningHash() []byte {
	if x != nil {
		return x.RunningHash
	}
	return nil
}

func (x *ConsensusTopicInfo) GetSequenceNumber() uint64 {
	if x != nil {
		return x.SequenceNumber
	}
	return 0
}

func (x *ConsensusTopicInfo) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *ConsensusTopicInfo) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *ConsensusTopicInfo) GetSubmitKey() *Key {
	if x != nil {
		return x.SubmitKey
	}
	return nil
}

func (x *ConsensusTopicInfo) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *ConsensusTopicInfo) GetAutoRenewAccount() *AccountID {
	if x != nil {
		return x.AutoRenewAccount
	}
	return nil
}

func (x *ConsensusTopicInfo) GetLedgerId() []byte {
	if x != nil {
		return x.LedgerId
	}
	return nil
}

var File_consensus_topic_info_proto protoreflect.FileDescriptor

var file_consensus_topic_info_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x03, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65,
	0x6d, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0e,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b,
	0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x28,
	0x0a, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f,
	0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x3c, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52,
	0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x49, 0x64, 0x42, 0x26,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consensus_topic_info_proto_rawDescOnce sync.Once
	file_consensus_topic_info_proto_rawDescData = file_consensus_topic_info_proto_rawDesc
)

func file_consensus_topic_info_proto_rawDescGZIP() []byte {
	file_consensus_topic_info_proto_rawDescOnce.Do(func() {
		file_consensus_topic_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_consensus_topic_info_proto_rawDescData)
	})
	return file_consensus_topic_info_proto_rawDescData
}

var file_consensus_topic_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_consensus_topic_info_proto_goTypes = []interface{}{
	(*ConsensusTopicInfo)(nil), // 0: proto.ConsensusTopicInfo
	(*Timestamp)(nil),          // 1: proto.Timestamp
	(*Key)(nil),                // 2: proto.Key
	(*Duration)(nil),           // 3: proto.Duration
	(*AccountID)(nil),          // 4: proto.AccountID
}
var file_consensus_topic_info_proto_depIdxs = []int32{
	1, // 0: proto.ConsensusTopicInfo.expirationTime:type_name -> proto.Timestamp
	2, // 1: proto.ConsensusTopicInfo.adminKey:type_name -> proto.Key
	2, // 2: proto.ConsensusTopicInfo.submitKey:type_name -> proto.Key
	3, // 3: proto.ConsensusTopicInfo.autoRenewPeriod:type_name -> proto.Duration
	4, // 4: proto.ConsensusTopicInfo.autoRenewAccount:type_name -> proto.AccountID
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_consensus_topic_info_proto_init() }
func file_consensus_topic_info_proto_init() {
	if File_consensus_topic_info_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_duration_proto_init()
	file_timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_consensus_topic_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusTopicInfo); i {
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
			RawDescriptor: file_consensus_topic_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_consensus_topic_info_proto_goTypes,
		DependencyIndexes: file_consensus_topic_info_proto_depIdxs,
		MessageInfos:      file_consensus_topic_info_proto_msgTypes,
	}.Build()
	File_consensus_topic_info_proto = out.File
	file_consensus_topic_info_proto_rawDesc = nil
	file_consensus_topic_info_proto_goTypes = nil
	file_consensus_topic_info_proto_depIdxs = nil
}
