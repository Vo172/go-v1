package logic

import (
	"context"

	"crmcore-customer-go/rpc/category/categoryclient"
	"crmcore-customer-go/rpc/category/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommonDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonDetailLogic {
	return &CommonDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommonDetailLogic) CommonDetail(in *categoryclient.DetailCommonReq) (*categoryclient.DeleteCommonResp, error) {
	// todo: add your logic here and delete this line

	return &categoryclient.DeleteCommonResp{}, nil
}
