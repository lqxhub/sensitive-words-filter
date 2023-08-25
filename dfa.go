package main

import (
	"container/list"
	"unicode"
)

func NewDfaManager() *DfaManager {
	dm := &DfaManager{}
	return dm
}

type DfaManager struct {
	root map[string]*DfaNode
}

//initialize the dfa
func (dm *DfaManager) init() {
	dm.root = make(map[string]*DfaNode)
}

func (dm *DfaManager) InitWordTree(sentences *list.List) {
	dm.init()
	for v := sentences.Front(); v != nil; v = v.Next() {
		dm.buildTree(v.Value.(string))
	}
}

func (dm *DfaManager) InitWordTreeSlice(sentences []string) {
	dm.init()
	for _, sentence := range sentences {
		dm.buildTree(sentence)
	}
}

//add a filter rule
func (dm *DfaManager) buildTree(sentence string) {
	var node *DfaNode
	runeSentence := []rune(sentence)
	for i, v := range runeSentence {
		node = dm.buildNode(node, string(v))
		if i == len(runeSentence)-1 { //这个词结束
			node.eof = true
		}
	}
}

func (dm *DfaManager) buildNode(node *DfaNode, word string) *DfaNode {
	if node == nil {
		node = dm.root[word]
		if node == nil {
			node = &DfaNode{}
			dm.root[word] = node
		}
	} else {
		node = node.insertNode(word)
	}

	return node
}

//do check
func (dm *DfaManager) Check(str string) bool {
	_str := make([]rune, 0, len(str))
	for _, v := range str {
		if unicode.Is(unicode.Han, v) || unicode.IsLetter(v) {
			_str = append(_str, v)
		}
	}
	for i := 0; i < len(_str); i++ {
		var node *DfaNode
		for _, v := range _str[i:] {
			if node == nil {
				node = dm.root[string(v)]
			} else {
				node = node.next(string(v))
			}
			if node == nil {
				break
			}
			if node.eof {
				return true
			}
		}
	}
	return false
}

// DfaNode ///////////////////////////////////////////////////////////////////////////////////////////////////////////////
type DfaNode struct {
	eof   bool
	next_ map[string]*DfaNode
}

func (dn *DfaNode) next(word string) *DfaNode {
	return dn.next_[word]
}

func (dn *DfaNode) insertNode(word string) *DfaNode {
	if dn.next_ == nil {
		dn.next_ = make(map[string]*DfaNode)
	}
	node := dn.next_[word]
	if node == nil {
		node = &DfaNode{}
		dn.next_[word] = node
	}
	return node
}
