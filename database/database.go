package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type CommunitySharedLibrary struct {
	data *sql.DB
}

/*
	 Connect -- Opens the connection to the database for the service to operate
		and then tests the connection by pinging the server.
*/
func Connect() *CommunitySharedLibrary {
	var err error
	csl := &CommunitySharedLibrary{}

	// all parameters are hard coded for clarity
	// ToDo: push operational parameters into a config file somewhere
	csl.data, err = sql.Open("postgres", "host=localhost port=5432 user=navionguy password=Lady5239 dbname=commlib")

	// If the open failed, hard stop
	if err != nil {
		panic(err)
	}

	err = csl.data.Ping()
	if err != nil {
		panic(err)
	}

	return csl
}

// Implementation of the Community Shared library DB

func (csl *CommunitySharedLibrary) Disconnect() {
	csl.data.Close()
}

// Add a collection to the library
func (csl *CommunitySharedLibrary) InsertCollection(name string, owner string) error {
	stmt := `INSERT INTO collections (name, owner) VALUES($1, $2)`

	_, err := csl.data.Exec(stmt, name, owner)

	return err
}

// Delete an entire collection from the library
func (csl *CommunitySharedLibrary) DeleteCollection(name string, owner string) error {
	// first I want to remove all the books in the collection
	stmt := `DELETE FROM books USING collections where collection IN (
		SELECT collection_id FROM collections WHERE name = $1 AND owner = $2)`

	_, err := csl.data.Exec(stmt, name, owner)

	if err != nil {
		return err
	}

	// then the collection itself
	stmt = `DELETE FROM collections WHERE name = $1 AND owner = $2`

	_, err = csl.data.Exec(stmt, name, owner)

	return err
}
