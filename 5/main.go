package main

import (
	"errors"
	"fmt"
)

func main() {
	messages := [3] string{"1","2","3"}
	messages[1] = "5"
	fmt.Println(messages[1])
	printMessages(messages)
}

func printMessages(messages [3] string) error {
	if len(messages) == 0 {
		return errors.New("empty array")

	}
	fmt.Println(messages)

	return nil
}