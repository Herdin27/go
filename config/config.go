package config

import (
	"api/learning/structs"
	"fmt"

	"github.com/jinzhu/gorm"
)

func DBinit() *gorm.DB{
	db, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@/%v", "root", "", "test"))
	if err != nil {
		return nil
	}
	db.AutoMigrate(structs.Person{})
	return db
}