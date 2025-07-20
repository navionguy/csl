package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Connect(t *testing.T) {
	db := Connect()
	defer db.data.Close()

	assert.NotNil(t, db, "Didn't get back a CSL object")
}

func Test_Insert_and_Delete(t *testing.T) {
	db := Connect()

	err := db.InsertCollection("Spy Novels", "Vince Flynn")
	assert.Nil(t, err, "Insert failed")

	err = db.DeleteCollection("Spy Novels", "Vince Flynn")
	assert.Nil(t, err, "Delete failed")
}
