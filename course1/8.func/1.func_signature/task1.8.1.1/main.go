package main

var CalculateCircleArea func(radius float64) float64
var CalculateRectangleArea func(width, height float64) float64
var CalculateTriangleArea func(base, height float64) float64

func main() {
	CalculateCircleArea = func(radius float64) float64 {
		return 3.14 * radius * radius
	}

	CalculateRectangleArea = func(width, height float64) float64 {
		return width * height
	}

	CalculateTriangleArea = func(base, height float64) float64 {
		return base * height / 2
	}
}
