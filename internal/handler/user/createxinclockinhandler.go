package user

import (
	"net/http"

	"github.com/lixvyang/xinclockin/internal/logic/user"
	"github.com/lixvyang/xinclockin/internal/svc"
	"github.com/lixvyang/xinclockin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateXinClockInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateClockinReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewCreateXinClockInLogic(r.Context(), svcCtx)
		resp, err := l.CreateXinClockIn(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
