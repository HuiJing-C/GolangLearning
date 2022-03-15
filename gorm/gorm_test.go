package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

type Product struct {
	Code  string `gorm:"primaryKey"`
	Price uint
}

func TestOne(t *testing.T) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s", "root", "root", "localhost", "3306", "mytest")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Migrate the schema
	if err = db.AutoMigrate(&Product{}); err != nil {
		panic(err)
	}
	// Create
	db.Create(&Product{Code: "D42", Price: 100})
}
