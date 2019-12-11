package model

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Chris-Ju/FISCO-BCOS-Demo/lib"
	"github.com/boltdb/bolt"
)

// CreateUser mothed
func CreateUser(user User) User {
	db, err := bolt.Open(DATAPATH, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		id, _ := b.NextSequence()
		user.Company.ID = int32(id)
		user.Company.Trusty = 0
		user.Company.Money = 0
		buf, err := json.Marshal(user)
		if err != nil {
			return err
		}
		return b.Put([]byte(user.Username), buf)
	})

	AddCompany(user.Company)

	return user
}

//QueryUserLogin Method
func QueryUserLogin(username string, password string) User {
	db, err := bolt.Open(DATAPATH, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	result := false
	user := User{}
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			err = json.Unmarshal(v, &user)
			if user.Username == username && user.Password == password {
				result = true
				break
			}
		}
		return nil
	})

	if !result {
		user.Company.ID = -1
	}
	return user
}

// GetAllCompanys method
func GetAllCompanys() []Company {
	b := lib.CallMethod("selectCompany")
	returnData := []CompanyReturn{}
	result := []Company{}

	err := json.Unmarshal(b, &returnData)
	if err != nil {
		fmt.Print(err)
	}
	for _, company := range returnData {
		result = append(result, Company{
			ID:     company.ID,
			Name:   company.Name,
			Trusty: company.Trusty,
			Money:  company.Money,
		})
	}
	return result
}

// AddCompany method
func AddCompany(company Company) bool {
	str := fmt.Sprintf("insertCompany %d \"%s\" %d %d", company.ID, company.Name, company.Trusty, company.Money)
	r := lib.SendtxMethod(str)
	if r == "0" {
		return false
	}
	return true
}

// SetCompany method
func SetCompany(id int32, company Company) bool {
	str := fmt.Sprintf("updateCompany %d \"%s\" %d %d", id, company.Name, company.Trusty, company.Money)
	r := lib.SendtxMethod(str)
	if r == "0" {
		return false
	}
	return true
}
