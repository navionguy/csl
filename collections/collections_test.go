package collections

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// common test data
type common_test struct {
	name  string
	owner string
}

var common_tests []common_test

// initialize the common_tests array
func init() {
	common_tests = append(common_tests, common_test{name: "Computer Science", owner: "George P. Burdell"})
	common_tests = append(common_tests, common_test{name: "Computer Science", owner: "Steve Gibson"})
	common_tests = append(common_tests, common_test{name: "Software Verification", owner: ""})
	common_tests = append(common_tests, common_test{name: "Science Fiction, Space-Futuristic", owner: "Steve Gibson"})
}

// some helper functions for testing

func clear_db() {
	for _, tt := range common_tests {
		coll := Collection{Name: tt.name, Owner: tt.owner}
		resp := httptest.NewRecorder()
		deleteCollection(resp, coll)
	}
}

func load_db() {
	// clear db to get to known state
	clear_db()

	for _, tt := range common_tests {
		coll := Collection{Name: tt.name, Owner: tt.owner}
		resp := httptest.NewRecorder()
		addCollectionData(resp, coll)
	}

}

func TestAddCollection(t *testing.T) {
	// first, delete any collections in the database
	// yeah, don't run this on a live database
	clear_db()

	for _, tt := range common_tests {
		coll := Collection{Name: tt.name, Owner: tt.owner}
		resp := httptest.NewRecorder()
		addCollectionData(resp, coll)

		assert.Equal(t, 200, resp.Code)
	}
}

func TestPostCollection(t *testing.T) {
	clear_db()

	for _, tt := range common_tests {
		col, _ := json.Marshal(Collection{Name: tt.name, Owner: tt.owner})
		req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/collection", strings.NewReader(string(col)))

		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		PostNewCollection(resp, req)

		assert.Equal(t, resp.Code, 200)
	}
}

// Delete collections via direct SQL call
func TestDeleteCollection(t *testing.T) {
	load_db()

	for _, tt := range common_tests {
		coll := Collection{Name: tt.name, Owner: tt.owner}
		resp := httptest.NewRecorder()
		deleteCollection(resp, coll)

		assert.Equal(t, 200, resp.Code)
	}

}

// delete collections via HTTP call
func TestRemoveCollection(t *testing.T) {
	clear_db()

	for _, tt := range common_tests {
		col, _ := json.Marshal(Collection{Name: tt.name, Owner: tt.owner})
		req := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/collection", strings.NewReader(string(col)))

		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		PostNewCollection(resp, req)

		assert.Equal(t, resp.Code, 200)
	}
}
