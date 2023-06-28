package main

import "fmt"

func mySet(arr ...string) map[string]int {
	res := make(map[string]int)
	for _, elem := range arr { // итерируемся по слайсу и добавляем элементы в мапу
		res[elem]++
	}
	return res
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	res := mySet(arr...)
	fmt.Println(res)
}
