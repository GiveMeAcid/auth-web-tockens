package main

import (

)
import (
	"os"
	"github.com/auth-web-tokens/settings"
	"log"
	"github.com/auth-web-tokens/services"
	"github.com/auth-web-tokens/server"
	"github.com/auth-web-tokens/models"
	"github.com/auth-web-tokens/services/config"
	"syscall"
	"runtime"
	"fmt"
	"os/signal"
)

func main() {
	if _, err := os.Stat(config.Config.PrivateKeyPath); err != nil {
		panic("Must specify private key")
	}
	if _, err := os.Stat(settings.Get().PublicKeyPath); err != nil {
		panic("Must specify public key")
	}

	log.Println("Initializing DB...")
	err := services.InitDB()
	if err != nil {
		panic(err)
	}
	log.Println("Migration DB...")
	models.Migrations(services.DB)

	HandleSignals()

	log.Println("Starting server...")
	parsedPort := config.Config.BasePort
	s := server.New(":" + parsedPort)

	if err := s.ListenAndServe(); err != nil {
		log.Println("Error " + err.Error())
	}
}

func HandleSignals() {
	sigChan := make(chan os.Signal)
	go func() {
		stacktrace := make([]byte, 1<<20)
		for sig := range sigChan {
			switch sig {
			case syscall.SIGQUIT:
				length := runtime.Stack(stacktrace, true)
				fmt.Println(string(stacktrace[:length]))
			case syscall.SIGINT:
				fallthrough
			case syscall.SIGTERM:
				fmt.Println("Shutting down...")
				os.Exit(0)
			}
		}
	}()
	signal.Notify(sigChan, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
}