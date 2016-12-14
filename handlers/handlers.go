package handlers

import (
	"net/http"
	"fmt"
	"github.com/auth-web-tokens/models"
	"encoding/json"
)

func Status(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Route is working...")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}