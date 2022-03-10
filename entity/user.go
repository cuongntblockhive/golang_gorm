package entity

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID 			string 			`json:"id" gorm:"type:UUID;primary_key;default:uuid_generate_v4();"`
    Name 		string			`json:"name"`
	Email		string 			`json:"email" gorm:"not null;unique;size:20;column:email_column;"`
	Age 		int				`json:"age" gorm:"default:18"`
	CreatedAt 	time.Time		`json:"createdAt"`
  	UpdatedAt 	time.Time		`json:"updatedAt"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleteAt" gorm:"index"`
	TestMigate 	string			`json:"testMigate" gorm:"-"`
  }
