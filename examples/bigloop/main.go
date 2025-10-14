package main

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i == j {
				print("a")
			} else if i < j {
				print("b")
			} else {
				print("c")
			}
		}
	}
}
