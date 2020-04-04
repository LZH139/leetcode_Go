package LongestWordInDictionary

import (
	 "container/list"
)

func longestWord(words []string) string {
	 t := Trie{
		root:  NodeConstructor("0"),
	}
	for i:=0;i<len(words);i++ {
		t.insert(words[i],i)
	}
	t.words = words

	return t.bfs()
}

type Node struct {
	s string
	children [26]*Node
	index int
}

func NodeConstructor(str string) *Node{
	return &Node{s:str,index:-1}
}

type Trie struct {
	root *Node
	words []string
}

func (t *Trie) insert(word string,index int){
	cur:=t.root
	length:=len(word)
	for i:=0;i<length;i++ {
		n := 25-(word[i]-'a')
		if cur.children[n] == nil {
			cur.children[n] = NodeConstructor(string(word[i]))
		}
		cur = cur.children[n]
	}
	cur.index = index
}

func (t *Trie) bfs() string {
	queue := list.New()
	queue.PushBack(t.root)
	var finalroot *Node
	for queue.Len()>0 {
		root:= queue.Front()
		queue.Remove(root)
		if cur,ok := root.Value.(*Node);ok{
			for i:=0;i<len(cur.children);i++ {
				if cur.children[i]!=nil && cur.children[i].index!=cur.index && cur.children[i].index!=-1{
					queue.PushBack(cur.children[i])
				}
				finalroot = cur
			}
		}

	}
	return t.words[finalroot.index]

}
