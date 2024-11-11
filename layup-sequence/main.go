package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	// Check if an argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number as an argument.")
		return
	}

	// Get the first argument
	arg := os.Args[1]

	// Convert the argument to an integer
	sequenceIterations, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Invalid number:", err)
		return
	}

	num1 := new(big.Int) // S(n - 2), the second to last number in our sequence
	num1.SetString("1", 10)

	num2 := new(big.Int) // S(n - 1), the latest number in our sequence
	num2.SetString("2", 10)

	newNum := new(big.Int) // the newly calculated number in our sequence

	if (sequenceIterations == 1) {
		fmt.Printf("%d\n", num1)
		return
	}

	for counter := 3; counter <= sequenceIterations; counter++ {
		if counter % 2 == 0 { // n is even
			newNum.Add(num2, num1)
		} else { // n is odd
			newNum.Mul(num2, big.NewInt(2))
			newNum.Sub(newNum, num1)
		}
		num1.Set(num2)
		num2.Set(newNum)
	}

	fmt.Println(num2)
}
