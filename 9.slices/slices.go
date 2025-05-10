package main

import "fmt"

func main() {
    // make(): Create slice with length 3 and capacity 5
    s := make([]int, 3, 5)
    s[0] = 10
    s[1] = 20
    s[2] = 30
    fmt.Println("Slice created with make:", s) // [10 20 30]

    // len(): Get the length of slice
    fmt.Println("Length:", len(s)) // 3

    // cap(): Get the capacity of slice
    fmt.Println("Capacity:", cap(s)) // 5

    // append(): Add elements
    s = append(s, 40, 50)
    fmt.Println("After append:", s) // [10 20 30 40 50]

    // append() with another slice
    more := []int{60, 70}
    s = append(s, more...)
    fmt.Println("Append another slice:", s) // [10 20 30 40 50 60 70]

    // copy(): Copy to another slice
    copied := make([]int, len(s))
    copy(copied, s)
    fmt.Println("Copied slice:", copied) // same as s

    // new(): Allocate pointer to a slice type (empty, nil slice)
    p := new([]int)
    fmt.Println("Pointer to empty slice using new():", p) // &[]

    // slice from slice
    sub := s[1:4]
    fmt.Println("Sub-slice (1:4):", sub) // [20 30 40]
}
