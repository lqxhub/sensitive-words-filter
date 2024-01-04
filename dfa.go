//敏感词检测

package sensitive_word_filter_filter

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"io"
	"os"
	"unicode"
)

//根据文件初始化 敏感词监测
func InitSensitiveWordWithFile(filePath string) (*SensitiveWordManager, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	sw := NewSensitiveWordManager()

	reader := bufio.NewReader(file)
	for {
		if line, _, err := reader.ReadLine(); err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read file error:%s", err.Error())
		} else {
			sw.buildTree(string(bytes.TrimFunc(line, func(r rune) bool { //去掉不可见字符
				return !unicode.IsPrint(r)
			})))
		}
	}
	return sw, nil

}

func NewSensitiveWordManager() *SensitiveWordManager {
	sw := &SensitiveWordManager{
		root: make(map[rune]*SensitiveWordNode),
	}
	return sw
}

func NewSensitiveWordManagerWithList(sentences *list.List) *SensitiveWordManager {
	sw := &SensitiveWordManager{
		root: make(map[rune]*SensitiveWordNode, sentences.Len()),
	}
	sw.InitWordTree(sentences)
	return sw
}

type SensitiveWordManager struct {
	root map[rune]*SensitiveWordNode
}

func (dm *SensitiveWordManager) InitWordTree(sentences *list.List) {
	for v := sentences.Front(); v != nil; v = v.Next() {
		dm.buildTree(v.Value.(string))
	}
}

func (dm *SensitiveWordManager) InitWordTreeSlice(sentences []string) {
	for _, sentence := range sentences {
		dm.buildTree(sentence)
	}
}

func (dm *SensitiveWordManager) buildTree(sentence string) {
	var node *SensitiveWordNode
	runeSentence := []rune(sentence)
	for i, v := range runeSentence {
		node = dm.buildNode(node, v)
		if i == len(runeSentence)-1 { //这个词结束
			node.eof = true
		}
	}
}

func (dm *SensitiveWordManager) buildNode(node *SensitiveWordNode, word rune) *SensitiveWordNode {
	if node == nil {
		node = dm.root[word]
		if node == nil {
			node = &SensitiveWordNode{}
			dm.root[word] = node
		}
	} else {
		node = node.insertNode(word)
	}

	return node
}

//检查是否含有敏感词
func (dm *SensitiveWordManager) HasSensitiveWords(str string) bool {
	_str := []rune(str)
	for i := 0; i < len(_str); i++ {
		var node *SensitiveWordNode
		for _, v := range _str[i:] {
			isSymbol := false
			if unicode.IsPunct(v) || unicode.IsSymbol(v) {
				isSymbol = true
			}
			if node == nil { //第一个字符
				node = dm.root[v]
			} else { //后面的字符
				_node := node.next(v)
				if _node == nil && isSymbol {
					//如果没有找到,并且是标点符号,则跳过标点符号检查
					//这里是为了防止标点符号被当做敏感词
					//例如: '曹--尼,,玛' 这里的`-`和`,`都会被跳过
					// www.baidu.com 这里的 `.` 也是敏感词的一部分,也会被检查出来
					continue
				}
				node = _node
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

// SensitiveWordNode ///////////////////////////////////////////////////////////////////////////////////////////////////
type SensitiveWordNode struct {
	eof   bool
	next_ map[rune]*SensitiveWordNode
}

func (dn *SensitiveWordNode) next(word rune) *SensitiveWordNode {
	return dn.next_[word]
}

func (dn *SensitiveWordNode) insertNode(word rune) *SensitiveWordNode {
	if dn.next_ == nil {
		dn.next_ = make(map[rune]*SensitiveWordNode)
	}
	node := dn.next_[word]
	if node == nil {
		node = &SensitiveWordNode{}
		dn.next_[word] = node
	}
	return node
}
