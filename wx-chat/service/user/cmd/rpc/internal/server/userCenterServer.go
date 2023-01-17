// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"wx-chat/service/user/cmd/rpc/internal/logic"
	"wx-chat/service/user/cmd/rpc/internal/svc"
	"wx-chat/service/user/cmd/rpc/pb"
)

type UserCenterServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedUserCenterServer
}

func NewUserCenterServer(svcCtx *svc.ServiceContext) *UserCenterServer {
	return &UserCenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UserCenterServer) FindUser(ctx context.Context, in *__.UserReq) (*__.UserRes, error) {
	l := logic.NewFindUserLogic(ctx, s.svcCtx)
	return l.FindUser(in)
}

func (s *UserCenterServer) AddUser(ctx context.Context, in *__.AddUserReq) (*__.CommonRes, error) {
	l := logic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}

func (s *UserCenterServer) UpdateUser(ctx context.Context, in *__.UserReq) (*__.CommonRes, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}