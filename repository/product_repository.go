package repository

import (
	"errors"
	"fmt"
	"github.com/jutionck/go-laundry-app-core/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(productId string) error
	FindById(productId string) (*model.Product, error)
	FindAll(page int, itemPerPage int, by string, values ...interface{}) ([]model.Product, error)
	FindByProductPriceId(productPriceId string) (*model.ProductPrice, error)
	FindByName(name string) (*model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func (c *productRepository) Create(product *model.Product) error {
	return c.db.Create(product).Error
}

func (c *productRepository) Update(product *model.Product) error {
	return c.db.Updates(product).Error
}

func (c *productRepository) Delete(productId string) error {
	return c.db.Delete(&model.Product{Id: productId}).Error
}

func (c *productRepository) FindById(productId string) (*model.Product, error) {
	product := model.Product{}
	result := c.db.Unscoped().Preload("ProductPrices").Find(&product, "id = ?", productId)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &product, nil
}

func (c *productRepository) FindAll(page int, itemPerPage int, by string, values ...interface{}) ([]model.Product, error) {
	var products []model.Product
	offset := itemPerPage * (page - 1)
	res := c.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Preload("ProductPrices").Find(&products)
	if by != "" {
		res = c.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Where(by, values...).Preload("ProductPrices").Find(&products)
	}
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return products, nil
}

func (c *productRepository) FindByProductPriceId(productPriceId string) (*model.ProductPrice, error) {
	productPrice := model.ProductPrice{}
	result := c.db.Unscoped().Preload("Product").Find(&productPrice, "id = ?", productPriceId)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &productPrice, nil
}

func (c *productRepository) FindByName(name string) (*model.Product, error) {
	product := model.Product{}
	err := c.db.Unscoped().Where("name=?", name).First(&product).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("productRepo.FindByName : %w", err)
	}

	return &product, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	repo := new(productRepository)
	repo.db = db
	return repo
}
