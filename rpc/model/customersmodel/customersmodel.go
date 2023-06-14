package customersmodel

import (
	"github.com/jmoiron/sqlx"
)

var _ CustomersModel = (*customCustomersModel)(nil)

type (
	CustomersModel interface {
		customersModel
	}
	customCustomersModel struct {
		*defaultCustomersModel
	}
)

func NewCustomersModel(db sqlx.DB) CustomersModel {
	return &customCustomersModel{
		defaultCustomersModel: newCustomersModel(db),
	}
}
