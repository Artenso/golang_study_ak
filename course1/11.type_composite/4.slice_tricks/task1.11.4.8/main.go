package main

func Shift(xs []int) (int, []int) {
	// xs[len(xs)-1:] - slice wich contains only last element
	// xs[:len(xs)-1] - slice wich contains all elements without last
	return xs[0], append(xs[len(xs)-1:], xs[:len(xs)-1]...)
}
