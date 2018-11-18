package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "jim",
		lastName:  "party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 99999,
		},
	}

	fmt.Printf("%v", jim)
}
