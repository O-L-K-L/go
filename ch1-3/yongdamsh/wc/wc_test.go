package wc

import "testing"

func TestCalculate(t *testing.T) {
	want := 2
	if got := Calculate("Hello, World!"); want != got {
		t.Errorf("Calculate() = %q, want %q", got, want)
	}
}
