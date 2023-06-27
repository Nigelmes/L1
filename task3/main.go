package main

import "fmt"

// вариант №1
func MySquare(out chan<- int, num int) { // функция принимает число и отправляет его квадрат в канал out (
	// канал out только для записи)
	out <- num * num
}
func main() {
	var sum int
	nums := []int{2, 4, 6, 8, 10}
	square := make(chan int, len(nums)) //  буферизованный канал

	for _, num := range nums { // в цикле запускаем функцию в горутине
		go MySquare(square, num)
	}
	for i := 0; i < len(nums); i++ {
		sum += <-square // складываем квадраты чисел в переменную sum высчитывая из буф канала
	}
	fmt.Printf("Sum of squares of numbers - [%d]\n", sum)
	// канал в этой программе можем не закрывать , потому что лишнего ничего не читаем
}

// вариант №2
//func MySquare(num int) int {
//	return num * num
//}
//
//func main() {
//	var sum int
//	wg := new(sync.WaitGroup)
//	nums := []int{2, 4, 6, 8, 10}
//	square := make(chan int, len(nums))
//
//	for _, num := range nums {
//		wg.Add(1)
//		go func(num int) { //в анонимной горутине записываем в канал результат выполнения функции MySquare
//			defer wg.Done()
//			square <- MySquare(num)
//		}(num)
//	}
//	go func() { // в отдельной горутине запускаем ожидание отработки всех горутин чтобы закрыть канал square
//		// и для того чтобы не блокировать main в ожидании
//		wg.Wait()
//		close(square)
//	}()
//
//	for a := range square { // for range автоматически считывает очередное значение из канала и проверяет, не закрыт ли тот. Если закрыт — выходит из цикла
//		sum += a
//	}
//	fmt.Printf("Sum of squares of numbers - [%d]\n", sum)
//}
