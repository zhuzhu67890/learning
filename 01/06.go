package main

import "fmt"

// 整型

func main(){
	var a = 10
	fmt.Printf("%d\n", a)
	fmt.Printf("%o\n", a) //八进制
	fmt.Printf("%b\n", a) //二进制
	fmt.Printf("%x\n", a) //十六进制
	//八进制
	b := 077
	fmt.Printf("%d\n", b)
	// 十六进制数
	c := 0xff
	fmt.Printf("%d\n", c)

	//查看变量的类型
	fmt.Printf("%T\n", a)
	//声明int8类型的变量
	e := int8(9)
	fmt.Printf("%T\n", e)

}
