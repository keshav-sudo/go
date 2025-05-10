package main

import "fmt"

func checkNumber(num int) {
    // If-Else statement
    if num%2 == 0 {
        fmt.Println("Even number")
    } else {
        fmt.Println("Odd number")
    }

    // If-Else-If statement
    if num < 10 {
        fmt.Println("Less than 10")
    } else if num == 10 {
        fmt.Println("Equal to 10")
    } else {
        fmt.Println("Greater than 10")
    }
}

func main() {
    checkNumber(7)  // Test with an odd number less than 10
    checkNumber(10) // Test with an even number equal to 10
    checkNumber(15) // Test with a number greater than 10
}
