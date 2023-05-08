package repo

import (
	"context"
	"maan/pkg/connections/database"
	"maan/pkg/model"
)

type ErrorScanRepo interface {
	SaveErrorScan(tx database.DbConn, ctx context.Context, errorScan *model.ErrorScan) error
}
