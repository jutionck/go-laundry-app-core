package main

import (
	"fmt"
	"github.com/jutionck/go-laundry-app-core/config"
	"github.com/jutionck/go-laundry-app-core/repository"
	"github.com/jutionck/go-laundry-app-core/usecase"
	"log"
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

	customerRepository := repository.NewCustomerRepository(db)
	customerUseCase := usecase.NewCustomerUseCase(customerRepository)

	//err := customerUseCase.UpdateCustomer(&model.Customer{
	//	Id:          "02d26b0a-8fb8-40ee-b7db-9d991bdb3a00",
	//	Name:        "Destry Faradila",
	//	PhoneNumber: "082180221161",
	//})
	//if err != nil {
	//	fmt.Println(err)
	//}

	customers, err := customerUseCase.FindAllCustomer(1, 5, "", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(customers)
}
