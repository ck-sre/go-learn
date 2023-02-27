package worder

import (
	"benchmarks-examples-doc/quoter"
	"fmt"
	"testing"
)

func BenchmarkUseCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quoter.Suntzu)
	}
}

func BenchmarkCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Counter(quoter.Suntzu)
	}
}

func ExampleCounter() {
	fmt.Println(Counter("hello there Kyushu"))
}

func TestCounter(t *testing.T) {
	n := Counter("there is peace in Japan")
	if n != 5 {
		t.Error("got", n, "want", 5)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("thirty five five")
	for k, v := range m {
		switch k {
		case "thirty":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "five":
			if v != 2 {
				t.Error("got", v, "want", 2)
			}

		}
	}

}
