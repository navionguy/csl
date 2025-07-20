package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/navionguy/csl/collections"
	"github.com/navionguy/csl/restapi"
)

func main() {
	fmt.Printf("Hello World!")

	requestMux()
}

// requestMux() -- creates a router and adds the needed routes
func requestMux() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// landing page will contain description of the full API
	myRouter.HandleFunc("/", restapi.DescribeAPI)

	// remove a collection
	myRouter.HandleFunc("/collection", collections.RemoveCollection).Methods(http.MethodDelete)
	// handle creation of a new collection
	myRouter.HandleFunc("/collection", collections.PostNewCollection)

	listenAndServe(myRouter)
}

// listenAndServe -- does, well, that...
func listenAndServe(router *mux.Router) {
	http.ListenAndServe(":8080", router)
}

// command line test for routes

/*
curl -i -X GET http://localhost:8080/
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"Name":"Spy Novels","Owner":"Vince Flynn"}' \
  http://localhost:8080/collection
curl --header "Content-Type: application/json" \
  --request DELETE \
  --data '{"Name":"Spy Novels","Owner":"Vince Flynn"}' \
  http://localhost:8080/collection
*/
