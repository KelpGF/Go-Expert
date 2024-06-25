package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
}

func changeName(u *User) {
	u.Name = "Kelvin"
}

func main() {
	var user User = User{
		Name: "",
	}
	changeName(&user)
	fmt.Println(user.Name)
}

func main2() {
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
