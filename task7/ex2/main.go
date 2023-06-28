// Пример конкурентной запси в map с помощью sync.Map
package main

import (
	"fmt"
	"sync"
)

func count(word string, counter *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done() // сброс счетчика
	for i := 0; i < len(word); i++ {
		counter.Store(word, i+1) // запись в sync.map
	}
}

func main() {
	wg := new(sync.WaitGroup)
	languages := []string{"Golang", "Java", "C++", "Python", "Ruby", "C", "Kotlin"}
	counter := new(sync.Map)

	for _, word := range languages {
		wg.Add(1) // добавляем счетчик для горутин
		go count(word, counter, wg)
	}
	wg.Wait()                                       // ждем пока все горутины отработают
	counter.Range(func(key, val interface{}) bool { //вывод всех элементов sync.map
		fmt.Println(key, val)
		return true
	})
}
