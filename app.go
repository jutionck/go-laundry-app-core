package main

import (
	"fmt"
	"github.com/jutionck/go-laundry-app-core/config"
	"github.com/jutionck/go-laundry-app-core/model"
	"github.com/jutionck/go-laundry-app-core/repository"
	"github.com/jutionck/go-laundry-app-core/usecase"
	"log"
	"time"
)

func main() {
	cfg := config.NewConfig()
	db := cfg.DbConn()
	defer func(cfg *config.Config) {
		err := cfg.DbClose()
		if err != nil {
			log.Println(err.Error())
		}
	}(&cfg)

	customerRepo := repository.NewCustomerRepository(db)
	productRepo := repository.NewProductRepository(db)
	billRepo := repository.NewBillRepository(db)

	billUseCase := usecase.NewBillUseCase(billRepo, customerRepo, productRepo)

	newBill01 := model.Bill{
		Date:       time.Now(),
		CustomerID: "2a01c39a-3607-4a39-8830-171b0c5a117a",
		BillDetails: []model.BillDetail{
			{
				Weight:         4,
				ProductPriceID: "4483460b-68d3-43c6-8300-5a76e50c7e69",
			},
		},
	}

	err := billUseCase.RegisterNewBill(&newBill01)
	if err != nil {
		fmt.Println(err)
	}
}
