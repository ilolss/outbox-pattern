// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: receiver.proto

package protogen

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

// ReceiverClient is the client API for Receiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReceiverClient interface {
	Receive(ctx context.Context, in *Task, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type receiverClient struct {
	cc grpc.ClientConnInterface
}

func NewReceiverClient(cc grpc.ClientConnInterface) ReceiverClient {
	return &receiverClient{cc}
}

func (c *receiverClient) Receive(ctx context.Context, in *Task, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/receiver.Receiver/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReceiverServer is the server API for Receiver service.
// All implementations should embed UnimplementedReceiverServer
// for forward compatibility
type ReceiverServer interface {
	Receive(context.Context, *Task) (*emptypb.Empty, error)
}

// UnimplementedReceiverServer should be embedded to have forward compatible implementations.
type UnimplementedReceiverServer struct {
}

func (UnimplementedReceiverServer) Receive(context.Context, *Task) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}

// UnsafeReceiverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReceiverServer will
// result in compilation errors.
type UnsafeReceiverServer interface {
	mustEmbedUnimplementedReceiverServer()
}

func RegisterReceiverServer(s grpc.ServiceRegistrar, srv ReceiverServer) {
	s.RegisterService(&Receiver_ServiceDesc, srv)
}

func _Receiver_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/receiver.Receiver/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverServer).Receive(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

// Receiver_ServiceDesc is the grpc.ServiceDesc for Receiver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Receiver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "receiver.Receiver",
	HandlerType: (*ReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Receive",
			Handler:    _Receiver_Receive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "receiver.proto",
}