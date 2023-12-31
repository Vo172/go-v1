// Code generated by goctl. DO NOT EDIT.
// Source: category.proto

package server

import (
	"context"

	"crmcore-customer-go/rpc/category/categoryclient"
	"crmcore-customer-go/rpc/category/internal/logic"
	"crmcore-customer-go/rpc/category/internal/svc"
)

type CategoryServer struct {
	svcCtx *svc.ServiceContext
	categoryclient.UnimplementedCategoryServer
}

func NewCategoryServer(svcCtx *svc.ServiceContext) *CategoryServer {
	return &CategoryServer{
		svcCtx: svcCtx,
	}
}

func (s *CategoryServer) CommonAdd(ctx context.Context, in *categoryclient.AddCommonReq) (*categoryclient.AddCommonResp, error) {
	l := logic.NewCommonAddLogic(ctx, s.svcCtx)
	return l.CommonAdd(in)
}

func (s *CategoryServer) CommonUpdate(ctx context.Context, in *categoryclient.UpdateCommonReq) (*categoryclient.UpdateCommonReq, error) {
	l := logic.NewCommonUpdateLogic(ctx, s.svcCtx)
	return l.CommonUpdate(in)
}

func (s *CategoryServer) CommonDelete(ctx context.Context, in *categoryclient.DeleteCommonReq) (*categoryclient.DeleteCommonResp, error) {
	l := logic.NewCommonDeleteLogic(ctx, s.svcCtx)
	return l.CommonDelete(in)
}

func (s *CategoryServer) CommonDetail(ctx context.Context, in *categoryclient.DetailCommonReq) (*categoryclient.DeleteCommonResp, error) {
	l := logic.NewCommonDetailLogic(ctx, s.svcCtx)
	return l.CommonDetail(in)
}

func (s *CategoryServer) CommonList(ctx context.Context, in *categoryclient.ListCommonReq) (*categoryclient.PageCommonResp, error) {
	l := logic.NewCommonListLogic(ctx, s.svcCtx)
	return l.CommonList(in)
}
