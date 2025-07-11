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

func requestMux() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// landing page will contain description of the full API
	myRouter.HandleFunc("/", restapi.DescribeAPI)

	// handle a request to list all collections in the library
	myRouter.HandleFunc("/collections", collections.ReturnAllCollections)
	http.ListenAndServe(":8080", nil)
}
