package hashmap

const (
	ChainingMaxLoadFactor float64 = 0.75
)

type node struct {
	key   string
	value any
	next  *node
}

type ChainingMap struct {
	buckets []*node
	len     int
}

func NewChainingMap() *ChainingMap {
	return &ChainingMap{
		buckets: make([]*node, 8),
		len:     0,
	}
}

func (h *ChainingMap) Get(key string) (any, bool) {
	bidx := Hash(key) % len(h.buckets)
	node := h.buckets[bidx]
	for node != nil {
		if node.key == key {
			return node.value, true
		}
		node = node.next
	}
	return nil, false
}

func (h *ChainingMap) Put(key string, value any) {
	insert := put(h.buckets, key, value)
	if insert {
		h.len++
	}
	if h.loadFactor() > ChainingMaxLoadFactor {
		h.grow()
	}
}

func (h *ChainingMap) Delete(key string) bool {
	bidx := Hash(key) % len(h.buckets)
	currnode := h.buckets[bidx]
	if currnode == nil {
		return false
	}
	if currnode.key == key {
		h.buckets[bidx] = currnode.next
		h.len--
		return true
	}
	for currnode.next != nil {
		if currnode.next.key == key {
			todelete := currnode.next
			currnode.next = todelete.next
			todelete.next = nil
			h.len--
			return true
		}
		currnode = currnode.next
	}
	return false
}

func (h *ChainingMap) loadFactor() float64 {
	return float64(h.len) / float64(len(h.buckets))
}

func (h *ChainingMap) grow() {
	newbuckets := make([]*node, 2*len(h.buckets))
	for _, bucket := range h.buckets {
		currnode := bucket
		for currnode != nil {
			put(newbuckets, currnode.key, currnode.value)
			currnode = currnode.next
		}
	}
	h.buckets = newbuckets
}

func put(buckets []*node, key string, value any) bool {
	bidx := Hash(key) % len(buckets)
	currnode := buckets[bidx]
	for currnode != nil {
		if currnode.key == key {
			currnode.value = value
			return false
		}
		currnode = currnode.next
	}

	newnode := &node{key, value, nil}
	newnode.next = buckets[bidx]
	buckets[bidx] = newnode
	return true
}
