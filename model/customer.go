package model

import (
	generateid "github.com/jutionck/generate-id"
	"gorm.io/gorm"
)

type Customer struct {
	Id          string `gorm:"primaryKey"`
	Name        string
	PhoneNumber string
	Bills       []Bill
	BaseModel   BaseModel `gorm:"embedded"`
}

func (u *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = generateid.GenerateId()
	return
}
