package main

import (
	"fmt"
	"unicode/utf8"
)

func reversefunc(arr string) string {
	lenght := utf8.RuneCountInString(arr) // получаем количество символов в строке,
	// len() - возвращает кол-во байт в строке, поэтому на не подходит
	res := []rune(arr) // создаем срез рун из строки
	for i := 0; i < lenght/2; i++ {
		res[i], res[lenght-1-i] = res[lenght-1-i], res[i] // в цикле переворачиваем элементы
	}
	return string(res) // возвращаем в виде строки
}

func main() {
	arr := "главрыба КОТ"
	fmt.Println(reversefunc(arr))
}
