package main

import (
	"fmt"
)

func main(){
	// if条件判断
	age := 15
	if age >= 18 {
		fmt.Println("新中国的成立了")
	} else {
		fmt.Println("要写作业了")
	}
	if age > 35 {
		fmt.Println("已经35岁了。")
	}else if age > 18 {
		fmt.Println("还没有到35岁")
	}else {
		fmt.Println("好好学习")
	}
}