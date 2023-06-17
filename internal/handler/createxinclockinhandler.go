package handler

import (
	"net/http"

	"github.com/lixvyang/xinclockin/internal/logic"
	"github.com/lixvyang/xinclockin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
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
