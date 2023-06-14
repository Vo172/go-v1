package product

import (
	"apiproduct/api/internal/svc"
	"apiproduct/api/internal/types"
	"context"
	"errors"
	_ "errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CustomerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCustomerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CustomerListLogic {
	return &CustomerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CustomerListLogic) CustomerList(req *types.CustomerReq) (resp []*types.CustomerReply, err error) {

	listUser, err := l.svcCtx.CustomersModel.FindAll(l.ctx)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	result := make([]*types.CustomerReply, 0, len(*listUser))
	for _, sourceObj := range *listUser {
		item := &types.CustomerReply{
			CustomerId:   sourceObj.CustomerId,
			CustomerName: sourceObj.CustomerName,
			Email:        sourceObj.Email,
			CreatedDate:  sourceObj.CustomerDate,
		}
		result = append(result, item)
	}
	return result, err
}
