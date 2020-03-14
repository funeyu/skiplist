package skiplist

import (
	"math/bits"
	"math/rand"
)

type Element interface {
	// 相等返回0， 大于other 返回 1 小于other 返回 -1
	Compare(other interface{}) int
}

type  minElement struct {}

func (m minElement) Compare(other interface{}) int{
	return 1
}

var MINELEMENT  = minElement{}

type SkipListNode struct {
	value Element
	pre *SkipListNode
	next *SkipListNode
	below *SkipListNode
}

type  SkipList struct {
	maxLevel int
	height int
	head []*SkipListNode
}

func Generate(maxLevel int) *SkipList {
	n := &SkipListNode{
		value: MINELEMENT,
	}
	var h []*SkipListNode
	h = append(h, n)
	return &SkipList {
		maxLevel: maxLevel,
		height: 1,
		head: h,
	}
}

func (s *SkipList) level() int {
	maxLevel := s.maxLevel
	level := maxLevel - 1
	x := rand.Uint64() & ((1 << uint(maxLevel-1)) - 1)
	zeroes := bits.TrailingZeros64(x)
	if zeroes <= maxLevel {
		level = zeroes
	}

	return level + 1
}

func (s *SkipList) isEmpty() bool {
	return s.head[0].next == nil
}

// 根据初始值新建一level，返回该level的 head节点指针
func (s *SkipList) newLevel(newNode *SkipListNode, belowHead *SkipListNode) *SkipListNode{
	head := &SkipListNode {
		value: MINELEMENT, next: newNode, below: belowHead, pre: nil,
	}

	return head
}

func (s *SkipList) Insert(n Element) {
	node := &SkipListNode{
		value: n,
	}
	if s.isEmpty() {
		s.head[0].next = node
		node.pre = s.head[0]
	} else {
		walked := make([]*SkipListNode, s.height)  // 记录下探到1层经历的节点
		c := s.head[s.height-1]
		for i := s.height; i > 0; i -- {
			for c != nil { // todo 换方式
				if c.value == n || c.next == nil {
					walked[i-1] = c
					c = c.below
					break
				}

				if c.next.value.Compare(n) <= 0 {
					c = c.next
				} else {
					walked[i-1] = c
					c = c.below
					break
				}
			}
		}
		l := s.level()
		for n:= 1; n <= l && n <= s.height + 1; n ++ {
			if n == s.height + 1 { // 代表新增的跳跃表已经到了突破之前的层数，需要新建一层
				head :=s.newLevel(node, s.head[s.height-1])
				s.head = append(s.head, head)
				s.height = s.height + 1
				break
			} else { // 双向链的insert操作
				next := walked[n-1].next
				walked[n-1].next = node
				node.next = next
				node.pre = walked[n-1]
				if next != nil {
					next.pre = node
				}
				node = &SkipListNode{
					value: node.value,
					below: node,
				}
			}
		}
	}
}

func (s *SkipList) Find(n Element) *SkipListNode{
	start := s.head[s.height -1]
	for i:= s.height; i > 0; i -- {
		for ; start.next != nil && start.next.value.Compare(n) <= 0; start = start.next {}
		if start.value == n {
			return start
		}
		start = start.below
	}
	if start !=nil && start.value == n {
		return start
	}
	return nil
}

func (s *SkipList) Delete(n Element) {
	found := s.Find(n)
	for found != nil {
		found.pre.next = found.next
		found = found.below
	}
}