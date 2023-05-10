package dao

import (
	"context"
	"maan/pkg/connections/database"
	"maan/pkg/connections/database/gorms"
	"maan/pkg/model"
	"maan/pkg/public"
)

type ScanRecordDao struct {
	baseDao
}

func NewScanRecordDao() *ScanRecordDao {
	return &ScanRecordDao{baseDao{conn: gorms.NewConn()}}
}

// 保存扫描结果

func (s *ScanRecordDao) SaveScanRecord(tx database.DbConn, ctx context.Context, scanRecord *model.ScanRecord) error {
	conn := s.getConn(tx)

	err := conn.Session(ctx).FirstOrCreate(scanRecord, model.ScanRecord{Content: scanRecord.Content}).Error

	if err != nil {
		return err
	}
	return nil
}

// 返回最近的扫描结果

func (s *ScanRecordDao) FindRecordByLimit(tx database.DbConn, ctx context.Context) (scanRecords []*model.ScanRecord, err error) {
	conn := s.getConn(tx)

	err = conn.Session(ctx).Order("updated_at desc").Where("mvss < 80").Limit(public.IntelligencePageSize).Find(&scanRecords).Error

	if err != nil {
		return nil, err
	}
	return scanRecords, nil
}
