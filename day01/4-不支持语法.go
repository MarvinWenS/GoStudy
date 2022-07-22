package main

import "fmt"

func main() {
	//不能这么表示
	//if 0 {
	//	fmt.Println("hello world")
	//}
	if true {
		fmt.Println("hello world")
	}
	//不能这么表示
	//if nil {
	//	fmt.Println("hello world")
	//}
	if false {
		fmt.Println("hello world")
	}
	//三目运算符不支持
	//a,b:=1,2
	//x:=a>b?0,-1
}
