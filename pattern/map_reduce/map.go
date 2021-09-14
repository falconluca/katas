package map_reduce

func MapToInt(strLst []string, action func(item string) int) []int {
	var result []int
	for _, strItem := range strLst {
		intItem := action(strItem)
		result = append(result, intItem)
	}
	return result
}

func MapToStr(strLst []string, action func(item string) string) []string {
	var result []string
	for _, strItem := range strLst {
		newStrItem := action(strItem)
		result = append(result, newStrItem)
	}
	return result
}
