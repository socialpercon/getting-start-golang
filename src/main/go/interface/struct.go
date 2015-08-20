package main

import "fmt"

//define a Rectangle struct that has a length and a width
type Rectangle struct {
	length, width int
}

//write a function Area that can apply to a Rectangle type
func (r Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	//define a new Rectangle instance with values for its properties
	r := Rectangle{length: 5, width: 3}
	fmt.Println("Rectangle details are: ", r)
	fmt.Println("Rectangle's area is: ", r.Area())
}
