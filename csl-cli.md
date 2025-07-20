# Community Shared Library - Command Line Interface

The Community Shared Library allows any group of people to maintain a catalog off all the books available within the community.  Most users will access the shared library using the web UI, but library administrators can use the CLI described here for automating many tasks, bulk loading of new book collections for example.

## Book Management

### csl-add-book - Add a book to a collection

csl-add-book -c *collection* [-h *holder*] -t *title* -a *author* -p *publish date* [-e *edition*] [-d *description*] [-g *genre*] [-n *notes*]

>collection and the optional *holder* identifies who posses the physical book.  If the values specified don't exist in the collections list, a new collection will be created.  
title, author, and published date are all required.  
edition, description, genre and notes are all option fields.

### csl-mod-book - Modify one or more fields for a books entry

csl-mod-book -c *collection* [-h *holder*] -t *title* -a *author*
-nc *new collection* [-nh *new holder* -p *publish date* -e *edition*\  
-d *description* -g *genre* -n *notes*]

>collection and the optional *holder* identifies who posses the physical book.  If the values specified don't exist in the collections list, a new collection will be created.  
New holder, title, author, and published date are all required.  
Publish date, edition, description, genre and notes are all option fields.  
If the new collection does not already exist, it will be created and assigned to the same owner as the old collection.

### csl-rmv-book - Remove a book entry from the users collection

csl-rmv-book -c *collection* [-h *holder*] -t *title* -a *author*

>Removes the specified book from a collection.

## Collection Management

### csl-add-coll - Add a new collection to the Library

csl-add-coll -n *collection* [-h *holder*]  

>collection - the name for the new collection  
holder *(optional)* - the name of the person or entity that holds the collection.  
If a collection already exists with a matching name and holder, the request will fail.

### csl-list-coll - List all collections in the library

>csl-list-coll
No parameters, returns all the collections in the library.

### csl-list-a-coll - List all books in a collection

csl-list-a-coll -n *collection* [-h *holder*]

>collection - the name of the collection to list  
holder *(optional)* - the name of the person or entity that holds the collection.
If the holder is not specified and multiple collections share the same name.  
All books in all collections will be returned.

### csl-ren-coll - Rename an existing collection

csl-ren-coll -n *collection* [-h *holder*] -to *new-name*  

>collection - the name of the collection  
holder *(optional)* - the name of person or entity that holds the collection.  
If the collection holder is not specified and multiple collections match by name, the request will fail.

### csl-rmv-coll - Remove an existing collection

csl-rmv-coll -n *collection* [-h *holder*]  

>collection - the name of the collection  
holder *(optional)* - the name of person or entity that holds the collection.  
If the collection holder is not specified and multiple collections match by name, the request will fail.  
All books in the collection will be removed from the library listings.

### csl-chg-coll-owner - change the owner of a collection

csl-chg-coll-owner -n *collection* -h *holder* -n *new holder*

>Updates the holder of a collection to a new identity.

## Searching the Library

### csl-search - Search the entire library

csl-search [-c *collection*] [-t *title*] [-a *author*] [-g *genre*]

>allows for searching the library using one or more search criteria.  
If no criteria is specified, the entire contents of the library will be returned.  
Search criteria can contain regular expressions.

## Using the Command Line Interface

Suppose a new member to the library wants to add her personal collection of books to the library.  
Since the web interface is still in development, she decides to use the command line interface.

To add her first book would be:

csl-add-book -c "Harry Potter Original Editions" -h "J.K. Rowling" -t "Harry Potter and the Philosopher's Stone" -a "J.K. Rowling" \  
 -p "June 26, 1997" -e "First" -d "Harry Potter Book 1" -g "Fantasy" -n "Author notes in margins"

The above command would create a new collection called "Harry Potter Original Editions" with "J.K. Rowling" as the holder of the book.
It would then create the entry for the book and associate it with the collection.
