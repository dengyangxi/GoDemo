package main

import "fmt"

func main() {

	fmt.Println("=========== 1. 通道取值，通道循环遍历 ==========")
	// 通道取值，通道循环遍历
	{
		//初始化一个通道
		var channel = make(chan int, 13)

		//往通道里面添加值
		channel <- 10
		channel <- 20
		channel <- 30

		//循环往通道里面添加值
		for i := 0; i < 10; i++ {
			//往通道里面添加值
			channel <- i * 10
		}

		//从通道中获取一个值
		var getChannelKey = <-channel
		fmt.Println("通道里面获取的值：", getChannelKey)

		//必须要关闭 管道，才能 range 循环
		close(channel)
		//循环打印通道
		for v := range channel {
			fmt.Println("通道里面获取的值：", v)
		}
	}

	fmt.Println("=========== 2. go协程 通讯 channel==========")
	//  go协程 通讯 channel
	{
		//初始化一个通道
		channel2 := make(chan string, 10)
		channel3 := make(chan string, 10)
		messageList := [...]string{"张三", "李四", "王五"}
		go sendMessage(messageList, channel2)
		go receiveMessage(channel2, channel3)
		//循环遍历 通道元素
		for k := range channel3 {
			//打印通道元素
			fmt.Println(k)
		}
	}

}

//往通道里面添加信息
func sendMessage(message [3]string, channel2 chan string) {
	for i := 0; i < len(message); i++ {
		channel2 <- message[i]
	}
	close(channel2)
}

func receiveMessage(channel2 chan string, channel3 chan string) {
	for v := range channel2 {
		//将通道 channel2 里面的信息取出来，加上 -同学， 存入通道 channel3
		channel3 <- v + "-同学"
	}
	close(channel3)
}
