// Concept: polymorphism through a slice of interface values
// Task: declare the Shape interface and write totalArea so one function sums every shape
// Expected output: rect 6.00
// circle 12.56
// total 18.56
// Hint: list the methods Rect and Circle already share in one interface; a []Shape can then
//       hold values of both types, and totalArea calls Area on each element (Go Tour: Methods 9-11)

package main

import "fmt"

type Shape interface {
	Area() float64
	Name() string
}

type Rect struct {
	W, H float64
}

func (r Rect) Area() float64 { return r.W * r.H }

func (r Rect) Name() string { return "rect" }

type Circle struct {
	R float64
}

func (c Circle) Area() float64 { return 3.14 * c.R * c.R }

func (c Circle) Name() string { return "circle" }

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	shapes := []Shape{Rect{W: 2, H: 3}, Circle{R: 2}}
	for _, shape := range shapes {
		fmt.Printf("%s %.2f\n", shape.Name(), shape.Area())
	}
	fmt.Printf("total %.2f\n", totalArea(shapes))
}
