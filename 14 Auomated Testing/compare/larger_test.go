package compare

import (
	"fmt"
	"testing"
)

func errorString(a int, b int, got int, want int) string {
	return fmt.Sprintf("Larger(%d, %d) = %d, want %d", a, b, got, want)
}

func TestFirstLarger(t *testing.T) {
	got := Larger(6, 4)
	want := 6
	if got != want {
		t.Error(errorString(6, 4, got, want))
	}
}

func TestSecondLarger(t *testing.T) {
	got := Larger(4, 6)
	want := 6
	if got != want {
		t.Error(errorString(4, 6, got, want))
	}
}
