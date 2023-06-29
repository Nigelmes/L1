package main

import "fmt"

type Sender interface { // есть интерфейс который отправляет файл только в XMl формате
	SendXML()
}

type JsonDoc struct { // структруа json

}

func (j *JsonDoc) ConvertToXml() { // метод который конвертирует json в xml
	fmt.Println("Convert json to xml")
}

type AdapterJsonDoc struct { // адаптер для интерфейса который отправляет только XML файл
	jjsondoc *JsonDoc
}

func (a *AdapterJsonDoc) SendXML() { // метод который удовлетворяет интерфейсу  Sender ,
	// который внутри себя конверирует json в XMl
	a.jjsondoc.ConvertToXml()
	fmt.Println("send XML document")
}

func WorkProcces(send Sender) { // какой то воркер который принимет интерфейс Sender и работает с ним
	send.SendXML()
}

func main() {
	jsondoc := &JsonDoc{}
	adapter := &AdapterJsonDoc{jjsondoc: jsondoc}
	WorkProcces(adapter)
}
