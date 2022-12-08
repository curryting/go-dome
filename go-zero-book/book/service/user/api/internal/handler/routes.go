// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "book/service/user/api/internal/handler/user"
	"book/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/search",
				Handler: user.SearchHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/userInfo",
				Handler: user.UserInfoHandler(serverCtx),
			},
		},
	)
}
