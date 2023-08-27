// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: email.proto

package srv

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	EmailSrv_SendCode_FullMethodName = "/EmailSrv/SendCode"
)

// EmailSrvClient is the client API for EmailSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailSrvClient interface {
	SendCode(ctx context.Context, in *Email, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type emailSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailSrvClient(cc grpc.ClientConnInterface) EmailSrvClient {
	return &emailSrvClient{cc}
}

func (c *emailSrvClient) SendCode(ctx context.Context, in *Email, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EmailSrv_SendCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailSrvServer is the server API for EmailSrv service.
// All implementations must embed UnimplementedEmailSrvServer
// for forward compatibility
type EmailSrvServer interface {
	SendCode(context.Context, *Email) (*emptypb.Empty, error)
	mustEmbedUnimplementedEmailSrvServer()
}

// UnimplementedEmailSrvServer must be embedded to have forward compatible implementations.
type UnimplementedEmailSrvServer struct {
}

func (UnimplementedEmailSrvServer) SendCode(context.Context, *Email) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCode not implemented")
}
func (UnimplementedEmailSrvServer) mustEmbedUnimplementedEmailSrvServer() {}

// UnsafeEmailSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailSrvServer will
// result in compilation errors.
type UnsafeEmailSrvServer interface {
	mustEmbedUnimplementedEmailSrvServer()
}

func RegisterEmailSrvServer(s grpc.ServiceRegistrar, srv EmailSrvServer) {
	s.RegisterService(&EmailSrv_ServiceDesc, srv)
}

func _EmailSrv_SendCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailSrvServer).SendCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailSrv_SendCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailSrvServer).SendCode(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailSrv_ServiceDesc is the grpc.ServiceDesc for EmailSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EmailSrv",
	HandlerType: (*EmailSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCode",
			Handler:    _EmailSrv_SendCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email.proto",
}