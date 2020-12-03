package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex.firstName, alex)
	// var alex person
	// fmt.Println(alex)
	// alex.firstName = "Alex"
	// fmt.Println(alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Parsons",
		contactInfo: contactInfo{
			email:   "jim.parsons@abc.com",
			zipcode: 780098,
		},
	}
	// fmt.Printf("Jim is: %+v", jim)
	jim.print()

	jimNew := &jim
	jimNew.updateFirstName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
