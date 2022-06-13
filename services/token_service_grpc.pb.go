// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: token_service.proto

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

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenServiceClient interface {
	//*
	// Creates a new Token by submitting the transaction
	CreateToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Updates the account by submitting the transaction
	UpdateToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Mints an amount of the token to the defined treasury account
	MintToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Burns an amount of the token from the defined treasury account
	BurnToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Deletes a Token
	DeleteToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Wipes the provided amount of tokens from the specified Account ID
	WipeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Freezes the transfer of tokens to or from the specified Account ID
	FreezeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Unfreezes the transfer of tokens to or from the specified Account ID
	UnfreezeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Flags the provided Account ID as having gone through KYC
	GrantKycToTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Removes the KYC flag of the provided Account ID
	RevokeKycFromTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Associates tokens to an account
	AssociateTokens(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Dissociates tokens from an account
	DissociateTokens(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Updates the custom fee schedule on a token
	UpdateTokenFeeSchedule(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//*
	// Retrieves the metadata of a token
	GetTokenInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Deprecated: Do not use.
	//*
	// (DEPRECATED) Gets info on NFTs N through M on the list of NFTs associated with a given account
	GetAccountNftInfos(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	//*
	// Retrieves the metadata of an NFT by TokenID and serial number
	GetTokenNftInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Deprecated: Do not use.
	//*
	// (DEPRECATED) Gets info on NFTs N through M on the list of NFTs associated with a given Token of type NON_FUNGIBLE
	GetTokenNftInfos(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// Pause the token
	PauseToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	//  Unpause the token
	UnpauseToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type tokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenServiceClient(cc grpc.ClientConnInterface) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) CreateToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/createToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) UpdateToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/updateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) MintToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/mintToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) BurnToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/burnToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) DeleteToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/deleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) WipeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/wipeTokenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) FreezeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/freezeTokenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) UnfreezeTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/unfreezeTokenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) GrantKycToTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/grantKycToTokenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) RevokeKycFromTokenAccount(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/revokeKycFromTokenAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) AssociateTokens(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/associateTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) DissociateTokens(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/dissociateTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) UpdateTokenFeeSchedule(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/updateTokenFeeSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) GetTokenInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.TokenService/getTokenInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *tokenServiceClient) GetAccountNftInfos(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.TokenService/getAccountNftInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) GetTokenNftInfo(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.TokenService/getTokenNftInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *tokenServiceClient) GetTokenNftInfos(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.TokenService/getTokenNftInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) PauseToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/pauseToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) UnpauseToken(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/proto.TokenService/unpauseToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
// All implementations must embed UnimplementedTokenServiceServer
// for forward compatibility
type TokenServiceServer interface {
	//*
	// Creates a new Token by submitting the transaction
	CreateToken(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Updates the account by submitting the transaction
	UpdateToken(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Mints an amount of the token to the defined treasury account
	MintToken(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Burns an amount of the token from the defined treasury account
	BurnToken(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Deletes a Token
	DeleteToken(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Wipes the provided amount of tokens from the specified Account ID
	WipeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Freezes the transfer of tokens to or from the specified Account ID
	FreezeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Unfreezes the transfer of tokens to or from the specified Account ID
	UnfreezeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Flags the provided Account ID as having gone through KYC
	GrantKycToTokenAccount(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Removes the KYC flag of the provided Account ID
	RevokeKycFromTokenAccount(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Associates tokens to an account
	AssociateTokens(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Dissociates tokens from an account
	DissociateTokens(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Updates the custom fee schedule on a token
	UpdateTokenFeeSchedule(context.Context, *Transaction) (*TransactionResponse, error)
	//*
	// Retrieves the metadata of a token
	GetTokenInfo(context.Context, *Query) (*Response, error)
	// Deprecated: Do not use.
	//*
	// (DEPRECATED) Gets info on NFTs N through M on the list of NFTs associated with a given account
	GetAccountNftInfos(context.Context, *Query) (*Response, error)
	//*
	// Retrieves the metadata of an NFT by TokenID and serial number
	GetTokenNftInfo(context.Context, *Query) (*Response, error)
	// Deprecated: Do not use.
	//*
	// (DEPRECATED) Gets info on NFTs N through M on the list of NFTs associated with a given Token of type NON_FUNGIBLE
	GetTokenNftInfos(context.Context, *Query) (*Response, error)
	// Pause the token
	PauseToken(context.Context, *Transaction) (*TransactionResponse, error)
	//  Unpause the token
	UnpauseToken(context.Context, *Transaction) (*TransactionResponse, error)
	mustEmbedUnimplementedTokenServiceServer()
}

// UnimplementedTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (UnimplementedTokenServiceServer) CreateToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedTokenServiceServer) UpdateToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToken not implemented")
}
func (UnimplementedTokenServiceServer) MintToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintToken not implemented")
}
func (UnimplementedTokenServiceServer) BurnToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnToken not implemented")
}
func (UnimplementedTokenServiceServer) DeleteToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (UnimplementedTokenServiceServer) WipeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WipeTokenAccount not implemented")
}
func (UnimplementedTokenServiceServer) FreezeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreezeTokenAccount not implemented")
}
func (UnimplementedTokenServiceServer) UnfreezeTokenAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfreezeTokenAccount not implemented")
}
func (UnimplementedTokenServiceServer) GrantKycToTokenAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantKycToTokenAccount not implemented")
}
func (UnimplementedTokenServiceServer) RevokeKycFromTokenAccount(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeKycFromTokenAccount not implemented")
}
func (UnimplementedTokenServiceServer) AssociateTokens(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssociateTokens not implemented")
}
func (UnimplementedTokenServiceServer) DissociateTokens(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DissociateTokens not implemented")
}
func (UnimplementedTokenServiceServer) UpdateTokenFeeSchedule(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTokenFeeSchedule not implemented")
}
func (UnimplementedTokenServiceServer) GetTokenInfo(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenInfo not implemented")
}
func (UnimplementedTokenServiceServer) GetAccountNftInfos(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountNftInfos not implemented")
}
func (UnimplementedTokenServiceServer) GetTokenNftInfo(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenNftInfo not implemented")
}
func (UnimplementedTokenServiceServer) GetTokenNftInfos(context.Context, *Query) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenNftInfos not implemented")
}
func (UnimplementedTokenServiceServer) PauseToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseToken not implemented")
}
func (UnimplementedTokenServiceServer) UnpauseToken(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnpauseToken not implemented")
}
func (UnimplementedTokenServiceServer) mustEmbedUnimplementedTokenServiceServer() {}

// UnsafeTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServiceServer will
// result in compilation errors.
type UnsafeTokenServiceServer interface {
	mustEmbedUnimplementedTokenServiceServer()
}

func RegisterTokenServiceServer(s grpc.ServiceRegistrar, srv TokenServiceServer) {
	s.RegisterService(&TokenService_ServiceDesc, srv)
}

func _TokenService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/createToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).CreateToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_UpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).UpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/updateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).UpdateToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_MintToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).MintToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/mintToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).MintToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_BurnToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).BurnToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/burnToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).BurnToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/deleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).DeleteToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_WipeTokenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).WipeTokenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/wipeTokenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).WipeTokenAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_FreezeTokenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).FreezeTokenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/freezeTokenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).FreezeTokenAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_UnfreezeTokenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).UnfreezeTokenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/unfreezeTokenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).UnfreezeTokenAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GrantKycToTokenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GrantKycToTokenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/grantKycToTokenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GrantKycToTokenAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_RevokeKycFromTokenAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).RevokeKycFromTokenAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/revokeKycFromTokenAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).RevokeKycFromTokenAccount(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_AssociateTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).AssociateTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/associateTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).AssociateTokens(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_DissociateTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).DissociateTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/dissociateTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).DissociateTokens(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_UpdateTokenFeeSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).UpdateTokenFeeSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/updateTokenFeeSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).UpdateTokenFeeSchedule(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GetTokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GetTokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/getTokenInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GetTokenInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GetAccountNftInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GetAccountNftInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/getAccountNftInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GetAccountNftInfos(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GetTokenNftInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GetTokenNftInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/getTokenNftInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GetTokenNftInfo(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_GetTokenNftInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).GetTokenNftInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/getTokenNftInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).GetTokenNftInfos(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_PauseToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).PauseToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/pauseToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).PauseToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_UnpauseToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).UnpauseToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TokenService/unpauseToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).UnpauseToken(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenService_ServiceDesc is the grpc.ServiceDesc for TokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createToken",
			Handler:    _TokenService_CreateToken_Handler,
		},
		{
			MethodName: "updateToken",
			Handler:    _TokenService_UpdateToken_Handler,
		},
		{
			MethodName: "mintToken",
			Handler:    _TokenService_MintToken_Handler,
		},
		{
			MethodName: "burnToken",
			Handler:    _TokenService_BurnToken_Handler,
		},
		{
			MethodName: "deleteToken",
			Handler:    _TokenService_DeleteToken_Handler,
		},
		{
			MethodName: "wipeTokenAccount",
			Handler:    _TokenService_WipeTokenAccount_Handler,
		},
		{
			MethodName: "freezeTokenAccount",
			Handler:    _TokenService_FreezeTokenAccount_Handler,
		},
		{
			MethodName: "unfreezeTokenAccount",
			Handler:    _TokenService_UnfreezeTokenAccount_Handler,
		},
		{
			MethodName: "grantKycToTokenAccount",
			Handler:    _TokenService_GrantKycToTokenAccount_Handler,
		},
		{
			MethodName: "revokeKycFromTokenAccount",
			Handler:    _TokenService_RevokeKycFromTokenAccount_Handler,
		},
		{
			MethodName: "associateTokens",
			Handler:    _TokenService_AssociateTokens_Handler,
		},
		{
			MethodName: "dissociateTokens",
			Handler:    _TokenService_DissociateTokens_Handler,
		},
		{
			MethodName: "updateTokenFeeSchedule",
			Handler:    _TokenService_UpdateTokenFeeSchedule_Handler,
		},
		{
			MethodName: "getTokenInfo",
			Handler:    _TokenService_GetTokenInfo_Handler,
		},
		{
			MethodName: "getAccountNftInfos",
			Handler:    _TokenService_GetAccountNftInfos_Handler,
		},
		{
			MethodName: "getTokenNftInfo",
			Handler:    _TokenService_GetTokenNftInfo_Handler,
		},
		{
			MethodName: "getTokenNftInfos",
			Handler:    _TokenService_GetTokenNftInfos_Handler,
		},
		{
			MethodName: "pauseToken",
			Handler:    _TokenService_PauseToken_Handler,
		},
		{
			MethodName: "unpauseToken",
			Handler:    _TokenService_UnpauseToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token_service.proto",
}
