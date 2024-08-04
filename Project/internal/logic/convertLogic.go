package logic

import (
	"context"
	"database/sql"

	"Project/internal/errorx"
	"Project/internal/svc"
	"Project/internal/types"
	conncheck "Project/pkg/connCheck"
	"Project/pkg/urlx"
	"Project/pkg/yhelee"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertLogic) Convert(req *types.ConvertReq) (resp *types.ConvertResp, err error) {
	// 1. 参数校验
	// 1.1 判断是否为空
	//   使用validator包，在进入业务逻辑前完成
	// 1.2 判断链接是否有效
	ok := conncheck.CheckUrl(req.LongUrl)
	logx.Debugw("url检查", logx.Field("长地址", req.LongUrl), logx.Field("结果", ok))
	if !ok {
		logx.Errorw("CheckUrl failed.", logx.Field("err", err))
		return nil, errorx.NewErrCode(errorx.InvalidParams, "长链接无效")
	}
	// 1.3 判断是否转链过
	// 1.3.1 生成md5值
	md5v := yhelee.GetMd5Value([]byte(req.LongUrl)) // []byte()是一个函数调用的形式,可以传string
	// 1.3.2 查询md5值判重
	reslut, err := l.svcCtx.ShortUrlDb.FindOneByMd5(l.ctx, sql.NullString{String: md5v, Valid: true}) // Valid = true 表示值有效
	if err != sqlx.ErrNotFound {
		if err == nil {
			return &types.ConvertResp{ShortUrl: reslut.Surl.String}, errorx.NewErrCode(errorx.InvalidParams, "此链接已转链过")
		} else {
			logx.Errorw("ShortUrlDb.FindOneByMd5 failed.", logx.Field("err", err))
			return nil, errorx.NewDefaultErrCode()
		}
	}
	// 1.4 输入的不能是一个短链接
	baseUrl, err := urlx.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("urlx.GetBasePath failed", logx.Field("url", req.LongUrl), logx.Field("err", err))
		return nil, err
	}
	logx.Debugw("取base地址", logx.Field("值", baseUrl))
	// 2. 取号
	// 3. 生成短链
	// 4. 入库
	// 5. 返回响应
	return
}
