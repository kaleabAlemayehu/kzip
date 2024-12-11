package huffman

type HuffBaseNode interface {
	IsLeaf() bool
	Weight() int
}

type HuffLeafNode struct {
	element byte
	weight  int
}

func NewLeafNode(el byte, w int) *HuffLeafNode {
	return &HuffLeafNode{element: el, weight: w}
}

func (h *HuffLeafNode) Value() byte {
	return h.element
}

func (h *HuffLeafNode) Weight() int {
	return h.weight
}

func (h *HuffLeafNode) IsLeaf() bool {
	return true
}

type HuffInternalNode struct {
	weight int
	left   HuffBaseNode
	right  HuffBaseNode
}

func NewInternalNode(l HuffBaseNode, r HuffBaseNode, w int) *HuffInternalNode {
	return &HuffInternalNode{weight: w, left: l, right: r}
}

func (h *HuffInternalNode) Weight() int {
	return h.weight
}

func (h *HuffInternalNode) Left() HuffBaseNode {
	return h.left
}

func (h *HuffInternalNode) Right() HuffBaseNode {
	return h.right
}

func (h *HuffInternalNode) IsLeaf() bool {
	return false
}

type HuffTree struct {
	root HuffBaseNode
}

func NewFreshTree(el byte, w int) *HuffTree {
	return &HuffTree{root: NewLeafNode(el, w)}
}

func NewTree(l HuffBaseNode, r HuffBaseNode, w int) *HuffTree {
	return &HuffTree{root: NewInternalNode(l, r, w)}
}

func (t *HuffTree) CompareTo(tree interface{}) int {
	that := tree.(HuffTree)
	if t.root.Weight() < that.root.Weight() {
		return -1
	} else if t.root.Weight() == that.root.Weight() {
		return 0
	} else {
		return 1
	}
}
