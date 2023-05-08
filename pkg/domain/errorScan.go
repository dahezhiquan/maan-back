package domain

import (
	"context"
	"maan/common/errs"
	"maan/pkg/dao"
	"maan/pkg/model"
	"maan/pkg/repo"
	"time"
)

type ErrorScanDomain struct {
	errorScanRepo repo.ErrorScanRepo
}

func NewErrorScanDomain() *ErrorScanDomain {
	return &ErrorScanDomain{
		errorScanRepo: dao.NewErrorScanDao(),
	}
}

func (e *ErrorScanDomain) SaveErrorScan(errorScan *model.ErrorScan) *errs.BError {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := e.errorScanRepo.SaveErrorScan(nil, ctx, errorScan)
	if err != nil {
		return errs.DBError
	}

	return nil
}
