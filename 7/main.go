package main

import "fmt"

func main() {
	messages := []string{
		"messages 1",
		"messages 2",
		"messages 3",
		"messages 4",
	}
	for i := range messages {
		fmt.Println(messages[i])
	}
}