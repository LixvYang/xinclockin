package user

import (
	"net/http"

	"github.com/lixvyang/xinclockin/internal/logic/user"
	"github.com/lixvyang/xinclockin/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserListLogic(r.Context(), svcCtx)
		resp, err := l.GetUserList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
