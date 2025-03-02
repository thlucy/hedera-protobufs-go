// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: get_by_solidity_id.proto

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
// Get the IDs in the format used by transactions, given the ID in the format used by Solidity. If
// the Solidity ID is for a smart contract instance, then both the ContractID and associated
// AccountID will be returned.
type GetBySolidityIDQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard info sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The ID in the format used by Solidity
	SolidityID string `protobuf:"bytes,2,opt,name=solidityID,proto3" json:"solidityID,omitempty"`
}

func (x *GetBySolidityIDQuery) Reset() {
	*x = GetBySolidityIDQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_by_solidity_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBySolidityIDQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBySolidityIDQuery) ProtoMessage() {}

func (x *GetBySolidityIDQuery) ProtoReflect() protoreflect.Message {
	mi := &file_get_by_solidity_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBySolidityIDQuery.ProtoReflect.Descriptor instead.
func (*GetBySolidityIDQuery) Descriptor() ([]byte, []int) {
	return file_get_by_solidity_id_proto_rawDescGZIP(), []int{0}
}

func (x *GetBySolidityIDQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetBySolidityIDQuery) GetSolidityID() string {
	if x != nil {
		return x.SolidityID
	}
	return ""
}

//*
// Response when the client sends the node GetBySolidityIDQuery
type GetBySolidityIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The Account ID for the cryptocurrency account
	AccountID *AccountID `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	//*
	// The file Id for the file
	FileID *FileID `protobuf:"bytes,3,opt,name=fileID,proto3" json:"fileID,omitempty"`
	//*
	// A smart contract ID for the instance (if this is included, then the associated accountID will
	// also be included)
	ContractID *ContractID `protobuf:"bytes,4,opt,name=contractID,proto3" json:"contractID,omitempty"`
}

func (x *GetBySolidityIDResponse) Reset() {
	*x = GetBySolidityIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_by_solidity_id_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBySolidityIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBySolidityIDResponse) ProtoMessage() {}

func (x *GetBySolidityIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_by_solidity_id_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBySolidityIDResponse.ProtoReflect.Descriptor instead.
func (*GetBySolidityIDResponse) Descriptor() ([]byte, []int) {
	return file_get_by_solidity_id_proto_rawDescGZIP(), []int{1}
}

func (x *GetBySolidityIDResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetBySolidityIDResponse) GetAccountID() *AccountID {
	if x != nil {
		return x.AccountID
	}
	return nil
}

func (x *GetBySolidityIDResponse) GetFileID() *FileID {
	if x != nil {
		return x.FileID
	}
	return nil
}

func (x *GetBySolidityIDResponse) GetContractID() *ContractID {
	if x != nil {
		return x.ContractID
	}
	return nil
}

var File_get_by_solidity_id_proto protoreflect.FileDescriptor

var file_get_by_solidity_id_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x62, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x49, 0x44, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x69, 0x74,
	0x79, 0x49, 0x44, 0x22, 0xd2, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x6f, 0x6c,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x25,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_by_solidity_id_proto_rawDescOnce sync.Once
	file_get_by_solidity_id_proto_rawDescData = file_get_by_solidity_id_proto_rawDesc
)

func file_get_by_solidity_id_proto_rawDescGZIP() []byte {
	file_get_by_solidity_id_proto_rawDescOnce.Do(func() {
		file_get_by_solidity_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_by_solidity_id_proto_rawDescData)
	})
	return file_get_by_solidity_id_proto_rawDescData
}

var file_get_by_solidity_id_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_get_by_solidity_id_proto_goTypes = []interface{}{
	(*GetBySolidityIDQuery)(nil),    // 0: proto.GetBySolidityIDQuery
	(*GetBySolidityIDResponse)(nil), // 1: proto.GetBySolidityIDResponse
	(*QueryHeader)(nil),             // 2: proto.QueryHeader
	(*ResponseHeader)(nil),          // 3: proto.ResponseHeader
	(*AccountID)(nil),               // 4: proto.AccountID
	(*FileID)(nil),                  // 5: proto.FileID
	(*ContractID)(nil),              // 6: proto.ContractID
}
var file_get_by_solidity_id_proto_depIdxs = []int32{
	2, // 0: proto.GetBySolidityIDQuery.header:type_name -> proto.QueryHeader
	3, // 1: proto.GetBySolidityIDResponse.header:type_name -> proto.ResponseHeader
	4, // 2: proto.GetBySolidityIDResponse.accountID:type_name -> proto.AccountID
	5, // 3: proto.GetBySolidityIDResponse.fileID:type_name -> proto.FileID
	6, // 4: proto.GetBySolidityIDResponse.contractID:type_name -> proto.ContractID
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_get_by_solidity_id_proto_init() }
func file_get_by_solidity_id_proto_init() {
	if File_get_by_solidity_id_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_get_by_solidity_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBySolidityIDQuery); i {
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
		file_get_by_solidity_id_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBySolidityIDResponse); i {
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
			RawDescriptor: file_get_by_solidity_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_get_by_solidity_id_proto_goTypes,
		DependencyIndexes: file_get_by_solidity_id_proto_depIdxs,
		MessageInfos:      file_get_by_solidity_id_proto_msgTypes,
	}.Build()
	File_get_by_solidity_id_proto = out.File
	file_get_by_solidity_id_proto_rawDesc = nil
	file_get_by_solidity_id_proto_goTypes = nil
	file_get_by_solidity_id_proto_depIdxs = nil
}
