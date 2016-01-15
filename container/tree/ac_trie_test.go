package tree

import (
	"testing"
)

func TestACTrie(t *testing.T) {
	tr := NewACTrie()
	tr.Insert("nihao", "hao", "hs", "hsr", "she", "he")
	tr.SetFail()
	tr.QueryFunc([]rune("sdmfhsgnshejfgnihaofhsrnihao"), func(s, e int) {
		t.Logf("start: %d, end: %d", s, e)
	})
}
