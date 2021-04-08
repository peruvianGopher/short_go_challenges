package main

import (
	"fmt"
)

func main() {
	var i interface{}
	var j []int = nil
	i = j
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}
