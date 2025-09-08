* Question : Write a program that swaps two variables using pointers.

==> we use two variable to swap the value of the pointer

If you’re just swapping small variables in one place, doing a direct swap in main() is simpler and just fine.
No real benefit in using pointers in this case.
Pointers become beneficial only when you want a reusable swap function or efficiently handle big data types.

In case of reuseable function pointer is beneficial and if small and in main function  use normal method because there is no real benefit