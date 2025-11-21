package pkg1

import "fmt"
import _ "init_project/test_init_order/pkg2"

const pkgName string = "pkg1"

var pkgNameValue string = getPkgName()

func init() {
	fmt.Println("init pkg1 222")
}
func init() {
	fmt.Println("init pkg1")
}

func getPkgName() string {
	fmt.Println("pkg1.pkgNameValue has been initialized")
	return pkgName
}
