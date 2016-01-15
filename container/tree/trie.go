package tree

import (
	"log"
)

// Trie Tree
type Trie struct {
	next map[rune]*Trie
	end  bool
}

// NewTrie new a Trie Tree.
func NewTrie() *Trie {
	return &Trie{
		next: make(map[rune]*Trie),
	}
}

// Insert insert word into trie tree.
func (t *Trie) Insert(vs ...string) {
	var i int
	for i = 0; i < len(vs); i++ {
		t.insert([]rune(vs[i]))
	}
}

func (t *Trie) insert(strs []rune) {
	var (
		n      *Trie
		ok     bool
		str    rune
		i      = 0
		strLen = len(strs)
	)
	if debug {
		log.Printf("trie: insert strs len: %d\n", strLen)
	}
	for i = 0; i < strLen; i++ {
		str = strs[i]
		if n, ok = t.next[str]; !ok {
			n = NewTrie()
			t.next[str] = n
			if debug {
				log.Printf("trie: insert new trie: %s\n", string(str))
			}
		}
		t = n
	}
	n.end = true
}

// Has tests whether the string s exists in trie tree.
func (t *Trie) Has(s string) bool {
	return t.has([]rune(s), true)
}

// HasPrefix tests whether the string s's prefix in trie tree.
func (t *Trie) HasPrefix(s string) bool {
	return t.has([]rune(s), false)
}

func (t *Trie) has(strs []rune, all bool) bool {
	var (
		ok     bool
		str    rune
		i      = 0
		strLen = len(strs)
	)
	if debug {
		log.Printf("trie: find strs len: %d\n", strLen)
	}
	for i = 0; i < strLen; i++ {
		str = strs[i]
		if t, ok = t.next[str]; !ok {
			if debug {
				log.Printf("trie: not find str: %s\n", string(str))
			}
			return false
		}
		if debug {
			log.Printf("trie: find str: %s\n", string(str))
		}
	}
	return !all || t.end
}
