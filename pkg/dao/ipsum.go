package dao

import (
	"context"
	"gorm.io/gorm"
	"maan/pkg/connections/database"
	"maan/pkg/connections/database/gorms"
	"maan/pkg/model"
)

type IpSumDao struct {
	baseDao
}

func NewIpSumDao() *IpSumDao {
	return &IpSumDao{baseDao{conn: gorms.NewConn()}}
}

func (i *IpSumDao) SaveHackerIp(tx database.DbConn, ctx context.Context, hackerIp *model.IpSum) error {
	conn := i.getConn(tx)
	err := conn.Session(ctx).Save(hackerIp).Error
	return err
}

func (i *IpSumDao) FindIpSumByIp(tx database.DbConn, ctx context.Context, ip string) (ipSum *model.IpSum, err error) {
	conn := i.getConn(tx)
	err = conn.Session(ctx).Where("hacker_ip = ?", ip).Find(&ipSum).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return ipSum, err
}
