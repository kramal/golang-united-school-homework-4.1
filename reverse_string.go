package reverse_string

func ReverseString(input string) (output string) {
	strLen := len(input)
	bstr := []rune(input)
	for i := 0; i < strLen/2; i++ {
		bstr[i], bstr[strLen-1-i] = bstr[strLen-1-i], bstr[i]
	}
	output = string(bstr)
	return output
}
