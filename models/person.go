package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Person struct {
	gorm.Model
	Users      []User
	Jobs       []Job `gorm:"many2many:person_job"`
	Emails     []Email
	FirstName  string
	LastName   string
	BirthDate  time.Time
	CitizenID  string
	PassportID string
}

type Job struct { //ตำแหน่งงาน
	gorm.Model
	JobID          uint
	OrgID          uint
	NameTH         string
	NameEN         string
	Description    string //รายละเอียดงาน
	ProfessionRate uint   //เงินค่าตำแหน่ง
}

type Org struct {
	gorm.Model
	OrgID     uint
	NameTH    string
	NameEN    string
	NameShort string `gorm:"type:char(3)"`
}

type Email struct {
	gorm.Model
	PersonID uint
	Email  string
}
