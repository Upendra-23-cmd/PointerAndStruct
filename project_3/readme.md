* Question : Write a program  to Implement a linked list node struct using pointers.

=> A linked list is linear data structure where element are connected using pointer.Unlike arrays(which uses indices and contiguous memory), linked lists use nodes scattered in memory but connected logically via pointers

Each node has two part :-
                            1.) Data -> stores actual data value
                            2.) Pointer/References -> stores the address of the next node

            example :- 
                        type Node struct {
                            data int        -> stores the data
                            next *node      -> stores the address of the next node
                        }

structure of linked list : => [10 | next] -> [20 | next] -> [30 | nil]

Types of linked list 

1.) Singly linked list
2.) Doubly linked list
3.) Circular linked list

