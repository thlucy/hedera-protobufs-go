// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: transaction_response.proto

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
// When the client sends the node a transaction of any kind, the node replies with this, which
// simply says that the transaction passed the precheck (so the node will submit it to the network)
// or it failed (so it won't). If the fee offered was insufficient, this will also contain the
// amount of the required fee. To learn the consensus result, the client should later obtain a
// receipt (free), or can buy a more detailed record (not free).
type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The response code that indicates the current status of the transaction.
	NodeTransactionPrecheckCode ResponseCodeEnum `protobuf:"varint,1,opt,name=nodeTransactionPrecheckCode,proto3,enum=proto.ResponseCodeEnum" json:"nodeTransactionPrecheckCode,omitempty"`
	//*
	// If the response code was INSUFFICIENT_TX_FEE, the actual transaction fee that would be
	// required to execute the transaction.
	Cost uint64 `protobuf:"varint,2,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_transaction_response_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionResponse) GetNodeTransactionPrecheckCode() ResponseCodeEnum {
	if x != nil {
		return x.NodeTransactionPrecheckCode
	}
	return ResponseCodeEnum_OK
}

func (x *TransactionResponse) GetCost() uint64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

var File_transaction_response_proto protoreflect.FileDescriptor

var file_transaction_response_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x13, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x59, 0x0a, 0x1b, 0x6e, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x1b,
	0x6e, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x65, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x42,
	0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73,
	0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_response_proto_rawDescOnce sync.Once
	file_transaction_response_proto_rawDescData = file_transaction_response_proto_rawDesc
)

func file_transaction_response_proto_rawDescGZIP() []byte {
	file_transaction_response_proto_rawDescOnce.Do(func() {
		file_transaction_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_response_proto_rawDescData)
	})
	return file_transaction_response_proto_rawDescData
}

var file_transaction_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transaction_response_proto_goTypes = []interface{}{
	(*TransactionResponse)(nil), // 0: proto.TransactionResponse
	(ResponseCodeEnum)(0),       // 1: proto.ResponseCodeEnum
}
var file_transaction_response_proto_depIdxs = []int32{
	1, // 0: proto.TransactionResponse.nodeTransactionPrecheckCode:type_name -> proto.ResponseCodeEnum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transaction_response_proto_init() }
func file_transaction_response_proto_init() {
	if File_transaction_response_proto != nil {
		return
	}
	file_response_code_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transaction_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
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
			RawDescriptor: file_transaction_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_response_proto_goTypes,
		DependencyIndexes: file_transaction_response_proto_depIdxs,
		MessageInfos:      file_transaction_response_proto_msgTypes,
	}.Build()
	File_transaction_response_proto = out.File
	file_transaction_response_proto_rawDesc = nil
	file_transaction_response_proto_goTypes = nil
	file_transaction_response_proto_depIdxs = nil
}
