package user

type Role struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type RoleRepo interface {
	Details(id int) (Role, error)
	// Save(role Role) error
}
