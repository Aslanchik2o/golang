package main

import "fmt"

func main() {
	fmt.Println(findMin(1, 2, 3, 4, 0, 5, 6, 54765765, -1))

}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0

	}
	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}

	return min
}package main

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