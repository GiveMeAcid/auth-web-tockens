package server

import (
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"net/http"
	"github.com/auth-web-tokens/controllers"
	"github.com/rs/cors"
	"log"
	"github.com/Andersen-soft/Solox/services/config"
	"runtime"
)

type ServerRouter struct {
	R *mux.Router
	N *negroni.Negroni
}

func CreateServerRouter() ServerRouter {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(controllers.NotFoundHandler)

	n := negroni.Classic()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		Debug:          true,
	})
	n.Use(c)

	recovery := negroni.Recovery{
		Logger:     log.New(config.LogFile, "\r\n", 0),
		PrintStack: false,
		StackAll:   false,
		StackSize:  1024 * 8,
	}

	recovery.ErrorHandlerFunc = logRecoveryError
	n.Use(recovery)

	n.Use(&negroni.Logger{
		ALogger: log.New(config.LogFile, "\r\n", 0),
	})

	n.UseHandler(r)

	sr := ServerRouter{
		R: r,
		N: n,
	}
	return sr
}

func logRecoveryError(err interface{}) {
	trace := make([]byte, 10024)
	count := runtime.Stack(trace, true)
	log.Printf("Error recoverd, stack trace, lines %d  trace: %s", count, string(trace))
}