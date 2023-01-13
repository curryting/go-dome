// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"book/service/user/rpc/internal/logic"
	"book/service/user/rpc/internal/svc"
	"book/service/user/rpc/pb"
)

type UserCenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserCenterServer
}

func NewUserCenterServer(svcCtx *svc.ServiceContext) *UserCenterServer {
	return &UserCenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UserCenterServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}