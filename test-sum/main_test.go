package main

import "testing"

func TestMySum(t *testing.T) {
	x := MySum(2, 3)
	if x != 5 {
		t.Error("Expected ", 5, "Got", x)
	}
}
