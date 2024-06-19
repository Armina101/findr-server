package internals

import "net/http"

type FindrAppService interface{

	FindrHomePage() http.HandlerFunc
	
	Register() http.HandlerFunc
	Login() http.HandlerFunc
	Logout() http.HandlerFunc
}

