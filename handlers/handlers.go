package handlers

import (
	"net/http"
	"fmt"
	"github.com/auth-web-tokens/models"
	"encoding/json"
	"github.com/gorilla/mux"
)

var users models.Users

func Status(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Route is working...")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(r.Body).Encode(&user)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range users {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(&models.User{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		//DB *gorm.DB
		user models.User
		//password = r.PostFormValue("password")
	)
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(user)
	//DB.Where(user).First(&user)
	//user.SetPassword(password)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range users {
		if item.Id == params["id"] {
			users = append(users[:index], users[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}