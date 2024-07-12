package main

func InsertToStart(xs []int, x ...int) []int {
	res := make([]int, 0, len(xs)+len(x))
	res = append(res, x...)
	res = append(res, xs...)
	return res
}
