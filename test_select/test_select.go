package test_select

import (
	"context"
	"fmt"
	"time"
)

func pending(ch chan<- int, data int) {
	ch <- data
}

func testSelect() { //select 用于监听多个channel的收/发动作，语法和switch类似，只不过switch和{}之间不能有表达式语句
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)
	ch3 := make(chan int, 3)
	ch_arr := []chan int{ch1, ch2, ch3}

	for i := 0; i < 9; i++ {
		pending(ch_arr[i%3], i)
	}

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	for {
		select {
		case x := <-ch1:
			fmt.Println("ch1 recieve ", x)
		case x := <-ch2:
			fmt.Println("ch2 recieve ", x)
		case x := <-ch3:
			fmt.Println("ch3 recieve ", x)
		case ch1 <- 100:
			fmt.Println("ch1 send 100")
		case ch2 <- 200:
			fmt.Println("ch2 send 200")
		case ch3 <- 300:
			fmt.Println("ch3 send 300")
		case <-ctx.Done(): // 一般配置这样的事件监听context超时异常
			fmt.Println("context timeout")
			if err := ctx.Done(); err != nil {
				fmt.Println("context timeout error:", err)
			}
			return
		default:
			fmt.Println("no channel ready")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func Run() {
	fmt.Println("test_select")
	testSelect()
}
