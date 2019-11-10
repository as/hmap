package hmap

type M struct {
	b   []*n
	len int
	cap int
}

type n struct {
	k string
	v string
	p *n
}

func (h *M) put(k, v string) {}
func (h *M) get(k string, del bool) (interface{}, bool) {}
func (h *M) ensure() {}}