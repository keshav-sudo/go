package main

import "fmt"

func main() {
    // String Example
    var str string = "Hello, World!"
    fmt.Println("With type:", str)

    var strWithoutType = "Hello, World!"
    fmt.Println("Without type:", strWithoutType)

    strShorthand := "Hello, World!"
    fmt.Println("Shorthand:", strShorthand)

    // Integer Example
    var num int = 42
    fmt.Println("With type:", num)

    var numWithoutType = 42
    fmt.Println("Without type:", numWithoutType)

    numShorthand := 42
    fmt.Println("Shorthand:", numShorthand)

    // Boolean Example
    var isTrue bool = true
    fmt.Println("With type:", isTrue)

    var isTrueWithoutType = true
    fmt.Println("Without type:", isTrueWithoutType)

    isTrueShorthand := true
    fmt.Println("Shorthand:", isTrueShorthand)

    // Float Example
    var pi float64 = 3.14
    fmt.Println("With type:", pi)

    var piWithoutType = 3.14
    fmt.Println("Without type:", piWithoutType)

    piShorthand := 3.14
    fmt.Println("Shorthand:", piShorthand)
}
