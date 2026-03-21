package hashmap

const (
	ProbingMaxLoadFactor = 0.7
)

type state int

const (
	occupied state = iota
	deleted
)

type item struct {
	key   string
	value any
	state state
}

type ProbingMap struct {
	slots      []*item
	len        int
	cap        int
	tombstones int
}

func NewProbingMap() *ProbingMap {
	return &ProbingMap{
		slots:      make([]*item, 8),
		len:        0,
		cap:        8,
		tombstones: 0,
	}
}

func (h *ProbingMap) Get(key string) (any, bool) {
	idx := h.probe(key, 1)
	if idx == -1 || h.slots[idx] == nil || h.slots[idx].state == deleted {
		return nil, false
	}
	return h.slots[idx].value, true
}

func (h *ProbingMap) Put(key string, value any) {
	if h.loadfactor() > ProbingMaxLoadFactor {
		h.grow()
	}
	idx := h.probe(key, 1)
	if h.slots[idx] == nil {
		h.slots[idx] = &item{key, value, occupied}
		h.len++
	} else {
		h.slots[idx].key = key
		h.slots[idx].value = value
		if h.slots[idx].state == deleted {
			h.tombstones--
			h.slots[idx].state = occupied
			h.len++
		}
	}
}

func (h *ProbingMap) Delete(key string) {
	idx := h.probe(key, 1)
	if idx != -1 && h.slots[idx] != nil && h.slots[idx].state == occupied {
		h.slots[idx].state = deleted
		h.slots[idx].key = ""
		h.slots[idx].value = nil
		h.len--
		h.tombstones++
	}
}

func (h *ProbingMap) loadfactor() float64 {
	return (float64(h.len) + float64(h.tombstones)) / float64(h.cap)
}

func (h *ProbingMap) probe(key string, step int) int {
	firsttomb := -1
	slotsprobed := 0
	inc := 0
	for {
		probeidx := (Hash(key) + inc) % len(h.slots)
		if h.slots[probeidx] == nil {
			if firsttomb != -1 {
				return firsttomb
			}
			return probeidx
		}
		if h.slots[probeidx].state == deleted && firsttomb == -1 {
			firsttomb = probeidx
		}
		if h.slots[probeidx].state == occupied && h.slots[probeidx].key == key {
			return probeidx
		}
		slotsprobed++
		if slotsprobed == len(h.slots) {
			break
		}
		inc += step
	}
	return -1
}

func (h *ProbingMap) grow() {
	newcap := 2 * h.cap
	oldslots := h.slots
	h.slots = make([]*item, 2*h.cap)
	h.cap = newcap
	h.tombstones = 0
	for _, s := range oldslots {
		if s != nil && s.state == occupied {
			probeidx := h.probe(s.key, 1)
			h.slots[probeidx] = &item{s.key, s.value, occupied}
		}
	}
}
