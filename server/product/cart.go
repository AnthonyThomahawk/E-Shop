package product

import "github.com/AnthonyThomahawk/E-Shop/server/user"

type Cart struct {
	Items []CartItem
}

type CartItem struct {
	Product  Product
	Quantity uint
}

// Represents the N-N table in the database
type UserProductCart struct {
	UserID    uint `gorm:"primaryKey;autoIncrement:false"`
	ProductID uint `gorm:"primaryKey;autoIncrement:false"`
	Quantity  uint

	User    user.User `gorm:"-" json:"-"`
	Product Product
}

type CartRepo interface {
	List(page, pageSize int, userID uint) ([]UserProductCart, error)
	Insert(userID, productID, quantity uint) error
	Update(userID, productID, quantity uint) error
	Delete(userID, productID uint) error
	// Upsert(id uint) error
	// Delete(id uint) error
}
