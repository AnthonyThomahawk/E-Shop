package database

import (
	"github.com/AnthonyThomahawk/E-Shop/server/user"
	"gorm.io/gorm"
)

type UserRepoGORM struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepoGORM {
	return &UserRepoGORM{db: db}
}

func (repo *UserRepoGORM) Details(id int) (user.User, error) {
	var usr user.User
	err := repo.db.Model(&user.User{}).First(&usr, id).Error
	return usr, err
}
