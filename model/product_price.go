package model

type ProductPrice struct {
	Id        string `gorm:"primaryKey"`
	Price     int
	IsActive  bool `gorm:"default:true"`
	ProductID string
	BaseModel BaseModel `gorm:"embedded"`
}
