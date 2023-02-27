package database

import (
	"github.com/AnthonyThomahawk/E-Shop/server/product"
	"gorm.io/gorm"
)

type ProductRepoGORM struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepoGORM {
	return &ProductRepoGORM{db: db}
}

func (repo *ProductRepoGORM) Details(id uint) (product.Product, error) {
	var prd product.Product
	err := repo.db.Model(&product.Product{}).Preload("Category").First(&prd).Error
	return prd, err
}

func (repo *ProductRepoGORM) List() ([]product.Product, error) {
	var prds []product.Product
	err := repo.db.Model(&product.Product{}).Preload("Category").Find(&prds).Error
	return prds, err
}

// TODO: implement upsert
func (repo *ProductRepoGORM) Insert(prd product.Product) error {
	return repo.db.Create(prd).Error
}
func (repo *ProductRepoGORM) Update(id uint) error {
	panic("not implemented")
}
func (repo *ProductRepoGORM) Delete(id uint) error {
	panic("not implemented")
}
