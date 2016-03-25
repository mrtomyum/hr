package models

import (
	"fmt"
	"testing"
)

//func TestModels(t *testing.T) {
//	users := MockUser()
//	for _, u := range users {
//		fmt.Println(u)
//	}
//	persons := MockPerson()
//	for _, p := range persons {
//		fmt.Println(p)
//	}
//	jobs := MockJob()
//	for _, t := range jobs {
//		fmt.Println(t)
//	}
//	orgs := MockOrg()
//	for _, o := range orgs {
//		fmt.Println(o)
//	}
//}

func TestGenMockData(t *testing.T) {
	db := SetupDB("sqlite3", "hr.db")
	db.DropTableIfExists(&User{})
	db.DropTableIfExists(&Person{})
	db.DropTableIfExists(&Job{})
	db.DropTableIfExists(&Org{})
	db.DropTableIfExists(&Email{})
	db.CreateTable(&User{})
	db.CreateTable(&Person{})
	db.CreateTable(&Job{})
	db.CreateTable(&Org{})
	db.CreateTable(&Email{})

	//users := MockUser()
	//for _, u := range users {
	//	db.Create(u)
	//}
	persons := MockPerson()
	for _, p := range persons {
		db.Create(p)
	}
	//jobs := MockJob()
	//for _, t := range jobs {
	//	db.Create(t)
	//}
	orgs := MockOrg()
	for _, o := range orgs {
		db.Create(o)
	}
	//emails := MockEmail()
	//for _, e := range emails {
	//	db.Create(e)
	//}
}

func TestUserPassword(t *testing.T) {
	db := SetupDB("sqlite3", "hr.db")
	u := User{}
	db.Debug().First(&u)
	fmt.Println(u)

	pass := "Thisisnewpass123"
	err := u.ChangePass(pass)
	if err != nil {
		t.Error("Fail ChangPass: ", err)
	}
	fmt.Println(u)
	db.Debug().Save(&u)
	db.Debug().First(&u)
	err = u.VerifyPass(pass)
	if err != nil {
		t.Error("Fail ComparePassword: ", err)
	}
	fmt.Println("Corrected Password!")
}
func TestUserWrongPassword(t *testing.T) {
	db := SetupDB("sqlite3", "hr.db")
	u := User{}
	db.Debug().First(&u)
	fmt.Println(u)

	pass := "Thisisnewpass123"
	err := u.ChangePass(pass)
	if err != nil {
		t.Error("Fail ChangPass: ", err)
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
}

func TestQueryEmailByPersonFirstName(t *testing.T) {
	db := SetupDB("sqlite3", "hr.db")
	expectEmail := "tom@nopadol.com"

	p := Person{}
	db.Debug().Where("first_name like ?", "%K%").Preload("Jobs").Preload("Users").Preload("Emails").Find(&p)
	fmt.Println(p.FirstName)
	for _, u := range p.Users {
		fmt.Println("User:", u)
	}
	fmt.Println("Jobs:", p.Jobs)
	fmt.Println("Emails:", p.Emails)

	count := 0
	for _, e := range p.Emails {
		if e.Email == expectEmail {
			count++
			fmt.Println(e.Email)
		}
		fmt.Println(e.Email)
	}
	if count == 0 {
		t.Error("Expected Kasem email =", expectEmail)
	}
	fmt.Println("TestQueryEmailByPersonFirstName = OK...")
}
