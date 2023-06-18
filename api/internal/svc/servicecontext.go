package svc

import (
	"crmcore-customer-go/api/internal/config"
	"crmcore-customer-go/api/internal/middleware"
	"crmcore-customer-go/rpc/category/category"
	_ "github.com/sijms/go-ora/v2"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	Keycloak *middleware.Keycloak
	Authen   rest.Middleware
	Category category.Category
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Keycloak: middleware.NewKeycloak(c.Keycloak.BasePath, c.Keycloak.ClientId, c.Keycloak.ClientSecret, c.Keycloak.Realm),
		Authen:   middleware.NewMiddleware(middleware.NewKeycloak(c.Keycloak.BasePath, c.Keycloak.ClientId, c.Keycloak.ClientSecret, c.Keycloak.Realm)).VerifyToken,
		Category: category.NewCategory(zrpc.MustNewClient(c.CommonRpc)),
	}
}
