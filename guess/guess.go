package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// number guessing game using binary search
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100

	fmt.Println("Please think of a number between", low, "and", high)
	fmt.Println("Press ENTER when ready")
	scanner.Scan()

	guessAmount := 0

	for {
		if low > high {
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
			guessAmount++
			high = guess - 1
		} else if response == "b" {
			guessAmount++
			low = guess + 1
		} else if response == "c" {
			guessAmount++
			fmt.Println("I won!")
			break
		} else {
			fmt.Println("Enter 'a', 'b' or 'c'")
		}
	}
	fmt.Println("Amount of guesses:", guessAmount)
}
