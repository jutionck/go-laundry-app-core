package model

import (
	generateid "github.com/jutionck/generate-id"
	"gorm.io/gorm"
)

type ProductPrice struct {
	Id        string `gorm:"primaryKey"`
	Price     int
	IsActive  bool `gorm:"default:true"`
	ProductID string
	Product   Product
	BaseModel BaseModel `gorm:"embedded"`
}

func (u *ProductPrice) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = generateid.GenerateId()
	return
}
