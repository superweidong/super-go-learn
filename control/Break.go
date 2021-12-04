package main

import "fmt"

func main() {

	//使用break结束所有循环
OutLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				fmt.Println("==", i, j)
				break OutLoop
			}
			fmt.Println(i, j)
		}
	}
}
