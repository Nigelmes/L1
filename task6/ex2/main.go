// Способ отмены №2, отменяем горутину через контекст отменяющей функцией cancel
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func say(ctx context.Context, id int, languages []string) {
	for {
		select {
		default:
			fmt.Printf("Worker #%d say - %s\n", id, languages[rand.Intn(len(languages))])
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	languages := []string{"Golang", "Java", "C++", "Python", "Ruby", "C", "Kotlin"}
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 3; i++ {
		go say(ctx, i, languages)
	}

	time.Sleep(2 * time.Second)
	cancel()
}
