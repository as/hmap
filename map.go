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

func (h *M) put(k, v string) {}
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
func (h *M) ensure() {}
