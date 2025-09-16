* Question : Write a program to Implement a pointer-based stack.

==> What is a Stack ? 
    :- A stack is a linear data structure that follow the rule :
                            LIFO that is Last In First Out
        
        Basic Operation in Stack 
                                1.) Push -> Add element at the top
                                2.) Pop  -> Remove the top element
                                3.) Peek/Top -> Look at the top element without removing it
                                4.) Isempty  -> Check if the stack has no element
    
*   Why we use struct in it ?
    :- Struct is the collection of similar data and in stack we store 
                                                        : data|nextnode value
                                    so, it hold data and point to the next node 
        And if we don't use struct than we have to use array or slice which are larger in size
        therefore we use struct here

    
*   Why we use pointer in it ?
    :- Pointer is the data type that store the memory address of another variable

    Now we use pointer here because Without pointer, we'd need an array or slice , which has a fixed size or costly resizing
    Pointer allow us to create a linked list, where each node dynamically points to the next 
    This makes the stack:
                            * Flexible(no size limit)
                            * Efficient(fast push/pop)
                            * Memory-safe(no unused reserved space like array)

*   How it works 
    