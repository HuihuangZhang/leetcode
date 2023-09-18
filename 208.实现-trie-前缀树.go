/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

package leetcode

// @lc code=start
// type trieNode struct {
// 	data map[rune]*trieNode
// 	isEnd bool
// }

// func newTrieNode() *trieNode {
// 	return &trieNode {
// 		data: make(map[rune]*trieNode),
// 		isEnd: false,
// 	} 
// }

// type Trie struct {
// 	nodes *trieNode
// }

// func Constructor() Trie {
// 	t := Trie{
// 		nodes: newTrieNode(),
// 	}
// 	return t
// }

// func (this *Trie) Insert(word string)  {
// 	node := this.nodes	
// 	for _, char := range word {
// 		rchar := rune(char)
// 		childNode, ok := node.data[rchar]
// 		if !ok {
// 			// 这里处理如果有新节点的情况
// 			childNode = newTrieNode()
// 			node.data[rchar] = childNode 
// 		}
// 		node = childNode
// 	}
// 	// 先插入 apple，再插入 app 的情况
// 	node.isEnd = true
// }

// func (this *Trie) Search(word string) bool {
// 	node := this.findNode(word)
// 	return node != nil && node.isEnd
// }

// func (this *Trie) StartsWith(prefix string) bool {
// 	node := this.findNode(prefix)
// 	return node != nil
// }

// func (this *Trie) findNode(word string) *trieNode {
// 	node := this.nodes
// 	for _, char := range word {
// 		rchar := rune(char)
// 		childNode, ok := node.data[rchar]
// 		if !ok {
// 			return nil
// 		}
// 		node = childNode
// 	}
// 	return node
// }

/// 因为字典树有个小写字母的限制，所以直接用 [26]xxx 可以提高
/// access performance。
/// 但因为每次都会初始化 26 个指针，所以会耗费一定的内存。
type trieNode struct {
	data [26]*trieNode
	isEnd bool
}

func newTrieNode() *trieNode {
	return &trieNode {
		data: [26]*trieNode{},
		isEnd: false,
	} 
}

type Trie struct {
	nodes *trieNode
}

func Constructor() Trie {
	t := Trie{
		nodes: newTrieNode(),
	}
	return t
}

func (this *Trie) Insert(word string)  {
	node := this.nodes	
	for _, rchar := range word {
		idx := int(rchar) - 97
		childNode := node.data[idx]
		if childNode == nil {
			// 这里处理如果有新节点的情况
			childNode = newTrieNode()
			node.data[idx] = childNode 
		}
		node = childNode
	}
	// 先插入 apple，再插入 app 的情况
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.findNode(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.findNode(prefix)
	return node != nil
}

func (this *Trie) findNode(word string) *trieNode {
	node := this.nodes
	for _, rchar := range word {
		idx := int(rchar) - 97
		childNode := node.data[idx]
		if childNode == nil {
			return nil
		}
		node = childNode
	}
	return node
}
/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

