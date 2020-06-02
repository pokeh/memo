package main

import "fmt"

func main() {
	// assertion: from an interface to a concrete type
	// v.(T)
	var eater Eater
	eater = Human{}
	human := eater.(Human)
	fmt.Printf("%T\n", human) //=> Human

	// conversion: between two conrete types
	// T(v)
	var someUint64 uint64
	var someInt32 int32 = 3
	someUint64 = uint64(someInt32)
	fmt.Printf("%T\n", someUint64) //=> uint64
}

type Eater interface {
	Eat(food string)
}

type Human struct{}

func (h Human) Eat(food string) {
	fmt.Println("mm", food)
}
