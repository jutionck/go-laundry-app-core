package main

import (
	"fmt"
	"github.com/jutionck/go-laundry-app-core/config"
	"github.com/jutionck/go-laundry-app-core/repository"
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

	//customerRepository := repository.NewCustomerRepository(db)
	//productPriceRepository := repository.NewProductRepository(db)
	billRepository := repository.NewBillRepository(db)

	// Transaction
	//customer, _ := customerRepository.FindById("bfeddb64-2a19-4c8a-a445-18892f05ae1a")
	//product, _ := productPriceRepository.FindById("d18d590e-00d6-411f-81bc-9d50caa5574e")
	//productPrice, _ := productPriceRepository.FindByProductPriceId("4483460b-68d3-43c6-8300-5a76e50c7e69")
	//fmt.Println("customer:", customer)
	//fmt.Println("product:", product)
	//fmt.Println("productPrice:", productPrice)
	//
	//newBill01 := model.Bill{
	//	Date:       time.Now(),
	//	CustomerID: customer.Id,
	//	BillDetails: []model.BillDetail{
	//		{
	//			Weight:         8,
	//			ProductPriceID: productPrice.Id,
	//		},
	//	},
	//}
	//
	//err := billRepository.Create(&newBill01)
	//if err != nil {
	//	fmt.Println(err)
	//}

	allTransaction, err := billRepository.FindAll(1, 10, "", "")
	if err != nil {
		fmt.Println(err)
	}
	for _, bill := range allTransaction {
		fmt.Println("BillID:", bill.Id)
		fmt.Println("Date:", bill.Date)
		fmt.Println("Customer:", bill.Customer.Name)
		fmt.Println("BillDetails:", bill.BillDetails)
	}
	// Insert
	//newProduct := model.Product{
	//	Name:     "Cuci Komplit",
	//	Duration: 3,
	//	ProductPrices: []model.ProductPrice{
	//		{
	//			Price: 8000,
	//		},
	//	},
	//}
	//err := productRepository.Create(&newProduct)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// bfeddb64-2a19-4c8a-a445-18892f05ae1a
	//oldProduct := model.Product{
	//	Id:          "bfeddb64-2a19-4c8a-a485-18892f05ae1a",
	//	Name:        "Jution CK",
	//	PhoneNumber: "082180221160",
	//}
	//err := productRepository.Update(&oldProduct)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//productId := "bfeddb64-2a19-4c8a-a445-18892f05ae1a"
	//err := productRepository.Delete(productId)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//product, err := productRepository.FindById(productId)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(product)

	// Find All
	//keyword := "name ilike ?"
	//value := "cuci"
	//products, err := productRepository.FindAll(1, 10, "", "")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, product := range products {
	//	fmt.Println("Name: ", product.Name)
	//	fmt.Println("Duration: ", product.Duration)
	//	for _, productPrice := range product.ProductPrices {
	//		fmt.Println("ProductID: ", productPrice.ProductID)
	//		fmt.Println("Price: ", productPrice.Price)
	//	}
	//}
}
