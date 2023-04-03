package main

import (
	"fmt"

)

type Animal interface{
	sayHi(who string) string
}

type Cat struct {

}

func (cat *Cat) sayHi(who string) string {
	return "this is cat, sayHi to " + who
}


func main(){

	fmt.Println("this is main")
	cat := &Cat{}
	fmt.Println(cat.sayHi("dog"))

	annimal := Animal(cat)
	fmt.Println(annimal.sayHi("human"))

}