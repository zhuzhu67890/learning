package main

import (
	"fmt"
	"strings"
)

func main(){
	s:= "i'm ok"
	fmt.Println(s)
	// 多行字符串
	s1 := `日照香炉生紫烟，遥看瀑布挂前川。
飞流直下三千尺，疑是银河落九天。`
	fmt.Println(s1)
	//字符串相关操作
	fmt.Println(len(s1))

	name := "理想"
	world := "dsb"
	ss := name + world
	fmt.Println(name + world)
	//ss1 := fmt.SPrintf("%s%s", name, world)
	fmt.Println(ss)
	//fmt.Println(ss1)
	// 分割
	ret := strings.Split(s1, "\\")
	fmt.Println(ret)
	//判断包含
	fmt.Println(strings.Contains(ss,"理想"))
	//前缀
	fmt.Println(strings.HasPrefix(ss,"理想"))
	//后缀
	fmt.Println(strings.HasSuffix(ss,"理想"))
	//判断位置
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4,"c"))
	fmt.Println(strings.Index(s4,"a"))

	//拼接
	fmt.Println(strings.Join(ret,"+"))

}
