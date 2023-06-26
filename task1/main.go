// В языке Go нет наследования — есть только композиция (встраивание)
package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Weight int
	Height int
}

func (h *Human) GetHumanInfo() {
	fmt.Printf("Name: [%s], Age: [%d], Weight: [%d], Height: [%d]\n", h.Name, h.Age, h.Weight, h.Height)
}

type Action struct {
	Human //встроили структуру Human в структуру Action, теперь можем использовать методы Human в структуре Action
}

func (a *Action) DoHomework() {
	fmt.Printf("Student [%s] doing homework\n", a.Name)
}

func main() {
	student1 := Action{
		Human{
			Name:   "Oleg",
			Age:    18,
			Weight: 72,
			Height: 180,
		},
	}
	student1.GetHumanInfo() // используем метод встроенного класса
	student1.DoHomework()
}
