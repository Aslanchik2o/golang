// изучение: panic(), recover(), delef()

package main

import "fmt"


func main(){
	defer handPanic()
	// defer()
 //   defer PrintMessage()
//	fmt.Println("main()")
//func main() {
//	messages := []string{
	//	"messages 1",
	//	"messages 2",
	//	"messages 3",
	//	"messages 4",
	//}
	//messages[4] = "messages 5"
//	fmt.Println(messages)
	messages := []string{
		"messages 1",
		"messages 2",
		"messages 3",
		"messages 4",
	}
	messages[4] = "messages 5"
//	fmt.Println(messages)
}

func handPanic(){
	if r := recover(); r!= nil{
		fmt.Println(r)
	}

	fmt.Println("handPanic() выполнилась успешно")
	
}
//func PrintMessage(){
//	fmt.Println("printMessage()")
//}

