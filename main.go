package main

import (
	"go-authentication/config"
	"go-authentication/controller"
	"go-authentication/router"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := config.DatabaseConnection()
	validate := validator.New()
	authController := controller.NewAuthControllerImpl(db, validate)

	routes := router.AuthRouter(authController)

	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = "0.0.0.0:8080" // Default
	}
	server := &http.Server{
		Addr:           addr,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
