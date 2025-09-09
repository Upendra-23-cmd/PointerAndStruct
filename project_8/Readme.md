* Question : write a program to Design a Car struct with Engine as a nested struct.
==> So here we will create a struct  where we group related data like car name , its engine but engine can we of various name therefore we create a struct where we store engine information


So what we do is we create two struct in one we store car information and in another we store engine information
that is 
        type Engine struct {
    	Name      string
    	Power     int
    	Engine_No int
        }

        type Car struct {
    	Name   string
    	Engine *Engine
        }

where we instead store a complete struct we store the pointer of it which help in managing the memory more efficently 
that throught this line " Engine *Engine "

here we learn how to use nested struct with pointer