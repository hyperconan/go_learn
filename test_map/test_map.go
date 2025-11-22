package test_map

import (
	"fmt"
	"sync"
	"time"
)

func basicUsage() {
	var m = make(map[int]string)
	m[1] = "a"
	m[2] = "b"
	fmt.Println(m)

	var m2 = make(map[int]string, 2) //声明大小，减少扩容带来的性能损耗
	m2[1] = "a"
	m2[2] = "b"
	m2[3] = "c" // 发生扩容

	value, ok := m2[4]     //读取不存在的key
	fmt.Println(value, ok) // 空,false

	delete(m2, 1)   // 删除key为1的元素
	fmt.Println(m2) // map[2:b 3:c]
}

func concurrentUsage() {
	// map 非并发安全，需要加锁
	m := map[int]string{
		1: "a",
		2: "b",
	}
	// 并发安全
	var lock sync.Mutex

	go func() {
		lock.Lock()
		m[3] = "c"
		fmt.Println(m)
		lock.Unlock()
	}()

	go func() {
		lock.Lock()
		m[3] = "d"
		fmt.Println(m)
		lock.Unlock()
	}()

	//休眠2秒
	select {
	case <-time.After(time.Millisecond * 200):
		fmt.Println(m)
	}
}

func Run() {
	fmt.Println("test_map")
	//basicUsage()
	concurrentUsage()
}
