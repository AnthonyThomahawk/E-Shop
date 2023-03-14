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

func (repo *UserRepoGORM) Load(email string) (*user.User, error) {
	var usr user.User
	err := repo.db.Model(&user.User{}).Where("email = ?", email).First(&usr).Error
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func (repo *UserRepoGORM) Save(user user.User) error {
	return repo.db.Create(&user).Error
}
