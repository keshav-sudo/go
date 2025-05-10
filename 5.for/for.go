package main

import "fmt"

func main() {
    // Normal for loop
    for i := 0; i < 5; i++ {
        fmt.Println("Normal loop:", i)
    }

    // While-style loop (Go doesn't have a 'while' keyword, but this mimics it)
    i := 0
    for i < 5 {
        fmt.Println("While-style loop:", i)
        i++
    }

    // Infinite loop
    for {
        fmt.Println("Infinite loop")
        break // Add break to prevent the loop from running forever
    }
}
