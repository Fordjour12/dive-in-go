package main

import (
	"fmt"
)

func greet(name string) {
	fmt.Printf("Hello World %s\n", name)
}

// function with return value
func jay(name string) string {
	return name
}

// variadic function
// By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters.
func sum(numbs ...int) int {
	total := 0
	for _, num := range numbs {
		total += num
	}
	return total
}

// Recursive function
func factorial(x int) int {

	if x == 0 {
		return 1
	}
	return x * factorial(x-1)

}

func main() {
	fmt.Println(jay("Jay"))
	greet("Jake Gyllenhaal")

	fmt.Println(factorial(5))

	// closure
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2))

	fmt.Println(sum(1, 2, 3, 4))

	// Anonymous function
	and := func() {
		fmt.Println("I am an anonymous function")
	}
	and()
}
