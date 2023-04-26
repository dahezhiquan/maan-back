package repo

import (
	"context"
	"maan/pkg/connections/database"
	"maan/pkg/model"
)

type NoticeRepo interface {
	FindNotice(tx database.DbConn, ctx context.Context) (notices []*model.Notice, err error)
}
