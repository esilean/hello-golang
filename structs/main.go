package main

import "fmt"

type contactInfo struct {
	phone string
	email string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	p := person{
		firstName: "leandro",
		lastName:  "bevilaqua",
		contact: contactInfo{
			phone: "11989462727",
		},
	}

	p.lastName = "bevila"
	p.contact.email = "le.bevilaqua"

	p.updateName("lea")
	p.toString()
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func (p person) toString() {
	fmt.Printf("%+v\n", p)
}
