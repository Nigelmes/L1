package main

import (
	"fmt"
	"time"
)

func MySleep(d time.Duration) {
	<-time.After(d) // блокируем горутину и ждем сигнал по истечению времени
}

func main() {
	fmt.Println("Golang")
	MySleep(3 * time.Second)
	fmt.Println("DevOps")
}
