package main

import "fmt"

const (
	_ = iota
	Num1
	Num2
	Num3
	Num4
	Num5
	Num6
	Num7
)
var name string
var age int

func main(){
	fmt.Println(Num1)
	fmt.Println(Num2)
	fmt.Println(Num3)
	fmt.Printf("%T", Num4)
	name = "刘子"
	age = 18
	fmt.Println(name)
	fmt.Println(age)


}
