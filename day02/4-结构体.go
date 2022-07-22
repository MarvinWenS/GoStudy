package main

import "fmt"

//c语言里面，我们可以使用typedef int MyInt
type MyInt int

////c语言结构体定义
//struct Person{
//
//}

//go语言结构体使用type+ struct来处理

type Student struct {
	name   string
	age    int
	gender string
	score  float64
}

func main() {
	var i, j MyInt
	i, j = 10, 20
	fmt.Println("i,j", i, j)
	//创建变量，并赋值
	liyi := Student{
		name:   "liyi",
		age:    20,
		gender: "女生",
		score:  8, //最后一个元素后面必须加上逗号，如果不加逗号则必须与}同一行
	}
	//使用结构体各个字段0
	fmt.Println("liyi:", liyi.name, liyi.age, liyi.gender, liyi.score)
	s1 := &liyi
	fmt.Println("liyi 使用指针打印:", s1.name, s1.age, s1.gender, s1.score)
	fmt.Println("liyi 使用指针（*s1）.name打印:", (*s1).name, (*s1).age, (*s1).gender, (*s1).score)
	//如果在定义期间对结构体赋值时，如果每个字段都赋值了，那么字段的名字可以省略不写
	//如果只对局部变量赋值，那么必须明确指定变量名字
	Duke := Student{
		name: "Duke",
		age:  28,
		//gender: "",
		//score:  0,
	}

	fmt.Println("Duke:", Duke)
}
