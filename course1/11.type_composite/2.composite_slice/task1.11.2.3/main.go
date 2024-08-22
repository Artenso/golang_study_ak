package main

func bitwiseXOR(n1, n2 int) int {
	return n1 ^ n2
}

func findSingleNumber(numbers []int) int {
	check := true
	for i, n1 := range numbers {
		for j, n2 := range numbers {
			check = true
			if i != j && bitwiseXOR(n1, n2) < 1 {
				check = false
				break
			}

		}
		if check {
			return n1
		}
	}
	return 0
}
