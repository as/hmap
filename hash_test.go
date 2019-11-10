package hmap_test

import (
	"github.com/as/hmap"
	"testing"
)

func TestBasic(t *testing.T){
	h := hmap.M{}
	h.Put("Hello", "a")
	v, ok := h.Get("Hello")
	if !ok || v != "a"{
		t.Fatal(v)
	}

	h.Put("Hello", "b")
	v, ok = h.Get("Hello")
	if !ok || v != "b"{
		t.Fatal(v)
	}

	h.Put("Hello", "c")
	v, ok = h.Get("Hello")
	if !ok || v != "c"{
		t.Fatal(v)
	}
	v, ok = h.Del("Hello")
	if !ok {
		t.Fatal(v)
	}
	v, ok = h.Del("Hello")
	if ok {
		t.Fatal(v)
	}
}