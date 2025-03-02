// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: custom_fees.proto

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
// A fraction of the transferred units of a token to assess as a fee. The amount assessed will never
// be less than the given minimum_amount, and never greater than the given maximum_amount.  The
// denomination is always units of the token to which this fractional fee is attached.
type FractionalFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The fraction of the transferred units to assess as a fee
	FractionalAmount *Fraction `protobuf:"bytes,1,opt,name=fractional_amount,json=fractionalAmount,proto3" json:"fractional_amount,omitempty"`
	//*
	// The minimum amount to assess
	MinimumAmount int64 `protobuf:"varint,2,opt,name=minimum_amount,json=minimumAmount,proto3" json:"minimum_amount,omitempty"`
	//*
	// The maximum amount to assess (zero implies no maximum)
	MaximumAmount int64 `protobuf:"varint,3,opt,name=maximum_amount,json=maximumAmount,proto3" json:"maximum_amount,omitempty"`
	//*
	// If true, assesses the fee to the sender, so the receiver gets the full amount from the token
	// transfer list, and the sender is charged an additional fee; if false, the receiver does NOT get
	// the full amount, but only what is left over after paying the fractional fee
	NetOfTransfers bool `protobuf:"varint,4,opt,name=net_of_transfers,json=netOfTransfers,proto3" json:"net_of_transfers,omitempty"`
}

func (x *FractionalFee) Reset() {
	*x = FractionalFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_fees_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FractionalFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FractionalFee) ProtoMessage() {}

func (x *FractionalFee) ProtoReflect() protoreflect.Message {
	mi := &file_custom_fees_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FractionalFee.ProtoReflect.Descriptor instead.
func (*FractionalFee) Descriptor() ([]byte, []int) {
	return file_custom_fees_proto_rawDescGZIP(), []int{0}
}

func (x *FractionalFee) GetFractionalAmount() *Fraction {
	if x != nil {
		return x.FractionalAmount
	}
	return nil
}

func (x *FractionalFee) GetMinimumAmount() int64 {
	if x != nil {
		return x.MinimumAmount
	}
	return 0
}

func (x *FractionalFee) GetMaximumAmount() int64 {
	if x != nil {
		return x.MaximumAmount
	}
	return 0
}

func (x *FractionalFee) GetNetOfTransfers() bool {
	if x != nil {
		return x.NetOfTransfers
	}
	return false
}

//*
// A fixed number of units (hbar or token) to assess as a fee during a CryptoTransfer that transfers
// units of the token to which this fixed fee is attached.
type FixedFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The number of units to assess as a fee
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	//*
	// The denomination of the fee; taken as hbar if left unset and, in a TokenCreate, taken as the id
	// of the newly created token if set to the sentinel value of 0.0.0
	DenominatingTokenId *TokenID `protobuf:"bytes,2,opt,name=denominating_token_id,json=denominatingTokenId,proto3" json:"denominating_token_id,omitempty"`
}

func (x *FixedFee) Reset() {
	*x = FixedFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_fees_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FixedFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixedFee) ProtoMessage() {}

func (x *FixedFee) ProtoReflect() protoreflect.Message {
	mi := &file_custom_fees_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixedFee.ProtoReflect.Descriptor instead.
func (*FixedFee) Descriptor() ([]byte, []int) {
	return file_custom_fees_proto_rawDescGZIP(), []int{1}
}

func (x *FixedFee) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *FixedFee) GetDenominatingTokenId() *TokenID {
	if x != nil {
		return x.DenominatingTokenId
	}
	return nil
}

//*
// A fee to assess during a CryptoTransfer that changes ownership of an NFT. Defines the fraction of
// the fungible value exchanged for an NFT that the ledger should collect as a royalty. ("Fungible
// value" includes both ℏ and units of fungible HTS tokens.) When the NFT sender does not receive
// any fungible value, the ledger will assess the fallback fee, if present, to the new NFT owner.
// Royalty fees can only be added to tokens of type type NON_FUNGIBLE_UNIQUE.
type RoyaltyFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The fraction of fungible value exchanged for an NFT to collect as royalty
	ExchangeValueFraction *Fraction `protobuf:"bytes,1,opt,name=exchange_value_fraction,json=exchangeValueFraction,proto3" json:"exchange_value_fraction,omitempty"`
	//*
	// If present, the fixed fee to assess to the NFT receiver when no fungible value is exchanged
	// with the sender
	FallbackFee *FixedFee `protobuf:"bytes,2,opt,name=fallback_fee,json=fallbackFee,proto3" json:"fallback_fee,omitempty"`
}

func (x *RoyaltyFee) Reset() {
	*x = RoyaltyFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_fees_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoyaltyFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoyaltyFee) ProtoMessage() {}

func (x *RoyaltyFee) ProtoReflect() protoreflect.Message {
	mi := &file_custom_fees_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoyaltyFee.ProtoReflect.Descriptor instead.
func (*RoyaltyFee) Descriptor() ([]byte, []int) {
	return file_custom_fees_proto_rawDescGZIP(), []int{2}
}

func (x *RoyaltyFee) GetExchangeValueFraction() *Fraction {
	if x != nil {
		return x.ExchangeValueFraction
	}
	return nil
}

func (x *RoyaltyFee) GetFallbackFee() *FixedFee {
	if x != nil {
		return x.FallbackFee
	}
	return nil
}

//*
// A transfer fee to assess during a CryptoTransfer that transfers units of the token to which the
// fee is attached. A custom fee may be either fixed or fractional, and must specify a fee collector
// account to receive the assessed fees. Only positive fees may be assessed.
type CustomFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Fee:
	//	*CustomFee_FixedFee
	//	*CustomFee_FractionalFee
	//	*CustomFee_RoyaltyFee
	Fee isCustomFee_Fee `protobuf_oneof:"fee"`
	//*
	// The account to receive the custom fee
	FeeCollectorAccountId *AccountID `protobuf:"bytes,3,opt,name=fee_collector_account_id,json=feeCollectorAccountId,proto3" json:"fee_collector_account_id,omitempty"`
}

func (x *CustomFee) Reset() {
	*x = CustomFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_fees_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomFee) ProtoMessage() {}

func (x *CustomFee) ProtoReflect() protoreflect.Message {
	mi := &file_custom_fees_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomFee.ProtoReflect.Descriptor instead.
func (*CustomFee) Descriptor() ([]byte, []int) {
	return file_custom_fees_proto_rawDescGZIP(), []int{3}
}

func (m *CustomFee) GetFee() isCustomFee_Fee {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (x *CustomFee) GetFixedFee() *FixedFee {
	if x, ok := x.GetFee().(*CustomFee_FixedFee); ok {
		return x.FixedFee
	}
	return nil
}

func (x *CustomFee) GetFractionalFee() *FractionalFee {
	if x, ok := x.GetFee().(*CustomFee_FractionalFee); ok {
		return x.FractionalFee
	}
	return nil
}

func (x *CustomFee) GetRoyaltyFee() *RoyaltyFee {
	if x, ok := x.GetFee().(*CustomFee_RoyaltyFee); ok {
		return x.RoyaltyFee
	}
	return nil
}

func (x *CustomFee) GetFeeCollectorAccountId() *AccountID {
	if x != nil {
		return x.FeeCollectorAccountId
	}
	return nil
}

type isCustomFee_Fee interface {
	isCustomFee_Fee()
}

type CustomFee_FixedFee struct {
	//*
	// Fixed fee to be charged
	FixedFee *FixedFee `protobuf:"bytes,1,opt,name=fixed_fee,json=fixedFee,proto3,oneof"`
}

type CustomFee_FractionalFee struct {
	//*
	// Fractional fee to be charged
	FractionalFee *FractionalFee `protobuf:"bytes,2,opt,name=fractional_fee,json=fractionalFee,proto3,oneof"`
}

type CustomFee_RoyaltyFee struct {
	//*
	// Royalty fee to be charged
	RoyaltyFee *RoyaltyFee `protobuf:"bytes,4,opt,name=royalty_fee,json=royaltyFee,proto3,oneof"`
}

func (*CustomFee_FixedFee) isCustomFee_Fee() {}

func (*CustomFee_FractionalFee) isCustomFee_Fee() {}

func (*CustomFee_RoyaltyFee) isCustomFee_Fee() {}

//*
// A custom transfer fee that was assessed during handling of a CryptoTransfer.
type AssessedCustomFee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The number of units assessed for the fee
	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	//*
	// The denomination of the fee; taken as hbar if left unset
	TokenId *TokenID `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	//*
	// The account to receive the assessed fee
	FeeCollectorAccountId *AccountID `protobuf:"bytes,3,opt,name=fee_collector_account_id,json=feeCollectorAccountId,proto3" json:"fee_collector_account_id,omitempty"`
	//*
	// The account(s) whose final balances would have been higher in the absence of this assessed fee
	EffectivePayerAccountId []*AccountID `protobuf:"bytes,4,rep,name=effective_payer_account_id,json=effectivePayerAccountId,proto3" json:"effective_payer_account_id,omitempty"`
}

func (x *AssessedCustomFee) Reset() {
	*x = AssessedCustomFee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_fees_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssessedCustomFee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssessedCustomFee) ProtoMessage() {}

func (x *AssessedCustomFee) ProtoReflect() protoreflect.Message {
	mi := &file_custom_fees_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssessedCustomFee.ProtoReflect.Descriptor instead.
func (*AssessedCustomFee) Descriptor() ([]byte, []int) {
	return file_custom_fees_proto_rawDescGZIP(), []int{4}
}

func (x *AssessedCustomFee) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *AssessedCustomFee) GetTokenId() *TokenID {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *AssessedCustomFee) GetFeeCollectorAccountId() *AccountID {
	if x != nil {
		return x.FeeCollectorAccountId
	}
	return nil
}

func (x *AssessedCustomFee) GetEffectivePayerAccountId() []*AccountID {
	if x != nil {
		return x.EffectivePayerAccountId
	}
	return nil
}

var File_custom_fees_proto protoreflect.FileDescriptor

var file_custom_fees_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x65, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01,
	0x0a, 0x0d, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x46, 0x65, 0x65, 0x12,
	0x3c, 0x0a, 0x11, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x66, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x61,
	0x78, 0x69, 0x6d, 0x75, 0x6d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6e,
	0x65, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6e, 0x65, 0x74, 0x4f, 0x66, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x73, 0x22, 0x66, 0x0a, 0x08, 0x46, 0x69, 0x78, 0x65, 0x64, 0x46, 0x65,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x15, 0x64, 0x65, 0x6e,
	0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x13, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x22, 0x89, 0x01,
	0x0a, 0x0a, 0x52, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x46, 0x65, 0x65, 0x12, 0x47, 0x0a, 0x17,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x66,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0c, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x78, 0x65, 0x64, 0x46, 0x65, 0x65, 0x52, 0x0b, 0x66, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x22, 0x82, 0x02, 0x0a, 0x09, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x46, 0x65, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x5f, 0x66, 0x65, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x69, 0x78, 0x65, 0x64, 0x46, 0x65, 0x65, 0x48, 0x00, 0x52, 0x08, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x46, 0x65, 0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x66, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x46, 0x65, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x46, 0x65, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x72, 0x6f, 0x79, 0x61, 0x6c, 0x74,
	0x79, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x46, 0x65, 0x65, 0x48, 0x00,
	0x52, 0x0a, 0x72, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x46, 0x65, 0x65, 0x12, 0x49, 0x0a, 0x18,
	0x66, 0x65, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x52, 0x15, 0x66, 0x65, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x05, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x22, 0xf0,
	0x01, 0x0a, 0x11, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x65, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x46, 0x65, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x08,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x07,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x18, 0x66, 0x65, 0x65, 0x5f, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x15, 0x66, 0x65, 0x65,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x4d, 0x0a, 0x1a, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x17, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x50, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68,
	0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_custom_fees_proto_rawDescOnce sync.Once
	file_custom_fees_proto_rawDescData = file_custom_fees_proto_rawDesc
)

func file_custom_fees_proto_rawDescGZIP() []byte {
	file_custom_fees_proto_rawDescOnce.Do(func() {
		file_custom_fees_proto_rawDescData = protoimpl.X.CompressGZIP(file_custom_fees_proto_rawDescData)
	})
	return file_custom_fees_proto_rawDescData
}

var file_custom_fees_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_custom_fees_proto_goTypes = []interface{}{
	(*FractionalFee)(nil),     // 0: proto.FractionalFee
	(*FixedFee)(nil),          // 1: proto.FixedFee
	(*RoyaltyFee)(nil),        // 2: proto.RoyaltyFee
	(*CustomFee)(nil),         // 3: proto.CustomFee
	(*AssessedCustomFee)(nil), // 4: proto.AssessedCustomFee
	(*Fraction)(nil),          // 5: proto.Fraction
	(*TokenID)(nil),           // 6: proto.TokenID
	(*AccountID)(nil),         // 7: proto.AccountID
}
var file_custom_fees_proto_depIdxs = []int32{
	5,  // 0: proto.FractionalFee.fractional_amount:type_name -> proto.Fraction
	6,  // 1: proto.FixedFee.denominating_token_id:type_name -> proto.TokenID
	5,  // 2: proto.RoyaltyFee.exchange_value_fraction:type_name -> proto.Fraction
	1,  // 3: proto.RoyaltyFee.fallback_fee:type_name -> proto.FixedFee
	1,  // 4: proto.CustomFee.fixed_fee:type_name -> proto.FixedFee
	0,  // 5: proto.CustomFee.fractional_fee:type_name -> proto.FractionalFee
	2,  // 6: proto.CustomFee.royalty_fee:type_name -> proto.RoyaltyFee
	7,  // 7: proto.CustomFee.fee_collector_account_id:type_name -> proto.AccountID
	6,  // 8: proto.AssessedCustomFee.token_id:type_name -> proto.TokenID
	7,  // 9: proto.AssessedCustomFee.fee_collector_account_id:type_name -> proto.AccountID
	7,  // 10: proto.AssessedCustomFee.effective_payer_account_id:type_name -> proto.AccountID
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_custom_fees_proto_init() }
func file_custom_fees_proto_init() {
	if File_custom_fees_proto != nil {
		return
	}
	file_basic_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_custom_fees_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FractionalFee); i {
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
		file_custom_fees_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FixedFee); i {
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
		file_custom_fees_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoyaltyFee); i {
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
		file_custom_fees_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomFee); i {
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
		file_custom_fees_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssessedCustomFee); i {
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
	file_custom_fees_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*CustomFee_FixedFee)(nil),
		(*CustomFee_FractionalFee)(nil),
		(*CustomFee_RoyaltyFee)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_custom_fees_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_custom_fees_proto_goTypes,
		DependencyIndexes: file_custom_fees_proto_depIdxs,
		MessageInfos:      file_custom_fees_proto_msgTypes,
	}.Build()
	File_custom_fees_proto = out.File
	file_custom_fees_proto_rawDesc = nil
	file_custom_fees_proto_goTypes = nil
	file_custom_fees_proto_depIdxs = nil
}
