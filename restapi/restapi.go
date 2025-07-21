package restapi

import (
	"fmt"
	"net/http"
)

func DescribeAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The Community Shared Library!\n")
	fmt.Fprintf(w, "A public place for sharing the books you love!!\n\n")
	fmt.Fprintf(w, "Coming soon...")
}
