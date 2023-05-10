package repo

import (
	"context"
	"maan/pkg/connections/database"
	"maan/pkg/model"
)

type ScanRecordRepo interface {
	SaveScanRecord(tx database.DbConn, ctx context.Context, scanRecord *model.ScanRecord) error
	FindRecordByLimit(tx database.DbConn, ctx context.Context) (scanRecords []*model.ScanRecord, err error)
}
