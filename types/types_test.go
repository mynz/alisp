package types

import "testing"

func testEnv(t *testing.T) {
	b := Boolean(true)
	if b.String() != "#t" {
		t.Fatalf("expect #t")
	}

}

func testList(t *testing.T) {
	p := Pair{Symbol("foo"), Number(3.14)}
	if !p.IsNull() {
		t.Fatalf("must bwe IsNull")
	}
}
