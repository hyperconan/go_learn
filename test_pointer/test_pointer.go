package test_pointer

import (
	"fmt"
	"unsafe"
)

func pointer_test() {
	var a uint8 = 15
	var a_pt *uint8 = &a
	var a_ppt **uint8 = &a_pt
	// a_pt 指向 a 的内存地址
	fmt.Println(*a_pt)
	fmt.Println(*a_ppt)
	fmt.Println(**a_ppt)
}

func pointer_str_test() {
	var a string = "hello world"
	var p *string = &a
	fmt.Println(*p)
}

func pointer_test_multi() {
	a := 2
	var p *int
	fmt.Println(&a) // 地址1
	p = &a
	fmt.Println(p, &a) // 地址1,地址1

	var pp **int
	pp = &p            //地址2
	fmt.Println(pp, p) //地址2，地址1
	**pp = 3
	fmt.Println(pp, *pp, p) //地址2，地址1，地址1
	fmt.Println(**pp, *p)   //3，3
	fmt.Println(a, &a)      // 3, 地址1
}

func pointer_move_test() {
	var a string = "Helloworld"
	up_A := uintptr(unsafe.Pointer(&a))
	up_A += 1
	var a2b []byte = []byte(a)
	fmt.Println(len(a2b))
	c := (*uint8)(unsafe.Pointer(up_A))
	fmt.Println(c)
	fmt.Println(*c)
}

func Run() {
	fmt.Println("test_pointer")
	pointer_test()
	pointer_str_test()
	pointer_test_multi()
	pointer_move_test()
}
