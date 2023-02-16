package cat

import (
	"fmt"
	"testing"
)

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsSum() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func TestYears(t *testing.T) {
	y := Years(10)
	if y != 70 {
		t.Error("got", y, "want", 70)
	}
}

func TestYearsSum(t *testing.T) {
	y := YearsSum(10)
	if y != 70 {
		t.Error("got", y, "want", 70)
	}
}

func ExampleYearsSum() {

}

func BenchMarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsSum(c *testing.B) {
	for i := 0; i < c.N; i++ {
		YearsSum(10)
	}
}
