[UNSOLVED]

Runtime Error. Your submission stopped because of a runtime error. ex: division by zero, array index out of bounds, uncaught exception You can try testing your code with custom input and try putting debug statements in your code.
Your submission encountered runtime error for the following input:
32 4 S 5 13 S 9 6 S 4 1 G 4 S 6 1 S 8 11 G 13 G 1 S 12 12 G 10 S 15 13 S 2 13 S 7 5 S 10 3 G 6 G 10 S 15 14 S 5 12 G 5 G 7 G 15 G 5 G 6 G 10 S 7 13 G 14 S 8 9 G 4 S 6 11 G 9 S 6 12 G 3

There are 1 lines in the input

Line 1 ( Corresponds to arg 1 ) : The line starts with a pair of number numOperations, capacity. capacity is the number your constructor is initialized with. 
Then numOperation operations follow. 
Each operation is either : 
 * G  : This corresponds to a function call get()
 * S   : This corresponds to a function call set(num1, num2)
Note that the function calls are made in order.  

The expected output for this testcase is:
1 -1 -1 -1 -1 3 12 5 14 12 -1 3 -1 -1 -1 -1 


==========================================================

https://www.interviewbit.com/problems/lru-cache/

Design and implement a data structure for LRU (Least Recently Used) cache. It should support the following operations: get and set.

    get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
    set(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity, it should invalidate the least recently used item before inserting the new item.

The LRU Cache will be initialized with an integer corresponding to its capacity. Capacity indicates the maximum number of unique keys it can hold at a time.

Definition of “least recently used” : An access to an item is defined as a get or a set operation of the item. “Least recently used” item is the one with the oldest access time.

    NOTE: If you are using any global variables, make sure to clear them in the constructor. 

Example :

Input : 
         capacity = 2
         set(1, 10)
         set(5, 12)
         get(5)        returns 12
         get(1)        returns 10
         get(10)       returns -1
         set(6, 14)    this pushes out key = 5 as LRU is full. 
         get(5)        returns -1 

