package model

type User struct {
	Username string  `json:"username,omitempty"`
	Password string  `json:"password,omitempty"`
	Company  Company `json:"company,omitempty"`
}

type Company struct {
	Id     int32  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Trusty bool   `json:"trusty,omitempty"`
	Money  int32  `json:"money,omitempty"`
}

type Account struct {
	Id       int32   `json:"id,omitempty"`
	Money    int32   `json:"money,omitempty"`
	ACompany Company `json:"ACompany,omitempty"`
	BCompany Company `json:"BCompany,omitempty"`
}

type Admin struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
