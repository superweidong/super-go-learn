package main

import (
	"fmt"
)

func main() {

	//break退出双层循环
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break
			}
		}
		break
	}
	fmt.Println("continue")

	//使用goto退出双层循环
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				goto gotoWhere
			}
		}
	}
gotoWhere:
	fmt.Println("goto where ")

	//使用goto统一处理错误
	a := 1
	if a < 1 {
		fmt.Printf("error")
	}
	if a > 1 {
		fmt.Printf("error")
	}

	if a < 1 {
		goto errorHandle
	}
	if a > 1 {
		goto errorHandle
	}
errorHandle:
	fmt.Printf("errorHandle")

}
