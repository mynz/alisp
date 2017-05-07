package eval

import "testing"

func TestNewEnv(t *testing.T) {
	env := NewEnv()
	env.Setup()

	if _, err := env.Get("car"); err != nil {
		t.Fatalf("cannot find car from primitive environment: %s", err)
	}
}
