package main

import "fmt"

func main() {
	//切片：slice，它的底层也是数组。可以动态改变长度
	//定义一个切片，包含多个地名
	//names := [10]string{"北京", "上海", "广州", "深圳"}   定长数组
	names := []string{"北京", "上海", "广州", "深圳"} //不定长数组
	for i, v := range names {
		fmt.Println("i:", i, "v:", v)
	}

	//1. 追加数组

	fmt.Println("追加元素前，name的长度：", len(names), ",容量：", cap(names))

	names = append(names, "海南")
	fmt.Println("names", names, "names:", names)

	fmt.Println("追加元素后，name的长度：", len(names), ",容量：", cap(names))
	//2. 对于一个切片，不仅有长度的概念len()，还有一个‘容量’的概念cap()
	nums := []int{}

	for i := 0; i < 50; i++ {
		nums = append(nums, i)
		fmt.Println("追加元素后，nums的长度：", len(nums), ",容量：", cap(nums))
	}

}
