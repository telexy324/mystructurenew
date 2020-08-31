package trie

type Node struct {
	last bool
	parent *Node
	children map[int32]*Node
}

type Trie struct {
	root *Node
}

func (trie *Trie) Init()  {
	trie.root = &Node{children: map[int32]*Node{}}
}

func (trie *Trie) Add(item string)  {
	currentNode := trie.root
	for _,r := range item {
		if _,ok:= currentNode.children[r];ok {
			currentNode = c
		}
	}
}