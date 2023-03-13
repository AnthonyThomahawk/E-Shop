package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	RoleID   uint

	Role Role `gorm:"-" json:"-"`
}

type UserRepo interface {
	Load(username string) error
}
