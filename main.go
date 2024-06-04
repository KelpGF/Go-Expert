package main

import "time"

func main() {
	// for i := 0; i < 10; i++ {
	// 	println(i)
	// }

	// numberArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for index, value := range numberArr {
	// 	println(index, value)
	// }

	// for {
	// 	println("Hello, World!")
	// 	time.Sleep(1 * time.Second)
	// }

	// condition := true
	// for condition {
	// 	println("Hello, World!")
	// 	time.Sleep(1 * time.Second)
	// 	condition = false
	// }

	condition := true

	for {
		println("Hello, World! kelvin")
		time.Sleep(1 * time.Second)

		condition = false
		if !condition {
			break
		}
	}
}
