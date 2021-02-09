package main

import (
	"log"
	"net/http"
)

func main()  {
http.HandleFunc("api/users/", RegistrationHandler)
http.HandleFunc("api/user/", RegistrationHandler)

http.ListenAndServe(port(), nil)
log.Printf("Listening at port %s", port())

}