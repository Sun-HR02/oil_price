package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	dsn := "root:hanjinhao@tcp(localhost:3306)/prices?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}
