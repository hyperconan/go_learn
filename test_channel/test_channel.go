package test_channel

import (
	"fmt"
	"time"
)

// channel 的三种操作方式 : 接收channel <- {data}，发送{data} <-channel，关闭close(channel)

func recieve(ch <-chan int) { // <-chan 限定了channel只能作recieve模式，recieve from(<-) channel也就是从channel中发送数据出去, {data} <-ch
	for v := range ch {
		fmt.Println("recieve:", v)
	}
}

func send(ch chan<- int) { // chan<- 限定了channel只能作send模式，send to(->) channel也就是send数据到channel ch<-{data}
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("send:", i)
	}
	close(ch)
}

func test_init_channel() {
	//var ch channel int -> 仅是声明，值为nil
	//ch := make(chan int, 3) // 创建一个有3个缓冲空间的channel
	ch := make(chan int) // 创建一个无缓冲的channel
	go send(ch)
	go recieve(ch)

	timer := time.After(1 * time.Second) // 每1秒触发一次 timer也是一个channel
	for {
		select { // 多路并发复用
		case v, ok := <-ch: // 当channel是关闭的 ok=false
			if ok {
				fmt.Println("select recieve:", v)
			} else {
				fmt.Println("channel closed")
				return
			}
		case <-timer:
			fmt.Println("超时")
			return
		default:
			fmt.Println("等待数据")
		}
	}
}

func Run() {
	fmt.Print("test_channel")
	test_init_channel()
}
