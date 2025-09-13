package main

import "fmt"

func main() {
	if 1 == 2 {
		fmt.Println("unreachable")
	} else {
		fmt.Println("reachable")
	}
}
