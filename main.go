package main

import (
	"fmt"
	"go-custom-data-types/organization"
)

func main() {

	p := organization.NewPerson("Banana", "Wilson", organization.NewEuropeanUnionIdentifier(666, "Germany"))
	err := p.SetTwitterHandler("@wilbana")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	// name1 := Name{First: "", Last: ""}
	// name2 := OtherName{First: "Banana", Last: "James"}

	// ssn := organization.NewSocialSecurityNumber("420-69-6666")
	// eu := organization.NewEuropeanUnionIdentifier("420-69-6666", "France")
	// eu2 := organization.NewEuropeanUnionIdentifier("420-69-6666", "France")

	// you can only use '==' if the memory layout is predictable
	// if the types are different but have the same field names, can't access '=='
	// you CAN compare two interfaces
	// portfolio := map[Name][]organization.Person{}
	// portfolio[name1] = []organization.Person{p}

	// if name1.equals(Name{}) {
	// 	println("same hat")
	// } else {
	// 	println("wrongdog")
	// }

	// fmt.Printf("%T\n", ssn)
	// fmt.Printf("%T\n", eu)
	// fmt.Printf("%T\n", eu2)

	fmt.Println(p.Country())
	fmt.Println(p.ID())
}

type Name struct {
	First  string
	Last   string
	Middle []string
}

// func (n Name) equals(otherName Name) bool {
// 	return n.First == otherName.First && n.Last == otherName.Last && len(n.Middle) == len(otherName.Middle)
// }

type OtherName struct {
	First string
	Last  string
}
