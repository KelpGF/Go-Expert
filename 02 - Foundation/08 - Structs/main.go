package main

import "fmt"

type Address struct {
	Street  string
	Number  int
	City    string
	Country string
}
type Contact struct {
	Phone string
	Email string
}

type Client struct {
	Name    string
	Age     int
	active  bool
	Address Address
	Contact
}

func (c Client) GetActive() bool {
	return c.active
}
func (c Client) Inactivate() {
	c.active = false
}

func main() {
	kelps := Client{
		Name:   "Kelps",
		Age:    22,
		active: true,
	}

	fmt.Printf("Name: %s \nAge: %d \nActive: %t\n", kelps.Name, kelps.Age, kelps.GetActive())

	kelps.Age = 23
	kelps.Inactivate() // This method is not changing the value of active in the struct, we will see how to fix this in the Pointers section
	fmt.Printf("\nName: %s \nAge: %d \nActive: %t\n", kelps.Name, kelps.Age, kelps.GetActive())

	kelps.Address.Street = "Main Street"
	kelps.Address.City = "Gotham"
	kelps.Address.Number = 123
	fmt.Printf("\nAddress: %s, %d - %s, %s\n", kelps.Address.Street, kelps.Address.Number, kelps.Address.City, kelps.Address.Country)

	kelps.Phone = "123-456-789"
	kelps.Contact.Email = "kelps@mail.com"
	fmt.Printf("\nPhone: %s\nEmail: %s\n", kelps.Phone, kelps.Email)
}
