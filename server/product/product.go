package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryID      uint
	SKU             string `gorm:"uniqueIndex"`
	Label           string
	Description     string
	Characteristics string
	Price           float64 // Price for 50 grams
	Stock           int64   // Stock in gramms
	ImageURL        string

	Category         Category          `gorm:"-" json:"-"`
	UserProductCarts []UserProductCart `gorm:"-" json:"-"`
}

type ProductRepo interface {
	Details(id int) (Product, error)
	List(page, pageSize int, categoryID *int) ([]Product, error)
	// Upsert(id uint) error
	// Delete(id uint) error
}
