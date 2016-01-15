package tree

import (
	"testing"
)

func TestTrie(t *testing.T) {
	tr := NewTrie()
	tr.Insert([]rune("我和你"))
	if tr.Has([]rune("我")) {
		t.Error("not found 我")
		t.FailNow()
	}
	if tr.Has([]rune("我和")) {
		t.Error("not found 我和")
		t.FailNow()
	}
	if !tr.Has([]rune("我和你")) {
		t.Error("not found 我和你")
		t.FailNow()
	}
	if tr.Has([]rune("和你")) {
		t.Error("found 我和你")
		t.FailNow()
	}
	if !tr.HasPrefix([]rune("我")) {
		t.Error("not found 我")
		t.FailNow()
	}
}
