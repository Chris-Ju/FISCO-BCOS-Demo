package router

import (
	"net/http"
	"strings"

	"github.com/Chris-Ju/FISCO-BCOS-Demo/controller"
	"github.com/Chris-Ju/FISCO-BCOS-Demo/lib"
	"github.com/gorilla/mux"
)

// Route define
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes type define
type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = lib.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{

	Route{
		"AddAccount",
		strings.ToUpper("Post"),
		"/api/account/",
		controller.AddAccount,
	},

	Route{
		"DeleteAccount",
		strings.ToUpper("Delete"),
		"/api/account/{accountID}",
		controller.DeleteAccount,
	},

	Route{
		"GetAccountByID",
		strings.ToUpper("Get"),
		"/api/account/{accountID}",
		controller.GetAccountByID,
	},

	Route{
		"GetAllArticle",
		strings.ToUpper("Get"),
		"/api/account/",
		controller.GetAllArticle,
	},

	Route{
		"TransferArticle",
		strings.ToUpper("Put"),
		"/api/account/{accountID}",
		controller.TransferArticle,
	},

	Route{
		"AdminLogin",
		strings.ToUpper("Post"),
		"/admin/login",
		controller.AdminLogin,
	},

	Route{
		"DeleteAccountByAdmin",
		strings.ToUpper("Delete"),
		"/admin/api/account/",
		controller.DeleteAccountByAdmin,
	},

	Route{
		"GetAccountsByAdmin",
		strings.ToUpper("Get"),
		"/admin/api/account/",
		controller.GetAccountsByAdmin,
	},

	Route{
		"GetCompanysByAdmin",
		strings.ToUpper("Get"),
		"/admin/api/company/",
		controller.GetCompanysByAdmin,
	},

	Route{
		"SetCompany",
		strings.ToUpper("Post"),
		"/admin/api/company/",
		controller.SetCompany,
	},

	Route{
		"CreateUser",
		strings.ToUpper("Post"),
		"/api/user/",
		controller.CreateUser,
	},

	Route{
		"QueryUser",
		strings.ToUpper("Get"),
		"/api/user/",
		controller.QueryUser,
	},

	Route{
		"UpdateUser",
		strings.ToUpper("Put"),
		"/api/user/",
		controller.UpdateUser,
	},
}
