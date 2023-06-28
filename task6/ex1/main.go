// Способ отмены №1, отменяем горутину пустой структурой
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func say(cancel <-chan struct{}, id int, languages []string) {
	for {
		select {
		default:
			fmt.Printf("Worker #%d say - %s\n", id, languages[rand.Intn(len(languages))])
		case <-cancel:
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	languages := []string{"Golang", "Java", "C++", "Python", "Ruby", "C", "Kotlin"}
	cancel := make(chan struct{})

	for i := 0; i < 3; i++ {
		go say(cancel, i, languages)
	}

	time.Sleep(2 * time.Second)
	cancel <- struct{}{}
}
