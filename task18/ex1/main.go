// решение №1 с использованием пакета atomic
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct { //  структура с одним полем total которое представляет значение счетчика
	total int64
}

func (c *Counter) increment() {
	atomic.AddInt64(&c.total, 1) // метод потокобезопасно увеличивает счетчик на 1
}

func (c *Counter) value() int64 {
	return atomic.LoadInt64(&c.total) // потокобезопасное чтение счетчика
}

func main() {
	wg := new(sync.WaitGroup)
	total := Counter{}

	for i := 0; i < 5; i++ { // в цикле запускаем 5 горутин
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ { // в каждой горутине инкрементируем счетчик 1000 раз
				total.increment()
			}
		}()
	}
	wg.Wait()                  // ждем пока отработают все горутины
	fmt.Println(total.value()) // вывод в stdoout
}
