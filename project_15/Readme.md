* Question : Write a program to Create a graph node struct for adjacency lists.

==> What is a graph node ?
    A graph is non-linear data structure consisting of vertices(nodes) and and edges(connection)
    A graph node (vertex) is just one element of the graph.
                                            Each node stores:
                                            Value / Data → e.g., 1, UserID, ServerName.
                                            Connections (adjacency list) → other nodes it is linked to.
    
    why we use struct here ?
    As we know a struct is a collection of related data So, in this graph we have similar data that is node and connection there we use struct

    if we had not used struct than we have to use other method to connect the node of the graph


    Why we use pointer ? 
    We use pointer to connect node where each node must point to the next node if we don't use pointer we have to copy the whole node whereas when we use pointer what we do we save the address of the node which save the space and the struct is light
