package lib

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	ORM *gorm.DB
}

func NewDatabase(config Config) Database {
	dbConfig := postgres.Config{
		DSN: config.Database.DSN(),
	}

	db, err := gorm.Open(postgres.New(dbConfig), &gorm.Config{})
	if err != nil {
		fmt.Println("error when connect with database", err)
	}
	return Database{
		ORM: db,
	}
}
