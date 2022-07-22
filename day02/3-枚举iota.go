package main

import "fmt"

//在go语言中没有枚举类型，但是可以使用const+iota（常量累加器）来进行模拟
//模拟表示一周的枚举
const (
	MONDAY    = iota //0
	TUESDAY   = iota //1
	WEDNESDAY = iota //2
	THURSDAY  = iota //3
	FRIDAY           //4 没有赋值，默认与上一行相同iota
	SATURDAY
	SUNDAY
	M, N = iota, iota //xonst属于预编译赋值，所以不需要:=进行自动推导
)

const (
	Jan  = iota + 1 //1
	Janu = iota + 1 //2
	Mar
	May
)

/*
1. iota 是常量组计数器
2. iota从0开始，每换行递增1
3. 常量组有个特点如果不赋值，默认与上一行表达式相同
4. 如果同一行出现了两个iota，那么两个iota的值是相同的
5. 每个常量组的iota是独立的，如果遇到const iota会重新清零
*/
func main() {
	//var number int
	//var name string
	//var flag bool
	//
	////可以使用变量组来将统一定义变量
	//var (
	//	number int
	//	name string
	//	flag bool
	//)
	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(WEDNESDAY)
	fmt.Println(THURSDAY)
	fmt.Println(FRIDAY)
	fmt.Println(SATURDAY)
	fmt.Println(SUNDAY)

	fmt.Println(M, N)

	fmt.Println(Jan)
	fmt.Println(Janu)
	fmt.Println(Mar)
	fmt.Println(May)
}
