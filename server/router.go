package server

import (
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"github.com/auth-web-tokens/controllers"
	"github.com/auth-web-tokens/repositories"
)

type Server struct {
	http.Server
}

func New(bindAddr string) *Server {
	sr := CreateServerRouter()

	s := &Server{
		Server: http.Server{
			Handler:        sr.N,
			Addr:           bindAddr,
			WriteTimeout:   15 * time.Second,
			ReadTimeout:    15 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}

	s.handleFunc(sr.R, "/login", controllers.Login).Methods("POST")
	s.handlePrivateFunc(sr.R, "/refresh_token", repositories.RequireTokenAuthentication, controllers.RefreshToken).Methods("GET")
	s.handlePrivateFunc(sr.R, "/logout", repositories.RequireTokenAuthentication, controllers.Logout).Methods("POST")

	return s
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func (s *Server) handleFunc(router *mux.Router, route string, fn HandlerFunc) *mux.Route {
	return router.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fn(w, r)
	})
}

func (s *Server) handlePrivateFunc(router *mux.Router, route string, required negroni.HandlerFunc, fn negroni.HandlerFunc) *mux.Route {
	return router.Handle(route,
		negroni.New(
			negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
				rw.Header().Set("Content-Type", "application/json")
				next(rw, r)
			}),
			negroni.HandlerFunc(required),
			negroni.HandlerFunc(fn),
		))
}