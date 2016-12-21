package controllers

import (
	"net/http"
	"strings"
	"github.com/auth-web-tokens/models/requests"
	"encoding/json"
	"github.com/auth-web-tokens/models"
	"github.com/auth-web-tokens/services/auth"
	"github.com/gorilla/context"
)

// authenticate
func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(requests.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	if strings.EqualFold(requestUser.Email, "") ||
		strings.EqualFold(requestUser.Password, "") ||
		strings.EqualFold(requestUser.UUID, "") == false {
		MakeResponseFail(w, http.StatusBadRequest, "Login is not valid")
		return
	}

	dbUser := new(models.User)

	responseStatus, token := auth.Login(requestUser, dbUser)

	w.WriteHeader(responseStatus)
	if responseStatus == 401 {
		MakeResponseFail(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	w.Header().Set("Content-Type", "json/application")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(requests.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	currentUser, ok := context.GetOk(r, "currentUser")
	if !ok {
		MakeResponseFail(w, http.StatusNotFound, requests.UserNotFound.Error())
		return
	}

	user := currentUser.(*models.User)

	w.Header().Set("Content-Type", "application/json")
	w.Write(auth.RefreshToken(requestUser, user))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := auth.Logout(r)
	if err != nil {
		MakeResponseFail(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

//func MakeResponseSuccess(w http.ResponseWriter, data interface{}) {
//	js, err := json.Marshal(data)
//	if err != nil {
//		log.Printf("[http] error encodind data %s    %+v", err, data)
//	}
//
//	w.WriteHeader(http.StatusOK)
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(js)
//}

func MakeResponseFail(w http.ResponseWriter, status int, message string) {
	js, _ := json.Marshal(map[string]string{"error": message})
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}