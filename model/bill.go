package model

import (
	generateid "github.com/jutionck/generate-id"
	"gorm.io/gorm"
	"time"
)

type Bill struct {
	Id          string `gorm:"primaryKey"`
	Date        time.Time
	CustomerID  string
	Customer    Customer
	BillDetails []BillDetail
	BaseModel   BaseModel `gorm:"embedded"`
}

func (u *Bill) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = generateid.GenerateId()
	return
}
