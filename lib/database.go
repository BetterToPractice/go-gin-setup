package lib

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	ORM *gorm.DB
}

func NewDatabase() Database {
	dbConfig := postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s",
			"localhost",
			"setup_user",
			"setup_password",
			"setup_db",
			5432,
			"Asia/Jakarta",
		),
	}

	db, err := gorm.Open(postgres.New(dbConfig), &gorm.Config{})
	if err != nil {
		fmt.Println("error when connect with database", err)
	}
	return Database{
		ORM: db,
	}
}
