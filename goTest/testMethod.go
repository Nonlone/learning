package main

import (
	"fmt"
	"goTest/main2"
	"reflect"
	"runtime"
	"strings"
)

func testMethodA(){
	fmt.Println("this is method a")
}

func main()  {

	main2.TestMethod2();
	t := main2.TestMethod2
	t()


	fmt.Println(GetPackageName(main))

}


func GetPackageName(temp interface{}) string {
    strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
    strs = strings.Split(strs[len(strs)-2], "/")
    return strs[len(strs)-1]
}