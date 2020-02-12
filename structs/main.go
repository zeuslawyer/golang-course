package main

import "fmt"

type contactInfo struct {
	zipCode int
	email   string
}

type person struct {
	firstName   string
	lastname    string
	contactInfo contactInfo // custom structs are not italicised the way primitive types are.
	// just contactInfo can also be used like javascript object notation if it was the property name instead of contact
}

func main() {
	jim := person{
		firstName:   "Jim",
		lastname:    "Samsoo",
		contactInfo: contactInfo{zipCode: 5455, email: "jimmysoo@jimmy.com"}}

	jimPointer := &jim   
	jimPointer.updateName("Barry")

	jim.print()
}
 
// receiver gets a pointer so that it updates the actual values, otherwise it will pass a copy of jim and that wont work!
// this is necessary when we want to mutate the value received in the receiver
func (personPointer *person) updateName(newFirstName string) *person {
	// (*personPointer).firstName = newFirstName
	return personPointer
}

func (p person) print() {
	fmt.Printf("%+v\n", p.contactInfo) // prints out the key + value (entire struct/obj). just %v will give you only the values
}
