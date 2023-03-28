package database

import (
	"gorm.io/gorm"

	"github.com/AnthonyThomahawk/E-Shop/server/product"
)

type CartRepoGORM struct {
	db *gorm.DB
}

func NewCartRepo(db *gorm.DB) *CartRepoGORM {
	return &CartRepoGORM{db: db}
}

func (repo *CartRepoGORM) List(page, pageSize int, userID uint) ([]product.UserProductCart, error) {
	var userProductCarts []product.UserProductCart
	err := repo.db.Model(&product.UserProductCart{}).
		Scopes(paginate(page, pageSize)).
		Where(&product.UserProductCart{UserID: userID}).
		InnerJoins("Product").
		Find(&userProductCarts).Error
	if err != nil {
		return nil, err
	}

	return userProductCarts, err
}
func (repo *CartRepoGORM) Insert(userID, productID, quantity uint) error {
	userProductCart := product.UserProductCart{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}

	return repo.db.Create(&userProductCart).Error
}

func (repo *CartRepoGORM) Update(userID, productID, quantity uint) error {
	result := repo.db.Model(&product.UserProductCart{}).
		Where(&product.UserProductCart{UserID: userID, ProductID: productID}).
		Update("quantity", quantity)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (repo *CartRepoGORM) Delete(userID, productID uint) error {
	return repo.db.Delete(&product.UserProductCart{}, map[string]any{
		"user_id":    userID,
		"product_id": productID,
	}).Error
}
