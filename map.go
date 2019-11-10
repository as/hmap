package hmap

type M struct {
	b   []*n
	len int
	cap int
}

func (h *M) Put(k, v string) {
	h.ensure()
	h.put(k, v)
}
func (h *M) Get(k string) (interface{}, bool) {
	if len(h.b) == 0 {
		return nil, false
	}
	return h.get(k, false)
}
func (h *M) Del(k string) (interface{}, bool) {
	if len(h.b) == 0 {
		return nil, false
	}
	return h.get(k, true)
}

type n struct {
	k string
	v string
	p *n
}

func (h *M) put(k, v string) {
	hv := hash(k) & (len(h.b) - 1)
/*
always add the new value to head since its simpler code. 
if this becomes concurrent we know only the head of the
list can change from underneath us. a cas spin on the
head pointer will ensure that head is updated. however,
the next pointer will also have to be updated.

*/
	p := n{k: k, v: v, p: h.b[hv]}
	h.b[hv] = &p
	h.len++
}
func (h *M) get(k string, del bool) (interface{}, bool) {
	hv := hash(k) & (len(h.b) - 1)
	prev := h.b[hv]
	for n := h.b[hv]; n != nil; n = n.p {
		if n.k == k {
			if del {
				prev.p = n.p
				n.p = nil
				if n == h.b[hv] {
					h.b[hv] = nil
				}
				h.len--
			}
			return n.v, true
		}
	}
	return nil, false
}
func (h *M) ensure() {
	if h.cap > h.len/2 {
		return
	}
	tmp := h.b
	if h.cap == 0 {
		h.cap = 32
	} else if h.cap > 4*h.len {
		// too big
		h.cap /= 2
	} else {
		h.cap *= 2
	}
	h.len = 0
	h.b = make([]*n, h.cap)
	for _, n := range tmp {
		for ; n != nil; n = n.p {
			h.put(n.k, n.v)
		}
	}
}
