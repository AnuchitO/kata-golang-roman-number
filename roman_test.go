package roman

import "testing"

func TestInput1ShouldBeI(t *testing.T) {
	r := roman(1)

	if r != "I" {
		t.Errorf("expected I but got '%s'", r)
	}
}

func TestInput2ShouldBeII(t *testing.T) {
	r := roman(2)

	if r != "II" {
		t.Errorf("expected II but got '%s'", r)
	}
}

func TestInput3ShouldBeIII(t *testing.T) {
	r := roman(3)

	if r != "III" {
		t.Errorf("expected III but got '%s'", r)
	}
}
