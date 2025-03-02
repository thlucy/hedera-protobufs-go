// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: transaction_get_fast_record.proto

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
// Get the tx record of a transaction, given its transaction ID. Once a transaction reaches
// consensus, then information about whether it succeeded or failed will be available until the end
// of the receipt period.  Before and after the receipt period, and for a transaction that was never
// submitted, the receipt is unknown.  This query is free (the payment field is left empty).
type TransactionGetFastRecordQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard info sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The ID of the transaction for which the record is requested.
	TransactionID *TransactionID `protobuf:"bytes,2,opt,name=transactionID,proto3" json:"transactionID,omitempty"`
}

func (x *TransactionGetFastRecordQuery) Reset() {
	*x = TransactionGetFastRecordQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_get_fast_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionGetFastRecordQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionGetFastRecordQuery) ProtoMessage() {}

func (x *TransactionGetFastRecordQuery) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_get_fast_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionGetFastRecordQuery.ProtoReflect.Descriptor instead.
func (*TransactionGetFastRecordQuery) Descriptor() ([]byte, []int) {
	return file_transaction_get_fast_record_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionGetFastRecordQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TransactionGetFastRecordQuery) GetTransactionID() *TransactionID {
	if x != nil {
		return x.TransactionID
	}
	return nil
}

//*
// Response when the client sends the node TransactionGetFastRecordQuery. If it created a new entity
// (account, file, or smart contract instance) then one of the three ID fields will be filled in
// with the ID of the new entity. Sometimes a single transaction will create more than one new
// entity, such as when a new contract instance is created, and this also creates the new account
// that it owned by that instance.
type TransactionGetFastRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The requested transaction records
	TransactionRecord *TransactionRecord `protobuf:"bytes,2,opt,name=transactionRecord,proto3" json:"transactionRecord,omitempty"`
}

func (x *TransactionGetFastRecordResponse) Reset() {
	*x = TransactionGetFastRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_get_fast_record_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionGetFastRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionGetFastRecordResponse) ProtoMessage() {}

func (x *TransactionGetFastRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_get_fast_record_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionGetFastRecordResponse.ProtoReflect.Descriptor instead.
func (*TransactionGetFastRecordResponse) Descriptor() ([]byte, []int) {
	return file_transaction_get_fast_record_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionGetFastRecordResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TransactionGetFastRecordResponse) GetTransactionRecord() *TransactionRecord {
	if x != nil {
		return x.TransactionRecord
	}
	return nil
}

var File_transaction_get_fast_record_proto protoreflect.FileDescriptor

var file_transaction_get_fast_record_proto_rawDesc = []byte{
	0x0a, 0x21, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x66, 0x61, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x1d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x3a, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52, 0x0d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x99, 0x01, 0x0a,
	0x20, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x46,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x46, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_get_fast_record_proto_rawDescOnce sync.Once
	file_transaction_get_fast_record_proto_rawDescData = file_transaction_get_fast_record_proto_rawDesc
)

func file_transaction_get_fast_record_proto_rawDescGZIP() []byte {
	file_transaction_get_fast_record_proto_rawDescOnce.Do(func() {
		file_transaction_get_fast_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_get_fast_record_proto_rawDescData)
	})
	return file_transaction_get_fast_record_proto_rawDescData
}

var file_transaction_get_fast_record_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transaction_get_fast_record_proto_goTypes = []interface{}{
	(*TransactionGetFastRecordQuery)(nil),    // 0: proto.TransactionGetFastRecordQuery
	(*TransactionGetFastRecordResponse)(nil), // 1: proto.TransactionGetFastRecordResponse
	(*QueryHeader)(nil),                      // 2: proto.QueryHeader
	(*TransactionID)(nil),                    // 3: proto.TransactionID
	(*ResponseHeader)(nil),                   // 4: proto.ResponseHeader
	(*TransactionRecord)(nil),                // 5: proto.TransactionRecord
}
var file_transaction_get_fast_record_proto_depIdxs = []int32{
	2, // 0: proto.TransactionGetFastRecordQuery.header:type_name -> proto.QueryHeader
	3, // 1: proto.TransactionGetFastRecordQuery.transactionID:type_name -> proto.TransactionID
	4, // 2: proto.TransactionGetFastRecordResponse.header:type_name -> proto.ResponseHeader
	5, // 3: proto.TransactionGetFastRecordResponse.transactionRecord:type_name -> proto.TransactionRecord
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_transaction_get_fast_record_proto_init() }
func file_transaction_get_fast_record_proto_init() {
	if File_transaction_get_fast_record_proto != nil {
		return
	}
	file_transaction_record_proto_init()
	file_basic_types_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transaction_get_fast_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionGetFastRecordQuery); i {
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
		file_transaction_get_fast_record_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionGetFastRecordResponse); i {
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
			RawDescriptor: file_transaction_get_fast_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_get_fast_record_proto_goTypes,
		DependencyIndexes: file_transaction_get_fast_record_proto_depIdxs,
		MessageInfos:      file_transaction_get_fast_record_proto_msgTypes,
	}.Build()
	File_transaction_get_fast_record_proto = out.File
	file_transaction_get_fast_record_proto_rawDesc = nil
	file_transaction_get_fast_record_proto_goTypes = nil
	file_transaction_get_fast_record_proto_depIdxs = nil
}
