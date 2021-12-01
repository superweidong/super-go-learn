package main

import (
	"container/list"
	"fmt"
)

//list为双向链表形式数据结构  list内数据类型没有强制规定  任何类型都可以共存
func main() {

	l := list.New()
	var ele = l.PushBack("B")
	//元素之前插入
	l.PushFront("A")
	//元素之后插入
	l.PushBack([]int{1, 2, 3})

	//指定元素位置的后置、前置插入
	l.InsertAfter("B+", ele)
	l.InsertBefore("B-", ele)

	l2 := list.New()
	l2.PushBack("3")
	l2.PushBack("4")

	//插入一个子list
	l.PushBackList(l2)
	//移除元素
	l.Remove(ele)

	//for循环list
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
