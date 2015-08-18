package main

import (
	"fmt"
	"github.com/socialpercon/fibonacci"
)

func main() {
	fmt.Println("Hello World. Fibonacci(10)=")
	for x := range fibonacci.Fibonacci(10) {
		fmt.Println(x)
	}
}
