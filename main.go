package main

import (
	"fmt"
	"time"

	x "github.com/oosawy/gosc/gox"
)

func message(c x.Context, props struct{ Name string }) x.Node {
	return x.N(fmt.Sprintf("Hello, %s! The time is %s.", props.Name, time.Now().Format("15:04:05")))
}

func app(_ x.Context) x.Node {
	return x.E("div", nil, x.E("h1", nil, x.E(message, x.P{"Name": "world"})))
}

func main() {
	tree := x.Render(x.E(app, nil))
	fmt.Printf("%+v\n", tree)
}
