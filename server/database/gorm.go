package database

import (
	"errors"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AnthonyThomahawk/E-Shop/server/product"
	"github.com/AnthonyThomahawk/E-Shop/server/user"
)

func prepareDataSourceName() (string, error) {
	dataSourceName := strings.Builder{}

	if host, found := os.LookupEnv("DB_HOST"); found {
		dataSourceName.WriteString("host=")
		dataSourceName.WriteString(host)
	} else {
		return "", errors.New("DB_HOST environment variable not found")
	}

	if port, found := os.LookupEnv("DB_PORT"); found {
		dataSourceName.WriteString(" port=")
		dataSourceName.WriteString(port)
	} else {
		return "", errors.New("DB_PORT environment variable not found")
	}

	if user, found := os.LookupEnv("DB_USER"); found {
		dataSourceName.WriteString(" user=")
		dataSourceName.WriteString(user)
	} else {
		return "", errors.New("DB_USER environment variable not found")
	}

	if dbname, found := os.LookupEnv("DB_NAME"); found {
		dataSourceName.WriteString(" dbname=")
		dataSourceName.WriteString(dbname)
	} else {
		return "", errors.New("DB_NAME environment variable not found")
	}

	if password, found := os.LookupEnv("DB_PASSWORD"); found {
		dataSourceName.WriteString(" password=")
		dataSourceName.WriteString(password)
	} else {
		return "", errors.New("DB_PASSWORD environment variable not found")
	}

	return dataSourceName.String(), nil
}

func Setup() (*gorm.DB, error) {
	dataSourceName, err := prepareDataSourceName()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&product.Product{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&product.Category{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&user.Role{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&product.UserProductCart{})
	if err != nil {
		return nil, err
	}

	err = SeedData(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
