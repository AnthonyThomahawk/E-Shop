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

	// TODO: Second iteration
	// AvgRating       *float32
	// Picture

	Category Category
}

type Category struct {
	ID          uint `gorm:"primarykey"`
	Description string
	Label       string

	Products []Product
}
