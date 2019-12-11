package model

import (
	"encoding/json"
	"fmt"

	"github.com/Chris-Ju/FISCO-BCOS-Demo/lib"
)

// AddAccount method
func AddAccount(account Account) bool {
	str := fmt.Sprintf("insertAccount %d %d %d %d", account.ID, account.ACompany.ID, account.BCompany.ID, account.Money)
	r := lib.SendtxMethod(str)
	if r == "0" {
		return false
	}
	return true
}

// DeleteAccount method
func DeleteAccount(id int32) bool {

	account := GetAccountByID(id)
	if account.ACompany.Money < account.Money {
		return false
	}
	str := fmt.Sprintf("removeAccount %d", id)
	lib.SendtxMethod(str)
	str = fmt.Sprintf("updateCompany %d \"%s\" %d %d", account.ACompany.ID, account.ACompany.Name, account.ACompany.Trusty, account.ACompany.Money-account.Money)
	lib.SendtxMethod(str)
	str = fmt.Sprintf("updateCompany %d \"%s\" %d %d", account.BCompany.ID, account.BCompany.Name, account.BCompany.Trusty, account.BCompany.Money+account.Money)
	lib.SendtxMethod(str)
	return true
}

// GetAccountByID method
func GetAccountByID(id int32) Account {
	accounts := GetAllAccounts()
	for _, account := range accounts {
		if account.ID == id {
			return account
		}
	}
	return Account{}
}

// GetAccountByCompanyID method
func GetAccountByCompanyID(id int32) []Account {
	accounts := GetAllAccounts()
	result := []Account{}
	for _, account := range accounts {
		if account.ACompany.ID == id || account.BCompany.ID == id {
			result = append(result, account)
		}
	}
	return result
}

// TransferAccount method
func TransferAccount(id int32, account Account) bool {
	accountOld := GetAccountByID(id)
	if accountOld.ID <= 0 {
		return false
	} else if accountOld.Money < account.Money {
		return false
	}
	if accountOld.Money == account.Money {
		DeleteAccountVip(id)
		return AddAccount(account)
	}
	account.ID = lib.GetRandom()
	AddAccount(account)
	str := fmt.Sprintf("updateAccount %d %d %d %d", id, accountOld.ACompany.ID, accountOld.BCompany.ID, accountOld.Money-account.Money)
	r := lib.SendtxMethod(str)
	if r == "0" {
		return false
	}
	return true

}

// DeleteAccountVip method
func DeleteAccountVip(id int32) bool {
	str := fmt.Sprintf("removeAccount %d", id)
	r := lib.SendtxMethod(str)
	if r == "0" {
		return false
	}
	return true
}

// GetAllAccounts method
func GetAllAccounts() []Account {
	companys := GetAllCompanys()
	b := lib.CallMethod("selectAccount")
	returnData := []AccountReturn{}
	result := []Account{}
	err := json.Unmarshal(b, &returnData)
	if err != nil {
		fmt.Print(err)
	}
	for i, account := range returnData {
		result = append(result, Account{
			ID:       account.ID,
			Money:    account.Money,
			ACompany: Company{ID: account.ACompany},
			BCompany: Company{ID: account.BCompany},
		})
		for _, company := range companys {
			if account.ACompany == company.ID {
				result[i].ACompany = company
			} else if account.BCompany == company.ID {
				result[i].BCompany = company
			}
		}
	}
	return result
}

// Financing method
func Financing(accountID int32, companyID int32) bool {
	account := GetAccountByID(accountID)
	if account.BCompany.ID != companyID {
		return false
	} else if account.ACompany.Trusty != 1 {
		return false
	}
	company := account.BCompany
	company.Money = company.Money + account.Money
	account.BCompany.ID = 1000
	account.BCompany.Name = "bank"
	account.BCompany.Trusty = 0
	account.BCompany.Money = 0
	DeleteAccountVip(accountID)
	AddAccount(account)
	return SetCompany(companyID, company)
}
