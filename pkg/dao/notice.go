package dao

import (
	"context"
	"gorm.io/gorm"
	"maan/pkg/connections/database"
	"maan/pkg/connections/database/gorms"
	"maan/pkg/model"
)

type NoticeDao struct {
	baseDao
}

func NewNoticeDao() *NoticeDao {
	return &NoticeDao{baseDao{conn: gorms.NewConn()}}
}

// 查询公告

func (n *NoticeDao) FindNotice(tx database.DbConn, ctx context.Context) (notices []*model.Notice, err error) {
	conn := n.getConn(tx)

	err = conn.Session(ctx).Order("created_at desc").Limit(5).Find(&notices).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return notices, err
}
