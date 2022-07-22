package main

import "fmt"

func main() {
	numsChan1 := make(chan int, 10)
	//写操作
	go func() {
		for i := 0; i < 50; i++ {
			numsChan1 <- i
			fmt.Println("写入数据：", i)
		}
		fmt.Println("数据全部写完毕，准备关闭管道！")
		close(numsChan1)
	}()
	//遍历管道时，只返回一个值
	//for range是不知道管道是否写完了，所以会一直在这里等待
	//在写入端，将管道关闭
	for v := range numsChan1 {
		fmt.Println("读取数据：", v)
	}
}
