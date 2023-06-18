package common

import (
	"context"

	"crmcore-customer-go/api/internal/svc"
	"crmcore-customer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommonUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonUpdateLogic {
	return &CommonUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommonUpdateLogic) CommonUpdate(req *types.UpdateCommonReq) (resp *types.UpdateCommonReq, err error) {
	// todo: add your logic here and delete this line

	return
}
