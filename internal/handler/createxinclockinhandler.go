package handler

import (
	"github.com/lixvyang/rebetxin-one/common/errorx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xinclockin/internal/logic"
	"xinclockin/internal/svc"
)

func CreateXinClockInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCreateXinClockInLogic(r.Context(), svcCtx)
		err := l.CreateXinClockIn()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
