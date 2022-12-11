package model

type Customer struct {
    Id string `gorm:"primaryKey"`
    Name string
    PhoneNumber string
    Bills []Bill
    BaseModel BaseModel `gorm:"embedded"`
}