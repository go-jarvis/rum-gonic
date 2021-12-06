package main

import (
	"fmt"
	"reflect"
)

type MyString string

func main() {
	p := Person{
		name: "zhangsan",
	}

	for _, v := range []interface{}{
		p, &p, "123", 234, MyString("121233"),
	} {
		typeName(v)
	}
}

type Person struct {
	name string
}

func (p Person) Name() string {

	pname := reflect.TypeOf(p).Name()

	return pname + "+with-context+" + p.name
}

type IPerson interface {
	Name() string
}

func typeName(v interface{}) {
	rt := reflect.TypeOf(v)

	rt = deref(rt)

	tname := rt.Name()

	fmt.Println(rt.Kind().String(), "(kind)=>(name):", tname)
}

func deref(rt reflect.Type) reflect.Type {
	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		break
	}
	return rt
}
