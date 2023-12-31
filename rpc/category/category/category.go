// Code generated by goctl. DO NOT EDIT.
// Source: category.proto

package category

import (
	"context"

	"crmcore-customer-go/rpc/category/categoryclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCommonReq     = categoryclient.AddCommonReq
	AddCommonResp    = categoryclient.AddCommonResp
	DeleteCommonReq  = categoryclient.DeleteCommonReq
	DeleteCommonResp = categoryclient.DeleteCommonResp
	DetailCommonReq  = categoryclient.DetailCommonReq
	DetailCommonResp = categoryclient.DetailCommonResp
	ListCommonReq    = categoryclient.ListCommonReq
	ListCommonResp   = categoryclient.ListCommonResp
	PageCommonResp   = categoryclient.PageCommonResp
	UpdateCommonReq  = categoryclient.UpdateCommonReq
	UpdateCommonResp = categoryclient.UpdateCommonResp

	Category interface {
		CommonAdd(ctx context.Context, in *AddCommonReq, opts ...grpc.CallOption) (*AddCommonResp, error)
		CommonUpdate(ctx context.Context, in *UpdateCommonReq, opts ...grpc.CallOption) (*UpdateCommonReq, error)
		CommonDelete(ctx context.Context, in *DeleteCommonReq, opts ...grpc.CallOption) (*DeleteCommonResp, error)
		CommonDetail(ctx context.Context, in *DetailCommonReq, opts ...grpc.CallOption) (*DeleteCommonResp, error)
		CommonList(ctx context.Context, in *ListCommonReq, opts ...grpc.CallOption) (*PageCommonResp, error)
	}

	defaultCategory struct {
		cli zrpc.Client
	}
)

func NewCategory(cli zrpc.Client) Category {
	return &defaultCategory{
		cli: cli,
	}
}

func (m *defaultCategory) CommonAdd(ctx context.Context, in *AddCommonReq, opts ...grpc.CallOption) (*AddCommonResp, error) {
	client := categoryclient.NewCategoryClient(m.cli.Conn())
	return client.CommonAdd(ctx, in, opts...)
}

func (m *defaultCategory) CommonUpdate(ctx context.Context, in *UpdateCommonReq, opts ...grpc.CallOption) (*UpdateCommonReq, error) {
	client := categoryclient.NewCategoryClient(m.cli.Conn())
	return client.CommonUpdate(ctx, in, opts...)
}

func (m *defaultCategory) CommonDelete(ctx context.Context, in *DeleteCommonReq, opts ...grpc.CallOption) (*DeleteCommonResp, error) {
	client := categoryclient.NewCategoryClient(m.cli.Conn())
	return client.CommonDelete(ctx, in, opts...)
}

func (m *defaultCategory) CommonDetail(ctx context.Context, in *DetailCommonReq, opts ...grpc.CallOption) (*DeleteCommonResp, error) {
	client := categoryclient.NewCategoryClient(m.cli.Conn())
	return client.CommonDetail(ctx, in, opts...)
}

func (m *defaultCategory) CommonList(ctx context.Context, in *ListCommonReq, opts ...grpc.CallOption) (*PageCommonResp, error) {
	client := categoryclient.NewCategoryClient(m.cli.Conn())
	return client.CommonList(ctx, in, opts...)
}
