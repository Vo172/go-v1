package common

import (
	"context"

	"crmcore-customer-go/api/internal/svc"
	"crmcore-customer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommonDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonDeleteLogic {
	return &CommonDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommonDeleteLogic) CommonDelete(req *types.DeleteCommonReq) (resp *types.DeleteCommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
