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
	// 1st way to assign values to structs.
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// 2nd way to assign values to structs.
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)

	// "%+v - all of the different field name"
	// fmt.Printf("%+v", alex)

	kante := person{
		firstName: "Ngolo",
		lastName:  "Kante",
		contactInfo: contactInfo{
			email:   "ngolo.kante@gmail.com",
			zipCode: 10192,
		},
	}

	// &variable: give me the memory address of the value this variable is pointing at
	// kantePointer := &kante
	// kantePointer.updateName("Jamie")
	kante.updateName("Jamie")
	kante.print()
}

// *pointer: give me the value this memory address is pointing at.
// *person: type description - it means we're working with a pointer to a person.
// *pointerToPerson: operator - it means we want to manipulate the value the pointer is referencing.
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
