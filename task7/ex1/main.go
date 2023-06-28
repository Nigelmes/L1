// Пример конкурентной запси в map с помощью sync.Mutex
package main

import (
	"fmt"
	"sync"
)

func count(mx *sync.Mutex, word string, counter map[string]int, wg *sync.WaitGroup) {
	defer wg.Done() // сброс счетчика
	for i := 0; i < len(word); i++ {
		mx.Lock()       // запираем мьютекс для записи, остальные горутины будут ждать
		counter[word]++ // записываем в мапу
		mx.Unlock()     // открываем мьютекс
	}
}

func main() {
	wg, mx := new(sync.WaitGroup), new(sync.Mutex) // объявляем sync.
	// mutex для обеспечения потокобезопасной записи в мапу
	languages := []string{"Golang", "Java", "C++", "Python", "Ruby", "C", "Kotlin"}
	counter := make(map[string]int)

	for _, word := range languages {
		wg.Add(1) // добавляем счетчик
		go count(mx, word, counter, wg)
	}
	wg.Wait() // ждем пока все горутины отработают
	fmt.Println(counter)
}
