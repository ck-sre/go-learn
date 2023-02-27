package joinerstr

import (
	"fmt"
	"strings"
	"testing"
)

func TestJoint(t *testing.T) {
	s := "Several bones together"
	sl := strings.Split(s, " ")
	s = Joint(sl)
	if s != "Sever bones together" {
		t.Error("got", s, "want", "Several bones together")
	}
}

func TestAddStr(t *testing.T) {
	s := "Several bones together"
	sl := strings.Split(s, " ")
	s = AddStr(sl)
	if s != "Sever bones together" {
		t.Error("got", s, "want", "Several bones together")
	}
}

func ExampleJoint() {
	s := "Latin dance excusable"
	e := strings.Split(s, " ")
	fmt.Println(AddStr(e))
	// Output:
	// "Latin dance excusable"
}

const t = `We should ask ourselves how much curry can we eat`

var lc = []string{"hello", "let's", "dance"}

func BenchmarkAddStr(b *testing.B) {
	sp := strings.Split(t, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddStr(sp)
	}
}

func BenchmarkJoint(b *testing.B) {
	sp := strings.Split(t, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Joint(sp)
	}
}
