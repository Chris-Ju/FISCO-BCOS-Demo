package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Chris-Ju/FISCO-BCOS-Demo/lib"
	"github.com/Chris-Ju/FISCO-BCOS-Demo/model"
)

// CreateUser Method
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	username := r.PostFormValue("username")
	password := lib.Md5(r.PostFormValue("password"))
	name := r.PostFormValue("name")
	user := model.User{}
	company := model.Company{}
	user.Username = username
	user.Password = password
	company.Name = name
	user.Company = company
	user = model.CreateUser(user)
	w.WriteHeader(http.StatusOK)
}

// QueryUser method
func QueryUser(w http.ResponseWriter, r *http.Request) {
	usernameA := r.URL.Query()["username"]
	passwordA := r.URL.Query()["password"]
	if len(usernameA) == 0 || len(passwordA) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username := usernameA[0]
	password := lib.Md5(passwordA[0])
	user := model.QueryUserLogin(username, password)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if user.Company.ID == -1 {
		w.WriteHeader(http.StatusForbidden)
	} else {
		w.WriteHeader(http.StatusOK)
		buf, _ := json.Marshal(user)
		w.Write(buf)
	}
}

// UpdateUser method
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
