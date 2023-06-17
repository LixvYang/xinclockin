package clockin

import (
	"net/http"

	"github.com/lixvyang/xinclockin/internal/logic/clockin"
	"github.com/lixvyang/xinclockin/internal/svc"
	"github.com/lixvyang/xinclockin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateContentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateContentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := clockin.NewCreateContentLogic(r.Context(), svcCtx)
		resp, err := l.CreateContent(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
