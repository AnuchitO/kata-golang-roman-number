package roman

func roman(number int) string {
	if number >= 10 {
		return "X" + roman(number - 10)
	}
	if number >= 9 {
		return "IX" + roman(number - 9)
	}
	if number >= 5 {
		return "V" + roman(number - 5)
	}
	if number >= 4 {
		return "IV" + roman(number - 4)
	}
	if number >= 1 {
		return "I" + roman(number - 1)
	}

	return ""
}
