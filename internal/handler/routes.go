// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"xinclockin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/clockin",
				Handler: CreateXinClockInHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/clockin/:id",
				Handler: DeleteXinClockInHandler(serverCtx),
			},
		},
	)
}