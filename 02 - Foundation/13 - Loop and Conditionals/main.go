package main

func main() {
	conditionals()
	// loop()
}

func loop() {
	// for loop
	for i := 0; i < 3; i++ {
		println(i)
	}

	// while loop
	j := 3
	for j < 6 {
		println(j)
		j++
	}

	// do while loop
	k := 6
	for {
		println(k)
		k++
		if k == 10 {
			break
		}
		continue
	}

	// range loop
	nums := []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		println(i, num)
	}

	// infinite loop
	// for {
	// 	println("Infinite loop")
	// }
}

func conditionals() {
	// if else
	a := 10
	b := 20
	c := 30

	if a == 10 {
		println("a is equal to 10")
	}

	if b != 21 {
		println("b is not equal to 21")
	} else {
		println("b is equal to 21")
	}

	if c == 31 {
		println("c is equal to 31")
	} else if c == 30 {
		println("c is equal to 30")
	} else {
		println("c is not equal to 30 or 31")
	}

	println(a == 10 && b == 20)
	println(a == 10 || b == 21)

	// switch case
	switch c {
	case 10:
		println("Equal to 10")
	case 20:
		println("Equal to 20")
	default:
		println("Not equal to 10 or 20")
	}
}
