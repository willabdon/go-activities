package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%5 == 0 {
			fmt.Printf("Pan\n")
		} else if i%3 == 0 {
			fmt.Printf("Pin\n")
		} else {
			fmt.Printf("%d\n", i)
		}
	}
}
