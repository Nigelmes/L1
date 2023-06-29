package main

import (
	"fmt"
	"strings"
)

func UniqueString(str string) bool { // проверка строки на уникальные символы
	counter := make(map[string]struct{}) // мапа для выеявления дубликатов
	for _, elem := range str {           // посимвольно итерируемся по строке
		if _, ok := counter[strings.ToLower(string(elem))]; ok { // проверям есть ли уже такое значение в мапе,
			// если есть возвращаем false
			return false
		}
		counter[strings.ToLower(string(
			elem))] = struct{}{} // все символы приводим в один регистр (в данном случае в нижний),
		// дабы избежать дубликатов в другом регистре
	}
	return true
}

func main() {
	arr := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, str := range arr {
		fmt.Println(str, "---", UniqueString(str))
	}
}
