package roman

func roman(number int) string {
	if number == 4 {
		return "IV"
	}
	if number >= 1 {
		return "I" + roman(number - 1)
	}

	return ""
}
