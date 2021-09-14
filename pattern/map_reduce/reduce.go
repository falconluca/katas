package map_reduce

func Reduce(lst []string, action func(item string) int) int {
	result := 0
	for _, strItem := range lst {
		result += action(strItem)
	}
	return result
}
