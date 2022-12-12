package manager

import (
	"fmt"
	"github.com/jutionck/go-laundry-app-core/config"
	"github.com/jutionck/go-laundry-app-core/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type InfraManager interface {
	SqlDb() *gorm.DB
}

type infraManager struct {
	db  *gorm.DB
	cfg config.Config
}

func (i *infraManager) SqlDb() *gorm.DB {
	return i.db
}

func (i *infraManager) initDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", i.cfg.Host, i.cfg.User, i.cfg.Password, i.cfg.DbName, i.cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	i.db = db
	env := os.Getenv("ENV")
	if env == "migration" {
		err := db.Debug().AutoMigrate(
			&model.Customer{},
			&model.Product{},
			&model.ProductPrice{},
			&model.BillDetail{},
			&model.Bill{},
		)
		if err != nil {
			panic(err)
			return
		}
	} else if env == "dev" {
		db.Debug()
	}
}

func NewInfra(config config.Config) InfraManager {
	infra := infraManager{cfg: config}
	infra.initDb()
	return &infra
}
