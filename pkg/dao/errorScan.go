package dao

import (
	"context"
	"maan/pkg/connections/database"
	"maan/pkg/connections/database/gorms"
	"maan/pkg/model"
)

type ErrorScanDao struct {
	baseDao
}

func NewErrorScanDao() *ErrorScanDao {
	return &ErrorScanDao{baseDao{conn: gorms.NewConn()}}
}

// 保存误报结果

func (e *ErrorScanDao) SaveErrorScan(tx database.DbConn, ctx context.Context, errorScan *model.ErrorScan) error {
	conn := e.getConn(tx)

	err := conn.Session(ctx).Save(errorScan).Error

	if err != nil {
		return err
	}
	return nil
}
