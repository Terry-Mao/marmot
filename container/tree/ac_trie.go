package tree

import (
	"log"
)

type ACTrie struct {
	next   map[rune]*ACTrie
	fail   *ACTrie
	length int
}

func NewACTrie() *ACTrie {
	return &ACTrie{
		next: make(map[rune]*ACTrie),
	}
}

func (t *ACTrie) initFail() {
	var (
		ok      bool
		str     rune
		move    int
		c, f, r *ACTrie
		queue   = []*ACTrie{t}
	)
	t.fail = nil
	for move = 0; move < len(queue); move++ {
		t = queue[move]             // bfs
		for str, c = range t.next { // get child node
			f = t.fail // parent fail
			for {
				if f == nil {
					r = t // set root
					break
				}
				if r, ok = f.next[str]; ok {
					break
				}
				f = f.fail
			}
			c.fail = r
			queue = append(queue, c)
		}
	}
}

func (t *ACTrie) QueryFunc(strs []rune, f func(int, int)) {
	var (
		ok      bool
		i, s, e int
		r, p    *ACTrie
		str     rune
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

func (t *ACTrie) Insert(strs []rune) {
	var (
		n      *ACTrie
		ok     bool
		str    rune
		i      = 0
		strLen = len(strs)
	)
	if debug {
		log.Printf("actrie: insert strs len: %d\n", strLen)
	}
	for i = 0; i < strLen; i++ {
		str = strs[i]
		if n, ok = t.next[str]; !ok {
			n = NewACTrie()
			t.next[str] = n
			if debug {
				log.Printf("actrie: insert new trie\n")
			}
		}
		t = n
	}
	n.length = strLen
}
