package model

import "gorm.io/gorm"

// 表名称

const TableNameNotice = "notice"

type Notice struct {
	gorm.Model
	Title string `gorm:"column:title;type:varchar(50);not null;comment:公告标题" json:"title"` // 公告标题
}

func (*Notice) TableName() string {
	return TableNameNotice
}
