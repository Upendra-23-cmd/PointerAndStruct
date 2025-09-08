* Question : Write a program to Demonstrate deep vs shallow copy with structs.

==> struct is a data type that is used to hold data of similar related data 

* Shallow copy of struct 
==> Shallow copy of struct ,here we copy the struct fields of the orginal data , but references are shared 

--> for example :- lets take a struct :
                        type original struct{
                            Name string
                            Books []string{"go basics", "docker guide"}
                        }
                
            =>   Now we create a shallow copy of this than it will be
                shallow := original

* Deep copy of struct
==> Deep copy of struct ,copy struct field,but refernce data copies are made new

for example :- we had taken the original struct
            ==>  Now we create a deep copy of the original struct
                deep := original{
                    Name : shallow.lib
                    Books: append([]string(nil), lib1.Books)
                }


when we update the value in both the copy
                for example shallow.Book[0]="concept"
                            deep.Book[1]="cleared"

                we check it by coding what will be the output
                