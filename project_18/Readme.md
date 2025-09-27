* Question : Write a program to Build a memory pool allocator demo.

==> A memory pool allocator is a technique where you pre allocate a chucnk of memory (a pool) and then reuse it for multiple allocations instead of calling malloc/new/make every time

:- Normally in Go , if you keep creating new objects/arrays , the garbage collector(GC) will eventually free them
:- But Frequent allocation and GC slow down preformanaces in high load systems 

:- A memory pool solves this by keeping a pool a pool of pre made ojects/buffer that you can :
                    1.) GET -> borrow memory when needed
                    2.) PUT -> return memory when done

This avoids constant allocation/deallocation → faster, less GC pressure, predictable memory usage

==> Why do we use struct for a memory pool
* A struct lets us group related data 
    A memoryPool need  to track:
                                1. blockSize -> size of ezch block 
                                2. free list -> available memory block
                                3. syncronization -> (sync.mutex) if multiple goroutines use it

==> why do we use pointer in memoryPool
* 1. Modified shared  state :- If you pass a struct by value , you get a copy .Modifying the copy doesn't affect the original Using pointer(*memoryPool)
means all function work on the same pool instance

Pointer = the “key” to the manager’s office → without a pointer, you only get a photocopy of the records, not the real data.