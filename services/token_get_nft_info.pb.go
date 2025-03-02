// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: token_get_nft_info.proto

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
// Represents an NFT on the Ledger
type NftID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The (non-fungible) token of which this NFT is an instance
	TokenID *TokenID `protobuf:"bytes,1,opt,name=tokenID,proto3" json:"tokenID,omitempty"`
	//*
	// The unique identifier of this instance
	SerialNumber int64 `protobuf:"varint,2,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
}

func (x *NftID) Reset() {
	*x = NftID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_get_nft_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftID) ProtoMessage() {}

func (x *NftID) ProtoReflect() protoreflect.Message {
	mi := &file_token_get_nft_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftID.ProtoReflect.Descriptor instead.
func (*NftID) Descriptor() ([]byte, []int) {
	return file_token_get_nft_info_proto_rawDescGZIP(), []int{0}
}

func (x *NftID) GetTokenID() *TokenID {
	if x != nil {
		return x.TokenID
	}
	return nil
}

func (x *NftID) GetSerialNumber() int64 {
	if x != nil {
		return x.SerialNumber
	}
	return 0
}

//*
// Applicable only to tokens of type NON_FUNGIBLE_UNIQUE. Gets info on a NFT for a given TokenID (of
// type NON_FUNGIBLE_UNIQUE) and serial number
type TokenGetNftInfoQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard info sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The ID of the NFT
	NftID *NftID `protobuf:"bytes,2,opt,name=nftID,proto3" json:"nftID,omitempty"`
}

func (x *TokenGetNftInfoQuery) Reset() {
	*x = TokenGetNftInfoQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_get_nft_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenGetNftInfoQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenGetNftInfoQuery) ProtoMessage() {}

func (x *TokenGetNftInfoQuery) ProtoReflect() protoreflect.Message {
	mi := &file_token_get_nft_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenGetNftInfoQuery.ProtoReflect.Descriptor instead.
func (*TokenGetNftInfoQuery) Descriptor() ([]byte, []int) {
	return file_token_get_nft_info_proto_rawDescGZIP(), []int{1}
}

func (x *TokenGetNftInfoQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TokenGetNftInfoQuery) GetNftID() *NftID {
	if x != nil {
		return x.NftID
	}
	return nil
}

//*
// UNDOCUMENTED
type TokenNftInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The ID of the NFT
	NftID *NftID `protobuf:"bytes,1,opt,name=nftID,proto3" json:"nftID,omitempty"`
	//*
	// The current owner of the NFT
	AccountID *AccountID `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	//*
	// The effective consensus timestamp at which the NFT was minted
	CreationTime *Timestamp `protobuf:"bytes,3,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	//*
	// Represents the unique metadata of the NFT
	Metadata []byte `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	//*
	// The ledger ID the response was returned from; please see <a href="https://github.com/hashgraph/hedera-improvement-proposal/blob/master/HIP/hip-198.md">HIP-198</a> for the network-specific IDs.
	LedgerId []byte `protobuf:"bytes,5,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
	//*
	// If an allowance is granted for the NFT, its corresponding spender account
	SpenderId *AccountID `protobuf:"bytes,6,opt,name=spender_id,json=spenderId,proto3" json:"spender_id,omitempty"`
}

func (x *TokenNftInfo) Reset() {
	*x = TokenNftInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_get_nft_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenNftInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenNftInfo) ProtoMessage() {}

func (x *TokenNftInfo) ProtoReflect() protoreflect.Message {
	mi := &file_token_get_nft_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenNftInfo.ProtoReflect.Descriptor instead.
func (*TokenNftInfo) Descriptor() ([]byte, []int) {
	return file_token_get_nft_info_proto_rawDescGZIP(), []int{2}
}

func (x *TokenNftInfo) GetNftID() *NftID {
	if x != nil {
		return x.NftID
	}
	return nil
}

func (x *TokenNftInfo) GetAccountID() *AccountID {
	if x != nil {
		return x.AccountID
	}
	return nil
}

func (x *TokenNftInfo) GetCreationTime() *Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *TokenNftInfo) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TokenNftInfo) GetLedgerId() []byte {
	if x != nil {
		return x.LedgerId
	}
	return nil
}

func (x *TokenNftInfo) GetSpenderId() *AccountID {
	if x != nil {
		return x.SpenderId
	}
	return nil
}

//*
// UNDOCUMENTED
type TokenGetNftInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The information about this NFT
	Nft *TokenNftInfo `protobuf:"bytes,2,opt,name=nft,proto3" json:"nft,omitempty"`
}

func (x *TokenGetNftInfoResponse) Reset() {
	*x = TokenGetNftInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_get_nft_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenGetNftInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenGetNftInfoResponse) ProtoMessage() {}

func (x *TokenGetNftInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_get_nft_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenGetNftInfoResponse.ProtoReflect.Descriptor instead.
func (*TokenGetNftInfoResponse) Descriptor() ([]byte, []int) {
	return file_token_get_nft_info_proto_rawDescGZIP(), []int{3}
}

func (x *TokenGetNftInfoResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TokenGetNftInfoResponse) GetNft() *TokenNftInfo {
	if x != nil {
		return x.Nft
	}
	return nil
}

var File_token_get_nft_info_proto protoreflect.FileDescriptor

var file_token_get_nft_info_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x66, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x55, 0x0a, 0x05, 0x4e, 0x66, 0x74, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x07, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x66, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x47, 0x65, 0x74, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x05, 0x6e,
	0x66, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x44, 0x52, 0x05, 0x6e, 0x66, 0x74, 0x49, 0x44, 0x22,
	0x82, 0x02, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x22, 0x0a, 0x05, 0x6e, 0x66, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x66, 0x74, 0x49, 0x44, 0x52, 0x05, 0x6e,
	0x66, 0x74, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0a, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x17, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74,
	0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x25,
	0x0a, 0x03, 0x6e, 0x66, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x03, 0x6e, 0x66, 0x74, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_get_nft_info_proto_rawDescOnce sync.Once
	file_token_get_nft_info_proto_rawDescData = file_token_get_nft_info_proto_rawDesc
)

func file_token_get_nft_info_proto_rawDescGZIP() []byte {
	file_token_get_nft_info_proto_rawDescOnce.Do(func() {
		file_token_get_nft_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_get_nft_info_proto_rawDescData)
	})
	return file_token_get_nft_info_proto_rawDescData
}

var file_token_get_nft_info_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_token_get_nft_info_proto_goTypes = []interface{}{
	(*NftID)(nil),                   // 0: proto.NftID
	(*TokenGetNftInfoQuery)(nil),    // 1: proto.TokenGetNftInfoQuery
	(*TokenNftInfo)(nil),            // 2: proto.TokenNftInfo
	(*TokenGetNftInfoResponse)(nil), // 3: proto.TokenGetNftInfoResponse
	(*TokenID)(nil),                 // 4: proto.TokenID
	(*QueryHeader)(nil),             // 5: proto.QueryHeader
	(*AccountID)(nil),               // 6: proto.AccountID
	(*Timestamp)(nil),               // 7: proto.Timestamp
	(*ResponseHeader)(nil),          // 8: proto.ResponseHeader
}
var file_token_get_nft_info_proto_depIdxs = []int32{
	4, // 0: proto.NftID.tokenID:type_name -> proto.TokenID
	5, // 1: proto.TokenGetNftInfoQuery.header:type_name -> proto.QueryHeader
	0, // 2: proto.TokenGetNftInfoQuery.nftID:type_name -> proto.NftID
	0, // 3: proto.TokenNftInfo.nftID:type_name -> proto.NftID
	6, // 4: proto.TokenNftInfo.accountID:type_name -> proto.AccountID
	7, // 5: proto.TokenNftInfo.creationTime:type_name -> proto.Timestamp
	6, // 6: proto.TokenNftInfo.spender_id:type_name -> proto.AccountID
	8, // 7: proto.TokenGetNftInfoResponse.header:type_name -> proto.ResponseHeader
	2, // 8: proto.TokenGetNftInfoResponse.nft:type_name -> proto.TokenNftInfo
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_token_get_nft_info_proto_init() }
func file_token_get_nft_info_proto_init() {
	if File_token_get_nft_info_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	file_timestamp_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_token_get_nft_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftID); i {
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
		file_token_get_nft_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenGetNftInfoQuery); i {
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
		file_token_get_nft_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenNftInfo); i {
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
		file_token_get_nft_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenGetNftInfoResponse); i {
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
			RawDescriptor: file_token_get_nft_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_get_nft_info_proto_goTypes,
		DependencyIndexes: file_token_get_nft_info_proto_depIdxs,
		MessageInfos:      file_token_get_nft_info_proto_msgTypes,
	}.Build()
	File_token_get_nft_info_proto = out.File
	file_token_get_nft_info_proto_rawDesc = nil
	file_token_get_nft_info_proto_goTypes = nil
	file_token_get_nft_info_proto_depIdxs = nil
}
