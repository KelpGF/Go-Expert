package main

import "fmt"

func main() {
	var myArray [3]int

	myArray[0] = 42
	myArray[1] = 27
	myArray[2] = 99

	fmt.Println(len(myArray))
	fmt.Println(myArray[0])
	fmt.Println(myArray[len(myArray)-1])

	for index, value := range myArray {
		fmt.Printf("Position: %d, Value: %d\n", index, value)
	}
}
