package models

import "time"

//func MockUser() []*User {
//	users := []*User{
//		&User{Name: "tom",PersonID: 1},
//		&User{Name: "jip", PersonID: 2},
//		&User{Name: "tam", PersonID: 3},
//		&User{Name: "noi", PersonID: 4},
//	}
//	return users
//}
func MockPerson() []*Person {
	persons := []*Person{
		&Person{
			Users: []User{
				{Name: "tom"},
				{Name:"tom2"},
			},
			Jobs: []Job{
				{JobID: 1, OrgID: 1, NameTH: "กรรมการผู้จัดการ"},
				{JobID: 1, OrgID: 2, NameTH: "ผู้อำนวยการขาย"},
			},
			Emails: []Email{
				{Email: "tom@nopadol.com"},
				{Email: "mrtomyum@gmail.com"},
			},
			FirstName: "Kasem",
			LastName: "Arnontavilas",
			BirthDate: time.Date(1974,10,4,0,0,0,0,time.UTC),
			CitizenID:"3509901377838",
		},

		&Person{
			Users: []User{{Name: "jip"}},
			Jobs: []Job{
				{JobID: 1, OrgID: 1, NameTH: "ผู้อำนวยการฝ่ายขาย", NameEN: "Sales Director"},
				{JobID: 1, OrgID: 3, NameTH: "ผู้อำนวยการฝ่ายบริหารสินค้า", NameEN: "Merchandise Director"},
			},
			Emails: []Email{
				{Email: "jip@nopadol.com"},
				{Email: "jipjiraporn@gmail.com"},
			},
			FirstName: "Jiraporn",
			LastName: "Arnontavilas",
			BirthDate: time.Date(1976,8,12,0,0,0,0,time.UTC),
		},

		&Person{
			Users: []User{{Name: "tam"}},
			Jobs: []Job{
				{NameTH: "ผู้อำนวยการสาขา2", NameEN: "Sales Director Branch 2"},
			},
			FirstName: "Tanan",
			LastName: "Arnontavilas",
		},
		&Person{
			Users: []User{{Name: "noi"}},
			Jobs: []Job{
				{NameTH: "ผู้จัดการฝ่าย IT"},
			},
			FirstName: "Satit",
			LastName: "Chomwattana",
		},
	}
	return persons
}

//func MockJob() []*Job {
//	jobs := []*Job{
//		&Job{NameTH: "ประธานกรรมการ", NameEN: "Board of Director"},
//		&Job{JobID: 1, NameTH: "รองประธานกรรมการ", NameEN: "Director"},
//		&Job{JobID: 2, NameTH: "กรรมการผู้จัดการ", NameEN: "Managing Director"},
//		&Job{JobID: 3, NameTH: "ผู้อำนวยการขาย", NameEN: "Sales Director"},
//		&Job{JobID: 3, NameTH: "ผู้อำนวยการบริหารสินค้า", NameEN: "Merchandise Director"},
//		&Job{JobID: 4, NameTH: "พนักงานขาย", NameEN: "Sales"},
//	}
//	return jobs
//}

func MockOrg() []*Org {
	orgs := []*Org{
		&Org{NameTH: "ผู้บริหาร", NameEN: "Management", NameShort: "MG"},
		&Org{OrgID: 1, NameTH: "ขาย", NameEN: "Sales", NameShort: "SA"},
		&Org{OrgID: 2, NameTH: "ขายปลีก", NameEN: "Retail", NameShort: "RT"},
		&Org{OrgID: 2, NameTH: "ขายส่ง", NameEN: "Wholesales", NameShort: "WS"},
		&Org{OrgID: 2, NameTH: "ออนไลน์", NameEN: "Online", NameShort: "OL"},
		&Org{OrgID: 1, NameTH: "สำนักงาน", NameEN: "Back Office", NameShort: "BO"},
		&Org{OrgID: 6, NameTH: "บุคคล", NameEN: "Human Resource", NameShort: "HR"},
		&Org{OrgID: 6, NameTH: "บัญชี", NameEN: "Accounting", NameShort: "HR"},
	}
	return orgs
}
//func MockEmail() []*Email {
//	emails := []*Email{
//		&Email{Email:"tom@nopadol.com", PersonID:1},
//		&Email{Email:"jip@nopadol.com", PersonID:2},
//		&Email{Email:"tam@nopadol.com", PersonID:3},
//		&Email{Email:"mrtomyum@gmail.com", PersonID:1},
//	}
//	return emails
//}