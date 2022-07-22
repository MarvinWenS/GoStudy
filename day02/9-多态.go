package main

import "fmt"

//实现go多态，需要实现定义接口
//人类的武器发起攻击，不同等级子弹效果不同

//定义接口类型是interface
type IAttack interface {
	//接口的函数可以有多个，但是只能由函数原型，不能有实现
	Attack()
}

//低等级
type HumanLowLevel struct {
	name  string
	level int
}

type HumanHighLevel struct {
	name  string
	level int
}

func (a *HumanLowLevel) Attack() {
	fmt.Println("我是：", a.name, ", 等级为:", a.level, "造成1000点伤害")
}

func (a *HumanHighLevel) Attack() {
	fmt.Println("我是：", a.name, ", 等级为:", a.level, "造成500点伤害")
}

//定义一个多态的通用接口，，传入不同的对象，调用同样的方法，实现不同的效果  ==>多态
func DoAttack(a IAttack) {
	a.Attack()
}
func main() {

	var player IAttack //定义一个包含Attack的接口变量
	lowlevel := HumanLowLevel{
		name:  "David",
		level: 1,
	}
	highlevel := HumanHighLevel{
		name:  "David",
		level: 10,
	}
	lowlevel.Attack()
	//对player赋值为lowlevel，接口需要使用指针类型来赋值
	player = &lowlevel
	player.Attack()

	player = &highlevel
	highlevel.Attack()
	player.Attack()
	fmt.Println("多态.......")
	DoAttack(&lowlevel)
	DoAttack(&highlevel)
}
