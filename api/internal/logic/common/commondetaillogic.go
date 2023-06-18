package common

import (
	"context"

	"crmcore-customer-go/api/internal/svc"
	"crmcore-customer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommonDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonDetailLogic {
	return &CommonDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommonDetailLogic) CommonDetail(req *types.DetailCommonReq) (resp *types.DeleteCommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
