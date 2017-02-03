package roman

type Pair struct {
	number int
	roman string
}

func roman(number int) string {
	pairs := []Pair{
		Pair{90, "XC"},
		Pair{50, "L"},
		Pair{40, "XL"},
		Pair{10, "X"},
		Pair{9, "IX"},
		Pair{5, "V"},
		Pair{4, "IV"},
		Pair{1, "I"},
	}

	for _, p := range pairs {
		if number >= p.number {
			return p.roman + roman(number - p.number)
		}
	}

	return ""
}
