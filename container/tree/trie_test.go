package tree

import (
	"testing"
)

func TestTrie(t *testing.T) {
	tr := NewTrie()
	tr.Insert("我和你")
	if tr.Has("我") {
		t.Error("not found 我")
		t.FailNow()
	}
	if tr.Has("我和") {
		t.Error("not found 我和")
		t.FailNow()
	}
	if !tr.Has("我和你") {
		t.Error("not found 我和你")
		t.FailNow()
	}
	if tr.Has("和你") {
		t.Error("found 我和你")
		t.FailNow()
	}
	if !tr.HasPrefix("我") {
		t.Error("not found 我")
		t.FailNow()
	}
}
