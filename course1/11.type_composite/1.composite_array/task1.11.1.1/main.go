package main

// func main() {
// 	xs := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
// 	fmt.Println(sum(xs))     // Вывод: 36
// 	fmt.Println(average(xs)) // Вывод: 4.5
// 	ys := [8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}
// 	fmt.Println(averageFloat(ys)) // Вывод: 5
// 	fmt.Println(reverse(xs))      // Вывод: [8 7 6 5 4 3 2 1]
// }

func sum(array [8]int) int {
	res := 0
	for _, num := range array {
		res += num
	}
	return res
}

func average(array [8]int) float64 {
	return float64(sum(array)) / float64(len(array))
}

func averageFloat(array [8]float64) float64 {
	res := 0.0
	for _, num := range array {
		res += num
	}
	return res / float64(len(array))
}

func reverse(array [8]int) [8]int {
	for i, j := 0, len(array)-1; i < len(array)/2; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}
