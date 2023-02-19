package main

import "testing"

func TestMySum(t *testing.T) {

	type tableTest struct {
		a []int
		r int
	}

	tableTests := []tableTest{
		tableTest{[]int{1, 2, 3}, 6},
		tableTest{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tableTests {
		y := MySum(v.a...)
		if y != v.r {
			t.Error("Expected ", v.r, " Got ", y)
		}
	}
}
