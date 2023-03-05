package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CategoryID      uint
	SKU             string
	Label           string
	Description     string
	Characteristics string
	Price           float64 // Price for 50 grams
	Stock           int64   // Stock in gramms
	ImageURL        string

	Category Category
}


type ProductRepo interface {
	//TODO: pagination
	// Details(id uint) (Product, error)
	List(page, pageSize int) ([]Product, error)
	// Upsert(id uint) error
	// Delete(id uint) error
}
