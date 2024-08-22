package main

func RemoveIDX(xs []int, idx int) []int {
	if idx < 0 || idx > len(xs)-1 {
		return []int{}
	}
	// this actions will protect xs from changes
	// and protect res from changes in xs
	res := make([]int, len(xs[:idx]), cap(xs))
	copy(res, xs[:idx])
	res = append(res, xs[idx+1:]...)
	return res
}
