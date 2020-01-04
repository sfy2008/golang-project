package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		// t.Errof("Hello=%q, want %q", got, want)
		t.Errorf("Hello=%v, want %v", got, want)
	}
}
