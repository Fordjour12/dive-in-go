package main

import (
	"fmt"
)	

func main() {
	fmt.Println("Hello World from Go!")

	// initialize a slice using make 
	mySliceInit := make([]int,2,5)
	fmt.Println(mySliceInit)

	// slicing an array
	myArray := [6]int{1,2,3,4,5,6} 
	myArray2 := []int{1,2,3,4,5,6}
	mySlice := myArray[2:5]
	mySlice2 := myArray2[2:5]

	fmt.Println("Printing mySlice and mySlice2:")
	fmt.Println(mySlice)
	fmt.Println(mySlice2)

	// append to a slice
	mySlice = append(mySlice, 7,8,9)
	fmt.Println("Appended 7,8,9 to mySlice")
	fmt.Println(mySlice)

	// copy a slice
	newSlice := make([]int, len(mySlice))
	copiedSlice := copy(newSlice, mySlice)
	fmt.Println("Copied slice:", copiedSlice)
	fmt.Println("Copied New slice:",newSlice)

	// sub-slicing a slice
	// subSlice := newSlice[2:5]
	// println("Sub-sliced slice:", subSlice)


	// accessing a slice
	element := mySlice[3]
	fmt.Println("Element at index 3:", element)

	// iterating over a slice
	for index, value := range mySlice {
		fmt.Println("Index:", index, "Value:", value)
	}


}