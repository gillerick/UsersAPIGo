package main

import (
	"./api"
	"log"
	"net/http"
	"os"
)

func main(){
	//Local http server and routes
	http.HandleFunc("/api/users", api.RegistrationHandler)
	http.HandleFunc("/api/register/", api.RegistrationHandler)
	http.ListenAndServe(port(), nil)
	log.Printf("Defaulting to port %s", port())
}

func port() string{
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" +port
}