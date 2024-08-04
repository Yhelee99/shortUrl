package logic

import (
	"context"

	"Project/internal/svc"
	"Project/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
