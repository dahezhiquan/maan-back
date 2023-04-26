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

type PictureDomain struct {
	pictureRepo repo.PictureRepo
}

func NewPictureDomain() *PictureDomain {
	return &PictureDomain{
		pictureRepo: dao.NewPictureDao(),
	}
}

// 获取图片信息 根据类型

func (p *PictureDomain) FindPic(picType string) ([]*model.Picture, *errs.BError) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	pics, err := p.pictureRepo.FindPic(nil, ctx, picType)
	if err != nil {
		zap.L().Error("Picture DB FindPic error", zap.Error(err))
		return nil, errs.DBError
	}
	if pics == nil {
		return nil, errs.PictureNotExistError
	}
	return pics, nil
}
