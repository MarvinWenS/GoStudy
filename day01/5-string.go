package main

import "fmt"

func main() {
	name := "duke"
	//原生输出字符串需要换行时，使用反引号``
	usage := `./a.out <option>
-h help
-a xxxx`
	fmt.Println("name:", name)
	fmt.Println("usage:", usage)
	//2.长度，访问
	//C++: name.length
	//GO:string没有.length方法/属性，可以使用自由函数len()进行处理
	//len:很常用
	len_name := len(name)
	fmt.Println("len_name:", len_name)
	//for循环不需要加（）
	for i := 0; i < len(name); i++ {
		fmt.Printf("i:%d,v:%c\n", i, name[i])
	}
	//3.字符串拼接
	i, j := "hello", "world"
	fmt.Println("i+j=", i+j)
	//使用const声明常量不能修改
	const add = "chengdu"
	//add="beijing" //不能修改
	fmt.Println("add:", add)

}
