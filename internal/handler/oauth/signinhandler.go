package oauth

import (
	"net/http"

	"github.com/lixvyang/xinclockin/internal/logic/oauth"
	"github.com/lixvyang/xinclockin/internal/svc"
	"github.com/lixvyang/xinclockin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SigninHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SigninReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := oauth.NewSigninLogic(r.Context(), svcCtx)
		resp, err := l.Signin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
