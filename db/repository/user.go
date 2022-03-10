package repository

import (
	"fmt"
	"gorm.io/gorm"
	"my_gorm/entity"
)

func CreateDummyData(tx *gorm.DB)(*entity.User , error){
	user := entity.User{Name: "John", Email: "1zdfgxc@gmail.com"}
	result := tx.Create(&user)
	return &user,result.Error
}

func GetListUsers(tx *gorm.DB)(*[]entity.User , error){
	var users []entity.User
	result := tx.Find(&users)
	return &users , result.Error
}

func GetUserById(tx *gorm.DB , id string)(*entity.User , error) {
	var user entity.User;
	result := tx.Find(&user , "id = ?" , id )
	return &user,result.Error
}