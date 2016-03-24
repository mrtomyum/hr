package models

import (
	"testing"
	"fmt"
)

func TestModels(t *testing.T) {
	users := MockUser()
	for _, u := range users {
		fmt.Println(u)
	}
	persons := MockPerson()
	for _, p := range persons {
		fmt.Println(p)
	}
	titles := MockTitle()
	for _, t := range titles {
		fmt.Println(t)
	}
	orgs := MockOrg()
	for _, o := range orgs {
		fmt.Println(o)
	}
}

func TestDB(t *testing.T) {
	db := SetupDB("sqlite3", "hr.db")
	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&Person{})
	db.DropTableIfExists(&Title{})
	db.DropTableIfExists(&Org{})
	db.CreateTable(&User{})
	db.CreateTable(&Person{})
	db.CreateTable(&Title{})
	db.CreateTable(&Org{})

	users := MockUser()
	for _, u := range users {
		db.Create(u)
	}
	persons := MockPerson()
	for _, p := range persons {
		db.Create(p)
	}
	titles := MockTitle()
	for _, t := range titles {
		db.Create(t)
	}
	orgs := MockOrg()
	for _, o := range orgs {
		db.Create(o)
	}
}

func TestUserPassword(t *testing.T){
	db := SetupDB("sqlite3", "hr.db")
	u := User{}
	db.Debug().First(&u)
	fmt.Println(u)

	pass := "Thisisnewpass123"
	err := u.ChangePass(pass)
	if err != nil {
		t.Error("Fail ChangPass: ",err)
	}
	fmt.Println(u)
	db.Debug().Save(&u)
	db.Debug().First(&u)
	err = u.VerifyPass(pass)
	if err != nil {
		t.Error("Fail ComparePassword: ",err)
	}
	fmt.Println("Corrected Password!")
}
func TestUserWrongPassword(t *testing.T){
	db := SetupDB("sqlite3", "hr.db")
	u := User{}
	db.Debug().First(&u)
	fmt.Println(u)

	pass := "Thisisnewpass123"
	err := u.ChangePass(pass)
	if err != nil {
		t.Error("Fail ChangPass: ",err)
	}
	fmt.Println(u)
	db.Debug().Save(&u)
	db.Debug().First(&u)
	err = u.VerifyPass("ThisIsWrongPassword")
	if err != nil {
		fmt.Println(err, "<<<Wrong Password should rise err OK...Pass!!")
	} else {
		t.Error("Expected Wrong pass fail but compare pass!!")
		t.FailNow()
	}
	fmt.Println("EOP")

}