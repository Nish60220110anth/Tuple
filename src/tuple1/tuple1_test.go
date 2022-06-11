package tuple1_test

import (
	"testing"
	"tuple/src/tuple1"
)

func TestAt(t *testing.T) {
	tuple1Var := tuple1.NewTuple1("test")
	got := (*tuple1Var).At(0)
	want := "test1"

	if got != want {
		t.Errorf("Func-At (path = tuple/src/tuple1) For Input tuple %s Got : \"%s\" Required \"%s\"", (*tuple1Var).String(), got, want)
	}
}
