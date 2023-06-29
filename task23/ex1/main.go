package main

import "fmt"

// способ №1
func DeleteIdx(str []string, idx int) []string { // принимаем срез и индекс который нужно удалить,
	// возвращаем новый срез
	str = append(str[:idx], str[idx+1:]...) // нарезаем срез append ом
	return str
}

func main() {
	str := []string{"go", "is", "awesome"}
	a := DeleteIdx(str, 2)
	fmt.Println(a)
}
