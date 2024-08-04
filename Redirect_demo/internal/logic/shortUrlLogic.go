package logic

import (
	"context"

	"demo/internal/svc"
	"demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortUrlLogic {
	return &ShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortUrlLogic) ShortUrl(req *types.Req) (resp *types.Resp, err error) {
	if req.ShortUrl == "1v32rtp" {
		// 如果短链接正确，返回长链接
		return &types.Resp{LongUrl: "https://github.com/Yhelee99/shortUrl"}, nil
	} else {
		// 标识符不正确，跳转到go中文网
		return &types.Resp{LongUrl: "https://go.p2hp.com/"}, nil
	}

	return
}
