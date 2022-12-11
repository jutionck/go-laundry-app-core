package model

type ProductPrice struct {
	Id        string
	Price     int
	IsActive  bool `gorm:"default:true"`
	ProductID string
	BaseModel BaseModel `gorm:"embedded"`
}
