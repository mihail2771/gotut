package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func main() {
	Rect := Rectangle{Width: 4.0, Height: 2.0}

	fmt.Printf("Площать прямоугольника со сторонами %f и %f равна %f\n", Rect.Height, Rect.Width, Rect.Area())
	fmt.Printf("Периметр прямоугольника со сторонами %f и %f равна %f\n", Rect.Height, Rect.Width, Rect.Perimeter())
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}
