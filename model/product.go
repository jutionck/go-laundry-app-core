package model

import (
	generateid "github.com/jutionck/generate-id"
	"gorm.io/gorm"
)

type Product struct {
	Id            string `gorm:"primaryKey"`
	Name          string
	Duration      int
	ProductPrices []ProductPrice
	BaseModel     BaseModel `gorm:"embedded"`
}

func (u *Product) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = generateid.GenerateId()
	return
}
