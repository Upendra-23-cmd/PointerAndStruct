* Question : Write a program to Build a Task struct with pointer references to dependencies.

==> In this Question we are need to create to created a task system where :
                            1.) Each task is represented by a struct
                            2.) A task may be depended on the other task before it can be execurted
                            3.) we use pointer so each task can directly references to the next Task

*   why we use struct here?
=>  A struct allow us to group
        : Task Name/ID
        : Pointer references to its dependency
    So, we get a complete representation of task in one place

*    why we use pointer here?
=>  Dependencies should not be copy of the other task but actual references
