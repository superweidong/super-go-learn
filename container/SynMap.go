package main

import (
	"fmt"
	"sync"
)

func main() {
	//currentTestMap
	//syncMapTest()
	var syncMap sync.Map
	syncMap.Store("A", 1)
	syncMap.Store("B", 2)
	//syncMapSearch(&syncMap)
	readData, result := syncMap.Load("A")
	fmt.Println(readData, result)
	syncMap.Delete("A")
	fmt.Println(syncMap.Load("A"))
}

func syncMapSearch(syncMap *sync.Map) {
	syncMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

func syncMapTest(syncMap *sync.Map) {
	//currentTestMap()
	go func() {
		for true {
			syncMap.Store("A", 1)
		}

	}()
	go func() {
		for true {
			fmt.Println(syncMap.Load("A"))
		}
	}()
	select {}
}

//对非同步map测试
func currentTestMap() {
	map1 := make(map[string]int)
	go func() {
		for true {
			map1["A"] = 1
		}

	}()
	go func() {
		for true {
			_ = map1["A"]
		}
	}()
	select {}
}
