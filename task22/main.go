package main

import (
	"fmt"
	"math/big"
)

func MiniCalc(num1, num2 *big.Int, operator string) *big.Int { // мини калькулятор
	res := new(big.Int) // создаем результирующую переменную
	switch operator {   // определяем переданынй опертор и в кейсах манипулируем числами с методоми пакета big
	case "+":
		return res.Add(num1, num2)
	case "-":
		return res.Sub(num1, num2)
	case "*":
		return res.Mul(num1, num2)
	case "/":
		return res.Div(num1, num2)
	}
	return res
}

func main() {
	num1, _ := new(big.Int).SetString(
		"7218882428714186575617218882428714186575617",
		10,
	) // задаем большое значение строкой , тк число слишком большое
	num2, _ := new(big.Int).SetString(
		"922337203685477580792233720368547758079223", 10,
	) // задаем большое значение строкой , тк число слишком большое

	fmt.Println(MiniCalc(num1, num2, "+"))
	fmt.Println(MiniCalc(num1, num2, "-"))
	fmt.Println(MiniCalc(num1, num2, "*"))
	fmt.Println(MiniCalc(num1, num2, "/"))
}
