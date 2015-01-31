package novel

import "testing"

func Test_IsValid(t *testing.T) {
	n := New()
	if n.IsValid() {
		t.Fatalf("[x] The novel should not be valid with an empty title!")
	}
	n.Title = "Foobar"
	if !n.IsValid() {
		t.Fatalf("[x] The novel should be valid with an title!")
	}
}
