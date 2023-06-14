package product

import (
	"context"

	"apiproduct/api/internal/svc"
	"apiproduct/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCustomerAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerAddLogic {
	return &CustomerAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CustomerAddLogic) CustomerAdd(req *types.AddCustomerReq) (resp *types.AddCustomerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
