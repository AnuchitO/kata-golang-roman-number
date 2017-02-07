package roman

type Pair struct {
	number int
	roman  string
}

func roman(number int) string {
	pairs := []Pair{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, p := range pairs {
		if number >= p.number {
			return p.roman + roman(number-p.number)
		}
	}

	return ""
}
