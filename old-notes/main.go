package main

import (
	"fmt"
	"go-custom-data-types/organization"
)

func main() {

	// var p organization.Identifiable = organization.Person{FirstName: "James", LastName: "Wilson"}

	// ^^ here we're saying p is of type Identifiable
	// if we type p. it doesnt give us the properties of Person, it gives us the method on Identifiable
	// because we've assigned this variable to match the interface Identifiable and not
	// the interface of our struct type Person

	p := organization.NewPerson("Banana", "Wilson", organization.NewEuropeanUnionIdentifier("123-45-6789", "Germany"))
	// now we can p. into the properties and the method
	// p.FirstName = "Collin"
	fmt.Println(p.ID())
	fmt.Println(p.Country())
	fmt.Println(p.Name.FullName())
	err := p.SetTwitterHandler("@wilbana")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}
	fmt.Println(p.TwitterHandler())
	fmt.Println(p.TwitterHandler().RedirectUrl())
	fmt.Printf("%T\n", organization.TwitterHandler("Test"))

	// but say we don't want someone to modify fields like this
	// made the fields in the struct lowercase
	// now when we p. we just get the ID method because the fields are not exported

	// so what can we do?
	// go has no concept of constructors but a common idiom to use is just creating a new function
	// called new + your type
}
