package main

import "fmt"

func printtype(elem interface{}) { // функция для печати типа интерфейса
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

func main() {
	arr := []interface{}{777, "Golang", 3.14, true}
	for _, elem := range arr { // итерируемся по слайсу интерфейса и передаем значения в функцию для определенния типа
		printtype(elem)
	}

}
