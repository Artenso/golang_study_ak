package main

// func main() {
// 	map1 := map[string]int{"apple": 3, "banana": 2}
// 	map2 := map[string]int{"orange": 5, "grape": 4}
// 	mergedMap := mergeMaps(map1, map2)
// 	for key, value := range mergedMap {
// 		fmt.Printf("%s: %d\n", key, value)
// 	}
// }

// mergeMaps merges two maps in one
// WARNING!
// if maps have same keys, only map2 key will be saved
func mergeMaps(map1, map2 map[string]int) map[string]int {
	for key, val := range map2 {
		map1[key] = val
	}
	return map1

	// if base maps should stay witout changes:
	//
	// res := make(map[string]int, len(map1)+len(map2))
	// for key, val := range map1 {
	// 	res[key] = val
	// }
	// for key, val := range map2 {
	// 	res[key] = val
	// }
	// return res
}
