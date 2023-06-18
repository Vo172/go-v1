package logic

import (
	"context"

	"crmcore-customer-go/rpc/category/categoryclient"
	"crmcore-customer-go/rpc/category/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommonAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonAddLogic {
	return &CommonAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommonAddLogic) CommonAdd(in *categoryclient.AddCommonReq) (*categoryclient.AddCommonResp, error) {
	// todo: add your logic here and delete this line

	return &categoryclient.AddCommonResp{}, nil
}
