package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

	var DB *gorm.DB

func ConnectToDb(){
	var err error
	dsn := "host=ruby.db.elephantsql.com user=flzgwqxk password=6EveCv6CIVTkF7Tp0EALii3YIqlitQw6 dbname=flzgwqxk port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db.")
	}
}
