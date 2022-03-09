package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/google/uuid"
)
type typetime time.Time 
type User struct {
	gorm.Model
    ID uuid.UUID `sql:"type:varchar(36);primary key"`
    Name string
  }
  
func main() {
	fmt.Println("Hellow word")
	dsn := "host=localhost user=postgres password=Abc12345 dbname=test_gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect Fail")
	} else {
		fmt.Println("Connect Successfully" , db)
		db.AutoMigrate(&User{})
		user := User{ID: uuid.New(), Name: "John"}
		result := db.Create(&user)
		if result != nil {
			fmt.Println("Create Successfully" , result)
		}
	}
	
}
