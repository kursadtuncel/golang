package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

//go's structs are typed collections of fields. they're useful for grouping data together to form records.

//this person struct type has name and age fields.
//newPerson constructs a new person struct with the given anme.
//go is a garbage collected languagee; you can safely return a pointer to local variable - it will only be cleaned-
//+ up by the garbage collector when there are no active references to it.

//An & prefix yields a pointer to the struct.
//You can also use dots with struct pointers - the pointers are automatically dereferenced.
//If a struct type is only used for a single value, we don’t have to give it a name. The value can have an anonymous struct type. This technique is commonly used for table-driven tests.
