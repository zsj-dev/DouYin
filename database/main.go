package main

import (
	"fmt"
	"log"

	"github.com/zsj-dev/DouYin/database/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Favorite{},
		model.Relation{},
		model.Video{},
		model.Comment{},
	)
	if err != nil {
		log.Panicf("数据库迁移异常: %v", err)
	}
	log.Println("数据库迁移完成")
}
func main() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		"root", "123456", "127.0.0.1", 3306, "douyin", "utf8mb4", "true", "Local",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("初始化 MySQL 异常: %v", err)
	}
	autoMigrate(db)
}
