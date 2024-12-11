package huffman

type HuffHeap []HuffLeafNode

func NewHuffHeap(hash_map map[byte]int) *HuffHeap {
	hh := make(HuffHeap, len(hash_map))
	i := 0
	for key, value := range hash_map {
		hh[i] = HuffLeafNode{
			element: key,
			weight:  value,
		}
		i++
	}
	return &hh
}

func (h *HuffHeap) Len() int {
	return len(*h)
}

func (h *HuffHeap) Less(i, j int) bool {
	return (*h)[i].weight < (*h)[j].weight
}

func (h *HuffHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *HuffHeap) Push(x any) {
	newLeaf := x.(HuffLeafNode)
	*h = append(*h, newLeaf)
}

func (h *HuffHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = item
	*h = old[0 : n-1]
	return item
}
