package main

import "fmt"

func handler() {
	defer func() {
		s := recover()
		fmt.Println("haha")
		fmt.Println(s)

	}()

	panic("!!!!")

}

func main() {
	handler()

	fmt.Println("main println")

}
