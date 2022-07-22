package main

import "fmt"

func main() {
	//go语言也有指针
	//结构体成员调用时   c语言 ptr->name  go语言 ptr.name
	//go语言使用指针是会使用内部的垃圾回收机制（GC（grabage collector）机制），开发人员不需要手动释放内存
	//C语言不允许返回栈上指针，go语言可以返回栈上指针，程序会在编译的时候就确定变量的分配位置
	//编译的时候，如果发现有必要的话，就将变量分配到堆上

	name := "lily"
	ptr := &name
	fmt.Println("name:", name)
	fmt.Println("ptr:", ptr)

	//02-使用new关键字定义
	name2ptr := new(string)
	*name2ptr = "Duke"
	fmt.Println("name2ptr:", *name2ptr)
	fmt.Println("name2ptr ptr:", name2ptr)

	//可以返回栈上的指针，编译器在编译程序时，会自动判断这段代码，将city变量分配在堆上,内存逃逸
	res := testPtr()
	fmt.Println("res:", *res, "res add:", res)
	//空指针，在c语言：null c++：nullptr go：nil
	//if语句即使一行代码也必须使用{}
	if res == nil {
		fmt.Println("res是空，nil")
	} else {
		fmt.Println("res非空")
	}
}

//定义一个函数返回一个string类型指针，go语言返回值写在参数后
func testPtr() *string {
	city := "成都"
	ptr := &city
	return ptr
}
