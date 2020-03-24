# 前缀树

## 一道面试题

将以下格式转换：

```
"{"A": 1, "B.A": 2, "B.B": 3, "CC.D.E": 4, "CC.D.F": 5}"
{"A":1,"B":{"A":2,"B":3},"CC":{"D":{"E":4,"F":5}}}
```

想了一下可以用前缀树解决：

<details><summary>main.go</summary>
<p>

```go
package main

import (
  "strconv"
  "fmt"
)

type node struct {
	children []*node
	data     int
	value    string
}

func (n *node) addChild(value string) *node {
	var child = n.search(value)
	if child == nil {
		child = &node{
			children: make([]*node, 0),
			value:    value,
		}
		n.children = append(n.children, child)
	}
	return child
}

func (n *node) search(value string) *node {
	for _, child := range n.children {
		if child.value == value {
			return child
		}
	}
	return nil
}

func (n *node) addPath(paths []string, data int) {
	child := n
	for i := 0; i < len(paths); i++ {
		child = child.addChild(paths[i])
	}
	child.data = data
}

func wrap(value string) string { return "'" + value + "'" }

func (n *node) flatten() string {
	res := ""

	if len(n.children) > 0 {
		r := ""
		if n.value != "" {
			r = wrap(n.value) + ":" + "{"
		} else {
			r = "{"
		}
		for _, child := range n.children {
			r += child.flatten() + ","
		}
		res += r[:len(r)-1] + "},"
	} else {
		res += wrap(n.value) + ":" + strconv.Itoa(n.data) + ","
	}

	return res[:len(res)-1]
}

func main(t *testing.T) {

	var strMap = map[string]int{
		"A": 1, "B.A": 2, "B.B": 3, "CC.D.E": 4, "CC.D.F": 5,
	}

	var trie = node{
		value:    "",
		children: make([]*node, 0),
	}

	for k, m := range strMap {
		trie.addPath(strings.Split(khttps://www.zhihu.com/question/27168319, "."), m)
	}

	fmt.Println(trie.flatten())
}
```
</p>
</details>

## 应用场景

之前用过前缀树在于web框架中的路由检索，其实前缀树又叫字典树，也可以用于字符串的匹配。那么对于字符串的匹配，用hash表好还是前缀树好呢？从工程上看什么场景下应该怎么去选择呢？

[为什么360面试官说Trie树没用？](https://www.zhihu.com/question/27168319)

为什么说前缀树是cache-friendly的呢？

为什么查询字符串时前缀树比hash table查询要快？

## 压缩

前缀树如何进行压缩？
前缀树的优缺点是什么？

## Leetcode

- [208. Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/#/description)
- [211. Add and Search Word - Data structure design](https://leetcode.com/submissions/detail/98755427/)
- [472. Concatenated Words](https://leetcode.com/problems/concatenated-words/#/description)
- [212. Word Search II](https://leetcode.com/problems/word-search-ii/#/description)
- [421. Maximum XOR of Two Numbers in an Array](https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/#/description)

## 扩展

- AC Automata on double array trie
- rank-select
- DFUDS trie
- louds trie
- patritia-trie
- recursive-patricia-trie
- [vector trie](https://github.com/chenhg5/CS-Interview/blob/master/dsa/immutable.md)

