package model

type BillDetail struct {
    Id string `gorm:"primaryKey"`
    Weight float32
    BillID string
    ProductPriceID string
    ProductPrice ProductPrice
    BaseModel BaseModel `gorm:"embedded"`
}