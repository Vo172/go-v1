package logic

import (
	"context"

	"crmcore-customer-go/rpc/category/categoryclient"
	"crmcore-customer-go/rpc/category/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommonUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonUpdateLogic {
	return &CommonUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommonUpdateLogic) CommonUpdate(in *categoryclient.UpdateCommonReq) (*categoryclient.UpdateCommonReq, error) {
	// todo: add your logic here and delete this line

	return &categoryclient.UpdateCommonReq{}, nil
}
