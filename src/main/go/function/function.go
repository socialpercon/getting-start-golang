package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

//func sum(x, y float64) float64 {
//	return x + y
//}

func main() {
	fmt.Println(sum(1, 2))
}
