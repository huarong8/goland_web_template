package database

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"gorm.io/driver/mysql"
	"time"

	"gitlab.feedtoken.tech/xdance_group/xdance/config"
	"gorm.io/gorm"
)

func SetUp() *gorm.DB{
	host := config.Ops.Mysql.Host
	port := config.Ops.Mysql.Port
	dbName := config.Ops.Mysql.Db
	user := config.Ops.Mysql.User
	passwd := config.Ops.Mysql.Passwd


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user,passwd, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil{
		panic(err)
	}
	iris.RegisterOnInterrupt(func(){
		defer sqlDB.Close()
	})
	sqlDB.SetMaxIdleConns(config.Ops.Mysql.MaxIdleConnects)
	sqlDB.SetMaxOpenConns(config.Ops.Mysql.MaxConnects)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.Ops.Mysql.ConnMaxIdleTime) * time.Second)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Ops.Mysql.ConnMaxLifetime) * time.Second)
	return db
}