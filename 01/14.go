package main

import (
	"fmt"
	"unicode"
)

func main(){

	s1 := "hello张线三在"
	count := 0
	for _, v := range s1{
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Print(count)
}
