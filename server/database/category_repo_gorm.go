package database


import (
	"github.com/AnthonyThomahawk/E-Shop/server/product"
	"gorm.io/gorm"
)

type CategoryRepoGORM struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepoGORM {
	return &CategoryRepoGORM{db: db}
}

func (repo *CategoryRepoGORM) Details(id uint) (product.Category, error) {
	var category product.Category
	err := repo.db.Model(&product.Category{}).First(&category).Error
	return category, err
}

func (repo *CategoryRepoGORM) List() ([]product.Category, error) {
	var categories []product.Category
	err := repo.db.Find(&categories).Error
	return categories, err
}

// TODO: implement upsert??
func (repo *CategoryRepoGORM) Insert(category product.Category) error {
	return repo.db.Create(category).Error
}
func (repo *CategoryRepoGORM) Update(id uint) error {
	panic("not implemented")
}
func (repo *CategoryRepoGORM) Delete(id uint) error {
	panic("not implemented")
}
