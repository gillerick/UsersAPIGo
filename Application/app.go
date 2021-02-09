package main

import (
	"Users/Application/api"
	"log"
	"net/http"
	"os"
)

func main()  {
http.HandleFunc("api/users/", api.RegistrationHandler)
http.HandleFunc("api/user/", api.RegistrationHandler)

http.ListenAndServe(port(), nil)
log.Printf("Listening at port %s", port())

}


func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0{
		port = "8080"
	}
	return ":"+port
}
