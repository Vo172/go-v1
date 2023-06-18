package svc

import (
	"crmcore-customer-go/rpc/category/internal/config"
	"crmcore-customer-go/rpc/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/sijms/go-ora/v2"
)

type ServiceContext struct {
	Config      config.Config
	CommonModel model.CommonModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn, err := sqlx.Open(c.Oracle.DriverName, c.Oracle.DataSource)
	if err != nil {
		fmt.Println("Lỗi khi kết nối cơ sở dữ liệu:", err)
	}
	return &ServiceContext{
		Config:      c,
		CommonModel: model.NewCommonModel(*sqlConn),
	}
}
