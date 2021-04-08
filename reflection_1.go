package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1
	j := 1
	no1 := &i
	no2 := &j

	if reflect.DeepEqual(no1, no2) {
		fmt.Println("equal")
		return
	}

	fmt.Println("not equal")
}
