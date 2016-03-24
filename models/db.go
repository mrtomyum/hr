package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func SetupDB(database, conn string) *gorm.DB{
	db, err := gorm.Open(database, conn)
	if err != nil {
		panic(err.Error())
	}
	db.DB().Ping()
	if err != nil {
		panic(err.Error())
	}
	db.SingularTable(true)
	return db
}
