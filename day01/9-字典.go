package main

import "fmt"

func main() {
	//1. 定义一个字典
	//学生id==>学生 学生名字 idNames
	var idNames map[int]string //定义一个map，此时map是不能直接赋值的，它是空的

	//2.分配空间,使用make，可以不指定长度，但是建议直接指定长度，性能更好
	//idNames = make(map[int]string) //可以使用
	idNames = make(map[int]string, 10)

	//3. 定义时，直接分配空间
	idNames1 := make(map[int]string, 10) //这是最常用的

	//4. 遍历map
	idNames[0] = "duke"
	idNames[1] = "david"
	for key, value := range idNames {
		fmt.Println("key:", key, ",value:", value)
	}

	idNames1[0] = "duke"
	idNames1[1] = "david"
	for key, value := range idNames1 {
		fmt.Println("key:", key, ",value:", value)
	}

	//5. 如何确定一个key是否存在map中
	//在map中不存在访问越界问题，他认为所有的key都是有效的，所以访问一个不存在的key不会崩溃，返回这个类型的零值
	//零值：bool false 数组 0 字符串 空
	name3 := idNames[9]
	fmt.Println("name3:", name3)

	//无法通过获取value来判断一个key是否存在，因此我们需要一个能够校验key是否存在的方式
	value, ok := idNames[10] //如果id=1是存在的，那么value就是对应key=1对应值，ok返回true，反之返回零值，ok返回false
	if ok {
		fmt.Println("id=10是存在的，value为：", value)
	} else {
		fmt.Println("id=10是不存在，value为：", value)
	}

	//6. 删除map中的元素
	//使用自由函数delete来删除指定的key
	fmt.Println("idNames删除前：", idNames)
	delete(idNames, 1)
	fmt.Println("idNames删除后：", idNames)
	delete(idNames, 1)
	fmt.Println("idNames删除无效的key：", idNames)
}
