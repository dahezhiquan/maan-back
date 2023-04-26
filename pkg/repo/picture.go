package repo

import (
	"context"
	"maan/pkg/connections/database"
	"maan/pkg/model"
)

type PictureRepo interface {
	FindPic(tx database.DbConn, ctx context.Context, picType string) (pics []*model.Picture, err error)
}
