package logic

import (
	"context"
	"database/sql"
	"errors"

	"Project/internal/errorx"
	"Project/internal/svc"
	"Project/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RedirectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedirectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedirectLogic {
	return &RedirectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedirectLogic) Redirect(req *types.RedirectReq) (resp *types.RedirectResp, err error) {
	res, err := l.svcCtx.ShortUrlDb.FindOneBySurl(l.ctx, sql.NullString{
		String: req.RedReq,
		Valid:  true,
	})

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, errorx.NewErrCode(errorx.NotFound, "未找到该短链")
		}
		logx.Errorw("ShortUrlDb.FindOneBySurl failed.", logx.Field("err", err))
		return nil, errorx.NewDefaultErrCode()
	}
	return &types.RedirectResp{RedResp: res.Lurl.String}, nil
}
