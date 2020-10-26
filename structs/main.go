package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jake := person{
		firstName: "Jacob",
		lastName:  "Nelson",
		// contactInfo: contactInfo{
		// 	email:   "spam@jacobnelson.tech",
		// 	zipCode: 78741,
		// },
	}
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}
