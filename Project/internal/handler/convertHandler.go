package handler

import (
	"net/http"

	"Project/internal/logic"
	"Project/internal/svc"
	"Project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ConvertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 解析请求参数
		var req types.ConvertReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Errorw("validator failed", logx.Field("err", err))
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// gozero 1.7.0 版本集成了validator库

		// // 校验请求参数
		// if err := validator.New().StructCtx(r.Context(), &req); err != nil {
		// 	fmt.Println("in")
		// 	logx.Errorw("validator failed", logx.Field("err", err))
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// 	return
		// }
		// fmt.Println("out")

		// 执行业务逻辑
		l := logic.NewConvertLogic(r.Context(), svcCtx)
		resp, err := l.Convert(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
