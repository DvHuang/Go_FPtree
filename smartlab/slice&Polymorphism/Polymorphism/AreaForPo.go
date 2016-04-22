package main

import (
	"fmt"
)

type shaper interface {
	Area() float32
}

type square struct {
	side float32
}

func (sq square) Area() float32 {
	return sq.side * sq.side
}

type rectangle struct {
	length, width float32
}

func (r rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := rectangle{5, 3}
	q := square{5}
	shapes := []shaper{r, q}

	for n, _ := range shapes {

		fmt.Println("shape deatils:", shapes[n])
		fmt.Println("Area of this shap is:", shapes[n].Area())
	}
}
