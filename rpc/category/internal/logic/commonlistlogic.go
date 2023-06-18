package logic

import (
	"context"
	"crmcore-customer-go/rpc/category/common/errorx"
	"errors"
	"fmt"
	"math"

	"crmcore-customer-go/rpc/category/categoryclient"
	"crmcore-customer-go/rpc/category/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommonListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonListLogic {
	return &CommonListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommonListLogic) CommonList(in *categoryclient.ListCommonReq) (*categoryclient.PageCommonResp, error) {
	listCommon, err := l.svcCtx.CommonModel.FindAll(l.ctx, in)
	if err != nil {
		logx.Error(fmt.Sprintf("[FIND_ALL_CATEGORY] error : %s"), err)
		return nil, errorx.NewDefaultError("Lấy danh sách không thành công")
	}

	count, err := l.svcCtx.CommonModel.Count(l.ctx, in)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	listData := make([]*categoryclient.ListCommonResp, 0, len(*listCommon))
	for _, sourceObj := range *listCommon {
		item := &categoryclient.ListCommonResp{
			Id:                 sourceObj.Id,
			Code:               sourceObj.Code,
			Name:               sourceObj.Name,
			Description:        sourceObj.Description.String,
			OrderNum:           sourceObj.OrderNum.Int32,
			Value:              sourceObj.Value.String,
			CommonCategoryCode: sourceObj.CommonCategoryCode,
			IsDefault:          sourceObj.IsDefault.Bool,
		}
		listData = append(listData, item)
	}
	var totalPage = int32(math.Ceil(float64(count) / float64(in.Size)))
	return &categoryclient.PageCommonResp{
		Content:      listData,
		TotalElement: count,
		TotalPages:   totalPage,
		Size:         in.Size,
		Number:       in.Page,
	}, err
}
