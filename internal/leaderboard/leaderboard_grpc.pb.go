// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: internal/leaderboard/leaderboard.proto

package leaderboard

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

// LeaderboardServiceClient is the client API for LeaderboardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeaderboardServiceClient interface {
	UpdateScore(ctx context.Context, in *UpdateScoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPlayerRank(ctx context.Context, in *GetPlayerRankRequest, opts ...grpc.CallOption) (*RankInfo, error)
	GetTopN(ctx context.Context, in *GetTopNRequest, opts ...grpc.CallOption) (*PlayerRankRangeResponse, error)
	GetPlayerRankRange(ctx context.Context, in *GetPlayerRankRangeRequest, opts ...grpc.CallOption) (*PlayerRankRangeResponse, error)
}

type leaderboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeaderboardServiceClient(cc grpc.ClientConnInterface) LeaderboardServiceClient {
	return &leaderboardServiceClient{cc}
}

func (c *leaderboardServiceClient) UpdateScore(ctx context.Context, in *UpdateScoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/leaderboard.LeaderboardService/UpdateScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderboardServiceClient) GetPlayerRank(ctx context.Context, in *GetPlayerRankRequest, opts ...grpc.CallOption) (*RankInfo, error) {
	out := new(RankInfo)
	err := c.cc.Invoke(ctx, "/leaderboard.LeaderboardService/GetPlayerRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderboardServiceClient) GetTopN(ctx context.Context, in *GetTopNRequest, opts ...grpc.CallOption) (*PlayerRankRangeResponse, error) {
	out := new(PlayerRankRangeResponse)
	err := c.cc.Invoke(ctx, "/leaderboard.LeaderboardService/GetTopN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderboardServiceClient) GetPlayerRankRange(ctx context.Context, in *GetPlayerRankRangeRequest, opts ...grpc.CallOption) (*PlayerRankRangeResponse, error) {
	out := new(PlayerRankRangeResponse)
	err := c.cc.Invoke(ctx, "/leaderboard.LeaderboardService/GetPlayerRankRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeaderboardServiceServer is the server API for LeaderboardService service.
// All implementations must embed UnimplementedLeaderboardServiceServer
// for forward compatibility
type LeaderboardServiceServer interface {
	UpdateScore(context.Context, *UpdateScoreRequest) (*emptypb.Empty, error)
	GetPlayerRank(context.Context, *GetPlayerRankRequest) (*RankInfo, error)
	GetTopN(context.Context, *GetTopNRequest) (*PlayerRankRangeResponse, error)
	GetPlayerRankRange(context.Context, *GetPlayerRankRangeRequest) (*PlayerRankRangeResponse, error)
	mustEmbedUnimplementedLeaderboardServiceServer()
}

// UnimplementedLeaderboardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLeaderboardServiceServer struct {
}

func (UnimplementedLeaderboardServiceServer) UpdateScore(context.Context, *UpdateScoreRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScore not implemented")
}
func (UnimplementedLeaderboardServiceServer) GetPlayerRank(context.Context, *GetPlayerRankRequest) (*RankInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerRank not implemented")
}
func (UnimplementedLeaderboardServiceServer) GetTopN(context.Context, *GetTopNRequest) (*PlayerRankRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopN not implemented")
}
func (UnimplementedLeaderboardServiceServer) GetPlayerRankRange(context.Context, *GetPlayerRankRangeRequest) (*PlayerRankRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerRankRange not implemented")
}
func (UnimplementedLeaderboardServiceServer) mustEmbedUnimplementedLeaderboardServiceServer() {}

// UnsafeLeaderboardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeaderboardServiceServer will
// result in compilation errors.
type UnsafeLeaderboardServiceServer interface {
	mustEmbedUnimplementedLeaderboardServiceServer()
}

func RegisterLeaderboardServiceServer(s grpc.ServiceRegistrar, srv LeaderboardServiceServer) {
	s.RegisterService(&LeaderboardService_ServiceDesc, srv)
}

func _LeaderboardService_UpdateScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServiceServer).UpdateScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leaderboard.LeaderboardService/UpdateScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServiceServer).UpdateScore(ctx, req.(*UpdateScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaderboardService_GetPlayerRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerRankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServiceServer).GetPlayerRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leaderboard.LeaderboardService/GetPlayerRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServiceServer).GetPlayerRank(ctx, req.(*GetPlayerRankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaderboardService_GetTopN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopNRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServiceServer).GetTopN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leaderboard.LeaderboardService/GetTopN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServiceServer).GetTopN(ctx, req.(*GetTopNRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaderboardService_GetPlayerRankRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerRankRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderboardServiceServer).GetPlayerRankRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/leaderboard.LeaderboardService/GetPlayerRankRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderboardServiceServer).GetPlayerRankRange(ctx, req.(*GetPlayerRankRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LeaderboardService_ServiceDesc is the grpc.ServiceDesc for LeaderboardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LeaderboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "leaderboard.LeaderboardService",
	HandlerType: (*LeaderboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateScore",
			Handler:    _LeaderboardService_UpdateScore_Handler,
		},
		{
			MethodName: "GetPlayerRank",
			Handler:    _LeaderboardService_GetPlayerRank_Handler,
		},
		{
			MethodName: "GetTopN",
			Handler:    _LeaderboardService_GetTopN_Handler,
		},
		{
			MethodName: "GetPlayerRankRange",
			Handler:    _LeaderboardService_GetPlayerRankRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/leaderboard/leaderboard.proto",
}
