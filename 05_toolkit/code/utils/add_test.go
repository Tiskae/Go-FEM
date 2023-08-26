package utils

import "testing"

func TestAdd(t *testing.T) {
	expected := 15
	actual := Add(1, 2, 3, 4, 5)

	if actual != expected {
		t.Errorf("Sum was incorrect! Expected: %d, Actual: %d", expected, actual)
	}
}
