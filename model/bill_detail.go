package model

import (
	generateid "github.com/jutionck/generate-id"
	"gorm.io/gorm"
)

type BillDetail struct {
	Id             string `gorm:"primaryKey"`
	Weight         float32
	BillID         string
	ProductPriceID string
	ProductPrice   ProductPrice
	BaseModel      BaseModel `gorm:"embedded"`
}

func (u *BillDetail) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = generateid.GenerateId()
	return
}
