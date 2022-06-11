package main

import (
	"fmt"
	"time"
)

func main() {

	//初始化 ，并不会执行 匿名函数
	closeFn := closurePackageFn(18, "张三")

	//执行函数
	year, name := closeFn()

	fmt.Printf("姓名是：[%v], 出生年份是：[%v]", year, name)
}

func closurePackageFn(age int, name string) func() (int, string) {
	//获取当前年份  - 传入 的岁数 = 出生年份
	year := time.Now().Year() - age

	//闭包函数
	return func() (int, string) {
		fmt.Println("姓名：", name, "年龄：", age, "出生年份", year)
		//返回 多个数据， 使用 , 号 隔开
		return year, name
	}
}
