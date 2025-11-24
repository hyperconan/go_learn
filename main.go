package main

import (
	_ "init_project/test_array_slice"
	_ "init_project/test_channel"
	_ "init_project/test_closure"
	_ "init_project/test_for"
	_ "init_project/test_goroutine"
	_ "init_project/test_if_switch"
	//"init_project/test_select"
	"init_project/test_select"

	_ "init_project/test_interface"
	_ "init_project/test_map"
	_ "init_project/test_range"
	_ "init_project/test_type_convert"

	//_ "init_project/test_init_order"
	_ "init_project/test_pointer"
	_ "init_project/test_struct"
	_ "init_project/test_type"
)

func main() {
	//test_pointer.Run()
	//test_struct.Run()
	//test_type.Run()
	//test_if_switch.Run()
	//test_for.Run()
	//test_closure.Run()
	//test_array_slice.Run()
	//test_map.Run()
	//test_range.Run()
	//test_type_convert.Run()
	//test_interface.Run()
	//test_init_order.Run()
	//test_goroutine.Run()
	//test_channel.Run()
	test_select.Run()
}
