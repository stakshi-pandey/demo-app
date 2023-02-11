package maths

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(1, 2) == 3 {
		t.Log("test passes")
	} else {
		t.Error("test failed...")
	}
}
