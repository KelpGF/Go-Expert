package main

import "fmt"

func main() {
	salaries := map[string]float64{
		"Kelvin": 1234.56,
		"Jenny":  2000.0,
	}

	fmt.Println("salaries:")
	fmt.Println(salaries)

	fmt.Println("\nKelvin:")
	fmt.Println(salaries["Kelvin"])

	delete(salaries, "Kelvin")
	fmt.Println("\nsalaries:")
	fmt.Println(salaries)

	salaries["Jelps"] = 9999.99
	fmt.Println("\nsalaries:")
	fmt.Println(salaries)

	for name, salary := range salaries {
		fmt.Printf("Name: %s, Salary: %.2f\n", name, salary)
	}

	fmt.Println("\nClear Maps")
	map1 := map[string]string{}
	map2 := make(map[string]string)
	fmt.Println(map1)
	fmt.Println(map2)
}
