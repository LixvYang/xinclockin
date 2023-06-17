package clockin

import (
	"context"

	"github.com/lixvyang/xinclockin/internal/svc"
	"github.com/lixvyang/xinclockin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContentLogic {
	return &CreateContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContentLogic) CreateContent(req *types.CreateContentReq) (resp *types.CreateContentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
