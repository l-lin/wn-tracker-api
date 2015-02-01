package novel

import "testing"

func Test_IsValid(t *testing.T) {
	n := New()
	if n.IsValid() {
		t.Fatalf("[x] The novel should not be valid with an empty userId, title and url!")
	}
	n.UserId = "UserId"
	if n.IsValid() {
		t.Fatalf("[x] The novel should be valid with an empty title and url!")
	}
	n.Title = "Title"
	if n.IsValid() {
		t.Fatalf("[x] The novel should be valid with an empty url!")
	}
	n.Url = "http://localhost"
	if !n.IsValid() {
		t.Fatalf("[x] The novel should be valid!")
	}
}
