//包名：每个go程序 都必须有一个包名 package main必须要有
//每个go程序，都是.go结尾的
//go程序中没有.h,没有.o，只有.go
//一个package,包名,相当于命名空间
package main

import "fmt"

//主函数，所有的函数必须使用func开头
//一个函数的返回值，不会放在func前，而是放在参数后面
//函数{必须与函数名同行，不能换行
func main() {
	//go语句不需要使用分号结尾
	fmt.Println("hello world")
}
