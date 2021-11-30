package main

import "fmt"

func main() {

	myMap := createMap()
	fmt.Println(myMap)

	myMap2 := makeCreateMap()
	fmt.Println(myMap2, len(myMap2))

	mapWithSlice := mapWithSliceValue()
	fmt.Println(mapWithSlice)

}

//value为切片类型的map
func mapWithSliceValue() map[string][]int {
	map1 := make(map[string][]int, 10)
	map1["slice"] = []int{1, 2, 3}
	return map1
}

func createMap() map[string]int {
	//构建一个map var mapName = map[keyType]valueType
	var myMap = map[string]int{"A": 1, "B": 2}
	//打印map
	fmt.Println(myMap)
	//根据key取出对应value
	fmt.Println(myMap["A"])
	//for循环遍历map
	for s, i := range myMap {
		fmt.Printf("key %s value %d", s, i)
		fmt.Println()
	}
	return myMap
}

//使用make()函数构建map
func makeCreateMap() map[string]int {
	//使用make()构建map  100为map长度
	var myMap2 = make(map[string]int, 100)
	myMap2["k1"] = 1
	myMap2["k2"] = 2
	return myMap2
}
