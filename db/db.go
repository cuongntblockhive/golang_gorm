package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"my_gorm/entity"
)

var db *gorm.DB
var err error

func Init(){
	dsn := "host=localhost user=postgres password=25011998 dbname=test_gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect Fail")
	} else {
		fmt.Println("Connect Successfully")
		db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
		db.AutoMigrate(&entity.User{})
	}
}

func Connect() *gorm.DB {
	if db == nil {
		Init()
	}
	return db
}



