package main

import "fmt"

func main() {
	names := [7]string{"北京", "上海", "广州", "深圳", "洛阳", "南京", "成都"} //不定长数组

	// 想基于names创建一个新的数组
	names1 := [3]string{}
	names1[0] = names[0]
	names1[1] = names[1]
	names1[2] = names[2]
	//1. 切片可以基于一个数组，灵活创建新的数组
	names2 := names[0:3] //左闭右开区间
	fmt.Println("names2:", names2)
	names[2] = "hello"
	fmt.Println("修改names2之后,names2", names2)
	fmt.Println("修改names2之后,names", names)

	//2. 如果从0开始截取，那么冒号左边的数组可以省略
	//3. 如果截取到数组的最后的一个元素，那么冒号右边的数组可以省略
	//4. 如果想从左到右全部使用，那么冒号左右两边都可以省略
	names3 := names[:5]
	fmt.Println("names3:", names3)
	names3 = names[5:]
	fmt.Println("names3:", names3)
	names3 = names[:]
	fmt.Println("names3:", names3)

	//5. 切片也可以基于一个字符串进行切片截取：取字符串的子串 helloworld
	sub1 := "helloworld"[5:7]
	fmt.Println("sub1:", sub1)

	//6. 可以在创建空切片的时候，使用make明确指定切片容量
	str := make([]string, 10, 20) //创建一个容量是20，当前长度为0的string类型切片,第三个参数不是必须，如果没有填写，默认与长度相同
	fmt.Println("str len:", len(str), "cap:", cap(str))
	copy(str, []string{"hellow", "world"})
	fmt.Println("str len:", len(str), "cap:", cap(str), "str:", str)
	//7. 如果想让切片，完全独立与原始数组，可以使用copy函数来完成
	namescopy := make([]string, len(names))
	copy(namescopy, names[:])
	namescopy[0] = "香港"
	fmt.Println("str namecopy :", namescopy)
	fmt.Println("str namecopy len:", len(namescopy), "cap:", cap(namescopy))
}
