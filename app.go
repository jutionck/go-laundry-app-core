package main

import (
	"database/sql"
	"fmt"
	"github.com/jutionck/go-laundry-app-core/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "jutioncandrakirana"
	dbPassword := "password"
	dbName := "db_enigma_laundry"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	enigmaDb, err := db.DB()
	defer func(enigmaDb *sql.DB) {
		err := enigmaDb.Close()
		if err != nil {
			panic(err)
		}

	}(enigmaDb)
	err = db.Debug().AutoMigrate(
		&model.Customer{},
		&model.Product{},
		&model.ProductPrice{},
		&model.BillDetail{},
		&model.Bill{})
	if err != nil {
		return
	}
}
