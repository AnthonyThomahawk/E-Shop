package database

import (
	"github.com/AnthonyThomahawk/E-Shop/server/user"
	"gorm.io/gorm"
)

type RoleRepoGORM struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) *RoleRepoGORM {
	return &RoleRepoGORM{db: db}
}

func(repo*RoleRepoGORM) Details(id int) (user.Role, error) {
	var role user.Role 
	err := repo.db.Model(&user.Role{}).First(&role, id).Error
	return role, err
}
