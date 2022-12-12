package usecase

import (
	"github.com/jutionck/go-laundry-app-core/model"
	"github.com/jutionck/go-laundry-app-core/repository"
)

type BillUseCase interface {
	RegisterNewBill(bill *model.Bill) error
	FindByBillId(billId string) (model.Bill, error)
	FindAllBill(page int, itemPerPage int, by string, values ...interface{}) ([]model.Bill, error)
}

type billUseCase struct {
	billRepo     repository.BillRepository
	customerRepo repository.CustomerRepository
	productRepo  repository.ProductRepository
}

func (b *billUseCase) RegisterNewBill(bill *model.Bill) error {
	customer, _ := b.customerRepo.FindById(bill.CustomerID)
	for _, prd := range bill.BillDetails {
		productPrice, _ := b.productRepo.FindByProductPriceId(prd.ProductPriceID)
		bill.BillDetails = []model.BillDetail{
			{
				Weight:         prd.Weight,
				ProductPriceID: productPrice.Id,
			},
		}
	}

	bill.CustomerID = customer.Id
	return b.billRepo.Create(bill)

}

func (b *billUseCase) FindByBillId(billId string) (model.Bill, error) {
	return b.billRepo.FindById(billId)
}

func (b *billUseCase) FindAllBill(page int, itemPerPage int, by string, values ...interface{}) ([]model.Bill, error) {
	return b.billRepo.FindAll(page, itemPerPage, by, values)
}

func NewBillUseCase(
	billRepo repository.BillRepository,
	customerRepo repository.CustomerRepository,
	productRepo repository.ProductRepository) BillUseCase {
	return &billUseCase{
		billRepo:     billRepo,
		customerRepo: customerRepo,
		productRepo:  productRepo,
	}
}
