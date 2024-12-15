package main

import "fmt"

type contactInfo struct {
	city    string
	zipcode string
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	govind := person{
		firstName: "Govind",
		lastName:  "L",
		contact: contactInfo{
			city:    "Chennai",
			zipcode: "33",
		},
	}

	// govind.print()

	//working with pointers

	// govindPointer := &govind
	// govindPointer.updateName("Govindaraajan")
	// govind.print()

	fmt.Println(*&govind.firstName)

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

func (personToPointer *person) updateName(newFirstName string) {
	(*personToPointer).firstName = newFirstName
}
