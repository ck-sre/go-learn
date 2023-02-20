package benchmarking

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	w := "Hello my name is Tomas"
	s := Greet("Tomas")
	if s != w {
		t.Error("got", s, "wanted", w)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Radul"))
	// Output:
	// Hello my name is Radul
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Alex")
	}
}
