package reverse_string

func ReverseString(input string) (output string) {
	strLen := len(input)

	if strLen == 0 {
		return ""
	}

	if strLen == 1 {
		return input
	}

	bstr := []rune(input)
	for i := 0; i < strLen/2; i++ {
		bstr[i], bstr[strLen-1-i] = bstr[strLen-1-i], bstr[i]
	}
	output = string(bstr)
	return output
}
