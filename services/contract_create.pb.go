// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: contract_create.proto

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
// Start a new smart contract instance. After the instance is created, the ContractID for it is in
// the receipt, and can be retrieved by the Record or with a GetByKey query. The instance will run
// the bytecode stored in a previously created file, referenced either by FileID or by the
// transaction ID of the transaction that created the file
//
//
// The constructor will be executed using the given amount of gas, and any unspent gas will be
// refunded to the paying account. Constructor inputs come from the given constructorParameters.
//  - The instance will exist for autoRenewPeriod seconds. When that is reached, it will renew
//    itself for another autoRenewPeriod seconds by charging its associated cryptocurrency account
//    (which it creates here). If it has insufficient cryptocurrency to extend that long, it will
//    extend as long as it can. If its balance is zero, the instance will be deleted.
//
//  - A smart contract instance normally enforces rules, so "the code is law". For example, an
//    ERC-20 contract prevents a transfer from being undone without a signature by the recipient of
//    the transfer. This is always enforced if the contract instance was created with the adminKeys
//    being null. But for some uses, it might be desirable to create something like an ERC-20
//    contract that has a specific group of trusted individuals who can act as a "supreme court"
//    with the ability to override the normal operation, when a sufficient number of them agree to
//    do so. If adminKeys is not null, then they can sign a transaction that can change the state of
//    the smart contract in arbitrary ways, such as to reverse a transaction that violates some
//    standard of behavior that is not covered by the code itself. The admin keys can also be used
//    to change the autoRenewPeriod, and change the adminKeys field itself. The API currently does
//    not implement this ability. But it does allow the adminKeys field to be set and queried, and
//    will in the future implement such admin abilities for any instance that has a non-null
//    adminKeys.
//
//  - If this constructor stores information, it is charged gas to store it. There is a fee in hbars
//    to maintain that storage until the expiration time, and that fee is added as part of the
//    transaction fee.
//
//  - An entity (account, file, or smart contract instance) must be created in a particular realm.
//    If the realmID is left null, then a new realm will be created with the given admin key. If a
//    new realm has a null adminKey, then anyone can create/modify/delete entities in that realm.
//    But if an admin key is given, then any transaction to create/modify/delete an entity in that
//    realm must be signed by that key, though anyone can still call functions on smart contract
//    instances that exist in that realm. A realm ceases to exist when everything within it has
//    expired and no longer exists.
//
//  - The current API ignores shardID, realmID, and newRealmAdminKey, and creates everything in
//    shard 0 and realm 0, with a null key. Future versions of the API will support multiple realms
//    and multiple shards.
//
//  - The optional memo field can contain a string whose length is up to 100 bytes. That is the size
//    after Unicode NFD then UTF-8 conversion. This field can be used to describe the smart contract.
//    It could also be used for other purposes. One recommended purpose is to hold a hexadecimal
//    string that is the SHA-384 hash of a PDF file containing a human-readable legal contract. Then,
//    if the admin keys are the public keys of human arbitrators, they can use that legal document to
//    guide their decisions during a binding arbitration tribunal, convened to consider any changes
//    to the smart contract in the future. The memo field can only be changed using the admin keys.
//    If there are no admin keys, then it cannot be changed after the smart contract is created.
type ContractCreateTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// the file containing the smart contract byte code. A copy will be made and held by the
	// contract instance, and have the same expiration time as the instance. The file is referenced
	// one of two ways:
	FileID *FileID `protobuf:"bytes,1,opt,name=fileID,proto3" json:"fileID,omitempty"`
	//*
	// the state of the instance and its fields can be modified arbitrarily if this key signs a
	// transaction to modify it. If this is null, then such modifications are not possible, and
	// there is no administrator that can override the normal operation of this smart contract
	// instance. Note that if it is created with no admin keys, then there is no administrator to
	// authorize changing the admin keys, so there can never be any admin keys for that instance.
	AdminKey *Key `protobuf:"bytes,3,opt,name=adminKey,proto3" json:"adminKey,omitempty"`
	//*
	// gas to run the constructor
	Gas int64 `protobuf:"varint,4,opt,name=gas,proto3" json:"gas,omitempty"`
	//*
	// initial number of tinybars to put into the cryptocurrency account associated with and owned
	// by the smart contract
	InitialBalance int64 `protobuf:"varint,5,opt,name=initialBalance,proto3" json:"initialBalance,omitempty"`
	//*
	// ID of the account to which this account is proxy staked. If proxyAccountID is null, or is an
	// invalid account, or is an account that isn't a node, then this account is automatically proxy
	// staked to a node chosen by the network, but without earning payments. If the proxyAccountID
	// account refuses to accept proxy staking , or if it is not currently running a node, then it
	// will behave as if  proxyAccountID was null.
	ProxyAccountID *AccountID `protobuf:"bytes,6,opt,name=proxyAccountID,proto3" json:"proxyAccountID,omitempty"`
	//*
	// the instance will charge its account every this many seconds to renew for this long
	AutoRenewPeriod *Duration `protobuf:"bytes,8,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	//*
	// parameters to pass to the constructor
	ConstructorParameters []byte `protobuf:"bytes,9,opt,name=constructorParameters,proto3" json:"constructorParameters,omitempty"`
	//*
	// shard in which to create this
	ShardID *ShardID `protobuf:"bytes,10,opt,name=shardID,proto3" json:"shardID,omitempty"`
	//*
	// realm in which to create this (leave this null to create a new realm)
	RealmID *RealmID `protobuf:"bytes,11,opt,name=realmID,proto3" json:"realmID,omitempty"`
	//*
	// if realmID is null, then this the admin key for the new realm that will be created
	NewRealmAdminKey *Key `protobuf:"bytes,12,opt,name=newRealmAdminKey,proto3" json:"newRealmAdminKey,omitempty"`
	//*
	// the memo that was submitted as part of the contract (max 100 bytes)
	Memo string `protobuf:"bytes,13,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *ContractCreateTransactionBody) Reset() {
	*x = ContractCreateTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractCreateTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractCreateTransactionBody) ProtoMessage() {}

func (x *ContractCreateTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_contract_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractCreateTransactionBody.ProtoReflect.Descriptor instead.
func (*ContractCreateTransactionBody) Descriptor() ([]byte, []int) {
	return file_contract_create_proto_rawDescGZIP(), []int{0}
}

func (x *ContractCreateTransactionBody) GetFileID() *FileID {
	if x != nil {
		return x.FileID
	}
	return nil
}

func (x *ContractCreateTransactionBody) GetAdminKey() *Key {
	if x != nil {
		return x.AdminKey
	}
	return nil
}

func (x *ContractCreateTransactionBody) GetGas() int64 {
	if x != nil {
		return x.Gas
	}
	return 0
}

func (x *ContractCreateTransactionBody) GetInitialBalance() int64 {
	if x != nil {
		return x.InitialBalance
	}
	return 0
}

func (x *ContractCreateTransactionBody) GetProxyAccountID() *AccountID {
	if x != nil {
		return x.ProxyAccountID
	}
	return nil
}

func (x *ContractCreateTransactionBody) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *ContractCreateTransactionBody) GetConstructorParameters() []byte {
	if x != nil {
		return x.ConstructorParameters
	}
	return nil
}

func (x *ContractCreateTransactionBody) GetShardID() *ShardID {
	if x != nil {
		return x.ShardID
	}
	return nil
}

func (x *ContractCreateTransactionBody) GetRealmID() *RealmID {
	if x != nil {
		return x.RealmID
	}
	return nil
}

func (x *ContractCreateTransactionBody) GetNewRealmAdminKey() *Key {
	if x != nil {
		return x.NewRealmAdminKey
	}
	return nil
}

func (x *ContractCreateTransactionBody) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

var File_contract_create_proto protoreflect.FileDescriptor

var file_contract_create_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf3, 0x03, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x6f, 0x64, 0x79, 0x12, 0x25, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x44, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x08, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x67, 0x61, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x49, 0x44, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x49,
	0x44, 0x12, 0x28, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x6c, 0x6d,
	0x49, 0x44, 0x52, 0x07, 0x72, 0x65, 0x61, 0x6c, 0x6d, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x10, 0x6e,
	0x65, 0x77, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4b, 0x65, 0x79, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65,
	0x79, 0x52, 0x10, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x61, 0x6c, 0x6d, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_create_proto_rawDescOnce sync.Once
	file_contract_create_proto_rawDescData = file_contract_create_proto_rawDesc
)

func file_contract_create_proto_rawDescGZIP() []byte {
	file_contract_create_proto_rawDescOnce.Do(func() {
		file_contract_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_create_proto_rawDescData)
	})
	return file_contract_create_proto_rawDescData
}

var file_contract_create_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contract_create_proto_goTypes = []interface{}{
	(*ContractCreateTransactionBody)(nil), // 0: proto.ContractCreateTransactionBody
	(*FileID)(nil),                        // 1: proto.FileID
	(*Key)(nil),                           // 2: proto.Key
	(*AccountID)(nil),                     // 3: proto.AccountID
	(*Duration)(nil),                      // 4: proto.Duration
	(*ShardID)(nil),                       // 5: proto.ShardID
	(*RealmID)(nil),                       // 6: proto.RealmID
}
var file_contract_create_proto_depIdxs = []int32{
	1, // 0: proto.ContractCreateTransactionBody.fileID:type_name -> proto.FileID
	2, // 1: proto.ContractCreateTransactionBody.adminKey:type_name -> proto.Key
	3, // 2: proto.ContractCreateTransactionBody.proxyAccountID:type_name -> proto.AccountID
	4, // 3: proto.ContractCreateTransactionBody.autoRenewPeriod:type_name -> proto.Duration
	5, // 4: proto.ContractCreateTransactionBody.shardID:type_name -> proto.ShardID
	6, // 5: proto.ContractCreateTransactionBody.realmID:type_name -> proto.RealmID
	2, // 6: proto.ContractCreateTransactionBody.newRealmAdminKey:type_name -> proto.Key
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_contract_create_proto_init() }
func file_contract_create_proto_init() {
	if File_contract_create_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_duration_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contract_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractCreateTransactionBody); i {
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
			RawDescriptor: file_contract_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_create_proto_goTypes,
		DependencyIndexes: file_contract_create_proto_depIdxs,
		MessageInfos:      file_contract_create_proto_msgTypes,
	}.Build()
	File_contract_create_proto = out.File
	file_contract_create_proto_rawDesc = nil
	file_contract_create_proto_goTypes = nil
	file_contract_create_proto_depIdxs = nil
}
