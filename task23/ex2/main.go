package main

import "fmt"

// способ №2
func DeleteIdx(str []string, idx int) []string { // принимаем срез и индекс который нужно удалить,
	// возвращаем новый срез
	var res []string
	for i, elem := range str { // итерируемся в цикле по срезу строк
		if i == idx { // если i равняемся индексу (idx) который хотим удалить ,
			// перескакиваем итерацию с помощью continue
			continue
		}
		res = append(res, elem) // добавляем элементы в новый срез
	}
	return res
}

func main() {
	str := []string{"go", "is", "awesome"}
	a := DeleteIdx(str, 0)
	fmt.Println(a)
}
