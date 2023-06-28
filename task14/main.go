package main

import "fmt"

func printtype(arr ...interface{}) {
	for _, elem := range arr { // итерируемся по переданным элементам
		switch elem.(type) { // ключевым словом (type) получаем тип элемента
		case int:
			fmt.Printf("[%v] - int\n", elem)
		case float64:
			fmt.Printf("[%v] - float\n", elem) //принтуем тип
		case string:
			fmt.Printf("[%v] - string\n", elem)
		case bool:
			fmt.Printf("[%v] - bool\n", elem)
		}
	}
}

func main() {
	arr := []interface{}{777, "Golang", 3.14, true}
	printtype(arr...)
}
