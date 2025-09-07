* Question : write a Program to Create a Student record system using structs with pointer receivers for updating grades.

=> In this question we will use the concept of nested struct to store data of some student

this is used to store fixed no. of student data and change there grade using pointer receiver


In this code we use a struct to store all the related details of the student (that is name ,roll no,grades)and created a new struct in which we store the detail of the class that in which class which student is there

Now when we change grades of the student we use pointer receiver because we need to change the actual data ,if we have used the value receiver we have made a copy of the update data and stored which is the waste of the memory therefore we directly change the original data therfore whenever we need the data we get the updated data not the old one