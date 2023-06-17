package user

import (
	"context"

	"github.com/lixvyang/xinclockin/internal/svc"
	"github.com/lixvyang/xinclockin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *CreateXinClockInLogic) CreateXinClockIn(req *types.CreateClockinReq) (resp *types.CreateClockinResp, err error) {
	// todo: add your logic here and delete this line

	return
}
