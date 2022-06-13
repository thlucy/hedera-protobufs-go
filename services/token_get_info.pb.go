// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: token_get_info.proto

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
// Gets information about Token instance
type TokenGetInfoQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard info sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither)
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The token for which information is requested. If invalid token is provided, INVALID_TOKEN_ID
	// response is returned.
	Token *TokenID `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenGetInfoQuery) Reset() {
	*x = TokenGetInfoQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_get_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenGetInfoQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenGetInfoQuery) ProtoMessage() {}

func (x *TokenGetInfoQuery) ProtoReflect() protoreflect.Message {
	mi := &file_token_get_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenGetInfoQuery.ProtoReflect.Descriptor instead.
func (*TokenGetInfoQuery) Descriptor() ([]byte, []int) {
	return file_token_get_info_proto_rawDescGZIP(), []int{0}
}

func (x *TokenGetInfoQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TokenGetInfoQuery) GetToken() *TokenID {
	if x != nil {
		return x.Token
	}
	return nil
}

//*
// The metadata about a Token instance
type TokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// ID of the token instance
	TokenId *TokenID `protobuf:"bytes,1,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	//*
	// The name of the token. It is a string of ASCII only characters
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//*
	// The symbol of the token. It is a UTF-8 capitalized alphabetical string
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	//*
	// The number of decimal places a token is divisible by. Always 0 for tokens of type
	// NON_FUNGIBLE_UNIQUE
	Decimals uint32 `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	//*
	// For tokens of type FUNGIBLE_COMMON - the total supply of tokens that are currently in
	// circulation. For tokens of type NON_FUNGIBLE_UNIQUE - the number of NFTs created of this
	// token instance
	TotalSupply uint64 `protobuf:"varint,5,opt,name=totalSupply,proto3" json:"totalSupply,omitempty"`
	//*
	// The ID of the account which is set as Treasury
	Treasury *AccountID `protobuf:"bytes,6,opt,name=treasury,proto3" json:"treasury,omitempty"`
	//*
	// The key which can perform update/delete operations on the token. If empty, the token can be
	// perceived as immutable (not being able to be updated/deleted)
	AdminKey *Key `protobuf:"bytes,7,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	//*
	// The key which can grant or revoke KYC of an account for the token's transactions. If empty,
	// KYC is not required, and KYC grant or revoke operations are not possible.
	KycKey *Key `protobuf:"bytes,8,opt,name=kycKey,proto3" json:"kycKey,omitempty"`
	//*
	// The key which can freeze or unfreeze an account for token transactions. If empty, freezing is
	// not possible
	FreezeKey *Key `protobuf:"bytes,9,opt,name=freezeKey,proto3" json:"freezeKey,omitempty"`
	//*
	// The key which can wipe token balance of an account. If empty, wipe is not possible
	WipeKey *Key `protobuf:"bytes,10,opt,name=wipeKey,proto3" json:"wipeKey,omitempty"`
	//*
	// The key which can change the supply of a token. The key is used to sign Token Mint/Burn
	// operations
	SupplyKey *Key `protobuf:"bytes,11,opt,name=supplyKey,proto3" json:"supplyKey,omitempty"`
	//*
	// The default Freeze status (not applicable, frozen or unfrozen) of Hedera accounts relative to
	// this token. FreezeNotApplicable is returned if Token Freeze Key is empty. Frozen is returned
	// if Token Freeze Key is set and defaultFreeze is set to true. Unfrozen is returned if Token
	// Freeze Key is set and defaultFreeze is set to false
	DefaultFreezeStatus TokenFreezeStatus `protobuf:"varint,12,opt,name=defaultFreezeStatus,proto3,enum=proto.TokenFreezeStatus" json:"defaultFreezeStatus,omitempty"`
	//*
	// The default KYC status (KycNotApplicable or Revoked) of Hedera accounts relative to this
	// token. KycNotApplicable is returned if KYC key is not set, otherwise Revoked
	DefaultKycStatus TokenKycStatus `protobuf:"varint,13,opt,name=defaultKycStatus,proto3,enum=proto.TokenKycStatus" json:"defaultKycStatus,omitempty"`
	//*
	// Specifies whether the token was deleted or not
	Deleted bool `protobuf:"varint,14,opt,name=deleted,proto3" json:"deleted,omitempty"`
	//*
	// An account which will be automatically charged to renew the token's expiration, at
	// autoRenewPeriod interval
	AutoRenewAccount *AccountID `protobuf:"bytes,15,opt,name=autoRenewAccount,proto3" json:"autoRenewAccount,omitempty"`
	//*
	// The interval at which the auto-renew account will be charged to extend the token's expiry
	AutoRenewPeriod *Duration `protobuf:"bytes,16,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	//*
	// The epoch second at which the token will expire
	Expiry *Timestamp `protobuf:"bytes,17,opt,name=expiry,proto3" json:"expiry,omitempty"`
	//*
	// The memo associated with the token
	Memo string `protobuf:"bytes,18,opt,name=memo,proto3" json:"memo,omitempty"`
	//*
	// The token type
	TokenType TokenType `protobuf:"varint,19,opt,name=tokenType,proto3,enum=proto.TokenType" json:"tokenType,omitempty"`
	//*
	// The token supply type
	SupplyType TokenSupplyType `protobuf:"varint,20,opt,name=supplyType,proto3,enum=proto.TokenSupplyType" json:"supplyType,omitempty"`
	//*
	// For tokens of type FUNGIBLE_COMMON - The Maximum number of fungible tokens that can be in
	// circulation. For tokens of type NON_FUNGIBLE_UNIQUE - the maximum number of NFTs (serial
	// numbers) that can be in circulation
	MaxSupply int64 `protobuf:"varint,21,opt,name=maxSupply,proto3" json:"maxSupply,omitempty"`
	//*
	// The key which can change the custom fee schedule of the token; if not set, the fee schedule
	// is immutable
	FeeScheduleKey *Key `protobuf:"bytes,22,opt,name=fee_schedule_key,json=feeScheduleKey,proto3" json:"fee_schedule_key,omitempty"`
	//*
	// The custom fees to be assessed during a CryptoTransfer that transfers units of this token
	CustomFees []*CustomFee `protobuf:"bytes,23,rep,name=custom_fees,json=customFees,proto3" json:"custom_fees,omitempty"`
	//*
	// The Key which can pause and unpause the Token.
	PauseKey *Key `protobuf:"bytes,24,opt,name=pause_key,json=pauseKey,proto3" json:"pause_key,omitempty"`
	//*
	// Specifies whether the token is paused or not. PauseNotApplicable is returned if pauseKey is not set.
	PauseStatus TokenPauseStatus `protobuf:"varint,25,opt,name=pause_status,json=pauseStatus,proto3,enum=proto.TokenPauseStatus" json:"pause_status,omitempty"`
	//*
	// The ledger ID the response was returned from; please see <a href="https://github.com/hashgraph/hedera-improvement-proposal/blob/master/HIP/hip-198.md">HIP-198</a> for the network-specific IDs.
	LedgerId []byte `protobuf:"bytes,26,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
}

func (x *TokenInfo) Reset() {
	*x = TokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_get_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfo) ProtoMessage() {}

func (x *TokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_token_get_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInfo.ProtoReflect.Descriptor instead.
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return file_token_get_info_proto_rawDescGZIP(), []int{1}
}

func (x *TokenInfo) GetTokenId() *TokenID {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *TokenInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TokenInfo) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TokenInfo) GetDecimals() uint32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *TokenInfo) GetTotalSupply() uint64 {
	if x != nil {
		return x.TotalSupply
	}
	return 0
}

func (x *TokenInfo) GetTreasury() *AccountID {
	if x != nil {
		return x.Treasury
	}
	return nil
}

func (x *TokenInfo) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *TokenInfo) GetKycKey() *Key {
	if x != nil {
		return x.KycKey
	}
	return nil
}

func (x *TokenInfo) GetFreezeKey() *Key {
	if x != nil {
		return x.FreezeKey
	}
	return nil
}

func (x *TokenInfo) GetWipeKey() *Key {
	if x != nil {
		return x.WipeKey
	}
	return nil
}

func (x *TokenInfo) GetSupplyKey() *Key {
	if x != nil {
		return x.SupplyKey
	}
	return nil
}

func (x *TokenInfo) GetDefaultFreezeStatus() TokenFreezeStatus {
	if x != nil {
		return x.DefaultFreezeStatus
	}
	return TokenFreezeStatus_FreezeNotApplicable
}

func (x *TokenInfo) GetDefaultKycStatus() TokenKycStatus {
	if x != nil {
		return x.DefaultKycStatus
	}
	return TokenKycStatus_KycNotApplicable
}

func (x *TokenInfo) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *TokenInfo) GetAutoRenewAccount() *AccountID {
	if x != nil {
		return x.AutoRenewAccount
	}
	return nil
}

func (x *TokenInfo) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *TokenInfo) GetExpiry() *Timestamp {
	if x != nil {
		return x.Expiry
	}
	return nil
}

func (x *TokenInfo) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *TokenInfo) GetTokenType() TokenType {
	if x != nil {
		return x.TokenType
	}
	return TokenType_FUNGIBLE_COMMON
}

func (x *TokenInfo) GetSupplyType() TokenSupplyType {
	if x != nil {
		return x.SupplyType
	}
	return TokenSupplyType_INFINITE
}

func (x *TokenInfo) GetMaxSupply() int64 {
	if x != nil {
		return x.MaxSupply
	}
	return 0
}

func (x *TokenInfo) GetFeeScheduleKey() *Key {
	if x != nil {
		return x.FeeScheduleKey
	}
	return nil
}

func (x *TokenInfo) GetCustomFees() []*CustomFee {
	if x != nil {
		return x.CustomFees
	}
	return nil
}

func (x *TokenInfo) GetPauseKey() *Key {
	if x != nil {
		return x.PauseKey
	}
	return nil
}

func (x *TokenInfo) GetPauseStatus() TokenPauseStatus {
	if x != nil {
		return x.PauseStatus
	}
	return TokenPauseStatus_PauseNotApplicable
}

func (x *TokenInfo) GetLedgerId() []byte {
	if x != nil {
		return x.LedgerId
	}
	return nil
}

//*
// Response when the client sends the node TokenGetInfoQuery
type TokenGetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The information requested about this token instance
	TokenInfo *TokenInfo `protobuf:"bytes,2,opt,name=tokenInfo,proto3" json:"tokenInfo,omitempty"`
}

func (x *TokenGetInfoResponse) Reset() {
	*x = TokenGetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_get_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenGetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenGetInfoResponse) ProtoMessage() {}

func (x *TokenGetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_get_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenGetInfoResponse.ProtoReflect.Descriptor instead.
func (*TokenGetInfoResponse) Descriptor() ([]byte, []int) {
	return file_token_get_info_proto_rawDescGZIP(), []int{2}
}

func (x *TokenGetInfoResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TokenGetInfoResponse) GetTokenInfo() *TokenInfo {
	if x != nil {
		return x.TokenInfo
	}
	return nil
}

var File_token_get_info_proto protoreflect.FileDescriptor

var file_token_get_info_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x65, 0x0a, 0x11, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xe4, 0x08, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x08, 0x74, 0x72, 0x65, 0x61,
	0x73, 0x75, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x08, 0x74, 0x72,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b,
	0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x22,
	0x0a, 0x06, 0x6b, 0x79, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x06, 0x6b, 0x79, 0x63, 0x4b,
	0x65, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x4b, 0x65, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65,
	0x79, 0x52, 0x09, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x07,
	0x77, 0x69, 0x70, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x77, 0x69, 0x70, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x4b, 0x65, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65,
	0x79, 0x52, 0x09, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x4a, 0x0a, 0x13,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x13, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x72, 0x65, 0x65,
	0x7a, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x4b, 0x79, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x4b, 0x79, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x4b, 0x79, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e,
	0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x52, 0x10, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x61,
	0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x28,
	0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x12, 0x2e, 0x0a, 0x09,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x0a,
	0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x53, 0x75, 0x70, 0x70, 0x6c,
	0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x53, 0x75, 0x70, 0x70,
	0x6c, 0x79, 0x12, 0x34, 0x0a, 0x10, 0x66, 0x65, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x0e, 0x66, 0x65, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65, 0x65, 0x52,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x09, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x70, 0x61, 0x75, 0x73,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x0c, 0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61, 0x75, 0x73, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0b, 0x70, 0x61, 0x75, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x49, 0x64, 0x22, 0x75, 0x0a,
	0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_get_info_proto_rawDescOnce sync.Once
	file_token_get_info_proto_rawDescData = file_token_get_info_proto_rawDesc
)

func file_token_get_info_proto_rawDescGZIP() []byte {
	file_token_get_info_proto_rawDescOnce.Do(func() {
		file_token_get_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_get_info_proto_rawDescData)
	})
	return file_token_get_info_proto_rawDescData
}

var file_token_get_info_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_token_get_info_proto_goTypes = []interface{}{
	(*TokenGetInfoQuery)(nil),    // 0: proto.TokenGetInfoQuery
	(*TokenInfo)(nil),            // 1: proto.TokenInfo
	(*TokenGetInfoResponse)(nil), // 2: proto.TokenGetInfoResponse
	(*QueryHeader)(nil),          // 3: proto.QueryHeader
	(*TokenID)(nil),              // 4: proto.TokenID
	(*AccountID)(nil),            // 5: proto.AccountID
	(*Key)(nil),                  // 6: proto.Key
	(TokenFreezeStatus)(0),       // 7: proto.TokenFreezeStatus
	(TokenKycStatus)(0),          // 8: proto.TokenKycStatus
	(*Duration)(nil),             // 9: proto.Duration
	(*Timestamp)(nil),            // 10: proto.Timestamp
	(TokenType)(0),               // 11: proto.TokenType
	(TokenSupplyType)(0),         // 12: proto.TokenSupplyType
	(*CustomFee)(nil),            // 13: proto.CustomFee
	(TokenPauseStatus)(0),        // 14: proto.TokenPauseStatus
	(*ResponseHeader)(nil),       // 15: proto.ResponseHeader
}
var file_token_get_info_proto_depIdxs = []int32{
	3,  // 0: proto.TokenGetInfoQuery.header:type_name -> proto.QueryHeader
	4,  // 1: proto.TokenGetInfoQuery.token:type_name -> proto.TokenID
	4,  // 2: proto.TokenInfo.tokenId:type_name -> proto.TokenID
	5,  // 3: proto.TokenInfo.treasury:type_name -> proto.AccountID
	6,  // 4: proto.TokenInfo.adminKey:type_name -> proto.Key
	6,  // 5: proto.TokenInfo.kycKey:type_name -> proto.Key
	6,  // 6: proto.TokenInfo.freezeKey:type_name -> proto.Key
	6,  // 7: proto.TokenInfo.wipeKey:type_name -> proto.Key
	6,  // 8: proto.TokenInfo.supplyKey:type_name -> proto.Key
	7,  // 9: proto.TokenInfo.defaultFreezeStatus:type_name -> proto.TokenFreezeStatus
	8,  // 10: proto.TokenInfo.defaultKycStatus:type_name -> proto.TokenKycStatus
	5,  // 11: proto.TokenInfo.autoRenewAccount:type_name -> proto.AccountID
	9,  // 12: proto.TokenInfo.autoRenewPeriod:type_name -> proto.Duration
	10, // 13: proto.TokenInfo.expiry:type_name -> proto.Timestamp
	11, // 14: proto.TokenInfo.tokenType:type_name -> proto.TokenType
	12, // 15: proto.TokenInfo.supplyType:type_name -> proto.TokenSupplyType
	6,  // 16: proto.TokenInfo.fee_schedule_key:type_name -> proto.Key
	13, // 17: proto.TokenInfo.custom_fees:type_name -> proto.CustomFee
	6,  // 18: proto.TokenInfo.pause_key:type_name -> proto.Key
	14, // 19: proto.TokenInfo.pause_status:type_name -> proto.TokenPauseStatus
	15, // 20: proto.TokenGetInfoResponse.header:type_name -> proto.ResponseHeader
	1,  // 21: proto.TokenGetInfoResponse.tokenInfo:type_name -> proto.TokenInfo
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_token_get_info_proto_init() }
func file_token_get_info_proto_init() {
	if File_token_get_info_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_custom_fees_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	file_timestamp_proto_init()
	file_duration_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_token_get_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenGetInfoQuery); i {
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
		file_token_get_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenInfo); i {
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
		file_token_get_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenGetInfoResponse); i {
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
			RawDescriptor: file_token_get_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_get_info_proto_goTypes,
		DependencyIndexes: file_token_get_info_proto_depIdxs,
		MessageInfos:      file_token_get_info_proto_msgTypes,
	}.Build()
	File_token_get_info_proto = out.File
	file_token_get_info_proto_rawDesc = nil
	file_token_get_info_proto_goTypes = nil
	file_token_get_info_proto_depIdxs = nil
}
