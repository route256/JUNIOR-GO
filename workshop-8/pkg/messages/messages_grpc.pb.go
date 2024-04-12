// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: api/messages.proto

package messages

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
	Messages_GetMessagesSummary_FullMethodName = "/messages.Messages/GetMessagesSummary"
	Messages_PullMessages_FullMethodName       = "/messages.Messages/PullMessages"
	Messages_PushMessages_FullMethodName       = "/messages.Messages/PushMessages"
	Messages_ExchangeMessages_FullMethodName   = "/messages.Messages/ExchangeMessages"
)

// MessagesClient is the client API for Messages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagesClient interface {
	GetMessagesSummary(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MessagesSummary, error)
	// server-to-client stream
	PullMessages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Messages_PullMessagesClient, error)
	// client-to-server stream
	PushMessages(ctx context.Context, opts ...grpc.CallOption) (Messages_PushMessagesClient, error)
	// bi-directional
	ExchangeMessages(ctx context.Context, opts ...grpc.CallOption) (Messages_ExchangeMessagesClient, error)
}

type messagesClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagesClient(cc grpc.ClientConnInterface) MessagesClient {
	return &messagesClient{cc}
}

func (c *messagesClient) GetMessagesSummary(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MessagesSummary, error) {
	out := new(MessagesSummary)
	err := c.cc.Invoke(ctx, Messages_GetMessagesSummary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) PullMessages(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Messages_PullMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messages_ServiceDesc.Streams[0], Messages_PullMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messagesPullMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messages_PullMessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type messagesPullMessagesClient struct {
	grpc.ClientStream
}

func (x *messagesPullMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messagesClient) PushMessages(ctx context.Context, opts ...grpc.CallOption) (Messages_PushMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messages_ServiceDesc.Streams[1], Messages_PushMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messagesPushMessagesClient{stream}
	return x, nil
}

type Messages_PushMessagesClient interface {
	Send(*Message) error
	CloseAndRecv() (*MessagesSummary, error)
	grpc.ClientStream
}

type messagesPushMessagesClient struct {
	grpc.ClientStream
}

func (x *messagesPushMessagesClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messagesPushMessagesClient) CloseAndRecv() (*MessagesSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessagesSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messagesClient) ExchangeMessages(ctx context.Context, opts ...grpc.CallOption) (Messages_ExchangeMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messages_ServiceDesc.Streams[2], Messages_ExchangeMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &messagesExchangeMessagesClient{stream}
	return x, nil
}

type Messages_ExchangeMessagesClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type messagesExchangeMessagesClient struct {
	grpc.ClientStream
}

func (x *messagesExchangeMessagesClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messagesExchangeMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessagesServer is the server API for Messages service.
// All implementations must embed UnimplementedMessagesServer
// for forward compatibility
type MessagesServer interface {
	GetMessagesSummary(context.Context, *emptypb.Empty) (*MessagesSummary, error)
	// server-to-client stream
	PullMessages(*emptypb.Empty, Messages_PullMessagesServer) error
	// client-to-server stream
	PushMessages(Messages_PushMessagesServer) error
	// bi-directional
	ExchangeMessages(Messages_ExchangeMessagesServer) error
	mustEmbedUnimplementedMessagesServer()
}

// UnimplementedMessagesServer must be embedded to have forward compatible implementations.
type UnimplementedMessagesServer struct {
}

func (UnimplementedMessagesServer) GetMessagesSummary(context.Context, *emptypb.Empty) (*MessagesSummary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessagesSummary not implemented")
}
func (UnimplementedMessagesServer) PullMessages(*emptypb.Empty, Messages_PullMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method PullMessages not implemented")
}
func (UnimplementedMessagesServer) PushMessages(Messages_PushMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method PushMessages not implemented")
}
func (UnimplementedMessagesServer) ExchangeMessages(Messages_ExchangeMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ExchangeMessages not implemented")
}
func (UnimplementedMessagesServer) mustEmbedUnimplementedMessagesServer() {}

// UnsafeMessagesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagesServer will
// result in compilation errors.
type UnsafeMessagesServer interface {
	mustEmbedUnimplementedMessagesServer()
}

func RegisterMessagesServer(s grpc.ServiceRegistrar, srv MessagesServer) {
	s.RegisterService(&Messages_ServiceDesc, srv)
}

func _Messages_GetMessagesSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).GetMessagesSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Messages_GetMessagesSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).GetMessagesSummary(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_PullMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagesServer).PullMessages(m, &messagesPullMessagesServer{stream})
}

type Messages_PullMessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type messagesPullMessagesServer struct {
	grpc.ServerStream
}

func (x *messagesPullMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Messages_PushMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagesServer).PushMessages(&messagesPushMessagesServer{stream})
}

type Messages_PushMessagesServer interface {
	SendAndClose(*MessagesSummary) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type messagesPushMessagesServer struct {
	grpc.ServerStream
}

func (x *messagesPushMessagesServer) SendAndClose(m *MessagesSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messagesPushMessagesServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Messages_ExchangeMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagesServer).ExchangeMessages(&messagesExchangeMessagesServer{stream})
}

type Messages_ExchangeMessagesServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type messagesExchangeMessagesServer struct {
	grpc.ServerStream
}

func (x *messagesExchangeMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messagesExchangeMessagesServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Messages_ServiceDesc is the grpc.ServiceDesc for Messages service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messages_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messages.Messages",
	HandlerType: (*MessagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessagesSummary",
			Handler:    _Messages_GetMessagesSummary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PullMessages",
			Handler:       _Messages_PullMessages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PushMessages",
			Handler:       _Messages_PushMessages_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ExchangeMessages",
			Handler:       _Messages_ExchangeMessages_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/messages.proto",
}
