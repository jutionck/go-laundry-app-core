package repository

import (
	"errors"
	"github.com/jutionck/go-laundry-app-core/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BillRepository interface {
	Create(bill *model.Bill) error
	FindById(billId string) (model.Bill, error)
	FindAll(page int, itemPerPage int, by string, values ...interface{}) ([]model.Bill, error)
}

type billRepository struct {
	db *gorm.DB
}

func (c *billRepository) Create(bill *model.Bill) error {
	return c.db.Create(bill).Error
}

func (c *billRepository) FindById(billId string) (model.Bill, error) {
	bill := model.Bill{}
	result := c.db.Unscoped().Preload(clause.Associations).Find(&bill, "id = ?", billId)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return bill, nil
		} else {
			return bill, err
		}
	}
	return bill, nil
}

func (c *billRepository) FindAll(page int, itemPerPage int, by string, values ...interface{}) ([]model.Bill, error) {
	var bills []model.Bill
	offset := itemPerPage * (page - 1)
	res := c.db.Order("created_at").Limit(itemPerPage).Offset(offset).Preload(clause.Associations).Find(&bills)
	if by != "" {
		res = c.db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Where(by, values...).Preload(clause.Associations).Find(&bills)
	}
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return bills, nil

}

func NewBillRepository(db *gorm.DB) BillRepository {
	repo := new(billRepository)
	repo.db = db
	return repo
}
