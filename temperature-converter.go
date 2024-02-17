package main

import "fmt"

const K = 373.0

func main() {
	C := K - 273.0
	fmt.Printf("boiling temperature of water in K = %g, boiling temperature of water in °C = %g", K, C)
}
