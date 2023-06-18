package logic

import (
	"context"

	"crmcore-customer-go/rpc/category/categoryclient"
	"crmcore-customer-go/rpc/category/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommonDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonDeleteLogic {
	return &CommonDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommonDeleteLogic) CommonDelete(in *categoryclient.DeleteCommonReq) (*categoryclient.DeleteCommonResp, error) {
	// todo: add your logic here and delete this line

	return &categoryclient.DeleteCommonResp{}, nil
}
