package main

import "fmt"

func main() {
	//变量定义：var
	//常量定义：const

	//01-先定义变量再赋值，var 变量名 数据类型
	var name string
	name = "wenzhao"
	var age int
	age = 10
	fmt.Println(name)
	fmt.Printf("name is :%s, %d", name, age) //跟C语言printf类似
	//02-定义时直接赋值
	var gender = "man"
	var gender1 string = "woman"
	fmt.Println("gender:", gender)
	fmt.Println("gender:", gender1)
	//03-定义直接赋值，使用自动推导（最常用）
	address := "北京"
	fmt.Println("address:", address)
	test(age, address)
	//04 - 平行赋值
	i, j := 10, 20 //同时定义两个变量
	fmt.Println(i, j)
	i, j = j, i //快速交换 i,j
	fmt.Println(i, j)
}
func test(a int, b string) {
	fmt.Println(a)
	fmt.Println(b)
}
