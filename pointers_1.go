package main

import (
	"fmt"
)

type person struct {
	n string
}

func main() {
	a := &person{
		n: "John",
	}
	b := a
	b.n = "Sam"
	fmt.Println(a.n)
}
