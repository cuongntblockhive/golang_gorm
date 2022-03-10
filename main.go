package main

import (
	"my_gorm/db"
	dbRepo "my_gorm/db/repository"
	"my_gorm/api/router"
)


  
func main() {
	db.Init()
	router.Init()
	ct := db.Connect()
	_ , _ = dbRepo.GetUserById(ct , "8d217b84-f581-4668-812b-9ef9ed0e601e")
}
