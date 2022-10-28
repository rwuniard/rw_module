package rw_module

import "fmt"

func SayHello() string {
	message := "Hello this is from rw_module. v0.2.0"
	fmt.Println(message)

	return message
}
