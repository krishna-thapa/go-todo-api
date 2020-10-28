package migration

import (
	"todoAPI/model"

	"github.com/jinzhu/gorm"
)
/*
 functions AutoMigrate and a reference to the User and Task models 
 to create the tables in the database in case it not exists.
*/

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.User{})
}
