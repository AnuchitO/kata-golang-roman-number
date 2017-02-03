package roman

import "testing"

func TestInput1ShouldBeI(t *testing.T) {
	r := roman(1)

	if r != "I" {
		t.Errorf("expected I but got '%s'", r)
	}
}
