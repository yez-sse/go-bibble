package main

import (
	"fmt"
	"log"
)

// SlashDict 采用前缀树的结构
// 这里时简化了，key只考虑a-z的小写字母，value只考虑int整数
type SlashDict struct {
	children [26]*SlashDict
	value    int
}

func NewSlashDict() SlashDict {
	return SlashDict{}
}

func (s *SlashDict) Insert(key string, val int) {
	node := s
	for i := 0; i < len(key); i++ {
		if key[i] == '/' {
			continue
		}
		idx := key[i] - 'a'
		if node.children[idx] != nil {
			node = node.children[idx]
			return
		}
		if node.children[idx] == nil {
			node.children[idx] = &SlashDict{}
			if i == len(key)-1 {
				node.children[idx].value = val
			}
		}
	}
}

func (s *SlashDict) Pop(key string) bool {
	node := s
	for i := 0; i < len(key); i++ {
		if key[i] == '/' {
			continue
		}
		idx := key[i] - 'a'
		if node.children[idx] != nil {
			node = node.children[idx]
			if i == len(key)-1 {
				fmt.Println(node.value)
				node = nil
				return true
			}
		} else {
			fmt.Printf("KeyError: Dict under key %s does noe have key %s", key[:i-1], key[i:])
			return false
		}
	}
	return false
}

func (s *SlashDict) Get(key string) (bool, *SlashDict) {
	node := s
	for i := 0; i < len(key); i++ {
		if key[i] == '/' {
			continue
		}
		idx := key[i] - 'a'
		if node.children[idx] != nil {
			node = node.children[idx]
			if i == len(key)-1 {
				return true, node
			}
		} else {
			log.Printf("KeyError: Dict under key %s does noe have key %s", key[:i-1], key[i:])
			return false, nil
		}
	}
	return false, nil
}

func (s *SlashDict) IsEnd(node *SlashDict) bool {
	for i := 0; i < 26; i++ {
		if node.children[i] != nil {
			return false
		}
	}
	return true
}

func (s *SlashDict) DeepKeys(node *SlashDict, path string) (ret []string) {
	for i := 0; i < 26; i++ {
		if node.children[i] != nil {
			path += string('a'+i) + "/"
			if node.IsEnd(node.children[i]) {
				ret = append(ret)
			} else {
				node.DeepKeys(node.children[i], path)
			}
		}
	}
	return
}

func main() {
	s := NewSlashDict()
	s.Insert("a/b/c/x", 1)
	s.Insert("a/b/c/y", 2)
	s.Insert("a/b/d", 3)
	s.Insert("e/f", 4)

	isTrue, node := s.Get("a/b/c")
	if isTrue {
		if s.IsEnd(node) {
			fmt.Println(node.value)
		} else {
			fmt.Println(node)
		}
	}

	isTrue, node = s.Get("a/b/d")
	if isTrue {
		if s.IsEnd(node) {
			fmt.Println(node.value)
		} else {
			fmt.Println(node)
		}
	}

	isTrue, node = s.Get("a/b/k")
	if isTrue {
		if s.IsEnd(node) {
			fmt.Println(node.value)
		} else {
			fmt.Println(node)
		}
	}

	s.Pop("a/b/c/y")

	isTrue, node = s.Get("e/f/g")
	if isTrue {
		if s.IsEnd(node) {
			fmt.Println(node.value)
		} else {
			fmt.Println(node)
		}
	}
}
