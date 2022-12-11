package model

import "time"

type Bill struct {
    Id string `gorm:"primaryKey"`
    Date time.Time
    CustomerID string
    BillDetails []BillDetail
    BaseModel BaseModel `gorm:"embedded"`
}