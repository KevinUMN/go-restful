package main

import (
	"log"
	"net/http"
	
	"github.com/KevinUMN/go-restful"
)

// This example shows how to use the OPTIONSFilter on a Container
//
// OPTIONS http://localhost:8080/users
//
// OPTIONS http://localhost:8080/users/1

func main() {
	wsContainer := restful.NewContainer()
	u := UserResource{}
	u.RegisterTo(wsContainer)
	
	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)
	
	// For use on the default container, you can write
	// restful.Filter(restful.OPTIONSFilter())
	
	log.Print("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
