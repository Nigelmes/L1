// var justString string
//
//	func someFunc() {
//		v := createHugeString(1 << 10)
//		justString = v[:100]
//	}
//
//	func main() {
//		someFunc()
//	}
package main

import (
	"fmt"
	"strings"
)

func createHugeString(lenght int) string {
	var str strings.Builder
	for i := 0; i < lenght; i++ {
		str.WriteString("Ф")
	}
	return str.String()
}

func someFunc() string { // из функции возвращаем строку вместо глобальной переменной
	//v := createHugeString(1 << 10)
	//return  v[:100]      могут возникнуть проблемы при нарезке слайса,
	//если в строке содержаться символы utf8 то они могут занимать несколько байт
	return createHugeString(100) // вместо создания большой строки и ее нарезки сразу создаем строку нужного размера
}

func main() {
	var justString string   // глобальную переменную делаем локальной
	justString = someFunc() //
	fmt.Println(justString)
}
