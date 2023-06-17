package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/lixvyang/xinclockin/internal/svc"
)

type DeleteXinClockInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteXinClockInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteXinClockInLogic {
	return &DeleteXinClockInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteXinClockInLogic) DeleteXinClockIn() error {
	// todo: add your logic here and delete this line

	return nil
}
