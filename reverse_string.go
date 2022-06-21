package reverse_string

func ReverseString(input string) (output string) {
	for _, itemRune := range input {
		output = string(itemRune) + output
	}
	return output
}
