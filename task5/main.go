package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func say(ctx context.Context, channel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case a := <-channel:
			fmt.Printf("Say - %s\n", a)
		case <-ctx.Done(): // канал отмены
			return
		}
	}
}

func sender(ctx context.Context, channel chan<- string, languages []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case channel <- languages[rand.Intn(len(languages))]: //отправляем в канал рандомное слово из слайса
		case <-ctx.Done(): //канал отмены
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := new(sync.WaitGroup) // waitgroup для ожидания горутин
	wg.Add(2)
	languages := []string{"Golang", "Java", "C++", "Python", "Ruby", "C", "Kotlin"}
	var N int
	if _, err := fmt.Scan(&N); err != nil {
		log.Fatal(err)
	}
	worktime := time.Duration(N) * time.Second // время таймаута
	channel := make(chan string)
	ctx, _ := context.WithTimeout(context.Background(), worktime) // контекст отмены по времени
	go sender(ctx, channel, languages, wg)                        // запуск отправителя в горутине
	go say(ctx, channel, wg)                                      // запуск говорителя в горутине
	wg.Wait()
}
