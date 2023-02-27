package product

type ProductRepo interface {
	//TODO: pagination
	// Details(id uint) (Product, error)
	List() ([]Product, error)
	// Upsert(id uint) error
	// Delete(id uint) error
}
