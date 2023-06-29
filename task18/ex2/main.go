// решение №2 с помощью Mutex
package main

import (
	"fmt"
	"sync"
)

type Counter struct { //  структура с полем total которое представляет значение счетчика и Mutex
	// для блокировки горутины
	mx    sync.RWMutex
	total int64
}

func (c *Counter) increment() {
	c.mx.Lock()         // лочим на запись мьютексом, только одна горутина может писать и читать
	defer c.mx.Unlock() //
	c.total++           // инкреминтируемся
}

func (c *Counter) value() int64 {
	c.mx.RLock()         // лочим на чтение, в этом случае все горутины могут читать
	defer c.mx.RUnlock() //
	return c.total       // возвращаем значение
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
