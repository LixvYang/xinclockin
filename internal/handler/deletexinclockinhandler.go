package handler

import (
	"net/http"

	"github.com/lixvyang/xinclockin/internal/logic"
	"github.com/lixvyang/xinclockin/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteXinClockInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDeleteXinClockInLogic(r.Context(), svcCtx)
		err := l.DeleteXinClockIn()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
