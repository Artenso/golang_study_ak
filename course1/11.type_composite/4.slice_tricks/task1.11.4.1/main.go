package main

func Cut(xs []int, start, end int) []int {
	if start > end {
		start, end = end, start
	}
	if end > len(xs)-1 {
		return []int{}
	}
	if start < 0 {
		return []int{}
	}
	return xs[start : end+1]
}
