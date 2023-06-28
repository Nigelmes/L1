package main

import "fmt"

func main() {
	x, y := 10, 5
	fmt.Printf("x = [%2d], y = [%2d]\n", x, y)
	x, y = y, x // магия!
	fmt.Printf("x = [%2d], y = [%2d]\n", x, y)
}
