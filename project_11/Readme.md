* Question : Write a program to Implement a Matrix struct with pointer-based operations.

==> So, Here we store a matrix in a struct using pointer

Now what is matrix, matrix is a collection of data stored in the form of rows and column

for example : A matrix M[2x3] stores data as,

                                [1 2 3]
                                [4 5 6]

what we do is we create a struct for this matrix where we take row , column and data 

that is for example :- 
                        type Matrix struct{
                            Row  int
                            cols int
                            data *[]float 
                        }

    The matrix is stored in the struct in this way because it store the data as 

                        row = 3
                        clos = 5
                        data = address of the slice
    
    if we had not used the pointer in the slice than it would have stored whole slice in the struct which can be huge for a large matix that is 

                         row = 3
                        clos = 5
                        data = [1,2,3,4,5]
    
    Now why we have used this struct here 
                                             because A matrix isn’t just numbers — it has metadata + data:
                                                            How many rows.
                                                            How many columns.
                                                            The actual numbers.
                        If we didn’t use a struct, we would have to manage three separate variables every time:
        
Benefits of using a struct here :-

1.) Encapsulation
                Bundles Rows, Cols, and Data together.
                You don’t need to manage them separately.
2.) Readability
                Functions become cleaner.
                Example: mat.Set(1,2,6) is way easier to understand than writing custom index math everywhere.
3.)Reusability
                Once you define a Matrix struct, you can reuse it in different programs, libraries, or operations (addition, multiplication, transpose).
4.)Extensibility
                Later, you can easily add more fields (like matrix name, precision type, metadata for sparse matrix, etc.) without breaking the design.
5.)OOP-style methods
                Go doesn’t have classes, but attaching methods to a struct (like Set, Get, Add) makes it behave like an object.
 
            This is why you see (m *Matrix) Set(...) — we’re giving behavior to our data.

* Math behind matrix storing value

                    { index := row * m.Cols + col }
    
