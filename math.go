package main

import (
	"fmt"
)

func main() {
	fmt.Println(Soma(10, 10))
}

func Soma(a int, b int) int {
	fmt.Println("Soma: ", a+b)
	return a + b
}
