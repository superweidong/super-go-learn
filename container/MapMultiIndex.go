package main

import "fmt"

type KeyVo struct {
	Name string
	Age  int
}

type ValueVo struct {
	Name    string
	Age     int
	Address string
}

//定义map
var mapper = make(map[KeyVo]*ValueVo)

//普通查询
func normalSearch(list []*ValueVo, name string, age int) {
	for _, vo := range list {
		if vo.Name == name && vo.Age == age {
			fmt.Println(vo)
			return
		}
	}
	fmt.Println("未找到元素")
}

//构建索引
func buildIndex(list []*ValueVo) {
	for _, vo := range list {
		key := KeyVo{Name: vo.Name, Age: vo.Age}
		mapper[key] = vo
	}
}

//索引查询
func indexSearch(name string, age int) {
	key := KeyVo{Name: name, Age: age}
	fmt.Println(mapper[key])
}

func main() {
	var list = []*ValueVo{{Name: "张三", Age: 10, Address: "北京"}, {Name: "李四", Age: 10, Address: "北京"}, {Name: "王五"}}
	normalSearch(list, "张三", 10)
	buildIndex(list)
	indexSearch("李四", 10)
	indexSearch("王五", 10)
}
