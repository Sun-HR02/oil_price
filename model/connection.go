package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"oil_price_show/conf"
)

func NewConnection() *gorm.DB {
	dburl, dbport, dbpasswd := conf.DBconf()
	dsn := "root:" + dbpasswd + "@tcp(" + dburl + ":" + dbport + ")/home?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
