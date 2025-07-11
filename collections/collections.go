package collections

import (
	"encoding/json"
	"net/http"
)

type Collection struct {
	Name  string `json:"Name"`
	Owner string `json:"Owner"`
	UUID  string `json:"UUID"`
}

var Collections = []Collection{
	{Name: "Computer Science", Owner: "George P. Burdell", UUID: "ad6cfdf7-d679-4ed6-8928-d18bc2d0f8ea"},
	{Name: "Software Verification", Owner: "", UUID: "11ee9e0f-6498-4baa-b046-49d01765496f"},
	{Name: "Science Fiction, Space-Futuristic", Owner: "Steve Gibson", UUID: "234491d5-2c96-402c-834e-bcec82bdf3b9"},
}

func ReturnAllCollections(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Collections)
}
