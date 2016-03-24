package models

func MockUser() []*User {
	users := []*User{
		&User{Name: "tom", PersonID: 1},
		&User{Name: "jip", PersonID: 2},
		&User{Name: "tam", PersonID: 3},
		&User{Name: "noi", PersonID: 4},
	}
	return users
}
func MockPerson() []*Person {
	persons := []*Person{
		&Person{FirstName: "Kasem", LastName: "Arnontavilas"},
		&Person{FirstName: "Jiraporn", LastName: "Arnontavilas"},
		&Person{FirstName: "Tanan", LastName: "Arnontavilas"},
		&Person{FirstName: "Satit", LastName: "Chomwattana"},
	}
	return persons
}

func MockTitle() []*Title {
	titles := []*Title{
		&Title{TH: "ประธานกรรมการ", EN: "Board of Director"},
		&Title{TitleID: 1, TH: "รองประธานกรรมการ", EN: "Director"},
		&Title{TitleID: 2, TH: "กรรมการผู้จัดการ", EN: "Managing Director"},
		&Title{TitleID: 3, TH: "ผู้อำนวยการขาย", EN: "Sales Director"},
		&Title{TitleID: 3, TH: "ผู้อำนวยการบริหารสินค้า", EN: "Merchandise Director"},
		&Title{TitleID: 4, TH: "พนักงานขาย", EN: "Sales"},
	}
	return titles
}

func MockOrg() []*Org {
	orgs := []*Org{
		&Org{TH: "ผู้บริหาร", EN: "Management", Short: "MG"},
		&Org{OrgID: 1, TH: "ขาย", EN: "Sales", Short: "SA"},
		&Org{OrgID: 2, TH: "ขายปลีก", EN: "Retail", Short: "RT"},
		&Org{OrgID: 2, TH: "ขายส่ง", EN: "Wholesales", Short: "WS"},
		&Org{OrgID: 2, TH: "ออนไลน์", EN: "Online", Short: "OL"},
		&Org{OrgID: 1, TH: "สำนักงาน", EN: "Back Office", Short: "BO"},
		&Org{OrgID: 6, TH: "บุคคล", EN: "Human Resource", Short: "HR"},
		&Org{OrgID: 6, TH: "บัญชี", EN: "Accounting", Short: "HR"},
	}
	return orgs
}