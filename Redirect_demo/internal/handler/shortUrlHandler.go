package handler

import (
	"net/http"

	"demo/internal/logic"
	"demo/internal/svc"
	"demo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShortUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewShortUrlLogic(r.Context(), svcCtx)
		resp, err := l.ShortUrl(&req)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// 原来返回响应数据的方式
			// httpx.OkJsonCtx(r.Context(), w, resp)

			// 返回重定向
			// 1. goZero方式
			// w.Header().Set("location", resp.LongUrl)
			// w.WriteHeader(http.StatusFound)
			// 2. http标准库方式
			http.Redirect(w, r, resp.LongUrl, http.StatusFound)
		}
	}
}
