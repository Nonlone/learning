package main

import (
	"fmt"
	. "goTest/helper"
	"goTest/helper/packagex"
	"goTest/main2"
)

func testMethodA(){
	fmt.Println("this is method a")
}

func main()  {

	main2.TestMethod2();
	t := main2.TestMethod2
	t()


	fmt.Println(packagex.GetPackageNameByFunc(EMPTY_FUNC))

}

