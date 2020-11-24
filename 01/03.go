package main

import "fmt"
//声明变量
//var name string
//var age int
//var isok bool
//批量声明
var (
	name string
	age int
	isok bool
)

func main(){
	name = "张三"
	age = 18
	isok = true
	//格式化打印输出
	fmt.Printf("我的名字是%s，年龄是%d", name, age)
	//输出后换行
	fmt.Println(name)
	//普通通打印输出
	fmt.Print(isok)

	//声明变量同时赋值
	var s1 string = "刘四海"
	fmt.Println(s1)
	//类型推导，根据值判断该变量是什么类型
	var s2 = "李四"
	fmt.Println(s2)
	//简短变量声明,只能在函数中使用
	s3 := "哈哈三"
	fmt.Println(s3)

	//匿名变量
	s4, _ := 33, 44
	fmt.Println(s4)
}

func foo()(int,string){
	return 10, "qimi"
}
func ab(){
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

}