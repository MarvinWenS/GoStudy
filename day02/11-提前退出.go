package main

import (
	"fmt"
	"runtime"
	"time"
)

//GOEXIT：提前退出go程
//return：返回当前函数
//exit：退出当前进程
func main() {
	go func() {
		go func() {
			func() {
				fmt.Println("这是子go程内部的函数")
				//return //这是返回当前函数 1位置会执行
				//os.Exit(-1) //退出进程 1、2、3不执行
				runtime.Goexit() //退出当前子go程，1不执行，2、3执行

			}()
			fmt.Println("子go程结束")   //1
			fmt.Println("go 22222") //4
		}()
		fmt.Println("111111111111   go 1111111") //5
	}()

	fmt.Println("这是主go程")
	time.Sleep(5 * time.Second) //2
	fmt.Println("Over")         //3
}
