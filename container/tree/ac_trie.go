package tree

import (
	"log"
)

// Aho Corasick Trie Tree
type ACTrie struct {
	next   map[rune]*ACTrie
	fail   *ACTrie
	length int
}

// NewACTrie new a Aho Corasick Trie Tree.
func NewACTrie() *ACTrie {
	return &ACTrie{
		next: make(map[rune]*ACTrie),
	}
}

// Insert insert words into ac trie tree, after Insert must call SetFail to
// build the fail table.
func (t *ACTrie) Insert(vs ...string) {
	var i int
	for i = 0; i < len(vs); i++ {
		t.insert([]rune(vs[i]))
	}
}

func (t *ACTrie) insert(strs []rune) {
	var (
		i      int
		ok     bool
		str    rune
		c      *ACTrie
		p      = t
		strLen = len(strs)
	)
	if debug {
		log.Printf("actrie: insert strs len: %d\n", strLen)
	}
	for i = 0; i < strLen; i++ {
		str = strs[i]
		if c, ok = p.next[str]; !ok {
			c = NewACTrie()
			p.next[str] = c
			if debug {
				log.Printf("actrie: insert new trie: %s\n", string(str))
			}
		}
		p = c
	}
	c.length = strLen
}

// SetFail build the fail table, must call after insert a word.
func (t *ACTrie) SetFail() {
	var (
		i       int
		ok      bool
		str     rune
		p, c, f *ACTrie
		queue   = []*ACTrie{t} // first node root
	)
	t.fail = nil
	for i = 0; i < len(queue); i++ {
		p = queue[i]                // bfs
		for str, c = range p.next { // get child node
			f = p.fail // parent fail
			for {
				if f == nil {
					c.fail = t // set root
					break
				}
				if c.fail, ok = f.next[str]; ok {
					break
				}
				f = f.fail
			}
			queue = append(queue, c)
		}
	}
}

// QueryFunc returns the index into strs of the every word point satisfying
// f(start, end).
func (t *ACTrie) QueryFunc(strs []rune, f func(int, int)) {
	var (
		ok      bool
		i, s, e int
		str     rune
		r, p    *ACTrie
		n       = t
	)
	for i, str = range strs {
		for {
			if n == nil {
				r = t // set root
				break
			}
			if r, ok = n.next[str]; ok {
				break
			}
			n = n.fail
		}
		n = r
		for p = n; p != nil; p = p.fail {
			if p.length > 0 {
				e = i + 1
				s = e - p.length
				f(s, e)
				if debug {
					log.Printf("actrie: find [%d:%d] = %s", s, e, string(strs)[s:e])
				}
			}
		}
	}
}
