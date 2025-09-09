* Question : write a program to Implement a binary tree node struct

==> First ,What is tree :- A tree is a hierarchical data structure that represent element in a parent child     relation

Tree has multiple parts which are : 
                                    1.) Root : The top most node of the tree
                                               A tree always have exactly one root
                                    2.) Node : A basic unit of tree that holds a value
                                               Each node can connect another node
                                    3.) Edge : The link/connection between parent and child nodes
                                    4.) Parent : A node that has one or more children
                                    5.) Child : A node that is directly connected below a parent
                                    6.) Leaf : A node that has no children
                                    7.) Height of tree : The number of edges in the longest path from the root   to a leaf 
                                    8.) Depth of Node : Distance from the root to that node 

* Now why we use struct and pointer in a tree
=> we use struct because we tree a collection of related data which is diffcult to manage them seprately therefore we use struct

* Now why the hell pointer is used
=> Here the pointer place a very crucial role as we know tree are connection of parent and children node now to store the child node we use nested struct if we won't use pointer what happen is the size exponetially becomes big as it has to load whole child struct in the parent struct for one or two node we see no problem but what for 1000 of node that the compiler fail to track the memory that is to be stored in that struct.

therefore we use pointer

so without pointer the memory
                               Root (10)
                            ├─ Left (5)
                            │   ├─ Left ( again child node )
                            │   └─ Right ( again child node )
                            └─ Right (15)
                                ├─ Left ( again child node )
                                └─ Right ( again child node )

 We saw the problem that why we use pointer but how does the pointer solve this problem
So, what a pointer does it saves the memory address of the child node now instead of loading entire struct it only load it address which is easy to manage and easy to update.

so with using pointer the memory becomes
                             Root (10)
                            ├─ Left  → address of node (5)
                            └─ Right → address of node (15)

                            Node(5)
                            ├─ Left  → address of node (3)
                            └─ Right → address of node (7)
