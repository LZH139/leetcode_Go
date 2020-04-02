package LongestWordInDictionary

import (
	 "container/list"
)

func longestWord(words []string) string {

}

type Node struct {
	s string
	children [26]*Node
	index int
}

func NodeConstructor(str string,i byte) *Node{
	return &Node{s:str,children:nil,index:i}
}

type Trie struct {
	root *Node
	words string
}

func (t *Trie) TrieConstructor(){
	t.root = NodeConstructor("0",0)
}

func (t *Trie) insert(word string){
	cur:=t.root
	length:=len(word)
	for i:=0;i<length;i++ {
		n := word[i]-'a'
		if cur.children[n] == nil {
			cur.children[n] = NodeConstructor(string(word[i]),word[0]-'a')
		}
		cur = cur.children[n]
	}
}

func (t *Trie) bfs() string {
	stringList := make([]string,26)
	queue := list.New()
	queue.PushBack(t.root)
	var flag bool
	for queue.Len()>0 {
		flag = false
		root:= queue.Front()
		queue.Remove(root)
		if cur,ok := root.Value.(*Node);ok{
			for i:=0;i<len(cur.children);i++ {
				if cur.children[i]!=nil {
					flag = true
					queue.PushBack(cur.children[i])
					stringList[i] += cur.children[i].s
				}
				if i == len(cur.children) -1 && !flag {

				}


			}
		}
	}


}
