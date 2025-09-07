* Question : write a go program to Implement a Vector2D or Vector3D struct with methods for addition subtraction, and dot product.

==> We first learn the concept what is vector,
 * A Vector is a quantity that has both magnitude and direction,Whereas Scalar is a qunatity that has only value
   that is magnitude

=> Now then what is Vector2D , Vector2D is a  Two dimensional vector that is it have two axis one is x(horizontal) and another is y(vertical)

Method to perform the following are(Taking two vector A and B) :-

* For Addition    : A = (x1,y1)
                  B = (x2,y2)
                  then the Addition of two vector is A+B which is ((x1+x2),(y1+y2))

* For Subtraction : A = (x1,y1)
                  B = (x2,y2)
                  then the Subtraction of two vector is A-B which is ((x2-x1),(y2-y1))

* For Dot Product : A = (x1,y1)
                  B = (x2,y2)
                  then the Dot Product of two vector is A.B which is ((x1 * x2)+(y1 * y2))

whereas Vector3D is a Three dimensional vector that has three axis one is x(horizontal) another is y(vertical) and the third one is z(depth)

Method to perform the following are(Taking two vector A and B) :-

* For Addition    : A = (x1,y1,z1)
                  B = (x2,y2,z2)
                  then the Addition of two vector is A+B which is ((x1+x2),(y1+y2),(z1+z2))

* For Subtraction : A = (x1,y1,z1)
                  B = (x2,y2,z2)
                  then the Subtraction of two vector is A-B which is ((x2-x1),(y2-y1),(z2-z1))

* For Dot Product : A = (x1,y1,z1)
                  B = (x2,y2,z2)
                  then the Dot Product of two vector is A.B which is ((x1 * x2)+(y1 * y2)+(z1 * z2))


=> Why we Use pointer and Struct in the code ?

* We use structs to represent vectors in a clean, object-like way.
* We use pointers when we want efficiency or need to mutate (change) the original vector.
* For small vectors (2D, 3D), value receivers are often fine for math functions, but scaling/transforms often
   use pointers.

So, what we do is we create a struct where we store the value of a vector and perform the action


Now we let's understand the code : 

* func (v1 *Vector2D) AddIn(v2 Vector2D) := In this we use a receiver,A receiver is a special parameter in Go methods that tells the compiler:
👉 “This function belongs to this type.”

It is written between the func keyword and the method name.  
* => func (r TypeName) MethodName(args...) returnType { ... }

r → the receiver variable (like a placeholder name: v, t, p, etc.)
TypeName → the type the method is attached to (struct, slice, map, or even a custom type).

There are two type of receiver, they are :-
                                            1.) Value receiver
                                            2.) Pointer receiver

