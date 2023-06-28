package main

import (
	"fmt"
	"log"
)

func setbit(num *int64, index uint, bit uint) { //  функция установки бита, index = позиция , bit = нужный бит
	if bit == 1 {
		*num |= int64(1) << index //  установка единицы в нужную позицию путем наложенмя маски и логического ИЛИ(+)
	} else if bit == 0 {
		*num &= ^(int64(1) << index) // установка нуля в нужную позицию путем наложенмя маски и логического И(*)
	} else {
		log.Fatal("bit must be 0 or 1")
	}
}

func main() {
	var num int64 = 1
	setbit(&num, 0, 0) //  меняем первый бит на ноль
	fmt.Println(num)
	setbit(&num, 3, 1) // меняем 4 бит на единцу
	fmt.Println(num)
}
