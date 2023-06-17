package svc

import (
	"github.com/lixvyang/xinclockin/internal/config"
	"github.com/lixvyang/xinclockin/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Admin  rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := new(ServiceContext)
	svc.Config = c
	svc.Admin = middleware.NewAdminMiddleware(c.AdminHeader.Xid).Handle

	return &ServiceContext{
		Config: c,
	}
}
