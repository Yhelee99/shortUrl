package logic

import (
	"Project/internal/errorx"
	"Project/internal/svc"
	"Project/internal/types"
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowUrlLogic {
	return &ShowUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowUrlLogic) ShowUrl(req *types.ShowReq) (resp *types.ShowResp, err error) {
	// 根据长链查短链
	reslut, err := l.svcCtx.ShortUrlDb.FindOneBySurl(l.ctx, sql.NullString{
		String: req.ShortUrl,
		Valid:  true,
	})

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx.NewErrCode(errorx.NotFound, "未找到该短链")
		}
		logx.Errorw("ShortUrlDb.FindOneBySurl failed.", logx.Field("err", err))
		return nil, errorx.NewDefaultErrCode()
	}

	lUrl := reslut.Lurl
	return &types.ShowResp{
		LongUrl: lUrl.String,
	}, nil
}
