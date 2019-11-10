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
