* Question :- Write a program to Create a Shape struct and extend it with Circle and Rectangle using embedding.

==> This question is little for a question saying to use nested struct

Instead in this what we have do is use the feature of oops that is inheritance but we know traditional inheritance is not used in go 

So, what we do is we use struct embeding to achieve compostion (which is one form of reuseability where one struct is embeded into another)

for example:  type A struct{
                                Name string
                                class int
                            }

                we embed this in another struct

                    type B struct{
                                    A
                                    Subject string
                                }

                * This is embeding

* what is the difference between embeded struct and nested struct

==> Key Differences (Embedded vs Nested)

==>Feature	                        Nested Struct	                                    Embedded Struct

* Declaration	                   Field has a name	                    | Struct is included without a field name
                                                                        |
* Accessing fields	            Must use field name → obj.Field.SubField|  Can access directly → obj.SubField
                                                                        |
* Code Reuse	                Provides structure but no promotion	    | Promotes fields/methods for reuse
                                                                        |
* Inheritance-like	            ❌ No inheritance-like behavior         | ✅ Supports inheritance-like behavior
                                                                        |
* Example                        Access	circle.ShapeField.Color	        |    circle.Color


