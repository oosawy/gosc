package gox

import (
	"context"
)

func Render(elm Element) Node {
	ctx := context.Background()

	if valid, _ := isComp(elm.typ); valid {
		node := callComp(ctx, elm.typ, elm.props)
		if elm, ok := node.(Element); ok {
			return Render(elm)
		}
		return node
	} else {
		if children, ok := elm.props["children"].([]Node); ok {
			var renderedChildren []Node
			for _, child := range children {
				if childElm, ok := child.(Element); ok {
					renderedChildren = append(renderedChildren, Render(childElm))
				} else {
					renderedChildren = append(renderedChildren, child)
				}
			}
			elm.props["children"] = renderedChildren
		}

		return elm
	}
}
