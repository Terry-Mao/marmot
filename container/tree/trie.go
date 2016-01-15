package tree

import (
	"log"
)

type Trie struct {
	next map[rune]*Trie
	end  bool
}

func NewTrie() *Trie {
	return &Trie{
		next: make(map[rune]*Trie),
	}
}

func (t *Trie) Insert(strs []rune) {
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
				log.Printf("trie: insert new trie\n")
			}
		}
		t = n
	}
	n.end = true
}

func (t *Trie) Has(strs []rune) bool {
	return t.has(strs, true)
}

func (t *Trie) HasPrefix(strs []rune) bool {
	return t.has(strs, false)
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
