package customersmodel

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
)

type (
	customersModel interface {
		FindAll(ctx context.Context) (*[]Customers, error)
		FindByCustomerName(username string) (*Customers, error)
	}

	defaultCustomersModel struct {
		db    sqlx.DB
		table string
	}
	Customers struct {
		CustomerId   int64  `db:"CUSTOMER_ID"`
		CustomerDate string `db:"CREATE_DATE"`
		Email        string `db:"EMAIL"`
		CustomerName string `db:"CUSTOMER_NAME"`
		Password     string `db:"PASSWORD"`
	}
)

func newCustomersModel(db sqlx.DB) *defaultCustomersModel {
	return &defaultCustomersModel{
		db:    db,
		table: "`customers`",
	}
}

func (c customCustomersModel) FindByCustomerName(username string) (*Customers, error) {
	var customer Customers
	var (
		userName string
		password string
	)
	query := "SELECT customer_name,password FROM customers WHERE customer_name = :1"
	rows, err := c.db.Query(query, username)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err2 := rows.Scan(&userName, &password)
		if err2 != nil {
			log.Fatal(err2)
		}
		customer.CustomerName = userName
		customer.Password = password
	}
	return &customer, nil
}

func (c customCustomersModel) FindAll(context.Context) (*[]Customers, error) {
	var customers []Customers
	var (
		customerId   int64
		customerName string
		email        string
		createDate   string
	)
	rows, err := c.db.Query("SELECT customer_id, customer_name, email, create_date FROM customers")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var customer Customers
		err := rows.Scan(&customerId, &customerName, &email, &createDate)
		if err != nil {
			log.Fatal(err)
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.CustomerDate = createDate
		customer.Email = email
		customers = append(customers, customer)
	}
	return &customers, nil
}
