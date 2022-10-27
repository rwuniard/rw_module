package rw_module

import "fmt"

func SayHello() string {
	message := "Hello this is from rw_module."
	fmt.Println(message)

	return message
}
