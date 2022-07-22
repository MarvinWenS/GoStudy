package main

import (
	"fmt"
	"os"
)

//从命令行输出参数，在switch中进行处理

func main() {
	//C: argc,**argv
	//GO：os.Args==>直接可以获取命令输入，是一个字符串切片
	cmds := os.Args
	//os.Args[0] 程序名
	//os.Args[1] 第一个参数
	for key, cmd := range cmds {
		fmt.Println("key:", key, "cmd:", cmd, "cmds len:", len(cmds))
	}
	if len(cmds) < 2 {
		fmt.Println("请正确输入参数！")
	}
	switch cmds[1] {
	case "hello":
		fmt.Println("Hello")
		//go语言中的switch，默认加上break了，不需要手动处理
		//如果想向下穿透的话，那么需要加上关键字:fallthrough
		fallthrough
	case "world":
		fmt.Println("world")

	default:
		fmt.Println("default called")

	}
}
