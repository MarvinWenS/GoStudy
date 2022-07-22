package main

import "fmt"

//Person类，绑定方法：Eat Run Laugh
//public private
/* C++语言
class Person{
public:
	string name;
	int age;
public:
	Eat(){
	    xxx
	}
}
*/
type Person struct {
	//成员属性
	name   string
	age    int
	gender string
	score  float64
}

//在类外面绑定方法
func (this *Person) Eat() {
	//fmt.Println("Person is eating")
	////类的方法可以使用自己的成员
	//fmt.Println("Name:", this.name, " is eating!")
	this.name = "Duke"
}

func (this Person) Eat2() {
	//fmt.Println("Person is eating")
	this.name = "liyi"
}
func main() {
	liyi := Person{
		name:   "Liyi",
		age:    30,
		gender: "女生",
		score:  10,
	}

	fmt.Println("Eat,使用*p Person,用指针修改name值...")
	fmt.Println("修改前liyi:", liyi) //liyi
	liyi.Eat()
	fmt.Println("修改后liyi:", liyi) //Duke
	fmt.Println("Eat2,使用p Person,但是不是指针修改name值...")
	fmt.Println("修改前liyi:", liyi) //Duke
	liyi.Eat2()
	fmt.Println("修改后liyi:", liyi) //Duke
}
