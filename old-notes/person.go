package organization

import (
	"errors"
	"fmt"
	"strings"
)

type Name struct {
	first string
	last  string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

// type Handler struct {
// 	handle string
// 	name    string
// }

// type aliases are essentially a reference to another type
// when you make an alias, you have access to all the same
// fields of that type and all the same methods

// type TwitterHandler = string
// removing the '=' allows us to satisfy the requirements of RedirectUrl
type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandle := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandle)
}

// functions are the only thing you can put in an interface

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Country() string
	Identifiable
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States of America"
}

type europeanUnionIdentifier struct {
	id      string
	country string
}

func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return europeanUnionIdentifier{
		id:      id,
		country: country,
	}
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeanUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

// Any type that implements this ID() method will be an Identifiable type
// Go implicitly inherits interfaces
// Nowhere in the code are we saying the Person type is a type of Identifiable.
// Example: We dont say type Person: Identifiable, Hashable, IIterable
// With go we don't need to do that at all.

// Instead Go just says, hey I have this interface with ID() string
// Hey, Person has a method on it with that same signature
// ergo, Person implements the type Identifiable

type Person struct {
	Name
	twitterHandler TwitterHandler
	// embedding the methods of the interface onto the type Person
	Identifiable
	Citizen
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
	}
}

func (p *Person) ID() string {
	return fmt.Sprintf("Person's Identifier: %s", p.Citizen.ID())
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handle should start with an '@' symbol")
	}

	p.twitterHandler = handler
	return nil
}

// when you're editing state within a custom type like this
// you need to use a pointer based receiver
// OR you need to return a new version of that type
// because when its being called a copy is being generated
// and so as you're changing the state,
// you're no longer changing the state of the original variable
// these pointer-based method receivers are really useful for modifying state
// for consistency, if you do have this common pattern is to make all
// of your methods pointer-based method receivers instead of value based
// for consistency more than anything
// even tho you're not technically wrong if you leave the read only ones value-based
//
// also use this pattern if youre not even modifying state but the object is just huge
// and holding a massive amount of data because as you're calling each method you'll
// be making copies of data over and over again. but if this is the case, try to
// break it down smaller and rearchitect it with other solutions
func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}

// a method declaration copies the fields of a type over to a new type
// whereas a type ALIAS copies the fields AND the methods
// they become that exact type
// type declaration is without '='
