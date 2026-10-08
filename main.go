package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

    // var a, b int
    // fmt.Scan(&a, &b)

    
	// quotient := a / b
	
	// remainder := a % b

    // fmt.Println(quotient)
    // fmt.Println(remainder)

    r := bufio.NewReader(os.Stdin)
	line, _ := r.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")
	// TODO: print the text of line, uppercased, with no whitespace around it.
	// Right now it prints the line untouched, which is wrong for every test.
    line = strings.ToUpper(strings.TrimSpace(line))

	fmt.Println(line)

}
