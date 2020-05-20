package main

import (
	"fmt"
)

func main() {
	c1 := customer{}
	c2 := new(customer)

	fmt.Println(c1)        //=> { 0}
	fmt.Println(c2)        //=> &{ 0}
	fmt.Println(c1 == *c2) //=> true
}

type customer struct {
	name string
	age  int
}
