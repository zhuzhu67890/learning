package main

import (
	"fmt"
)

const pi = 3.41234
const (
	code1 = 16
	code2 = 18
	error = 404

)
// 批量声明常量时，如果某一行声明 后没有赋值，默认就和上一行一致
const (
	n1 =100
	n2
	n3
)
// iota
const (
	a1 = iota
	a2 = iota
	a3 = iota

)

func main(){

	fmt.Print(code1)
	fmt.Println(code2)
	fmt.Println(error)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}