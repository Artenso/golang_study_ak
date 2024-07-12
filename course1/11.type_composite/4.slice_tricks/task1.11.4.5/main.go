package main

func FilterDividers(xs []int, divider int) []int {
	if divider == 0 {
		return []int{}
	}

	res := make([]int, 0, len(xs))
	for _, x := range xs {
		if x%divider == 0 {
			res = append(res, x)
		}
	}
	return res
}
