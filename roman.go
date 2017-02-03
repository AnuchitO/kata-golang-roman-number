package roman

func roman(number int) string {
	if number >= 1 {
		return "I" + roman(number - 1)
	}

	return ""
}
