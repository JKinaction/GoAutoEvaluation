package common

import (
	"domo1/util/model"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		port,
		database)

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

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
