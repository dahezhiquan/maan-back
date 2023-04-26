package domain

import (
	"context"
	"go.uber.org/zap"
	"maan/common/errs"
	"maan/pkg/dao"
	"maan/pkg/model"
	"maan/pkg/repo"
	"time"
)

type NoticeDomain struct {
	noticeRepo repo.NoticeRepo
}

func NewNoticeDomain() *NoticeDomain {
	return &NoticeDomain{
		noticeRepo: dao.NewNoticeDao(),
	}
}

// 获取最近公告

func (n *NoticeDomain) FindNotice() ([]*model.Notice, *errs.BError) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	notices, err := n.noticeRepo.FindNotice(nil, ctx)
	if err != nil {
		zap.L().Error("Notice DB FindNotice error", zap.Error(err))
		return nil, errs.DBError
	}
	if notices == nil {
		return nil, errs.NoticeNotExistError
	}
	return notices, nil
}
