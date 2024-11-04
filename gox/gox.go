package gox

import "reflect"

// The component Type must be a string or Component
type Type any

type P map[string]any

type Node interface {
	node() any
}

type node struct {
	value any
}

func (n node) node() any {
	return n.value
}

func N(value any) Node {
	return node{value}
}

type Element struct {
	typ   Type
	props P
}

func (e Element) node() any {
	return e
}

// Component must be func(c Context, props any) Node
type Component any

func E(typ Type, props P, children ...Node) Element {
	kind := reflect.TypeOf(typ).Kind()
	switch kind {
	case reflect.String, reflect.Func:
		if props == nil {
			props = make(P)
		}
		props["children"] = children
		return Element{typ: typ, props: props}
	default:
		panic("typ must be a string or Component")
	}
}
