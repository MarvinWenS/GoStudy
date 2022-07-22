package main

import (
	"fmt"
	"time"
)

func main() {
	//1. 当缓存写满的时候，写阻塞，当被读取后，再恢复写入
	//2. 当缓冲区读取完毕，读阻塞
	//3. 如果管道没有使用make分配空间，那么管道默认是nil的，读取、写入都会阻塞
	//4. 对于管道，读与写的次数对等

	//numsChan := make(chan int, 10)
	var name chan string //默认时nil的
	name = make(chan string, 10)
	go func() {
		fmt.Println("names:", <-name)
	}()
	name <- "hello" //由于name是nil的，写操作会阻塞在这里
	time.Sleep(1 * time.Second)

	numsChan1 := make(chan int, 10)
	//写操作
	go func() {
		for i := 0; i < 50; i++ {
			numsChan1 <- i
			fmt.Println("写入数据：", i)
		}
	}()

	//读
	func() {
		for i := 0; i < 60; i++ {
			data := <-numsChan1
			fmt.Println("读取数据：", data)
		}
	}()
}
