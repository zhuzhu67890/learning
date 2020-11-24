package main

import "fmt"

func main(){

	//for i := 1; i<= 10;i++ {
	//	fmt.Println(i)
	//
	//}
	//变种1
	//var i = 5
	//for ; i<10; i++ {
	//	fmt.Println(i)
	//}
	//变种2
	//var i = 5
	//for i < 10 {
	//	fmt.Print(i)
	//	i++
	//}

	// for range循环
	s := "hello张三"
	for i,v := range s{
		fmt.Printf("%d %c\n",i, v)
	}
}
