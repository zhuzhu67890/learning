package main

import "fmt"

func main(){
	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) //把字符串强制转换成了一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3)) //把rune切片强制转换成字符串

	c1 := "红"
	c2 := '红'  //rune(int32)

	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3:%T  c4:%T\n", c3, c4)

	// 类型转换
	n1 := 10
	var f float64

	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T", f)


}