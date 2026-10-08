package main

import "fmt"

func main() {
    // var first, second string
    // fmt.Scan(&first)
    // fmt.Scan(&second)
    // // Print one greeting per line, first name first.
    // fmt.Println("Hello,", first)
    // fmt.Println("Hello,", second)

    // var first, second int
    // fmt.Scan(&first, &second)

    // fmt.Println(first + second)
    // fmt.Println(first - second)
    // fmt.Println(first == second)

    var a, b int
    fmt.Scan(&a, &b)

    
	quotient := a / b
	
	remainder := a % b

    fmt.Println(quotient)
    fmt.Println(remainder)

}
