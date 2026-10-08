package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

    // r := bufio.NewReader(os.Stdin)
	// line, _ := r.ReadString('\n')
	// line = strings.TrimRight(line, "\r\n")
	// // TODO: print the text of line, uppercased, with no whitespace around it.
	// // Right now it prints the line untouched, which is wrong for every test.
    // line = strings.ToUpper(strings.TrimSpace(line))

	// fmt.Println(line)


    r := bufio.NewReader(os.Stdin)
	name, _ := r.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")
	qtyLine, _ := r.ReadString('\n')
	qty, _ := strconv.Atoi(strings.TrimSpace(qtyLine))
	priceLine, _ := r.ReadString('\n')
	price, _ := strconv.ParseFloat(strings.TrimSpace(priceLine), 64)
	total := float64(qty) * price

	// TODO: replace this Println with one fmt.Printf that lays the
	// four values out in the receipt columns described in the exercise.
	fmt.Printf("%-12s %3d x %6.2f = %8.2f",name, qty, price, total)

}
