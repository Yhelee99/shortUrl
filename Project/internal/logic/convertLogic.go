package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"Project/internal/errorx"
	"Project/internal/svc"
	"Project/internal/types"
	"Project/model"
	"Project/pkg/basex"
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

	if ok := conncheck.CheckUrl(req.LongUrl); !ok {
		logx.Errorw("CheckUrl failed.", logx.Field("err", err))
		return nil, errorx.NewErrCode(errorx.InvalidParams, "长链接无效")
	}
	// 1.3 判断是否转链过
	// 1.3.1 生成md5值
	md5v := yhelee.GetMd5Value([]byte(req.LongUrl)) // []byte()是一个函数调用的形式,可以传string
	// 1.3.2 查询md5值判重
	reslut, err := l.svcCtx.ShortUrlDb.FindOneByMd5(l.ctx, sql.NullString{String: md5v, Valid: true}) // Valid = true 表示值有效
	if !errors.Is(err, sqlx.ErrNotFound) {
		if err == nil {
			return nil, errorx.NewErrCode(errorx.InvalidParams, fmt.Sprintf("此链接已被转链为：%v", l.svcCtx.Config.Domain+"/redirect/"+reslut.Surl.String))
		} else {
			logx.Errorw("ShortUrlDb.FindOneByMd5 failed.", logx.Field("err", err))
			return nil, errorx.NewDefaultErrCode()
		}
	}
	// 1.4 输入的不能是一个短链接
	baseUrl, err := urlx.GetBasePath(req.LongUrl)
	if err != nil {
		logx.Errorw("urlx.GetBasePath failed", logx.Field("url", req.LongUrl), logx.Field("err", err))
		return nil, errorx.NewErrCode(errorx.InvalidParams, "请输入长链接")
	}
	logx.Debugw("取base地址", logx.Field("值", baseUrl))

	var short string
	for {
		// 2. 取号
		num, err := l.svcCtx.Sequence.GetNumb()
		if err != nil {
			return nil, errorx.NewDefaultErrCode()
		}
		logx.Debugw("取号成功", logx.Field("值", num))

		// 3. 生成短链
		// 3.1 10进制转62进制
		short = basex.IntToString(num)
		logx.Debugw("10进制转62进制", logx.Field("值", short))
		// 3.2 预防破解,考虑安全性   ---通过打乱编码表（basestring）实现
		// 3.3 添加屏蔽词
		if _, ok := l.svcCtx.BlackList[short]; !ok {
			break
		}
	}

	// 4.1 入库
	_, err = l.svcCtx.ShortUrlDb.Insert(l.ctx, &model.ShortUrlMap{
		Lurl: sql.NullString{String: req.LongUrl, Valid: true},
		Md5:  sql.NullString{String: md5v, Valid: true},
		Surl: sql.NullString{String: short, Valid: true},
	})
	if err != nil {
		logx.Errorw("ShortUrlDb.Insert failed.", logx.Field("err", err))
		return nil, errorx.NewDefaultErrCode()
	}
	// 4.2 存入布隆过滤器
	if err := l.svcCtx.Filter.Add([]byte(short)); err != nil {
		logx.Errorw("BloomFilter.Add() failed", logx.Field("err", err))
	}

	// 5. 返回响应
	shortUrl := l.svcCtx.Config.Domain + "/redirect/" + short

	return &types.ConvertResp{ShortUrl: shortUrl}, nil
}
