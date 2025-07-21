# Community Shared Library - Restful Interface

## Collection Management

### List all Collections

You can request a list of all, or some of the collections in the system.  
To do this you can provide the name of the collection and or owner/holder of the collection.  
For both values, you can use regular expressions to search for multiple collections.  
If both parameters are blank or missing.  You will get the full list of collections.

Below is an example of curl command to search for all collections with the name ***Computer Science***.

    curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"Name":"Computer Science","Owner":""}' \
    http://localhost:8080/collection  

### Creating a new Collection

The new collection must have name for the collection can have a named Owner.
Both of which should be strings.

Below is an example of curl command to create a new collection.

    curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"Name":"Spy Novels","Owner":"Vince Flynn"}' \
    http://localhost:8080/collection  

If a collection with that name and owner already exists, the request will return with a status of ***StatusConflict***.  

### Deleting an existing Collection  

Deleting a collection requires the collection name and owner, if there is one.  
Any books that are in the passed collection will have their entries deleted as well.  

Below is an example curl command to delete a collection.

    curl --header "Content-Type: application/json" \
    --request DELETE \
    --data '{"Name":"Spy Novels","Owner":"Vince Flynn"}' \
    http://localhost:8080/collection

If the no collection can be found with the parameters given, the request will return a status of ***StatusBadRequest***.  

### Modifying an exiting Collection

If you need to change the name or owner of a collection, you need to supply the current name and owner, then either  
the new collection name, the new owner name, or both.  
All of the books in the collection, will be remain in the updated collection.  

Below is an example curl command to modify a collection.

    curl --header "Content-Type: application/json" \
    --request PATCH \
    --data '{"Name":"Spy Novels","Owner":"Vince Flynn", "NewName": "War on Terror", "NewOwner": "Tom Clancy"}' \
    http://localhost:8080/collection

If the no collection can be found with the parameters given, the request will return a status of ***StatusBadRequest***.  

If a collection with the new name and new owner already exists, the request will return with a status of ***StatusConflict***.  

## Book Management

### Adding a book to a Collection

To add a book to a collection, the following values are required:  

- CollectionName
- CollectionOwner
- Title
- Author
- PublicationDate

Optional values are:  

- Edition
- Description
- Genre
- Notes

Below is an example curl command to add a book to a collection.

    curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"CollectionName":"Great Leaders","CollectionOwner":"Don Hinds",\
    "Title": "King of Kings", "Author": "Asfa-Wossen Asserate",\
    "PublicationDate": "2017", "Notes": "English edition"}' \
    http://localhost:8080/books  

### Removing a book from a Collection or Collections

To remove a book from a collection or from the entire library:  

All values are optional and can be expressed as regular expressions.  
Caution!  Use of regular expressions might remove unintended books.  
Best practice is to search first and adjust you search until it returns  
only the books you want.

- CollectionName
- CollectionOwner
- Title
- Author
- PublicationDate
- Edition
- Description
- Genre
- Notes

Below is an example curl command to add a book to a collection.

    curl --header "Content-Type: application/json" \
    --request DELETE \
    --data '{"CollectionName":"Historical Non-Fiction","CollectionOwner":"Don Hinds",\
    "Title": "Lee's Cavalrymen", "Author": "Edward G. Longacre"}\
    http://localhost:8080/books  

### Modifying a book in the library

All values are optional and can be expressed as regular expressions.  

- CollectionName
- CollectionOwner
- Title
- Author
- PublicationDate
- Edition
- Description
- Genre
- Notes

You can optionally change any value by passing the new value using:

- NewColName
- NewColOwner
- NewTitle
- NewAuthor
- NewPubDate
- NewEdition
- NewDescription
- NewGenre
- NewNotes

Below is an example curl command to modify a book in the library.

    curl --header "Content-Type: application/json" \
    --request PATCH \
    --data '{"CollectionName":"Great Leaders","CollectionOwner":"Don Hinds",\
    "Title": "King of Kings", "Author": "Asfa-Wossen Asserate",\
    "NewPubDate": "2014", "NewNotes": "German edition"}' \
    http://localhost:8080/books  


### Search for a book in the Library

All values are optional and can be expressed as regular expressions.   

- CollectionName
- CollectionOwner
- Title
- Author
- PublicationDate
- Edition
- Description
- Genre
- Notes

Below is an example curl command to search for a book in the library.

    curl --header "Content-Type: application/json" \
    --request GET \
    --data '{"CollectionOwner":"Don Hinds","Title": "Emerson's Essays", "Author": "Ralph Waldo Emerson",\
    "PublicationDate": "1841" http://localhost:8080/books  
