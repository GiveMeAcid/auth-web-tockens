package controllers

import (
	"net/http"
	"strings"
	"github.com/auth-web-tokens/models/requests"
	"github.com/auth-web-tokens/services"
	"encoding/json"
	"github.com/auth-web-tokens/models"
)

// authenticate
func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(requests.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	dbUser := new(models.User)
	responseStatus, token := services.Login(requestUser, dbUser)
}

// returns users list
func (*MainController) GetUsersAction(w http.ResponseWriter, r *http.Request) {

	var isAuthorized, user = checkAuth(r)

	if (!isAuthorized) {
		services.ToJSON(w, services.MakeErrorResponse("You are not authorized"), http.StatusUnauthorized)
		return
	}

	// only admins can see users
	if (user.Role != models.USER_ROLE_ADMIN) {
		services.ToJSON(w, services.MakeErrorResponse("Access denied"), http.StatusForbidden)
		return
	}

	// makes empty array, uses to store users from database
	users := make([]*models.User, 0)

	services.DB.First(&users).Debug()

	services.ToJSON(w, users, http.StatusOK)
}

// register a new user
func (*MainController) PostRegisterAction(w http.ResponseWriter, r *http.Request) {

	var (
		email = r.PostFormValue("email")
		password = r.PostFormValue("password")
		confirmPassword = r.PostFormValue("confirm_password")
	)

	// validate entered data
	if (len(email) == 0 || len(password) == 0 || len(confirmPassword) == 0) {
		services.ToJSON(w, services.MakeErrorResponse("Fill all required fields"), http.StatusBadRequest)
		return
	}

	if (!services.IsEmailValid(email)) {
		services.ToJSON(w, services.MakeErrorResponse("Email is invalid"), http.StatusBadRequest)
		return
	}

	if ((strings.Compare(password, confirmPassword)) != 0) {
		services.ToJSON(w, services.MakeErrorResponse("Fields 'password' and 'confirm_password' must be the same"), http.StatusBadRequest)
		return
	}

	user := models.User{Email:email}

	// check is user with entered email is not exists in the database
	services.DB.Where(user).First(&user)

	if (user.Id != 0) {
		services.ToJSON(w, services.MakeErrorResponse("User with entered email is exists"), http.StatusBadRequest)
		return;
	}

	user.SetPassword(password)
	user.Role = models.USER_ROLE_USER

	// insert new user
	services.DB.Save(&user)

	// returns created user
	services.ToJSON(w, user, http.StatusOK)
}
