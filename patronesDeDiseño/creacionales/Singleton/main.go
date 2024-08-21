package main

import "fmt"

func main() {
	for x := 0; x <= 15; x++ {
		go getInstance()
	}
	fmt.Scanln()
}
