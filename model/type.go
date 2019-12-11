package model

// User define
type User struct {
	Username string  `json:"username,omitempty"`
	Password string  `json:"password,omitempty"`
	Company  Company `json:"company,omitempty"`
}

// Company define
type Company struct {
	ID     int32  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Trusty int32  `json:"trusty,omitempty"`
	Money  int32  `json:"money,omitempty"`
}

// CompanyReturn define
type CompanyReturn struct {
	ID     int32  `json:"a,omitempty"`
	Name   string `json:"b,omitempty"`
	Trusty int32  `json:"c,omitempty"`
	Money  int32  `json:"d,omitempty"`
}

// Account define
type Account struct {
	ID       int32   `json:"id,omitempty"`
	Money    int32   `json:"money,omitempty"`
	ACompany Company `json:"ACompany,omitempty"`
	BCompany Company `json:"BCompany,omitempty"`
}

// AccountReturn define
type AccountReturn struct {
	ID       int32 `json:"a,omitempty"`
	ACompany int32 `json:"b,omitempty"`
	BCompany int32 `json:"c,omitempty"`
	Money    int32 `json:"d,omitempty"`
}

// Admin define
type Admin struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// DATAPATH define
const DATAPATH string = "data/db.db"
const groupID uint = 1
const URL string = "http://localhost:8545"
