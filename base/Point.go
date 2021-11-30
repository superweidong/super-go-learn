package main

import "fmt"
//指针
func main() {
	//定义字符串
	str:= "hello"
	//&符取地址值  赋值给指针
	var address *string = &str
	fmt.Printf("%T %[1]v", address)
	//*address  取出指针对应的值
	strTmp := *address
	println("------> point address ",address)
	println("------> point address value ", strTmp)

	//指针实现值交换
	x,y :="A","B"
	println("------>swap value before", x, y)
	swapValue(&x, &y)
	println("------>swap value after", x, y)

	//指针实现地址值交换
	m,n :="M","N"
	println("------>out swap addr before", &m, &n)
	println("------>out swap addr before", m, n)
	swapAddr(&m, &n)
	println("------>out swap addr after", &m, &n)
	println("------>out swap addr after", m, n)

	println(dummy(9))

}

// 本函数测试入口参数和返回值情况
func dummy(b int) int {
	// 声明一个变量c并赋值
	var c int
	c = b
	return c
}

//地址交换 地址交换只在方法内部 因为传参方式类似java 传递的是地址值的临时变量
func swapAddr(m,n *string)  {
	println("------>inner swap before", m, n)
	//将m的地址赋值给tmp
	tmp := m
	//将n的地址赋值给m
	m = n
	n = tmp
	println("------>inner swap after", m, n)
	//2.地址交换简单写法
	m,n = n,m
}
//值交换
func swapValue(x,y *string)  {
	//x,y 形参均为指针类型  *x取出值  赋值给临时变量tmp
	tmp := *x
	//将y的值赋值给x
	*x = *y
	//将tmp的值赋值给y
	*y = tmp
}

