package main

import "fmt"

func main() {
	//标签 label1
	//goto label1 ===>下次进入循环时，i不会保存之前的状态，重新从0开始计算，重新来过
	//break label1 ===>continue 回跳到指定位置，但是会记录当前i的状态
	//continue label1 ===>直接跳出指定位置
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				//goto LABEL1
				//continue LABEL1
				break LABEL1
			}

			fmt.Println("i,j:", i, j)
		}
	}
	fmt.Println("Over!")
}
