package main

import "fmt"

func fibonacci() func() int {

	current := 0
	previous := 0
	temp := 0

	return func() int {

		if current == 0 && previous == 0 {
			current = 1
			return 0
		} else {
			temp = current
			current = current + previous
			previous = temp
		}

		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
