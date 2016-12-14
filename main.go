package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/auth-web-tokens/handlers"
	"github.com/auth-web-tokens/services"
)
func init() {
	services.InitDB()
}

func main() {
	router := mux.NewRouter()

	//router.Method = true; // sets processing of incorrect request methods
	//router.MethodNotAllowed = http.HandlerFunc(methodNotAllowedHandler);
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	router.HandleFunc("/", handlers.Status).Methods("GET")
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3030", router))
}

// error handler if page is not found
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	services.ToJSON(w, services.MakeErrorResponse("Requested page in not found"), http.StatusNotFound)
}

// error handler if requested method is not correct
//func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
//	services.ToJSON(w, services.MakeErrorResponse("Method '" + r.Method + "' not allowed"), http.StatusMethodNotAllowed)
//}