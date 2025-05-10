package main

import "fmt"

func main() {
    // len(): length of slice
    nums := []int{10, 20, 30}
    fmt.Println("Length:", len(nums)) // 3

    // cap(): capacity of slice
    fmt.Println("Capacity:", cap(nums)) // 3

    // append(): add elements to slice
    nums = append(nums, 40)
    fmt.Println("After append:", nums) // [10 20 30 40]

    // make(): create a map
    m := make(map[string]int)
    m["a"] = 1
    fmt.Println("Map with make:", m) // map[a:1]

    // delete(): remove key from map
    delete(m, "a")
    fmt.Println("Map after delete:", m) // map[]

    // new(): allocate memory
    p := new(int)
    *p = 100
    fmt.Println("Value from new pointer:", *p) // 100

    // copy(): copy slice data
    a := []int{1, 2, 3}
    b := make([]int, 3)
    copy(b, a)
    fmt.Println("Copied slice:", b) // [1 2 3]
}
