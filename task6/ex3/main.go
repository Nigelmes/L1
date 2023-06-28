// Способ отмены №3, отменяем горутину через контекст таймаут
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func say(ctx context.Context, id int, languages []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Printf("Worker #%d say - %s\n", id, languages[rand.Intn(len(languages))])
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := new(sync.WaitGroup)
	languages := []string{"Golang", "Java", "C++", "Python", "Ruby", "C", "Kotlin"}
	timeout := 2 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go say(ctx, i, languages, wg)
	}
	wg.Wait()
}
