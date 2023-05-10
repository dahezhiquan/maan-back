package domain

import (
	"context"
	"maan/common/errs"
	"maan/pkg/dao"
	"maan/pkg/model"
	"maan/pkg/repo"
	"time"
)

type ScanRecordDomain struct {
	scanRecordRepo repo.ScanRecordRepo
}

func NewScanRecordDomain() *ScanRecordDomain {
	return &ScanRecordDomain{
		scanRecordRepo: dao.NewScanRecordDao(),
	}
}

func (s *ScanRecordDomain) SaveScanRecord(scanRecord *model.ScanRecord) *errs.BError {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := s.scanRecordRepo.SaveScanRecord(nil, ctx, scanRecord)
	if err != nil {
		return errs.DBError
	}
	return nil
}

func (s *ScanRecordDomain) FindScanRecord() ([]*model.ScanRecord, *errs.BError) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	scanRecords, err := s.scanRecordRepo.FindRecordByLimit(nil, ctx)
	if err != nil {
		return nil, errs.DBError
	}
	return scanRecords, nil
}
