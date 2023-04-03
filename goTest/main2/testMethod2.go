package main2

import (
	"fmt"
	. "goTest/helper"
	"goTest/helper/packagex"
)


func TestMethod2(){
	fmt.Println("this is method2")
	fmt.Println(packagex.GetPackageNameByFunc(EMPTY_FUNC))
}

