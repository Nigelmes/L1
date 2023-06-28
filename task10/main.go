package main

import "fmt"

func main() {
	temperature := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := make(map[int][]float64) // мапа в качестве ключа будет шаг в 10, и слайс вещественных чисел

	for _, t := range temperature {
		res[(int(t)/10)*10] = append(res[(int(t)/10)*10],
			t) // вычисляем шаг по формуле и добавляем в слайс значения которое удовлетворят этому шагу
	}
	fmt.Println(res) // printl выведет отсортированную мапу в порядке возрастания
}
