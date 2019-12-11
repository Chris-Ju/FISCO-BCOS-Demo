package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Chris-Ju/FISCO-BCOS-Demo/model"
)

// AdminLogin method
func AdminLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	username := r.URL.Query()["username"][0]
	password := r.URL.Query()["password"][0]
	if username == "admin" && password == "admin" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

// DeleteAccountByAdmin method
func DeleteAccountByAdmin(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query()["accountID"][0])
	flag := model.DeleteAccountVip(int32(id))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if flag {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// GetAccountsByAdmin method
func GetAccountsByAdmin(w http.ResponseWriter, r *http.Request) {
	accounts := model.GetAllAccounts()
	buf, _ := json.Marshal(accounts)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

// GetCompanysByAdmin method
func GetCompanysByAdmin(w http.ResponseWriter, r *http.Request) {
	buf, _ := json.Marshal(model.GetAllCompanys())
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

// SetCompany method
func SetCompany(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	name := r.PostFormValue("name")
	trusty, _ := strconv.Atoi(r.PostFormValue("trusty"))
	money, _ := strconv.Atoi(r.PostFormValue("money"))
	company := model.Company{ID: int32(id), Name: name, Money: int32(money), Trusty: int32(trusty)}
	flag := model.SetCompany(company.ID, company)
	if flag {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}
