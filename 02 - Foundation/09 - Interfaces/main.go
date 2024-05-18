package main

import "fmt"

type ActiveControl interface {
	GetActive() bool
}

type Client struct {
	Name   string
	Age    int
	active bool
}

func (c Client) GetActive() bool {
	return c.active
}

func main() {
	kelps := Client{
		Name:   "Kelps",
		Age:    22,
		active: true,
	}

	fmt.Printf("Name: %s \nAge: %d \nActive: %t\n", kelps.Name, kelps.Age, kelps.GetActive())
}

func GetActivate(a ActiveControl) bool {
	return a.GetActive()
}
