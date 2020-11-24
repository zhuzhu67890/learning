package main

import "fmt"

const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
	c5
)

func main(){
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println(c5)
	fmt.Print()

}