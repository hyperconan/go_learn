package pkg2

import "fmt"

const pkgName string = "pkg2"

var pkgNameValue string = getPkgName()

func init() {
	fmt.Println("init pkg2")
}

func getPkgName() string {
	fmt.Println("pkg2.pkgNameValue has been initialized")
	return pkgName
}
