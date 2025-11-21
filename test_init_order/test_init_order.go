package test_init_order

import "fmt"
import _ "init_project/test_init_order/pkg1"

const pkgNameValue string = "main"

var mainVar = getMainVar()

func init() {
	fmt.Println("init main")
}

func main() {
	fmt.Println("main method invoked")
}

func getMainVar() string {
	fmt.Println("main.pkgNameValue has been initialized")
	return pkgNameValue
}
