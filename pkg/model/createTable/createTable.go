package main

import (
	"log"
	"secureQR/pkg/connections/database/gorms"
	"secureQR/pkg/model"
)

/*
创建数据表
*/
func main() {
	db := gorms.GetDB()
	if db == nil {
		log.Fatal("创建数据表的获取DB操作出现了错误")
	}
	// 自动迁移，漏洞表
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Notice{})
	if err != nil {
		log.Fatal("创建表错误")
		return
	}
	// 设置表的备注
	db.Exec("ALTER TABLE notice COMMENT '公告'")
}
