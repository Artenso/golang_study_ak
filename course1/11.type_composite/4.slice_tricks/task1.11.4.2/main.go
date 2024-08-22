package main

func InsertAfterIDX(xs []int, idx int, x ...int) []int {
	if idx < -1 || idx > len(xs)-1 {
		return []int{}
	}
	tmp := make([]int, len(xs[idx+1:]))
	copy(tmp, xs[idx+1:])
	return append(append(xs[:idx+1], x...), tmp...)
}
