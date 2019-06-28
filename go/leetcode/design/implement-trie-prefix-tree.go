package main

type Trie struct {
	IsEnd bool
	Child [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Child:[26]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	root := this
	for i:=0; i<len(word); i++ {
		val := root.Child[word[i]-'a']
		if  val != nil {
			root = val
		} else {
			tree := &Trie{Child:[26]*Trie{}}
			root.Child[word[i]] = tree
			root = tree
		}
	}
	if !root.IsEnd {
		root.IsEnd = true
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	root := this
	for i:=0; i<len(word); i++ {
		val := root.Child[word[i]-'a']
		if  val != nil {
			root = val
		} else {
			return false
		}
	}
	return root.IsEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for i:=0; i<len(prefix); i++ {
		val := root.Child[prefix[i]-'a']
		if  val != nil {
			root = val
		} else {
			return false
		}
	}
	return true
}
