package packagex

import (
	"reflect"
	"runtime"
	"strings"
)



func GetPackageName() string {
	return GetPackageNameByFunc(func() {})
}

func GetPackageNameByFunc(temp interface{}) string {
	return getFuncPath(temp, func(length int) int { return 0 }, func(length int) int {
		if length > 0 {
			return length - 1
		}
		return length
	})
}

func getFuncPath(temp interface{}, nameLengthFunc func(length int) int, namePathLengthFunc func(length int) int) string {
	funcName := runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()
	funcNameArr := strings.Split(funcName, ".")
	funcNamePathArr := strings.Split(funcNameArr[nameLengthFunc(len(funcNameArr))], "/")
	return funcNamePathArr[namePathLengthFunc(len(funcNamePathArr))]
}
