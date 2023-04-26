package model

import "gorm.io/gorm"

// 表名称

const TableNamePicture = "picture"

type Picture struct {
	gorm.Model
	Src  string `gorm:"column:src;type:varchar(1000);not null;comment:图片地址" json:"src"` // 图片地址
	Type string `gorm:"column:type;type:varchar(30);not null;comment:图片类型" json:"type"` // 图片类型
}

func (*Picture) TableName() string {
	return TableNamePicture
}
