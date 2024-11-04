package gox

import (
	"context"
	"fmt"
)

func Render(c context.Context, elm Element) string {
	ctx := WithContext(c)

	return renderElem(ctx, elm)
}

func renderComp(ctx Context, comp Component, props P) Node {
	c := WithContext(ctx)
	return callComp(c, comp, props)
}

func renderElem(ctx Context, elm Element) string {
	if valid, _ := isComp(elm.typ); valid {
		node := renderComp(ctx, elm.typ, elm.props)
		return renderNode(ctx, node)
	}

	if children, ok := elm.props["children"].([]Node); ok {
		var renderedChildren string
		for _, child := range children {
			renderedChildren += renderNode(ctx, child)
		}

		return fmt.Sprintf("<%s>%s</%s>", elm.typ, renderedChildren, elm.typ)
	}

	return fmt.Sprintf("<%s></%s>", elm.typ, elm.typ)
}

func renderNode(ctx Context, node Node) string {
	switch n := node.(type) {
	case Element:
		return renderElem(ctx, n)
	default:
		switch n := node.node().(type) {
		case string:
			return n
		default:
			return fmt.Sprint(n)
		}
	}
}
