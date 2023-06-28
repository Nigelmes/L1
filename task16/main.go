package main

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) < 2 { // базовый случай для выхода из рекурсии
		return arr
	}
	left, right := 0, len(arr)-1 // индексы границ массива
	pivot := (left + right) / 2  //опорный индекс
	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left])   //рекурсивный вызов левой части
	QuickSort(arr[left+1:]) // рекурсивный вызов правой части

	return arr
}

func main() {
	arr := []int{5, 6, 11, 1, 4}
	fmt.Println(QuickSort(arr))

}
