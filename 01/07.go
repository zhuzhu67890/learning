package main

import "fmt"

//浮点数

func main(){

	//math.MaxFloat32  //float 32最大值
	f1 := 1.234325325
	// 默认go语言中的小数都是float64类型
	fmt.Printf("%T\n", f1)
	// 显示声明float32类型
	f2 := float32(1.2432525)
	fmt.Printf("%T\n", f2)
	// float32类型的值不能直接赋值给float64
	// f1 = f2


}
