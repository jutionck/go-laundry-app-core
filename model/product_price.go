package model

type ProductPrice struct {
    Id string
    Price int
    IsActive bool
    ProductID string
    BaseModel BaseModel `gorm:"embedded"`
}