package main

import (
	"fmt"
)

func main() {

	var a string
	var b interface{}

	a = "test"

	b = a

	c, _ := b.(string)

	fmt.Println(c)

}
