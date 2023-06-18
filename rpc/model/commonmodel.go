package model

import (
	"context"
	types "crmcore-customer-go/rpc/category/categoryclient"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
)

var (
	sysCommonField            = builder.RawFieldNames(&Common{})
	sysCommonRows             = strings.ReplaceAll(strings.Join(sysCommonField, ","), "`", "")
	sysJobRowsExpectAutoSet   = strings.Join(stringx.Remove(sysCommonField, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysJobRowsWithPlaceHolder = strings.Join(stringx.Remove(sysCommonField, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	CommonModel interface {
		FindAll(ctx context.Context, req *types.ListCommonReq) (*[]Common, error)
		Count(ctx context.Context, req *types.ListCommonReq) (int32, error)
		//FindByCustomerName(username string) (*Common, error)
	}
	customCommonModel struct {
		db    sqlx.DB
		table string
	}
	Common struct {
		Id                 string         `db:"ID"`
		Code               string         `db:"CODE"`
		Name               string         `db:"NAME"`
		Description        sql.NullString `db:"DESCRIPTION"`
		OrderNum           sql.NullInt32  `db:"ORDER_NUM"`
		Value              sql.NullString `db:"VALUE"`
		CommonCategoryCode string         `db:"COMMON_CATEGORY_CODE"`
		IsDefault          sql.NullBool   `db:"IS_DEFAULT"`
	}
)

func NewCommonModel(db sqlx.DB) CommonModel {
	return &customCommonModel{
		db:    db,
		table: "common",
	}
}

func (c customCommonModel) FindAll(ctx context.Context, req *types.ListCommonReq) (*[]Common, error) {
	var commons []Common
	var start = req.Page * req.Size
	var end = (req.Page + 1) * req.Size
	//fmt.Println(fmt.Sprintf("TEXT : %s", sysCommonRows))
	var strSelect = fmt.Sprintf(" SELECT %s ", sysCommonRows)
	var strWhere = " IS_ACTIVE = 1 "
	if req.Code != "" {
		strWhere = fmt.Sprintf("%s AND LOWER(CODE) = '%s' ", strWhere, strings.ToLower(req.Code))
	}
	if req.Name != "" {
		strWhere = fmt.Sprintf("%s AND LOWER(NAME) LIKE '%%%s%%' ", strWhere, strings.ToLower(req.Name))
	}
	if req.CommonCategoryCode != "" {
		strWhere = fmt.Sprintf("%s AND LOWER(COMMON_CATEGORY_CODE) = '%s' ", strWhere, strings.ToLower(req.CommonCategoryCode))
	}
	var query = fmt.Sprintf("%s FROM ( %s ,ROWNUM num FROM  %s WHERE %s AND ROWNUM <= %d) WHERE num > %d ", strSelect, strSelect, c.table, strWhere, end, start)
	err := c.db.Select(&commons, query)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	return &commons, nil
}

func (c *customCommonModel) Count(ctx context.Context, req *types.ListCommonReq) (int32, error) {
	var count []int32
	var strWhere = " IS_ACTIVE = 1 "
	if req.Code != "" {
		strWhere = fmt.Sprintf("%s AND LOWER(CODE) = '%s' ", strWhere, strings.ToLower(req.Code))
	}
	if req.Name != "" {
		strWhere = fmt.Sprintf("%s AND LOWER(NAME) LIKE '%%%s%%' ", strWhere, strings.ToLower(req.Name))
	}
	if req.CommonCategoryCode != "" {
		strWhere = fmt.Sprintf("%s AND LOWER(COMMON_CATEGORY_CODE) = '%s' ", strWhere, strings.ToLower(req.CommonCategoryCode))
	}
	var query = fmt.Sprintf("SELECT count(*) as count FROM %s WHERE %s ", c.table, strWhere)
	err := c.db.Select(&count, query)
	switch err {
	case nil:
		return count[0], nil
	default:
		return 0, err
	}
}
