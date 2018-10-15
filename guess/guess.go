package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
* Todo:
* 1. Print out number of tries
* 2. Tell when user is lying
 */

func main() {
	// number guessing game using binary search
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100

	fmt.Println("Please think of a number between", low, "and", high, "100")
	fmt.Println("Press ENTER when ready")
	scanner.Scan()

	for {
		if low == high {
			fmt.Println("You are lying")
			break
		}
		fmt.Println("low:", low, "\nhigh", high)

		guess := (low + high) / 2
		fmt.Println("I guess the number is", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too high")
		fmt.Println("(b) too low")
		fmt.Println("(c) correct")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I won!")
			break
		} else {
			fmt.Println("Enter 'a', 'b' or 'c'")
		}
	}
}
