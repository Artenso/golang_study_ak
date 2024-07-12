package main

func MaxDifference(numbers []int) int {
	if len(numbers) <= 1 {
		return 0
	}
	min := numbers[0]
	max := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
			continue
		}
		if num > max {
			max = num
		}

	}
	return max - min
}
