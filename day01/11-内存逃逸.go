package main

import "fmt"

//本来在栈上的内存，分配至堆中了

//定义一个函数返回一个string类型指针，go语言返回值写在参数后
func testPtr2() *string {
	city := "成都"
	ptr := &city

	name := "Duke"
	p0 := &name
	fmt.Println("p0:", p0, "*p0", *p0)
	return ptr
}

func main() {
	p1 := testPtr2()
	fmt.Println("p1:", p1, "*p1:", *p1)
}
