package tree

import (
	"testing"
)

func TestACTrie(t *testing.T) {
	tr := NewACTrie()
	tr.Insert([]rune("nihao"))
	tr.Insert([]rune("hao"))
	tr.Insert([]rune("hs"))
	tr.Insert([]rune("hsr"))
	tr.initFail()
	tr.QueryFunc([]rune("sdmfhsgnshejfgnihaofhsrnihao"), func(s, e int) {
		t.Logf("start: %d, end: %d", s, e)
	})
}
