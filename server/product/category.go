package product

type Category struct {
	ID          uint `gorm:"primarykey"`
	Description string
	Label       string

	Products []Product `gorm:"-"`
}


type CategoryRepo interface {
	// Details(id uint) (Product, error)
	List() ([]Category, error)
	// Upsert(id uint) error
	// Delete(id uint) error
}
