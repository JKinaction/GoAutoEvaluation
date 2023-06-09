package common

import (
	"domo1/config"
	"domo1/util/model"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := config.Conf.Mysql.Host
	port := config.Conf.Mysql.Port
	database := config.Conf.Mysql.Database
	username := config.Conf.Mysql.Username
	password := config.Conf.Mysql.Password
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		port,
		database)
	// args := fmt.Sprintf("root:123@tcp(mysql:3306)/gae?charset=utf8&parseTime=true")
	logrus.Info(args)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)           // 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxOpenConns(100)          //设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) //设置了连接可复用的最大时间
	//数据库迁移
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Question{})
	db.AutoMigrate(&model.InputAnswer{})
	DB = db
	logrus.Info("数据库初始化成功")
}

func GetDB() *gorm.DB {
	return DB
}
