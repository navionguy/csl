package collections

import (
	//"database"
	"encoding/json"
	"io"
	"net/http"

	"github.com/navionguy/csl/database"
)

type Collection struct {
	Name  string `json:"Name"`
	Owner string `json:"Owner"`
	UUID  string `json:"UUID"`
}

// PostNewCollection -- Create a new, empty collection
// first up, try to retrieve the request body
func PostNewCollection(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	coll, err := extractCollectionData(w, body)

	if err != nil {
		return
	}

	addCollectionData(w, coll)
}

// try to add the collection to the library
func addCollectionData(w http.ResponseWriter, collection Collection) {
	csl := database.Connect()
	defer csl.Disconnect()

	err := csl.InsertCollection(collection.Name, collection.Owner)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func RemoveCollection(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	coll, err := extractCollectionData(w, body)

	if err != nil {
		return
	}

	deleteCollection(w, coll)
}

func deleteCollection(w http.ResponseWriter, collection Collection) {
	csl := database.Connect()
	defer csl.Disconnect()

	err := csl.DeleteCollection(collection.Name, collection.Owner)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// common routines

// see if the request body is recognizable json
func extractCollectionData(w http.ResponseWriter, body []byte) (Collection, error) {
	var collection Collection
	err := json.Unmarshal(body, &collection)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	return collection, err
}
