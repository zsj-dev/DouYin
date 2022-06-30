package initialization

import (
	"fmt"
	"log"

	"github.com/zsj-dev/DouYin/basic-service/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func RegisterMySQL() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		//"root", "1234", "127.0.0.1", 3306, "byte-douyin", "utf8mb4", "true", "Local",
		"root", "123456", "127.0.0.1", 3306, "byte-douyin", "utf8mb4", "true", "Local",
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panicf("初始化 MySQL 异常: %v", err)
	}
	conf.MySQL = db
}
