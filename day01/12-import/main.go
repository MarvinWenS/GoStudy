package main

import "12-import/sub" //sub是文件名，也是包名

func main() {
	res := sub.Sub(2, 2) //包名.函数去调用
	//如果一个包里面的函数向对外提供访问权限，那么一定要首字母大写
	//大写字母开头的函数相当于public
	//小写字母开头的函数相当于private
}
