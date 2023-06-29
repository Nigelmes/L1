package main

import (
	"fmt"
	"strings"
)

func reversefunc(arr string) string {
	res := strings.Fields(arr) // разбиваем строку на массив строк
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i] // в цикле переворачиваем элементы
	}
	return strings.Join(res, " ") // собираем обратно массив строк в строку и возвращаем
}

func main() {
	arr := "snow dog sun"
	fmt.Println(reversefunc(arr))
}
