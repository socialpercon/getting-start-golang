package main

import "fmt"

func sum(x, y interface{}) interface{} {
	switch x.(type) {
	case int:
		x_value := x.(int)
		y_value := y.(int)
		return x_value + y_value
	case float64:
		x_value := x.(float64)
		y_value := y.(float64)
		return x_value + y_value
	}
	return 0
}

//func sum(x, y float64) float64 {
//	return x + y
//}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1.4, 2.5))
}
