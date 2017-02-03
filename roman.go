package roman

func roman(number int) string {
	if number == 9 {
		return "IX"
	}
	if number >= 5 {
		return "V" + roman(number - 5)
	}
	if number == 4 {
		return "IV"
	}
	if number >= 1 {
		return "I" + roman(number - 1)
	}

	return ""
}
