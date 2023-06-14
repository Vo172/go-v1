package svc

import (
	"apiproduct/api/internal/config"
	"apiproduct/api/internal/middleware"
	"apiproduct/rpc/model/customersmodel"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/sijms/go-ora/v2"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	CustomersModel customersmodel.CustomersModel
	Keycloak       *middleware.Keycloak
	Authen         rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := sqlx.Open(c.Oracle.DriverName, c.Oracle.DataSource)
	if err != nil {
		fmt.Println("Lỗi khi kết nối cơ sở dữ liệu:", err)
	}
	return &ServiceContext{
		Config:         c,
		CustomersModel: customersmodel.NewCustomersModel(*db),
		Keycloak:       middleware.NewKeycloak(c.Keycloak.BasePath, c.Keycloak.ClientId, c.Keycloak.ClientSecret, c.Keycloak.Realm),
		Authen:         middleware.NewMiddleware(middleware.NewKeycloak(c.Keycloak.BasePath, c.Keycloak.ClientId, c.Keycloak.ClientSecret, c.Keycloak.Realm)).VerifyToken,
	}
}
