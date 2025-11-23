package test_goroutine

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func simple() {
	go func() {
		fmt.Println("run goroutine in closure")
	}()
	go func(string) {
	}("gorouine: closure params")
	go say("in goroutine: world")
	say("hello")
}

type Counter interface {
	Increment()
	GetCount() int
}

type UnsafeCounter struct {
	count int
}

// 增加计数
func (c *UnsafeCounter) Increment() {
	c.count += 1
}

// 获取当前计数
func (c *UnsafeCounter) GetCount() int {
	return c.count
}

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() // 避免因代码修改或异常导致的锁未释放问题。 如果在
	c.count += 1
}

func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func test_counter() {
	safeCoun := SafeCounter{}
	//unsafeCoun := UnsafeCounter{}

	coun := &safeCoun // var coun Counter = &unsafeCoun

	// 启动100个goroutine同时增加计数
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				coun.Increment()
			}
		}()
	}

	// 等待一段时间确保所有goroutine完成
	time.Sleep(time.Second)
	// 输出最终计数
	fmt.Printf("Final count: %d\n", coun.GetCount())
}

func Run() {
	fmt.Println("test_goroutine")
	//simple()
	test_counter()
}
