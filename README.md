# sw-median
# How to run
You need to provide two flag
1. file location [-f] required
2. window size [-w] default is 3 (optional)

- run with default sliding window size
```go run main.go run -f=./test3.csv```

- run with custom sliding window size
```go run main.go run -f=./test3.csv -w=1000```

# Approach
1. one queue to track the sliding window lenght and item index
2. two priority queue to calculate median, (max priority queue, min priority queue). 

# Run time
assume: n = total number,  w = window size
I used two priority queue and maximum queue lenght should be <= (w/2)+1.
Priority queue takes O(logN) for every insert and delete operation.
But to get first item from the queue, time is  O(1)

to get median, run time = O(1)
to add delay, run time = O(n log w)

# Why custome queue in (/ds/queue)
Before I write this code, I look through some Go data structure library. They implement queue using slice.
So there is a problem is slice. 
For example- ```s := []int{1,2,3,4,5,6,7,8,9,10}```
then, if I want take 4 element from index 4, then code is ```s1 := s[4:4+4]``
But the problem is
```
re-slicing a slice doesn't make a copy of the underlying array. The full array will be kept in memory until it is no longer referenced. Occasionally this can cause the program to hold all the data in memory when only a small piece of it is needed.
(https://blog.golang.org/go-slices-usage-and-internals)
```

So that, I write queue without slice.