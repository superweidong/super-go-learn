package main

import "fmt"

func main() {

	//使用continue 跳出内部循环  继续执行下一次外部循环
OutLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				fmt.Println("==", i, j)
				continue OutLoop
			}
			fmt.Println(i, j)
		}
	}
}
