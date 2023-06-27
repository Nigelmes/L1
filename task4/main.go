package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(ctx context.Context, say <-chan string, id int) {
	for {
		select {
		case a := <-say:
			fmt.Printf("Worker #%d say - %s\n", id, a)
		case <-ctx.Done(): // канал отмены
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var N int
	languages := []string{"Golang", "Java", "C++", "Python", "Ruby", "C", "Kotlin"}
	if _, err := fmt.Scan(&N); err != nil {
		log.Fatal(err)
	}
	channel := make(chan string, N)                         // канал для отправки слов
	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст для отмены горутин

	for i := 0; i < N; i++ { // в цикле запускаем N воркеров
		go worker(ctx, channel, i)
	}

	go func() { // горутина для бесконечной отправки сообщений в канал
		for {
			select {
			case channel <- languages[rand.Intn(len(languages))]: //отправляем в канал рандомное слово из слайса
			case <-ctx.Done(): //канал отмены
				return
			}
		}
	}()

	quit := make(chan os.Signal, 1)                      // канал для завершения приложения через ctr+c
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT) // обработчик сигнала
	<-quit                                               // ждем сигнал

	cancel()       // отменяем горутины
	close(channel) // закрываем канал
}
