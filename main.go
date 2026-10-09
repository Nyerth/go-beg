package main

import (
	"fmt"
	
)

func main() {
	// Exercise 01
    // var first, second string
    // fmt.Scan(&first)
    // fmt.Scan(&second)
    // // Print one greeting per line, first name first.
    // fmt.Println("Hello,", first)
    // fmt.Println("Hello,", second)

	// Exercise 02
    // var first, second int
    // fmt.Scan(&first, &second)

    // fmt.Println(first + second)
    // fmt.Println(first - second)
    // fmt.Println(first == second)

	// Exercise 03
    // var a, b int
    // fmt.Scan(&a, &b)

    
	// quotient := a / b
	
	// remainder := a % b

    // fmt.Println(quotient)
    // fmt.Println(remainder)

	// Exercise 04
    // r := bufio.NewReader(os.Stdin)
	// line, _ := r.ReadString('\n')
	// line = strings.TrimRight(line, "\r\n")
	// // TODO: print the text of line, uppercased, with no whitespace around it.
	// // Right now it prints the line untouched, which is wrong for every test.
    // line = strings.ToUpper(strings.TrimSpace(line))

	// fmt.Println(line)

	// Exercise 05
    // r := bufio.NewReader(os.Stdin)
	// name, _ := r.ReadString('\n')
	// name = strings.TrimRight(name, "\r\n")
	// qtyLine, _ := r.ReadString('\n')
	// qty, _ := strconv.Atoi(strings.TrimSpace(qtyLine))
	// priceLine, _ := r.ReadString('\n')
	// price, _ := strconv.ParseFloat(strings.TrimSpace(priceLine), 64)
	// total := float64(qty) * price

	// // TODO: replace this Println with one fmt.Printf that lays the
	// // four values out in the receipt columns described in the exercise.
	// fmt.Printf("%-12s %3d x %6.2f = %8.2f",name, qty, price, total)

	// Exercise 06
	// var n int
	// fmt.Scan(&n)
	
	// switch { 
	// 	case n % 3 == 0 && n % 5 == 0:
	// 		fmt.Println("FizzBuzz")
	// 	case n % 3 == 0:
	// 		fmt.Println("Fizz")
	// 	case n % 5 == 0:
	// 		fmt.Println("Buzz")
	// 	default:				
	// 		fmt.Println(n)
	// }
	// fmt.Println(n)

	// Exercise 07
	var n int
    fmt.Scan(&n)
    // TODO: accumulate the sum of 1..n with a loop, and print that instead of the placeholder.
	if n >= 0 && n <= 100000 {
		sum := 0
		for i := 1; i <= n; i++ {
			sum += i
		}
		fmt.Println(sum)
	}


}
