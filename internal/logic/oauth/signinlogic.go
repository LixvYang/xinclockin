package oauth

import (
	"context"

	"github.com/lixvyang/xinclockin/internal/svc"
	"github.com/lixvyang/xinclockin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SigninLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSigninLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SigninLogic {
	return &SigninLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SigninLogic) Signin(req *types.SigninReq) (resp *types.SigninResp, err error) {
	// todo: add your logic here and delete this line

	return
}
