package gox

import (
	"context"
	"fmt"
)

func Render(c context.Context, elm Element) string {
	ctx := WithContext(c)

	return elm.render(ctx)
}

func (n primNode) render(ctx Context) string {
	switch n.value.(type) {
	case string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return fmt.Sprintf("%v", n.value)
	case bool:
		return ""
	default:
		panic("invalid type")
	}
}

func (c Children) render(ctx Context) string {
	var renderedChildren string
	for _, child := range c {
		renderedChildren += child.render(ctx)
	}
	return renderedChildren
}

func (e tagElement) render(ctx Context) string {
	children := e.children.render(ctx)
	return fmt.Sprintf("<%s>%s</%s>", e.tag, children, e.tag)
}

func (e compElement) render(ctx Context) string {
	node := e.comp(ctx, e.props, e.children)
	return node.render(ctx)
}
