package main

import (
	"fmt"
	"sync"
)

func MySquare(num int, wg *sync.WaitGroup) {
	defer wg.Done() // сброс счетчика
	fmt.Printf("Square [%d] = %d\n", num, num*num)
}

func main() {
	wg := new(sync.WaitGroup) //тип позволяет определить группу горутин, которые должны выполняться вместе как одна группа
	nums := []int{2, 4, 6, 8, 10}
	for _, num := range nums {
		wg.Add(1) // счетчик горутин
		go MySquare(num, wg)
	}
	wg.Wait() // ожидания выполнения всех горутин в группе
}
