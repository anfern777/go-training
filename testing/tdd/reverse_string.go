package tdd

func ReverseString(input string) string {
	var output []rune
	parsedInput := []rune(input)
	for i := len(parsedInput) - 1; i >= 0; i-- {
		output = append(output, parsedInput[i])
	}
	return string(output)
}
