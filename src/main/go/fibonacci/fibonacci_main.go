package main

import (
	"fmt"
	"github.com/socialpercon/getting-start-golang/src/main/go/fibonacci"
)

func main() {
	for x := range fibonacci.Fibonacci(10) {
		fmt.Println(x)
	}

}
