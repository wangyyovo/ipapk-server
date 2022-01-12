package models

import (
	"github.com/wangyyovo/ipapk-server/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var orm *gorm.DB

func InitDB() error {
	var err error

	gormConfig := gorm.Config{}

	gormConfig.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	orm, err = gorm.Open(mysql.Open(conf.AppConfig.Database))
	if err != nil {
		return err
	}

	orm.AutoMigrate(&Bundle{})
	return nil
}
