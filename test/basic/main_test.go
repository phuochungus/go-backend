package basic

import "testing"

func TestAddOne(t *testing.T) {
	got := AddOne(1)
	if got != 2 {
		t.Errorf("AddOne(1) = %d; want 2", got)
	}
}
