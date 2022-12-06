// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package usercenter

import (
	"context"

	"user-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserInfoReq  = pb.GetUserInfoReq
	GetUserInfoResp = pb.GetUserInfoResp

	UserCenter interface {
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	}

	defaultUserCenter struct {
		cli zrpc.Client
	}
)

func NewUserCenter(cli zrpc.Client) UserCenter {
	return &defaultUserCenter{
		cli: cli,
	}
}

func (m *defaultUserCenter) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewUserCenterClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}
