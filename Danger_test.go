package Danger

import "testing"

func TestHello(t *testing.T) {
	want := "Danger, Danger."
	if got := Danger(); got != want {
		t.Errorf("Danger() = %q, want %q", got, want)
	}
}
