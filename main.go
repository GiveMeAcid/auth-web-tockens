package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/auth-web-tokens/services"
	"github.com/auth-web-tokens/models"
	"github.com/auth-web-tokens/controllers"
	"fmt"
)

func init() {
	db()
}

func main() {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	mainController := controllers.MainController{}
	router.HandleFunc("/login", mainController.PostLoginAction).Methods("POST")
	router.HandleFunc("/register", mainController.PostRegisterAction).Methods("POST")
	router.HandleFunc("/users", mainController.GetUsersAction).Methods("GET")
	log.Fatal(http.ListenAndServe(":3030", router))
}

func db() {
	fmt.Println("****************")
	fmt.Println("Database connection...")
	services.InitDB()
	fmt.Println("****************")
	fmt.Println("Run migrations...")
	fmt.Println("****************")
	models.Migrations()
	fmt.Println("Run fixtures...")
	models.Fixtures()
	fmt.Println("****************")
}

// error handler if page is not found
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	services.ToJSON(w, services.MakeErrorResponse("bad request"), http.StatusNotFound)
}

//error handler if requested method is not correct
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	services.ToJSON(w, services.MakeErrorResponse("Method '" + r.Method + "' not allowed"), http.StatusMethodNotAllowed)
}