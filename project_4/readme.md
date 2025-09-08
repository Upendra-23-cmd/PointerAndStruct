* Question : Write a program to Build a book library where each book is a struct, managed with pointers.

==> In this question what we do is we create a struct in which we keep the data of the book( like book name, Author, UID etc)
and we store this data in a slice through  pointer

Than why we use pointer ,Because in if we store the whole data it would be huge and less manageable therefore instead of store whole data we use its address which light and manageable

* why we use []*Book{book1, book2, book3}
[] → means a slice (a dynamic array in Go).
=>    *Book → means the slice will hold pointers to Book structs, not full Book structs.
=>    {book1, book2, book3} → these are the actual pointers (addresses) of the books we already created.

So together:
[]*Book{book1, book2, book3}
creates a slice of book pointers, holding the addresses of book1, book2, and book3.
