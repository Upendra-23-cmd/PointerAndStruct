* Question : Write a Program to Implement a pointer-based queue.

==> A queue is a linear data structure that follows the rule :
                            FIFO that is First In First Out
            which means first element inserted is the first one to come out
    
    Operation performed in the Queue are : 
                                            1.) Enqueue -> Add an Element at the rear(end)
                                            2.) Dequeue -> Remove an Element from the top
                                            3.) Peak    -> See the first element of the queue
                                            4.) IsEmpty -> Check if the queue is empty or not
    
    Why we use struct in it 
->  Here we have handle multiple variable if we handle data seperately therefore we use struct which help to collect similar data together which help build a linked list connected with the previous node

    why we use pointer in it 
->  Without pointers we need arrays, which have fixed size 
    Pointer allow dynamic memory allocation , so queue can grow or shrink

