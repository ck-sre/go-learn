package avg_centred

import (
	"fmt"
	"testing"
)

func TestAvgCentred(t *testing.T) {
	type testStruct struct {
		args   []int
		answer float64
	}

	testStructs := []testStruct{
		{[]int{11, 2, 5, 3, 2}, 5},
		{[]int{11, 2, 5, 7, 2}, 7},
		{[]int{11, 2, 13, 3, 2}, 9},
	}
	for _, v := range testStructs {
		ef := AvgCentre(v.args)
		if ef != v.answer {
			t.Error("Got ", ef, "want ", v.answer)
		}
	}
	fmt.Println()
}

func BenchmarkAvgCentre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AvgCentre([]int{12, 43, 52, 4})
	}
}

func ExampleAvgCentre() {
	fmt.Println(AvgCentre([]int{4, 5, 1, 4}))
	//Output:
	//	3
}
