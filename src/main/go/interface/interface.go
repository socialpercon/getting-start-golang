package main

import "fmt"

type Shaper interface {
	Area() int
}

type Rectangle struct {
	length, width int
}

type Square struct {
	side int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

//the Square type also has an Area function and therefore,
//it too, implements the Shaper interface
func (sq Square) Area() int {
	return sq.side * sq.side
}

func main() {
	r := Rectangle{length: 5, width: 3}
	fmt.Println("Rectangle details are: ", r)
	var s Shaper
	//equivalent to "s = Shaper(r)"
	//since Go identifies r matches the Shaper interface
	s = r
	fmt.Println("Area of Shaper(Rectangle): ", s.Area())
	fmt.Println()

	q := Square{side: 5}
	fmt.Println("Square details are: ", q)
	//equivalent to "s = Shaper(q)
	s = q
	fmt.Println("Area of Shaper(Square): ", s.Area())
	fmt.Println()
}
