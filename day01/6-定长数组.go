package main

import "fmt"

func main() {
	//1-定义，定义一个具有10个数组的数组
	//C语言定义：int num[10]={1,2,3,4}
	//go语言定义：
	//	num:=[10]int{1,2,3,4}(常用方式)
	//  var nums = [10]int{1, 2, 3, 4}
	//  var nums [10]int=[10]{1, 2, 3, 4}
	nums := [10]int{1, 2, 3, 4}
	//2-遍历 方式一
	for i := 0; i < len(nums); i++ {
		fmt.Println("i:", i, "num[i]:", nums[i])
	}
	//2-遍历 方式二 for range key:数组下标 value：数组的值
	for key, value := range nums {
		//key=0     value=1  ==>  nums[0]  value是nums[key]的一个副本
		//value全程只是一个变量，不断被重新赋值，修改它不会修改原始数据
		fmt.Println("key:", key, "value:", value, "&value:", &value, "&nums[key]", &nums[key])
	}

	//go语言中，如果想忽略key或者value，可以使用_ ,如果全部忽略 不能使用:= 直接使用=
	for _, value := range nums {
		fmt.Println("忽略key，value:", value)
	}
}
