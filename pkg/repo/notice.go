package repo

import (
	"context"
	"secureQR/pkg/connections/database"
	"secureQR/pkg/model"
)

type NoticeRepo interface {
	FindNotice(tx database.DbConn, ctx context.Context) (notices []*model.Notice, err error)
}
