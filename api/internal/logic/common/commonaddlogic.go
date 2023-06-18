package common

import (
	"context"

	"crmcore-customer-go/api/internal/svc"
	"crmcore-customer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommonAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonAddLogic {
	return &CommonAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommonAddLogic) CommonAdd(req *types.AddCommonReq) (resp *types.AddCommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
