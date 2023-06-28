// Способ отмены №4, отменяем горутину через булевый указатель
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func say(stop *bool, id int, languages []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if *stop {
			fmt.Println("stop with pointer")
			return
		}
		fmt.Printf("Worker #%d say - %s\n", id, languages[rand.Intn(len(languages))])

	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := new(sync.WaitGroup)
	languages := []string{"Golang", "Java", "C++", "Python", "Ruby", "C", "Kotlin"}
	stop := new(bool)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go say(stop, i, languages, wg)
	}
	time.Sleep(2 * time.Second)
	*stop = true
	wg.Wait()
}
