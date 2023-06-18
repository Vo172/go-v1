package common

import (
	"context"
	"crmcore-customer-go/api/internal/common/errorx"
	"crmcore-customer-go/api/internal/svc"
	"crmcore-customer-go/api/internal/types"
	"crmcore-customer-go/rpc/category/categoryclient"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
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

func (l *CommonListLogic) CommonList(req *types.ListCommonReq, r *http.Request) (resp *types.PageCommonResp, err error) {
	listCommon, err := l.svcCtx.Category.CommonList(l.ctx, &categoryclient.ListCommonReq{
		Code:               req.Code,
		Name:               req.Name,
		CommonCategoryCode: req.CommonCategoryCode,
		Page:               req.Page,
		Size:               req.Size,
	})
	if err != nil {
		logx.Error(fmt.Sprintf("[FIND_ALL_CATEGORY] error : %s"), err)
		return nil, errorx.NewDefaultError("Lấy danh sách không thành công")
	}
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var listData []*types.ListCommonResp

	for _, item := range listCommon.Content {
		listData = append(listData, &types.ListCommonResp{
			Id:                 item.Id,
			Code:               item.Code,
			Name:               item.Name,
			Description:        item.Description,
			OrderNum:           item.OrderNum,
			Value:              item.Value,
			CommonCategoryCode: item.CommonCategoryCode,
			IsDefault:          item.IsDefault,
		})
	}
	return &types.PageCommonResp{
		Content:      listData,
		TotalElement: listCommon.TotalElement,
		TotalPages:   listCommon.TotalPages,
		Size:         req.Size,
		Number:       req.Page,
	}, err
}
