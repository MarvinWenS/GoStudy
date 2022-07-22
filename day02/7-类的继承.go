package main

import "fmt"

//在go语言中，权限都是通过首字母的大小写来控制
//1. import ==> 如果包名不同，那么只有大写字母开头的才是public的
//2. 对于类的成员、方法==> 只有大写开头的才能在其他包中使用
type Human struct {
	//成员属性
	name   string
	age    int
	gender string
}

//在类外面绑定方法
func (this *Human) Eat() {
	fmt.Println("this is :", this.name)
}

//定义一个学生类，去嵌套一个Human
type Student1 struct {
	hum    Human //包含Human类型的变量,此时属于类的嵌套
	school string
	score  float64
}

//定义一个老师，去继承Human
type Teacher struct {
	Human          //直接写Human类型，没有字段名字
	subject string //学科
}

func main() {
	s1 := Student1{
		hum: Human{
			name:   "Liyi",
			age:    18,
			gender: "女生",
		},
		school: "昌平一中",
		score:  0,
	}
	fmt.Println("s1.name:", s1.hum.name)
	fmt.Println("s1.school:", s1.school)

	t1 := Teacher{}
	t1.subject = "语文"
	t1.name = "荣老师" //下面几个字段都是继承字段
	t1.age = 35
	t1.gender = "女生"
	fmt.Println("t1:", t1)
	t1.Eat()

	//继承的时候，虽然我们没有定义字段名字，但是会自动创建一个默认的同名字段
	//这是为了在子类中依然可以操作父类，因为：子类父类可能出现同名的字段或函数
	fmt.Println("t1.Human.name:", t1.Human.name)
}
