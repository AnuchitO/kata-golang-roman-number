package roman

import "testing"

func TestCaseRomanNumber(t *testing.T) {
	romans := map[int]string{
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		9: "IX",
		10: "X",
		11: "XI",
		40: "XL",
	}

	for number, expected := range romans {
		assertEqual(t, roman(number), expected)
	}
}

func assertEqual(t *testing.T, r, expected string) {
	if r != expected{
		t.Errorf("expected '%s' but got '%s'", expected, r)
	}
}
