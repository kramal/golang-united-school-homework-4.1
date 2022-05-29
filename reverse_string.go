package reverse_string

func ReverseString(input string) (output string) {
	bstr := []rune(input)
	strLen := len(bstr)

	if strLen == 0 {
		return ""
	}

	if strLen == 1 {
		return input
	}

	for i := 0; i < strLen/2; i++ {
		bstr[i], bstr[strLen-1-i] = bstr[strLen-1-i], bstr[i]
	}
	output = string(bstr)
	return output
}
