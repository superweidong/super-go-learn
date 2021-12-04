package main

import "fmt"

/**
 * range时间遍历 slice、map、string
 */
func main() {

	//遍历数组
	//rangeArr()
	//遍历切片
	//rangeSlice()
	//遍历map
	//rangeMap()
	//遍历字符串
	str := "ABCDE"
	for i, val := range str {
		fmt.Printf("下标 %d 值 ox%x", i, val)
		fmt.Println()
	}

}

func rangeMap() {
	m1 := map[string]string{"A": "1", "B": "2"}
	for key, value := range m1 {
		fmt.Printf("key %s val %s", key, value)
		fmt.Println()
	}
}

func rangeSlice() {
	sli := []string{"A", "B", "C"}
	sli2 := sli[1:2]
	fmt.Println(sli2)
}

func rangeArr() {
	arr := []int{1, 2, 3}
	for i, val := range arr {
		fmt.Printf("下标：%d 数值：%d", i, val)
		fmt.Println()
	}
}
