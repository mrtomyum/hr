package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Impl struct {
	DB *gorm.DB
}

func (i Impl) InitDB() {
	var err error
	i.DB, err = gorm.Open("sqlite3", "hr.db")
	if err != nil {
		panic(err.Error())
	}
	i.DB.SingularTable(true)
}

func SetupDB(database, location string) *gorm.DB{
	db, err := gorm.Open(database, location)
	if err != nil {
		panic(err.Error())
	}
	db.SingularTable(true)
	return db
}