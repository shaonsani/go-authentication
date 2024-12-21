package main

import (
 "go-authentication/config"
 "go-authentication/controller"
 "go-authentication/router"
 "net/http"
 "time"
 "os"

 "github.com/go-playground/validator/v10"
)

func main() {
 // Database
 db := config.DatabaseConnection()

 validate := validator.New()

 // Controller
 authController := controller.NewAuthControllerImpl(db, validate)

 // Router
 routes := router.AuthRouter(authController)

 addr := os.Getenv("SERVER_ADDR") // e.g., ":8080" or "0.0.0.0:8080"
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