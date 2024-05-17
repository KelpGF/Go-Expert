package main

import "fmt"

const stringFormat = "len=%d cap=%d %v\n"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sLen := len(s)
	s2 := make([]int, 10) // []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	fmt.Printf(stringFormat, len(s2), cap(s2), s2)
	fmt.Println("\n\nComplete Slice")
	fmt.Printf(stringFormat, len(s), cap(s), s)

	fmt.Println("\n\nEmpty Slice")
	fmt.Printf(stringFormat, len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf(stringFormat, len(s[0:0]), cap(s[0:0]), s[0:0])

	fmt.Println("\n\nFirsts position")
	fmt.Printf(stringFormat, len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf(stringFormat, len(s[0:4]), cap(s[0:4]), s[0:4])

	fmt.Println("\n\nLasts position")
	fmt.Printf(stringFormat, len(s[2:]), cap(s[2:]), s[2:])
	fmt.Printf(stringFormat, len(s[2:10]), cap(s[2:10]), s[2:10])
	fmt.Printf(stringFormat, len(s[2:sLen]), cap(s[2:sLen]), s[2:sLen])

	s = append(s, 11)
	fmt.Println("\n\nAppend")
	fmt.Printf(stringFormat, len(s), cap(s), s)

	for index, value := range s {
		fmt.Printf("Position: %d, Value: %d\n", index, value)
	}
}
