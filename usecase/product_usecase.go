package usecase

import (
	"fmt"
	"github.com/jutionck/go-laundry-app-core/model"
	"github.com/jutionck/go-laundry-app-core/repository"
)

type ProductUseCase interface {
	RegisterNewProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(productId string) error
	FindByProductId(productId string) (*model.Product, error)
	FindAllProduct(page int, itemPerPage int, by string, values ...interface{}) ([]model.Product, error)
	FindByProductPriceId(productPriceId string) (*model.ProductPrice, error)
}

type productUseCase struct {
	productRepo repository.ProductRepository
}

func (p *productUseCase) RegisterNewProduct(product *model.Product) error {
	prd, err := p.productRepo.FindByName(product.Name)
	if err != nil {
		return fmt.Errorf("productUseCase.RegisterNewProduct : %w", err)
	}

	if prd != nil {
		return fmt.Errorf("product with name %s already exists", product.Name)
	}
	err = p.productRepo.Create(product)
	if err != nil {
		return fmt.Errorf("productUseCase.RegisterNewProduct : %w", err)
	}
	return nil
}

func (p *productUseCase) UpdateProduct(product *model.Product) error {
	prd, err := p.productRepo.FindById(product.Id)
	if err != nil {
		return fmt.Errorf("productUseCase.UpdateProduct : %w", err)
	}

	if prd == nil {
		return fmt.Errorf("product with ID %s not exists", product.Id)
	}

	prd, err = p.productRepo.FindByName(product.Name)
	if err != nil {
		return fmt.Errorf("productUseCase.UpdateProduct : %w", err)
	}

	if prd != nil && prd.Id != product.Id {
		return fmt.Errorf("product with phone number %s already exists", product.Name)
	}
	err = p.productRepo.Update(product)
	if err != nil {
		return fmt.Errorf("productUseCase.UpdateProduct : %w", err)
	}
	return nil
}

func (p *productUseCase) DeleteProduct(productId string) error {
	prd, err := p.productRepo.FindById(productId)
	if err != nil {
		return fmt.Errorf("productUseCase.DeleteProduct : %w", err)
	}

	if prd == nil {
		return fmt.Errorf("product with ID %s not exists", productId)
	}

	err = p.productRepo.Delete(productId)
	if err != nil {
		return fmt.Errorf("productUseCase.DeleteProduct : %w", err)
	}
	return nil
}

func (p *productUseCase) FindByProductId(productId string) (*model.Product, error) {
	product, err := p.productRepo.FindById(productId)
	if err != nil {
		return nil, fmt.Errorf("FindByCustomerId : %w", err)
	}
	return product, nil
}

func (p *productUseCase) FindAllProduct(page int, itemPerPage int, by string, values ...interface{}) ([]model.Product, error) {
	return p.productRepo.FindAll(page, itemPerPage, by, values)
}

func (p *productUseCase) FindByProductPriceId(productPriceId string) (*model.ProductPrice, error) {
	return p.productRepo.FindByProductPriceId(productPriceId)
}

func NewProductUseCase(productRepo repository.ProductRepository) ProductUseCase {
	return &productUseCase{productRepo: productRepo}
}
