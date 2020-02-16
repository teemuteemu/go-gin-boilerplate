package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"hello/config"
)

func Connect(config *config.Config) *gorm.DB {
	connectionStr := fmt.Sprintf(
		"host=%s sslmode=disable port=%d dbname=%s user=%s password=%s",
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBUser,
		config.DBPassword,
	)

	db, err := gorm.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}

	return db
}
