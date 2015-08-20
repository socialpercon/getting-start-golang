package main

import "fmt"

func foo(a interface{}) {
	fmt.Println(a.(float64))
}

func main() {
	var a int = 10
	var b float64 = float64(a)
	fmt.Println(a)
	fmt.Println(b)
	foo(a)
}
