package main

import "fmt"

func fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func fibRec(n int) int {
	if n <= 1 {
		return 1
	}
	return fibRec(n-1) + fibRec(n-2)
}

func Fibonacci(n int) chan int {
	c := make(chan int)
	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			c <- a
		}
		close(c)
	}()
	return c
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fib(i), fibRec(i))
	}

	for x := range Fibonacci(10) {
		fmt.Println(x)
	}

}
