// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: chess.proto

package pb

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
	ChessService_Create_FullMethodName = "/ChessService/Create"
	ChessService_Move_FullMethodName   = "/ChessService/Move"
)

// ChessServiceClient is the client API for ChessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChessServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Move(ctx context.Context, opts ...grpc.CallOption) (ChessService_MoveClient, error)
}

type chessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChessServiceClient(cc grpc.ClientConnInterface) ChessServiceClient {
	return &chessServiceClient{cc}
}

func (c *chessServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, ChessService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chessServiceClient) Move(ctx context.Context, opts ...grpc.CallOption) (ChessService_MoveClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChessService_ServiceDesc.Streams[0], ChessService_Move_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chessServiceMoveClient{stream}
	return x, nil
}

type ChessService_MoveClient interface {
	Send(*MoveRequest) error
	Recv() (*MoveResponse, error)
	grpc.ClientStream
}

type chessServiceMoveClient struct {
	grpc.ClientStream
}

func (x *chessServiceMoveClient) Send(m *MoveRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chessServiceMoveClient) Recv() (*MoveResponse, error) {
	m := new(MoveResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChessServiceServer is the server API for ChessService service.
// All implementations must embed UnimplementedChessServiceServer
// for forward compatibility
type ChessServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Move(ChessService_MoveServer) error
	mustEmbedUnimplementedChessServiceServer()
}

// UnimplementedChessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChessServiceServer struct {
}

func (UnimplementedChessServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedChessServiceServer) Move(ChessService_MoveServer) error {
	return status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedChessServiceServer) mustEmbedUnimplementedChessServiceServer() {}

// UnsafeChessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChessServiceServer will
// result in compilation errors.
type UnsafeChessServiceServer interface {
	mustEmbedUnimplementedChessServiceServer()
}

func RegisterChessServiceServer(s grpc.ServiceRegistrar, srv ChessServiceServer) {
	s.RegisterService(&ChessService_ServiceDesc, srv)
}

func _ChessService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChessServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChessService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChessServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChessService_Move_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChessServiceServer).Move(&chessServiceMoveServer{stream})
}

type ChessService_MoveServer interface {
	Send(*MoveResponse) error
	Recv() (*MoveRequest, error)
	grpc.ServerStream
}

type chessServiceMoveServer struct {
	grpc.ServerStream
}

func (x *chessServiceMoveServer) Send(m *MoveResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chessServiceMoveServer) Recv() (*MoveRequest, error) {
	m := new(MoveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChessService_ServiceDesc is the grpc.ServiceDesc for ChessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChessService",
	HandlerType: (*ChessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ChessService_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Move",
			Handler:       _ChessService_Move_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chess.proto",
}
