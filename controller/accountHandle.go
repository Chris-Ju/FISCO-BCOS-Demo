package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Chris-Ju/FISCO-BCOS-Demo/lib"
	"github.com/Chris-Ju/FISCO-BCOS-Demo/model"
)

// AddAccount method
func AddAccount(w http.ResponseWriter, r *http.Request) {
	randomID := lib.GetRandom()
	money, _ := strconv.Atoi(r.PostFormValue("money"))
	ACompany, _ := strconv.Atoi(r.PostFormValue("ACompany"))
	BCompany, _ := strconv.Atoi(r.PostFormValue("BCompany"))

	account := model.Account{
		ID:       randomID,
		Money:    int32(money),
		ACompany: model.Company{ID: int32(ACompany)},
		BCompany: model.Company{ID: int32(BCompany)},
	}

	flag := model.AddAccount(account)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if flag {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

// DeleteAccount method, Pay
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	id := lib.GetID(r.URL.Path)
	flag := model.DeleteAccount(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if flag {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// GetAccountByID method
func GetAccountByID(w http.ResponseWriter, r *http.Request) {
	id := lib.GetID(r.URL.Path)
	account := model.GetAccountByID(id)
	buf, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

// GetAllAccount method
func GetAllAccount(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query()["id"][0])
	accounts := model.GetAccountByCompanyID(int32(id))
	buf, _ := json.Marshal(accounts)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

// TransferAccount method
func TransferAccount(w http.ResponseWriter, r *http.Request) {
	id := lib.GetID(r.URL.Path)
	money, _ := strconv.Atoi(r.PostFormValue("money"))
	ACompany, _ := strconv.Atoi(r.PostFormValue("ACompany"))
	BCompany, _ := strconv.Atoi(r.PostFormValue("BCompany"))

	account := model.Account{}
	account.ID = int32(id)
	account.Money = int32(money)
	account.ACompany.ID = int32(ACompany)
	account.BCompany.ID = int32(BCompany)

	flag := model.TransferAccount(account.ID, account)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if flag {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// Financing method
func Financing(w http.ResponseWriter, r *http.Request) {
	accountID := lib.GetID(r.URL.Path)
	companyID, _ := strconv.Atoi(r.URL.Query()["companyID"][0])
	flag := model.Financing(int32(accountID), int32(companyID))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if flag {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
