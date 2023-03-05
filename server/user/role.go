package user

type Role struct {
	ID   uint `gorm:"primarykey"`
	Name string
}
