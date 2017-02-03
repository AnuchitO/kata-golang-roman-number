package roman

import "testing"

func TestInput1ShouldBeI(t *testing.T) {
	r := roman(1)

	assertEqual(t, r, "I")
}

func TestInput2ShouldBeII(t *testing.T) {
	r := roman(2)

	assertEqual(t, r, "II")
}

func TestInput3ShouldBeIII(t *testing.T) {
	r := roman(3)

	assertEqual(t, r, "III")
}

func assertEqual(t *testing.T, r, expected string) {
	if r != expected{
		t.Errorf("expected '%s' but got '%s'", expected, r)
	}
}
