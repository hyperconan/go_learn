package test_range

import (
	"fmt"
	"reflect"
)

func test_range_str() {
	s1 := "你好a"
	for idx, v := range s1 {
		fmt.Printf("idx:%d value:%d char:%c \n", idx, s1[idx], s1[idx])
		fmt.Println(reflect.TypeOf(s1[idx]))                                                   // s1[idx] 获取到byte值
		fmt.Printf("idx:%d value:%d char:%c byte len:%d\n", idx, v, v, len([]byte(string(v)))) // v 获取到rune值
		fmt.Println(reflect.TypeOf(v), "\n")                                                   // v 获取到rune值
	}
}

func test_range_arr_slice() {
	arr := [...]int{1, 2, 3}
	sli := []int{1, 2, 3}
	for idx, v := range arr {
		fmt.Printf("idx:%d value:%d \n", idx, v)
	}
	for idx, v := range sli {
		fmt.Printf("idx:%d value:%d \n", idx, v)
	}
}

func producer(ch chan int) {
	i := 1
	capSize := cap(ch)
	for i < capSize {
		ch <- i
		i++
	}
	close(ch)
}

func test_range_channel() {
	ch := make(chan int, 10) // 创建拥有10个容量的channel
	go producer(ch)          // 启动一个goroutine生产数据
	for v := range ch {      // channel没有索引，遍历时只能拿到channel的每一个值
		fmt.Printf("value:%d \n", v)
	}
}

func Run() {
	//test_range_str()
	//test_range_arr_slice()
	test_range_channel()
}
