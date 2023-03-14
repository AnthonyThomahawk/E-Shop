package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string
	PasswordHash string
	RoleID   uint

	Role Role `gorm:"-" json:"-"`
}

type UserRepo interface {
	Load(email string) (*User, error)
	Save(user User) error
}
