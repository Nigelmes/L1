package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64 // для инкапсуляции полей переменные пишем с нижнего регистра
}

func NewPoint(x float64, y float64) *Point { // конструктор для структуры Point
	return &Point{x: x, y: y}
}

func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y),
		2)) // расчет расстояния между двумя точками по формуле
}

func main() {
	point1 := NewPoint(5.5, 6.6)
	point2 := NewPoint(1.1, 2.5)
	fmt.Println(Distance(point1, point2))
}
