package repository

import (
	"errors"
	"github.com/jutionck/go-laundry-app-core/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	Update(customer *model.Customer) error
	Delete(customerId string) error
	FindById(customerId string) (model.Customer, error)
	FindAll(page int, itemPerPage int, by string, values ...interface{}) ([]model.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) Create(customer *model.Customer) error {
	return c.db.Create(customer).Error
}

func (c *customerRepository) Update(customer *model.Customer) error {
	return c.db.Updates(customer).Error
}

func (c *customerRepository) Delete(customerId string) error {
	//return c.db.Delete(&model.Customer{}, "id="+"'"+customerId+"'").Error
	return c.db.Delete(&model.Customer{Id: customerId}).Error
}

func (c *customerRepository) FindById(customerId string) (model.Customer, error) {
	customer := model.Customer{}
	// UnScoped() -> akan men-select delete is NULL
	result := c.db.Unscoped().Find(&customer, "id = ?", customerId)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) FindAll(page int, itemPerPage int, by string, values ...interface{}) ([]model.Customer, error) {
	var customers []model.Customer
	offset := itemPerPage * (page - 1)
	res := c.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Find(&customers)
	if by != "" {
		res = c.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Where(by, values...).Find(&customers)
	}
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customers, nil

}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
