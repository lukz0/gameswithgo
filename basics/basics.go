package main

import "fmt"

func main() {
	var torf bool = true

	helloWorld := "Hello World"

	fmt.Println(torf)
	fmt.Println(helloWorld)

	//infinite loop
	/*for {
		fmt.Println("Hello World")
	}*/
	// While
	/*i := 0
	for i < 10 {
		fmt.Println("Hello World:", i)
		i += 1
	}*/

	// For
	/*for i := 0; i < 10; i++ {
		fmt.Println("Hello World:", i)
	}*/
	test := 5
	if test > 5 {
		fmt.Println("Test is bigger than 5")
	} else if test < 5 {
		fmt.Println("Test is smaller than 5")
	} else {
		fmt.Println("Test is exactly 5!")
	}
}
