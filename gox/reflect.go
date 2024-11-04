package gox

import (
	"errors"
	"reflect"
)

func isComp(value any) (bool, error) {
	compType := reflect.TypeOf(value)
	if compType.Kind() != reflect.Func {
		return false, errors.New("gox: component must be a function")
	}

	numIn := compType.NumIn()
	if numIn != 1 && numIn != 2 {
		return false, errors.New("gox: component function must have 1 or 2 parameters")
	}

	if compType.In(0) != reflect.TypeOf((*Context)(nil)).Elem() {
		return false, errors.New("gox: first parameter of component function must be gox.Context")
	}

	if compType.NumOut() != 1 {
		return false, errors.New("gox: component function must return a single value")
	}

	if !compType.Out(0).Implements(reflect.TypeOf((*Node)(nil)).Elem()) {
		return false, errors.New("gox: component function must return a Node")
	}

	return true, nil
}

func callComp(ctx Context, component any, props P) Node {
	if valid, err := isComp(component); !valid {
		panic(err)
	}

	compValue := reflect.ValueOf(component)
	compType := compValue.Type()

	var args []reflect.Value
	args = append(args, reflect.ValueOf(ctx))

	if compType.NumIn() == 2 {
		args = append(args, decodeProps(props, compType.In(1)))
	}

	result := compValue.Call(args)
	node := result[0].Interface().(Node)

	return node
}

func decodeProps(props P, typ reflect.Type) reflect.Value {
	if typ.Kind() != reflect.Struct {
		panic("gox: target must be a struct")
	}

	target := reflect.New(typ).Elem()
	propsValue := reflect.ValueOf(props)

	for i := 0; i < target.NumField(); i++ {
		field := target.Type().Field(i)
		propValue := propsValue.MapIndex(reflect.ValueOf(field.Name))

		if propValue.IsValid() && target.Field(i).CanSet() {
			target.Field(i).Set(propValue.Elem().Convert(field.Type))
		}
	}

	return target
}
