package main

import "fmt"

//1. 函数返回值在参数列表之后
//2. 如果有多个返回值，需要使用圆括号包裹，多个参数之间使用,分割
func testFunc(a int, b int, c string) (int, string, bool) {

	return a + b, c, true
}

func testFunc1(a, b int, c string) (res int, str string, b1 bool) {

	//直接使用返回值的变量名字参与运算
	res = a + b
	str = c
	b1 = true
	//当返回值有名字的时候，可以直接简写return
	return
}

//如果返回值只有一个参数，并且没有名字，那么不需要加圆括号
func testFunc2() int {
	return 10

}
func main() {
	v1, s1, _ := testFunc(10, 20, "hello")
	fmt.Println("v1:", v1, "s1:", s1)

	v2, s2, _ := testFunc1(10, 10, "hello")
	fmt.Println("v2:", v2, "s2:", s2)

	v3 := testFunc2()
	fmt.Println("v3:", v3)
}
