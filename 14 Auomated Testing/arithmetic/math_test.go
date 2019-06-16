package arithmetic

import (
	
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 3) != 4 {
		t.Error("1+3 didn't equal 4")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(5, 6) != -1 {
		t.Error("5-6 didn't equal -1")
	}
}
