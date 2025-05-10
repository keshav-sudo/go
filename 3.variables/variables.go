package main

import "fmt"

func main() {
    var char rune = 'A'
    fmt.Println("With type:", char)

    var charWithoutType = 'A'
    fmt.Println("Without type:", charWithoutType)

    charShorthand := 'A'
    fmt.Println("Shorthand:", charShorthand)
}
