package handler

import (
	"net/http"

	"Project/internal/logic"
	"Project/internal/svc"
	"Project/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShowUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShowUrlLogic(r.Context(), svcCtx)
		resp, err := l.ShowUrl(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
