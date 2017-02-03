package roman

import "testing"

func TestInput1ShouldBeI(t *testing.T) {
	assertEqual(t, roman(1), "I")
}

func TestInput2ShouldBeII(t *testing.T) {
	assertEqual(t, roman(2), "II")
}

func TestInput3ShouldBeIII(t *testing.T) {
	assertEqual(t, roman(3), "III")
}

func TestInput4ShouldBeIV(t *testing.T) {
	assertEqual(t, roman(4), "IV")
}

func TestInput5ShouldBeIV(t *testing.T) {
	assertEqual(t, roman(5), "V")
}

func TestInput6ShouldBeVI(t *testing.T) {
	assertEqual(t, roman(6), "VI")
}

func TestInput9ShouldBeIX(t *testing.T) {
	assertEqual(t, roman(9), "IX")
}

func TestInput10ShouldBeX(t *testing.T) {
	assertEqual(t, roman(10), "X")
}

func assertEqual(t *testing.T, r, expected string) {
	if r != expected{
		t.Errorf("expected '%s' but got '%s'", expected, r)
	}
}
