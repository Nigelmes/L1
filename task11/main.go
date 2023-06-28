package main

import "fmt"

func getIntersection(nums1, nums2 []int) []int {
	counter := make(map[int]int) // мапа счетчик для одого из слайса
	var res []int
	for _, num := range nums1 {
		counter[num]++ // добавляем все значения из первого слайса в мапу и подсчитываем дубликаты
	}
	for _, num := range nums2 {
		if count, ok := counter[num]; ok && count > 0 { // итерируемся по второму слайсу и ищем пересечения,
			// могут быть несколько пересечений , смотрим по count
			counter[num] -= 1 // уменьшаем счетчик
			res = append(res, num)
		}
	}
	return res
}

func main() {
	nums1 := []int{2, 9, 44, 555, 1, 67, 0, 743, 88, 88, 88}
	nums2 := []int{88, 157, 96, 33, 2, 45, 555}
	res := getIntersection(nums1, nums2)
	fmt.Println(res)
}
