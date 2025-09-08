package main

import "fmt"

func Generic[T any](arg T) (*T, error) { return &arg, nil }

func main() {
	h, _ := Generic("hello")
	fmt.Println(*h)
}

/**
func Generic_string(arg string) (*string, error) { return &arg, nil }

func main() {
	h, _ := Generic_string("hello")
	fmt.Println(*h)
}
**/
