package main

import (
	"fmt"
	"strconv"
)

func main() {

	//定义数组
	var arr [10]string
	//带默认值的定义
	arr2 := []string{"1", "2", "3"}
	var arr3 = []string{"3"}
	println(arr2)
	println(arr3)
	//普通for循环赋值
	for i := range arr {
		//将i转为字符串
		arr[i] = strconv.Itoa(i) + "A"
	}
	//普通for
	normalFor(arr)
	//切片截取
	sliceSplit(arr)
	//make函数
	makeSlice()
	//copy
	copySlice()
	//remove array
	removeArray(arr)
}

func removeArray(arr [10]string) {
	remIdx := 2
	var arr2 []string
	arr2 = append(arr[:remIdx], arr[remIdx+1:]...)
	fmt.Println("remove ", arr2)
}

//copy()实现往切片复制
func copySlice() {
	c1 := []int{1, 2, 3}
	c2 := []int{4, 5}
	copy(c2, c1)
	fmt.Println(c1[:], c2[:])

	c3 := []int{1, 2, 3}
	c4 := []int{4, 5}
	copy(c3, c4)
	fmt.Println(c3[:], c4[:])

	c5 := []int{9, 10}
	//将c5引用给refC6   c5变动 refC6也变动
	refC6 := c5
	refC7 := make([]int, 3)
	//如果数copy出来的 不随原数组变动而变动
	copy(refC7, c5)
	c5[0] = 11
	fmt.Println(c5[:], refC6[:], refC7[:])

}

//append()实现往切片追加元素
func appendSlice() {
	strList := make([]int, 2, 3)
	strList = append(strList, 4, 5, 6)
	fmt.Println(strList)
	//追加一个切片
	strList = append(strList, []int{7, 8, 9}...)
	fmt.Println(strList)
	//头部追加
	strList = append([]int{11, 12, 13}, strList...)
	fmt.Println(strList)
}

//make()函数构建切片 make( []Type, size, cap )  size:分配元素数  cap:为预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。
func makeSlice() {
	strList := make([]int, 2, 3)
	fmt.Println(strList[:], len(strList))
}

//切片就是一个数组子集  通过下标截取一定位置的元素
func sliceSplit(arr [10]string) {
	//输出从0 到数组长度的所有数据
	fmt.Println(arr[0:len(arr)])
	//省略结束位置 默认数组长度结束  同理省略开始位置
	fmt.Println(arr[0:])
	fmt.Println(arr[:])
	//截取1-5
	fmt.Println(arr[1:5])
}

func normalFor(arr [10]string) {
	//增强for循环遍历
	for index, value := range arr {
		fmt.Printf("arr index %d item %s", index, value)
		fmt.Println()
	}
	//增强for循环  省略index 遍历
	for _, value := range arr {
		fmt.Printf("%s ", value)
	}
	fmt.Println()
}
