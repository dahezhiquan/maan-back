package dao

import (
	"context"
	"gorm.io/gorm"
	"maan/pkg/connections/database"
	"maan/pkg/connections/database/gorms"
	"maan/pkg/model"
)

type PictureDao struct {
	baseDao
}

func NewPictureDao() *PictureDao {
	return &PictureDao{baseDao{conn: gorms.NewConn()}}
}

// 根据图片类型 查询图片信息

func (p *PictureDao) FindPic(tx database.DbConn, ctx context.Context, picType string) (pics []*model.Picture, err error) {
	conn := p.getConn(tx)

	err = conn.Session(ctx).Order("created_at desc").Limit(4).Where("type = ?", picType).Find(&pics).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return pics, err
}
