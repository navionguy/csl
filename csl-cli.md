# Community Shared Library - Command Line Interface

The Community Shared Library allows any group of people to maintain a catalog off all the books available within the community.  Most users will access the shared library using the web UI, but library administrators can use the CLI described here for automating many tasks, bulk loading of new book collections for example.

## Book Management

### csl-add-book - Add a book to a collection

csl-add-book -c *collection* [-h *holder*] -t *title* -a *author* -p *publish date* [-e *edition*] [-d *description*] [-g *genre*] [-n *notes*]

>collection and the optional *holder* identifies who posses the physical book.  If the values specified don't exist in the collections list, a new collection will be created.  
title, author, and published date are all required.  
edition, description, genre and notes are all option fields.

### csl-mod-book - Modify one or more fields for a books entry

csl-mod-book -c *collection* [-h *holder*] -t *title* -a *author* [-nc *new collection* -nh *new holder* -p *publish date* -e *edition* -d *description* -g *genre* -n *notes*]

>collection and the optional *holder* identifies who posses the physical book.  If the values specified don't exist in the collections list, a new collection will be created.  
title, author, and published date are all required.  
edition, description, genre and notes are all option fields.

### csl-rmv-book - Remove a book entry from the users collection

csl-rmv-book -c *collection* [-h *holder*] -t *title* -a *author*

>Removes the specified book.

## Collection Management

### csl-add-coll - Add a new collection to the Library

csl-add-coll -n *collection* [-h *holder*]  

>collection - the name for the new collection  
holder *(optional)* - the name of person or entity that holds the collection.  
If a collection already exists with a matching name and holder, the request will fail.

### csl-list-coll - List all collections in the library

>csl-list-coll
No parameters, returns all the collections in the library.

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

### csl-chg-coll-owner - change the owner of a collection

csl-chg-coll-owner -n *collection* -h *holder*

## Searching the Library

### csl-search - Search the entire library

csl-search [-c *collection*]
