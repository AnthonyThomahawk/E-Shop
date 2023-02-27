package database

import (
	"github.com/AnthonyThomahawk/E-Shop/server/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&product.Product{})
	if err != nil {
		return nil, err
	}

	err = SeedData(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
