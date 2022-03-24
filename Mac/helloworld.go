package main

import "fmt"

var count int = 4
var string1 string = "你好"

type person struct {
	name string
	age  int
}

func main() {
	var P person
	P.name = "brian"
	P.age = 30
	fmt.Printf("%s is %d year old \n", P.name, P.age)

}
