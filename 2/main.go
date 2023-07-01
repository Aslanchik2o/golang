package main

import "fmt"

func main() {
 massage := "Я скоро стану GOLANG NINJA"
	fmt.Println(massage)
}
вот так создается переменая 
без указание ее она не будет работать 
cd 1 - нужен когда у тебя не работает go run main.go
переменую int можно использовать только с числами 

|| - или
&& - и
= - присваевание 
== - проверка 

package main

import (
	"fmt"
)

func main() {
massage, entered := enterTheClub(18)
fmt.Println(massage)
fmt.Println(entered)
//	var message string
//	message = sayHello("МАксим", 21)
	//printMessage(message)
	//printMessage("вызов 1")
	//printMessage("вызов 2")
	//printMessage("вызов 3")

}
//func printMessage(message string) {
//	fmt.Println(message)
//}

//func sayHello(name string, age int) string {
//	return fmt.Sprintf("Привет, %s ! Тебе %d лет", name, age)

//}

func enterTheClub(age int) (string, bool) {
    if age >=18 {
	  return "Входи, только осторожно", true 
	}  else{
		return "тебе нет 18-ти", false
	}