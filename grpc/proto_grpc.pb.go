// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: grpc/proto.proto

package proto

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

const (
	Mutual_RequestToken_FullMethodName = "/DSMutualExclusion.Mutual/RequestToken"
	Mutual_ReleaseToken_FullMethodName = "/DSMutualExclusion.Mutual/ReleaseToken"
)

// MutualClient is the client API for Mutual service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MutualClient interface {
	RequestToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	ReleaseToken(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error)
}

type mutualClient struct {
	cc grpc.ClientConnInterface
}

func NewMutualClient(cc grpc.ClientConnInterface) MutualClient {
	return &mutualClient{cc}
}

func (c *mutualClient) RequestToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, Mutual_RequestToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutualClient) ReleaseToken(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error) {
	out := new(ReleaseResponse)
	err := c.cc.Invoke(ctx, Mutual_ReleaseToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MutualServer is the server API for Mutual service.
// All implementations must embed UnimplementedMutualServer
// for forward compatibility
type MutualServer interface {
	RequestToken(context.Context, *TokenRequest) (*TokenResponse, error)
	ReleaseToken(context.Context, *ReleaseRequest) (*ReleaseResponse, error)
	mustEmbedUnimplementedMutualServer()
}

// UnimplementedMutualServer must be embedded to have forward compatible implementations.
type UnimplementedMutualServer struct {
}

func (UnimplementedMutualServer) RequestToken(context.Context, *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestToken not implemented")
}
func (UnimplementedMutualServer) ReleaseToken(context.Context, *ReleaseRequest) (*ReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseToken not implemented")
}
func (UnimplementedMutualServer) mustEmbedUnimplementedMutualServer() {}

// UnsafeMutualServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MutualServer will
// result in compilation errors.
type UnsafeMutualServer interface {
	mustEmbedUnimplementedMutualServer()
}

func RegisterMutualServer(s grpc.ServiceRegistrar, srv MutualServer) {
	s.RegisterService(&Mutual_ServiceDesc, srv)
}

func _Mutual_RequestToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutualServer).RequestToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mutual_RequestToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutualServer).RequestToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mutual_ReleaseToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutualServer).ReleaseToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mutual_ReleaseToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutualServer).ReleaseToken(ctx, req.(*ReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mutual_ServiceDesc is the grpc.ServiceDesc for Mutual service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mutual_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DSMutualExclusion.Mutual",
	HandlerType: (*MutualServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestToken",
			Handler:    _Mutual_RequestToken_Handler,
		},
		{
			MethodName: "ReleaseToken",
			Handler:    _Mutual_ReleaseToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
