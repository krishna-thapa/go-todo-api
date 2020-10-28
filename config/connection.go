package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*
ORM Gorm the connection with the database is very easy, just need to have a
String connection and use the Open function, if any error occurs a panic
error is launched and stops the application, otherwise, a successful connection is returned.
*/

var DB *gorm.DB

// Update with latest version: https://gorm.io/docs/connecting_to_the_database.html
func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "host=postgrestodo port=5432 user=admin dbname=tododb password=123  sslmode=disable")

	if err != nil {
		panic(err.Error())
	}

	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
