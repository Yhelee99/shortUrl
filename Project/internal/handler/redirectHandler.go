package handler

import (
	"net/http"

	"Project/internal/logic"
	"Project/internal/svc"
	"Project/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RedirectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RedirectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRedirectLogic(r.Context(), svcCtx)
		resp, err := l.Redirect(&req)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			http.Redirect(w, r, resp.RedResp, http.StatusFound)
		}

	}
}
