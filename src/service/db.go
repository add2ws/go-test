package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func init() {
	//dsn := "root:root@tcp(localhost:3308)/mssgs?charset=utf8&parseTime=True&loc=Local"
	dsn := "root:root@123@tcp(172.17.0.2:3306)/mssgs?charset=utf8&parseTime=True&loc=Local"
	//dsn := "root:root@123@tcp(sons.top:8806)/mssgs?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接异常：" + err.Error())
	}
	DBConn = db
}
