package svc

import "net/http"

type UserHandler interface {
	Signup(http.ResponseWriter, *http.Request)
}
