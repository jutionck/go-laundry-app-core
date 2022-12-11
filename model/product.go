package model

type Product struct {
    Id string `gorm:"primaryKey"`
    Name string
    Duration int
    ProductPrices []ProductPrice
    BaseModel BaseModel `gorm:"embedded"`
}