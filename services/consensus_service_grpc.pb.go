// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConsensusServiceClient is the client API for ConsensusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsensusServiceClient interface {
	//*
	// Create a topic to be used for consensus.
	// If an autoRenewAccount is specified, that account must also sign this transaction.
	// If an adminKey is specified, the adminKey must sign the transaction.
	// On success, the resulting TransactionReceipt contains the newly created TopicId.
	// Request is [ConsensusCreateTopicTransactionBody](#proto.ConsensusCreateTopicTransactionBody)
	CreateTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Update a topic.
	// If there is no adminKey, the only authorized update (available to anyone) is to extend the expirationTime.
	// Otherwise transaction must be signed by the adminKey.
	// If an adminKey is updated, the transaction must be signed by the pre-update adminKey and post-update adminKey.
	// If a new autoRenewAccount is specified (not just being removed), that account must also sign the transaction.
	// Request is [ConsensusUpdateTopicTransactionBody](#proto.ConsensusUpdateTopicTransactionBody)
	UpdateTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Delete a topic. No more transactions or queries on the topic (via HAPI) will succeed.
	// If an adminKey is set, this transaction must be signed by that key.
	// If there is no adminKey, this transaction will fail UNAUTHORIZED.
	// Request is [ConsensusDeleteTopicTransactionBody](#proto.ConsensusDeleteTopicTransactionBody)
	DeleteTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Retrieve the latest state of a topic. This method is unrestricted and allowed on any topic by any payer account.
	// Deleted accounts will not be returned.
	// Request is [ConsensusGetTopicInfoQuery](#proto.ConsensusGetTopicInfoQuery)
	// Response is [ConsensusGetTopicInfoResponse](#proto.ConsensusGetTopicInfoResponse)
	GetTopicInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	//*
	// Submit a message for consensus.
	// Valid and authorized messages on valid topics will be ordered by the consensus service, gossipped to the
	// mirror net, and published (in order) to all subscribers (from the mirror net) on this topic.
	// The submitKey (if any) must sign this transaction.
	// On success, the resulting TransactionReceipt contains the topic's updated topicSequenceNumber and
	// topicRunningHash.
	// Request is [ConsensusSubmitMessageTransactionBody](#proto.ConsensusSubmitMessageTransactionBody)
	SubmitMessage(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type consensusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsensusServiceClient(cc grpc.ClientConnInterface) ConsensusServiceClient {
	return &consensusServiceClient{cc}
}

func (c *consensusServiceClient) CreateTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.ConsensusService/createTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusServiceClient) UpdateTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.ConsensusService/updateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusServiceClient) DeleteTopic(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.ConsensusService/deleteTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusServiceClient) GetTopicInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.ConsensusService/getTopicInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusServiceClient) SubmitMessage(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.ConsensusService/submitMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsensusServiceServer is the server API for ConsensusService service.
// All implementations must embed UnimplementedConsensusServiceServer
// for forward compatibility
type ConsensusServiceServer interface {
	//*
	// Create a topic to be used for consensus.
	// If an autoRenewAccount is specified, that account must also sign this transaction.
	// If an adminKey is specified, the adminKey must sign the transaction.
	// On success, the resulting TransactionReceipt contains the newly created TopicId.
	// Request is [ConsensusCreateTopicTransactionBody](#proto.ConsensusCreateTopicTransactionBody)
	CreateTopic(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Update a topic.
	// If there is no adminKey, the only authorized update (available to anyone) is to extend the expirationTime.
	// Otherwise transaction must be signed by the adminKey.
	// If an adminKey is updated, the transaction must be signed by the pre-update adminKey and post-update adminKey.
	// If a new autoRenewAccount is specified (not just being removed), that account must also sign the transaction.
	// Request is [ConsensusUpdateTopicTransactionBody](#proto.ConsensusUpdateTopicTransactionBody)
	UpdateTopic(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Delete a topic. No more transactions or queries on the topic (via HAPI) will succeed.
	// If an adminKey is set, this transaction must be signed by that key.
	// If there is no adminKey, this transaction will fail UNAUTHORIZED.
	// Request is [ConsensusDeleteTopicTransactionBody](#proto.ConsensusDeleteTopicTransactionBody)
	DeleteTopic(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Retrieve the latest state of a topic. This method is unrestricted and allowed on any topic by any payer account.
	// Deleted accounts will not be returned.
	// Request is [ConsensusGetTopicInfoQuery](#proto.ConsensusGetTopicInfoQuery)
	// Response is [ConsensusGetTopicInfoResponse](#proto.ConsensusGetTopicInfoResponse)
	GetTopicInfo(context.Context, *Query) (*Response, error)
	//*
	// Submit a message for consensus.
	// Valid and authorized messages on valid topics will be ordered by the consensus service, gossipped to the
	// mirror net, and published (in order) to all subscribers (from the mirror net) on this topic.
	// The submitKey (if any) must sign this transaction.
	// On success, the resulting TransactionReceipt contains the topic's updated topicSequenceNumber and
	// topicRunningHash.
	// Request is [ConsensusSubmitMessageTransactionBody](#proto.ConsensusSubmitMessageTransactionBody)
	SubmitMessage(context.Context, *Transaction) (*TransactionResponse, error)
	mustEmbedUnimplementedConsensusServiceServer()
}

// UnimplementedConsensusServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConsensusServiceServer struct {
}

func (UnimplementedConsensusServiceServer) CreateTopic(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (UnimplementedConsensusServiceServer) UpdateTopic(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopic not implemented")
}
func (UnimplementedConsensusServiceServer) DeleteTopic(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}
func (UnimplementedConsensusServiceServer) GetTopicInfo(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopicInfo not implemented")
}
func (UnimplementedConsensusServiceServer) SubmitMessage(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitMessage not implemented")
}
func (UnimplementedConsensusServiceServer) mustEmbedUnimplementedConsensusServiceServer() {}

// UnsafeConsensusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsensusServiceServer will
// result in compilation errors.
type UnsafeConsensusServiceServer interface {
	mustEmbedUnimplementedConsensusServiceServer()
}

func RegisterConsensusServiceServer(s grpc.ServiceRegistrar, srv ConsensusServiceServer) {
	s.RegisterService(&ConsensusService_ServiceDesc, srv)
}

func _ConsensusService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConsensusService/createTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).CreateTopic(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsensusService_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConsensusService/updateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).UpdateTopic(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsensusService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConsensusService/deleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).DeleteTopic(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsensusService_GetTopicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).GetTopicInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConsensusService/getTopicInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).GetTopicInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsensusService_SubmitMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServiceServer).SubmitMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ConsensusService/submitMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServiceServer).SubmitMessage(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsensusService_ServiceDesc is the grpc.ServiceDesc for ConsensusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsensusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ConsensusService",
	HandlerType: (*ConsensusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createTopic",
			Handler:    _ConsensusService_CreateTopic_Handler,
		},
		{
			MethodName: "updateTopic",
			Handler:    _ConsensusService_UpdateTopic_Handler,
		},
		{
			MethodName: "deleteTopic",
			Handler:    _ConsensusService_DeleteTopic_Handler,
		},
		{
			MethodName: "getTopicInfo",
			Handler:    _ConsensusService_GetTopicInfo_Handler,
		},
		{
			MethodName: "submitMessage",
			Handler:    _ConsensusService_SubmitMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consensus_service.proto",
}
