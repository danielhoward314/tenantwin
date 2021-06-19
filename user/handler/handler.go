package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	js "github.com/danielhoward314/tenantwin/user/serializer/json"
	"github.com/danielhoward314/tenantwin/user/svc"
)

type userHandlerImpl struct {
	l           *log.Logger
	userService svc.UserService
}

func NewUserHandler(l *log.Logger, userService svc.UserService) svc.UserHandler {
	return &userHandlerImpl{
		l:           l,
		userService: userService,
	}
}

func (h *userHandlerImpl) serializer(contentType string) svc.UserSerializer {
	// leaves room to move to switch on content type header for different serializers
	return &js.User{}
}

func (h *userHandlerImpl) Signup(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	serializer := h.serializer(contentType)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userToCreate, err := serializer.Decode(reqBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = h.userService.Signup(userToCreate)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	resBody, err := serializer.Encode(userToCreate)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusCreated)
	w.Write(resBody)
}
