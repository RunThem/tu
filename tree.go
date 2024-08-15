package tu

import (
	"fmt"
	"iter"
)

type node[K, V any] struct {
	left   *node[K, V]
	right  *node[K, V]
	parent *node[K, V]
	key    K
	val    V
	height int
}

func (mod *node[K, V]) getHeight() int {
	if mod == nil {
		return 0
	}

	return mod.height
}

func (mod *node[K, V]) maxHeight() int {
	return max(mod.left.getHeight(), mod.right.getHeight())
}

func (mod *node[K, V]) string(prefix string, isTail bool, str *string) {
	if mod.right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		mod.right.string(newPrefix, false, str)
	}

	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += fmt.Sprintf("%v: %v\n", mod.key, mod.val)

	if mod.left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}

		mod.left.string(newPrefix, true, str)
	}
}

type Tree[K, V any] struct {
	root *node[K, V]
	len  int
	cmp  func(K, K) int
}

func NewTree[K, V any](cmp func(K, K) int) *Tree[K, V] {
	if cmp == nil {
		return nil
	}

	return &Tree[K, V]{cmp: cmp}
}

func (mod *Tree[K, V]) Len() int {
	return mod.len
}

func (mod *Tree[K, V]) IsEmpty() bool {
	return mod.len == 0
}

func (mod *Tree[K, V]) Clear() {
	mod.root = nil
	mod.len = 0
}

func (mod *Tree[K, V]) IsExist(key K) bool {
	n := mod.root
	for n != nil {
		resutl := mod.cmp(key, n.key)
		if resutl == 0 {
			break
		}

		if resutl < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}

	return n != nil
}

func (mod *Tree[K, V]) At(key K) V {
	var val V

	n := mod.root
	for n != nil {
		resutl := mod.cmp(key, n.key)
		if resutl == 0 {
			break
		}

		if resutl < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}

	if n != nil {
		val = n.val
	}

	return val
}

func (mod *Tree[K, V]) Next(key K) V {
	var val V

	n := mod.root
	for n != nil {
		resutl := mod.cmp(key, n.key)
		if resutl == 0 {
			break
		}

		if resutl < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}

	if n != nil {
		val = mod.rangeNext(n).val
	}

	return val
}

func (mod *Tree[K, V]) Prev(key K) V {
	var val V

	n := mod.root
	for n != nil {
		resutl := mod.cmp(key, n.key)
		if resutl == 0 {
			break
		}

		if resutl < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}

	if n != nil {
		val = mod.rangePrev(n).val
	}

	return val
}

func (mod *Tree[K, V]) Re(key K, val V) {
	n := mod.root
	for n != nil {
		resutl := mod.cmp(key, n.key)
		if resutl == 0 {
			break
		}

		if resutl < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}

	if n != nil {
		n.val = val
	}
}

func (mod *Tree[K, V]) Pop(key K) V {
	return mod.pop(key)
}

func (mod *Tree[K, V]) Put(key K, val V) {
	mod.put(key, val)
}

func (mod *Tree[K, V]) String() string {
	str := "Tree Tree\n"
	if !mod.IsEmpty() {
		mod.root.string("", true, &str)
	}

	return str
}

func (mod *Tree[K, V]) Range(order bool) iter.Seq2[K, V] {
	var fn iter.Seq2[K, V]
	if order {
		fn = func(yield func(K, V) bool) {
			for n := mod.rangeFrist(); n != nil; n = mod.rangeNext(n) {
				if !yield(n.key, n.val) {
					return
				}
			}
		}
	} else {
		fn = func(yield func(K, V) bool) {
			for n := mod.rangeLast(); n != nil; n = mod.rangePrev(n) {
				if !yield(n.key, n.val) {
					return
				}
			}
		}
	}

	return fn
}

func (mod *Tree[K, V]) Map(fn func(key K, val V) V) *Tree[K, V] {
	tree := NewTree[K, V](mod.cmp)

	for k, v := range mod.Range(true) {
		tree.Put(k, fn(k, v))
	}

	return tree
}

func (mod *Tree[K, V]) Filter(fn func(key K, val V) bool) *Tree[K, V] {
	tree := NewTree[K, V](mod.cmp)

	for k, v := range mod.Range(true) {
		if fn(k, v) {
			tree.put(k, v)
		}
	}

	return tree
}

func (mod *Tree[K, V]) IsAny(fn func(key K, val V) bool) bool {
	for k, v := range mod.Range(true) {
		if fn(k, v) {
			return true
		}
	}

	return false
}

func (mod *Tree[K, V]) IsAll(fn func(key K, val V) bool) bool {
	for k, v := range mod.Range(true) {
		if !fn(k, v) {
			return false
		}
	}

	return true
}

func (mod *Tree[K, V]) Copy() *Tree[K, V] {
	tree := NewTree[K, V](mod.cmp)

	for k, v := range mod.Range(true) {
		tree.put(k, v)
	}

	return tree
}

/*
 * api
 *
 */
func (mod *Tree[K, V]) rangeFrist() *node[K, V] {
	n := mod.root

	if n == nil {
		return nil
	}

	for n.left != nil {
		n = n.left
	}

	return n
}

func (mod *Tree[K, V]) rangeNext(n *node[K, V]) *node[K, V] {
	if n.right != nil {
		n = n.right
		for n.left != nil {
			n = n.left
		}
	} else {
		for {
			last := n
			n = n.parent

			if n == nil || n.left == last {
				break
			}
		}
	}

	return n
}

func (mod *Tree[K, V]) rangeLast() *node[K, V] {
	n := mod.root

	if n == nil {
		return nil
	}

	for n.right != nil {
		n = n.right
	}

	return n
}

func (mod *Tree[K, V]) rangePrev(n *node[K, V]) *node[K, V] {
	if n.left != nil {
		n = n.left
		for n.right != nil {
			n = n.right
		}
	} else {
		for {
			last := n
			n = n.parent

			if n == nil || n.right == last {
				break
			}
		}
	}

	return n
}

func (mod *Tree[K, V]) put(key K, val V) {
	link := &mod.root

	var parent *node[K, V]
	for *link != nil {
		parent = *link

		result := mod.cmp(key, parent.key)
		if result == 0 {
			return // tree put node already exists.
		}

		if result < 0 {
			link = &(parent.left)
		} else {
			link = &(parent.right)
		}
	}

	n := &node[K, V]{left: nil, right: nil, parent: parent, key: key, val: val, height: 1}
	*link = n

	for n = n.parent; n != nil; n = n.parent {
		lh := n.left.getHeight()
		rh := n.right.getHeight()
		h := max(lh, rh) + 1
		diff := lh - rh

		if n.height == h {
			break
		}

		n.height = h

		if diff <= -2 {
			n = mod.fix(true, n)
		} else if diff >= 2 {
			n = mod.fix(false, n)
		}
	}

	mod.len++
}

func (mod *Tree[K, V]) pop(key K) V {
	var val V
	var p *node[K, V]

	n := mod.root
	for n != nil {
		result := mod.cmp(key, n.key)
		if result == 0 {
			break
		}

		if result < 0 {
			n = n.left
		} else {
			n = n.right
		}
	}

	if n == nil {
		return val
	} else {
		val = n.val
	}

	if n.left != nil && n.right != nil {
		p = mod.popAnd(n)
	} else {
		p = mod.popOr(n)
	}

	if p != nil {
		for n != nil {

			lh := n.left.getHeight()
			rh := n.right.getHeight()
			h := n.maxHeight() + 1
			diff := lh - rh

			if n.height != h {
				n.height = h
			} else if diff >= -1 && diff <= -1 {
				break
			}

			if diff <= -2 {
				n = mod.fix(true, n)
			} else if diff >= 2 {
				n = mod.fix(false, n)
			}

			n = n.parent
		}
	}

	mod.len--

	return val
}

func (mod *Tree[K, V]) popAnd(n *node[K, V]) *node[K, V] {
	old := n
	n = n.right

	for l := n.left; l != nil; n = l {
	}

	ch := n.right
	p := n.parent

	if ch != nil {
		ch.parent = p
	}

	mod.replace(p, n, ch)

	if n.parent == old {
		p = n
	}

	n.left, n.right, n.parent, n.height = old.left, old.right, old.parent, old.height

	mod.replace(old.parent, old, n)
	old.left.parent = n

	if old.right != nil {
		old.right.parent = n
	}

	return p
}

func (mod *Tree[K, V]) popOr(n *node[K, V]) *node[K, V] {
	ch := n.left
	if ch == nil {
		ch = n.right
	}

	p := n.parent
	mod.replace(p, n, ch)

	if ch != nil {
		ch.parent = p
	}

	return p
}

func (mod *Tree[K, V]) fix(op bool, n *node[K, V]) *node[K, V] {
	var th *node[K, V]

	// left
	if op {
		th = n.right
	} else {
		th = n.left
	}

	lh := th.left.getHeight()
	rh := th.right.getHeight()

	// left
	if op {
		if lh > rh {
			th = mod.rotate(false, th)
			th.right.height = th.right.maxHeight() + 1
			th.height = th.maxHeight() + 1
		}

		n = mod.rotate(true, n)
		n.left.height = n.left.maxHeight() + 1
		n.height = n.maxHeight() + 1
	} else { // right
		if lh < rh {
			th = mod.rotate(true, th)
			th.left.height = th.left.maxHeight() + 1
			th.height = th.maxHeight() + 1
		}

		n = mod.rotate(false, n)
		n.right.height = n.right.maxHeight() + 1
		n.height = n.maxHeight() + 1
	}

	return n
}

func (mod *Tree[K, V]) rotate(op bool, n *node[K, V]) *node[K, V] {
	var th *node[K, V]
	p := n.parent

	// left
	if op {
		th = n.right
		n.right = th.left

		if th.left != nil {
			th.left.parent = n
		}

		th.left, th.parent = n, p
	} else { // right
		th = n.left
		n.left = th.right

		if th.right != nil {
			th.right.parent = n
		}

		th.right, th.parent = n, p
	}

	mod.replace(p, n, th)

	n.parent = th

	return th
}

func (mod *Tree[K, V]) replace(p, o, n *node[K, V]) {
	if p == nil {
		mod.root = n
		return
	}

	if p.left == o {
		p.left = n
	} else {
		p.right = n
	}
}
