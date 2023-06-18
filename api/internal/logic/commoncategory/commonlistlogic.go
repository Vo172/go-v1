package commoncategory

import (
	"context"

	"crmcore-customer-go/api/internal/svc"
	"crmcore-customer-go/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonListLogic {
	return &CommonListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommonListLogic) CommonList(req *types.ListCommonCategoryReq) (resp *types.ListCommonCategoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
