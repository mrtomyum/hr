package models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
	Users []User
}

type Title struct { //ตำแหน่งงาน
	gorm.Model
	TitleID uint
	TH string
	EN string
}

type Org struct {
	gorm.Model
	OrgID uint
	TH string
	EN string
	Short string `gorm:"type:char(3)"`
}

