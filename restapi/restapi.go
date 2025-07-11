package restapi

import (
	"fmt"
	"net/http"
)

func DescribeAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The Community Shared Library!")
	fmt.Fprintf(w, "Below is a description of the libraries restful API.")
	fmt.Fprintf(w, "Coming soon...")
}
