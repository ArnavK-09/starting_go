package main

// imports,
import "fmt"

// basic type
type Person struct {
	name string
	age  uint8
}

// struct functions
func (p *Person) bday() {
	p.age += 1
	fmt.Println(p.name, " Birthday today! Age: ", p.age)
}

// embed types
type User struct {
	year   uint
	person Person
}
type User2 struct {
	year   uint
	Person // anonymous field
}

// main
func main() {
	// creating variables with struct
	p1 := Person{"Go", 3}
	var p2 Person // init with base data type value
	p2.name = "Golang"
	p2.age = 5
	p3 := &Person{
		name: "User",
		age:  69}
	p4 := new(Person) // returns pointer
	// print
	fmt.Println(p1, p2, p3, p4)

	// functions
	p1.bday()
	p2.bday()
	p3.bday()

	// embed structs
	user := User{}
	user.person.name = "test"
	user.person.bday()
	// or
	user2 := User2{}
	user2.name = "test2" // directly
	user2.Person.name = "test3"
	user2.Person.bday() // indirect
	user2.bday()

}
