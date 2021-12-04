package main

import "fmt"

func main() {

	str := "hello"
	switch str {
	//一分支单值
	case "hello":
		fmt.Println("hello9")
		break
	//一分支多值
	case "world", "nihao":
		fmt.Println("world9")
		break
	default:
		fmt.Println("nil")
	}

	//分支表达式
	i := 10
	switch {
	case i < 10:
		fmt.Printf("i < 10")
	}
}
