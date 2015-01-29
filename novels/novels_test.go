package novels

import "testing"

func Test_NovelToJson(t *testing.T) {
	expected := "{\"Id\":1,\"Name\":\"foobar\"}"
	novels := NovelToJson()
	if expected != novels {
		t.Log("Expected: " + expected + ". Actual was: " + novels)
		t.FailNow()
	}
}
