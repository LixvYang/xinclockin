package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"xinclockin/internal/svc"
)

type CreateXinClockInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateXinClockInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateXinClockInLogic {
	return &CreateXinClockInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateXinClockInLogic) CreateXinClockIn() error {
	// todo: add your logic here and delete this line

	return nil
}
