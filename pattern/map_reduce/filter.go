package map_reduce

func IntFilter(lst []int, action func(item int) bool) []int {
	var result []int
	for _, intItem := range lst {
		if action(intItem) {
			result = append(result, intItem)
		}
	}
	return result
}
